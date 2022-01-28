// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - ragù               v0.2.3
// source: pkg/management/management.proto

package management

import (
	context "context"
	core "github.com/kralicky/opni-monitoring/pkg/core"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ManagementClient is the client API for Management service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagementClient interface {
	CreateBootstrapToken(ctx context.Context, in *CreateBootstrapTokenRequest, opts ...grpc.CallOption) (*core.BootstrapToken, error)
	RevokeBootstrapToken(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListBootstrapTokens(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*core.BootstrapTokenList, error)
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*core.ClusterList, error)
	DeleteCluster(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CertsInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CertsInfoResponse, error)
	GetCluster(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*core.Cluster, error)
	EditCluster(ctx context.Context, in *EditClusterRequest, opts ...grpc.CallOption) (*core.Cluster, error)
	CreateRole(ctx context.Context, in *core.Role, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRole(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRole(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*core.Role, error)
	CreateRoleBinding(ctx context.Context, in *core.RoleBinding, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteRoleBinding(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetRoleBinding(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*core.RoleBinding, error)
	ListRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*core.RoleList, error)
	ListRoleBindings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*core.RoleBindingList, error)
	SubjectAccess(ctx context.Context, in *core.SubjectAccessRequest, opts ...grpc.CallOption) (*core.ReferenceList, error)
}

type managementClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementClient(cc grpc.ClientConnInterface) ManagementClient {
	return &managementClient{cc}
}

func (c *managementClient) CreateBootstrapToken(ctx context.Context, in *CreateBootstrapTokenRequest, opts ...grpc.CallOption) (*core.BootstrapToken, error) {
	out := new(core.BootstrapToken)
	err := c.cc.Invoke(ctx, "/management.Management/CreateBootstrapToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) RevokeBootstrapToken(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/RevokeBootstrapToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) ListBootstrapTokens(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*core.BootstrapTokenList, error) {
	out := new(core.BootstrapTokenList)
	err := c.cc.Invoke(ctx, "/management.Management/ListBootstrapTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*core.ClusterList, error) {
	out := new(core.ClusterList)
	err := c.cc.Invoke(ctx, "/management.Management/ListClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) DeleteCluster(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) CertsInfo(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CertsInfoResponse, error) {
	out := new(CertsInfoResponse)
	err := c.cc.Invoke(ctx, "/management.Management/CertsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) GetCluster(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*core.Cluster, error) {
	out := new(core.Cluster)
	err := c.cc.Invoke(ctx, "/management.Management/GetCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) EditCluster(ctx context.Context, in *EditClusterRequest, opts ...grpc.CallOption) (*core.Cluster, error) {
	out := new(core.Cluster)
	err := c.cc.Invoke(ctx, "/management.Management/EditCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) CreateRole(ctx context.Context, in *core.Role, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) DeleteRole(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) GetRole(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*core.Role, error) {
	out := new(core.Role)
	err := c.cc.Invoke(ctx, "/management.Management/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) CreateRoleBinding(ctx context.Context, in *core.RoleBinding, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/CreateRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) DeleteRoleBinding(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/management.Management/DeleteRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) GetRoleBinding(ctx context.Context, in *core.Reference, opts ...grpc.CallOption) (*core.RoleBinding, error) {
	out := new(core.RoleBinding)
	err := c.cc.Invoke(ctx, "/management.Management/GetRoleBinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) ListRoles(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*core.RoleList, error) {
	out := new(core.RoleList)
	err := c.cc.Invoke(ctx, "/management.Management/ListRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) ListRoleBindings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*core.RoleBindingList, error) {
	out := new(core.RoleBindingList)
	err := c.cc.Invoke(ctx, "/management.Management/ListRoleBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) SubjectAccess(ctx context.Context, in *core.SubjectAccessRequest, opts ...grpc.CallOption) (*core.ReferenceList, error) {
	out := new(core.ReferenceList)
	err := c.cc.Invoke(ctx, "/management.Management/SubjectAccess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServer is the server API for Management service.
// All implementations must embed UnimplementedManagementServer
// for forward compatibility
type ManagementServer interface {
	CreateBootstrapToken(context.Context, *CreateBootstrapTokenRequest) (*core.BootstrapToken, error)
	RevokeBootstrapToken(context.Context, *core.Reference) (*emptypb.Empty, error)
	ListBootstrapTokens(context.Context, *emptypb.Empty) (*core.BootstrapTokenList, error)
	ListClusters(context.Context, *ListClustersRequest) (*core.ClusterList, error)
	DeleteCluster(context.Context, *core.Reference) (*emptypb.Empty, error)
	CertsInfo(context.Context, *emptypb.Empty) (*CertsInfoResponse, error)
	GetCluster(context.Context, *core.Reference) (*core.Cluster, error)
	EditCluster(context.Context, *EditClusterRequest) (*core.Cluster, error)
	CreateRole(context.Context, *core.Role) (*emptypb.Empty, error)
	DeleteRole(context.Context, *core.Reference) (*emptypb.Empty, error)
	GetRole(context.Context, *core.Reference) (*core.Role, error)
	CreateRoleBinding(context.Context, *core.RoleBinding) (*emptypb.Empty, error)
	DeleteRoleBinding(context.Context, *core.Reference) (*emptypb.Empty, error)
	GetRoleBinding(context.Context, *core.Reference) (*core.RoleBinding, error)
	ListRoles(context.Context, *emptypb.Empty) (*core.RoleList, error)
	ListRoleBindings(context.Context, *emptypb.Empty) (*core.RoleBindingList, error)
	SubjectAccess(context.Context, *core.SubjectAccessRequest) (*core.ReferenceList, error)
	mustEmbedUnimplementedManagementServer()
}

// UnimplementedManagementServer must be embedded to have forward compatible implementations.
type UnimplementedManagementServer struct {
}

func (UnimplementedManagementServer) CreateBootstrapToken(context.Context, *CreateBootstrapTokenRequest) (*core.BootstrapToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBootstrapToken not implemented")
}
func (UnimplementedManagementServer) RevokeBootstrapToken(context.Context, *core.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeBootstrapToken not implemented")
}
func (UnimplementedManagementServer) ListBootstrapTokens(context.Context, *emptypb.Empty) (*core.BootstrapTokenList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBootstrapTokens not implemented")
}
func (UnimplementedManagementServer) ListClusters(context.Context, *ListClustersRequest) (*core.ClusterList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedManagementServer) DeleteCluster(context.Context, *core.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (UnimplementedManagementServer) CertsInfo(context.Context, *emptypb.Empty) (*CertsInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CertsInfo not implemented")
}
func (UnimplementedManagementServer) GetCluster(context.Context, *core.Reference) (*core.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCluster not implemented")
}
func (UnimplementedManagementServer) EditCluster(context.Context, *EditClusterRequest) (*core.Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCluster not implemented")
}
func (UnimplementedManagementServer) CreateRole(context.Context, *core.Role) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}
func (UnimplementedManagementServer) DeleteRole(context.Context, *core.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}
func (UnimplementedManagementServer) GetRole(context.Context, *core.Reference) (*core.Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedManagementServer) CreateRoleBinding(context.Context, *core.RoleBinding) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoleBinding not implemented")
}
func (UnimplementedManagementServer) DeleteRoleBinding(context.Context, *core.Reference) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleBinding not implemented")
}
func (UnimplementedManagementServer) GetRoleBinding(context.Context, *core.Reference) (*core.RoleBinding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleBinding not implemented")
}
func (UnimplementedManagementServer) ListRoles(context.Context, *emptypb.Empty) (*core.RoleList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (UnimplementedManagementServer) ListRoleBindings(context.Context, *emptypb.Empty) (*core.RoleBindingList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoleBindings not implemented")
}
func (UnimplementedManagementServer) SubjectAccess(context.Context, *core.SubjectAccessRequest) (*core.ReferenceList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubjectAccess not implemented")
}
func (UnimplementedManagementServer) mustEmbedUnimplementedManagementServer() {}

// UnsafeManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementServer will
// result in compilation errors.
type UnsafeManagementServer interface {
	mustEmbedUnimplementedManagementServer()
}

func RegisterManagementServer(s grpc.ServiceRegistrar, srv ManagementServer) {
	s.RegisterService(&Management_ServiceDesc, srv)
}

func _Management_CreateBootstrapToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBootstrapTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).CreateBootstrapToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/CreateBootstrapToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).CreateBootstrapToken(ctx, req.(*CreateBootstrapTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_RevokeBootstrapToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).RevokeBootstrapToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/RevokeBootstrapToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).RevokeBootstrapToken(ctx, req.(*core.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_ListBootstrapTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).ListBootstrapTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/ListBootstrapTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).ListBootstrapTokens(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).DeleteCluster(ctx, req.(*core.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_CertsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).CertsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/CertsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).CertsInfo(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_GetCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).GetCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/GetCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).GetCluster(ctx, req.(*core.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_EditCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).EditCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/EditCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).EditCluster(ctx, req.(*EditClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Role)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).CreateRole(ctx, req.(*core.Role))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).DeleteRole(ctx, req.(*core.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).GetRole(ctx, req.(*core.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_CreateRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.RoleBinding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).CreateRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/CreateRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).CreateRoleBinding(ctx, req.(*core.RoleBinding))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_DeleteRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).DeleteRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/DeleteRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).DeleteRoleBinding(ctx, req.(*core.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_GetRoleBinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.Reference)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).GetRoleBinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/GetRoleBinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).GetRoleBinding(ctx, req.(*core.Reference))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/ListRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).ListRoles(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_ListRoleBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).ListRoleBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/ListRoleBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).ListRoleBindings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_SubjectAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(core.SubjectAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).SubjectAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Management/SubjectAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).SubjectAccess(ctx, req.(*core.SubjectAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Management_ServiceDesc is the grpc.ServiceDesc for Management service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Management_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "management.Management",
	HandlerType: (*ManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBootstrapToken",
			Handler:    _Management_CreateBootstrapToken_Handler,
		},
		{
			MethodName: "RevokeBootstrapToken",
			Handler:    _Management_RevokeBootstrapToken_Handler,
		},
		{
			MethodName: "ListBootstrapTokens",
			Handler:    _Management_ListBootstrapTokens_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _Management_ListClusters_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _Management_DeleteCluster_Handler,
		},
		{
			MethodName: "CertsInfo",
			Handler:    _Management_CertsInfo_Handler,
		},
		{
			MethodName: "GetCluster",
			Handler:    _Management_GetCluster_Handler,
		},
		{
			MethodName: "EditCluster",
			Handler:    _Management_EditCluster_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _Management_CreateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _Management_DeleteRole_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _Management_GetRole_Handler,
		},
		{
			MethodName: "CreateRoleBinding",
			Handler:    _Management_CreateRoleBinding_Handler,
		},
		{
			MethodName: "DeleteRoleBinding",
			Handler:    _Management_DeleteRoleBinding_Handler,
		},
		{
			MethodName: "GetRoleBinding",
			Handler:    _Management_GetRoleBinding_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _Management_ListRoles_Handler,
		},
		{
			MethodName: "ListRoleBindings",
			Handler:    _Management_ListRoleBindings_Handler,
		},
		{
			MethodName: "SubjectAccess",
			Handler:    _Management_SubjectAccess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/management/management.proto",
}
