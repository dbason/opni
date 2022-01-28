package management

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kralicky/opni-monitoring/pkg/config/v1beta1"
	"github.com/kralicky/opni-monitoring/pkg/core"
	"github.com/kralicky/opni-monitoring/pkg/logger"
	"github.com/kralicky/opni-monitoring/pkg/pkp"
	"github.com/kralicky/opni-monitoring/pkg/rbac"
	"github.com/kralicky/opni-monitoring/pkg/storage"
	"github.com/kralicky/opni-monitoring/pkg/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	DefaultTokenTTL = 2 * time.Minute
)

type Server struct {
	UnimplementedManagementServer
	ManagementServerOptions
	config       *v1beta1.ManagementSpec
	logger       *zap.SugaredLogger
	rbacProvider rbac.Provider
}

var _ ManagementServer = (*Server)(nil)

type ManagementServerOptions struct {
	tokenStore   storage.TokenStore
	clusterStore storage.ClusterStore
	rbacStore    storage.RBACStore
	tlsConfig    *tls.Config
}

type ManagementServerOption func(*ManagementServerOptions)

func (o *ManagementServerOptions) Apply(opts ...ManagementServerOption) {
	for _, op := range opts {
		op(o)
	}
}

func TokenStore(tokenStore storage.TokenStore) ManagementServerOption {
	return func(o *ManagementServerOptions) {
		o.tokenStore = tokenStore
	}
}

func ClusterStore(tenantStore storage.ClusterStore) ManagementServerOption {
	return func(o *ManagementServerOptions) {
		o.clusterStore = tenantStore
	}
}

func RBACStore(rbacStore storage.RBACStore) ManagementServerOption {
	return func(o *ManagementServerOptions) {
		o.rbacStore = rbacStore
	}
}

func TLSConfig(config *tls.Config) ManagementServerOption {
	return func(o *ManagementServerOptions) {
		o.tlsConfig = config
	}
}

func NewServer(conf *v1beta1.ManagementSpec, opts ...ManagementServerOption) *Server {
	options := ManagementServerOptions{}
	options.Apply(opts...)
	if options.tokenStore == nil {
		panic("token store is required")
	}
	if options.clusterStore == nil {
		panic("cluster store is required")
	}
	if options.rbacStore == nil {
		panic("rbac store is required")
	}
	return &Server{
		ManagementServerOptions: options,
		config:                  conf,
		logger:                  logger.New().Named("mgmt"),
		rbacProvider:            storage.NewRBACProvider(options.rbacStore, options.clusterStore),
	}
}

func (m *Server) ListenAndServe(ctx context.Context) error {
	if m.config.GRPCListenAddress == "" {
		return errors.New("GRPCListenAddress not configured")
	}
	lg := m.logger
	listener, err := util.NewProtocolListener(ctx, m.config.GRPCListenAddress)
	if err != nil {
		return err
	}
	m.logger.With(
		"address", listener.Addr().String(),
	).Info("management gRPC server starting")
	srv := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	RegisterManagementServer(srv, m)
	go func() {
		<-ctx.Done()
		srv.Stop()
	}()
	if httpAddr := m.config.HTTPListenAddress; httpAddr != "" {
		m.logger.With(
			"address", httpAddr,
		).Info("management HTTP server starting")
		mux := runtime.NewServeMux()
		if err := RegisterManagementHandlerFromEndpoint(ctx, mux, listener.Addr().String(),
			[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}); err != nil {
			return err
		}
		server := &http.Server{
			Addr:    httpAddr,
			Handler: mux,
		}
		defer func() {
			if err := server.Close(); err != nil {
				lg.Errorw("failed to close http gateway", "error", err)
			}
		}()
		go func() {
			if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				lg.With(
					zap.Error(err),
				).Error("http gateway exited with error")
			}
		}()
	}
	return srv.Serve(listener)
}

func (m *Server) CertsInfo(ctx context.Context, _ *emptypb.Empty) (*CertsInfoResponse, error) {
	resp := &CertsInfoResponse{
		Chain: []*core.CertInfo{},
	}
	for _, tlsCert := range m.tlsConfig.Certificates[:1] {
		for _, der := range tlsCert.Certificate {
			cert, err := x509.ParseCertificate(der)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			resp.Chain = append(resp.Chain, &core.CertInfo{
				Issuer:      cert.Issuer.String(),
				Subject:     cert.Subject.String(),
				IsCA:        cert.IsCA,
				NotBefore:   cert.NotBefore.Format(time.RFC3339),
				NotAfter:    cert.NotAfter.Format(time.RFC3339),
				Fingerprint: pkp.NewSha256(cert).Encode(),
			})
		}
	}
	return resp, nil
}

func (s *Server) SubjectAccess(ctx context.Context, sar *core.SubjectAccessRequest) (*core.ReferenceList, error) {
	return s.rbacProvider.SubjectAccess(ctx, sar)
}
