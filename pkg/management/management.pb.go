// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	ragù          v0.2.3
// source: pkg/management/management.proto

package management

import (
	core "github.com/kralicky/opni-monitoring/pkg/core"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBootstrapTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ttl *durationpb.Duration `protobuf:"bytes,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *CreateBootstrapTokenRequest) Reset() {
	*x = CreateBootstrapTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBootstrapTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBootstrapTokenRequest) ProtoMessage() {}

func (x *CreateBootstrapTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBootstrapTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateBootstrapTokenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBootstrapTokenRequest) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

type CertsInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain []*core.CertInfo `protobuf:"bytes,1,rep,name=chain,proto3" json:"chain,omitempty"`
}

func (x *CertsInfoResponse) Reset() {
	*x = CertsInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertsInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertsInfoResponse) ProtoMessage() {}

func (x *CertsInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertsInfoResponse.ProtoReflect.Descriptor instead.
func (*CertsInfoResponse) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{1}
}

func (x *CertsInfoResponse) GetChain() []*core.CertInfo {
	if x != nil {
		return x.Chain
	}
	return nil
}

type ListClustersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatchLabels *core.LabelSelector `protobuf:"bytes,1,opt,name=matchLabels,proto3" json:"matchLabels,omitempty"`
}

func (x *ListClustersRequest) Reset() {
	*x = ListClustersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClustersRequest) ProtoMessage() {}

func (x *ListClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClustersRequest.ProtoReflect.Descriptor instead.
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{2}
}

func (x *ListClustersRequest) GetMatchLabels() *core.LabelSelector {
	if x != nil {
		return x.MatchLabels
	}
	return nil
}

type EditClusterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster *core.Reference   `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Labels  map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EditClusterRequest) Reset() {
	*x = EditClusterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_management_management_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditClusterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditClusterRequest) ProtoMessage() {}

func (x *EditClusterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_management_management_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditClusterRequest.ProtoReflect.Descriptor instead.
func (*EditClusterRequest) Descriptor() ([]byte, []int) {
	return file_pkg_management_management_proto_rawDescGZIP(), []int{3}
}

func (x *EditClusterRequest) GetCluster() *core.Reference {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *EditClusterRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_pkg_management_management_proto protoreflect.FileDescriptor

var file_pkg_management_management_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a,
	0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x03,
	0x74, 0x74, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x36, 0x0a, 0x11, 0x43, 0x65, 0x72, 0x74,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x00, 0x3a, 0x00,
	0x22, 0x43, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x42, 0x00, 0x3a, 0x00, 0x22, 0xb1, 0x01, 0x0a, 0x12, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x00,
	0x12, 0x3a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x64,
	0x69, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x39, 0x0a, 0x0b,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x00, 0x32, 0xb4, 0x15, 0x0a, 0x0a, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0xb6, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x27, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0xba,
	0x3e, 0x3b, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x24, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x22,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x22, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x22, 0x2a, 0x22, 0x7d, 0x28, 0x00, 0x30,
	0x00, 0x12, 0xa0, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x5b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0xba, 0x3e, 0x39, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x22, 0x7b, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x22, 0x7d,
	0x28, 0x00, 0x30, 0x00, 0x12, 0x9b, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0xba, 0x3e, 0x31, 0x12, 0x13, 0x0a,
	0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x10, 0x01, 0x42, 0x1a, 0x7b, 0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x7d, 0x28, 0x00,
	0x30, 0x00, 0x12, 0x9a, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12,
	0x14, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0xba, 0x3e, 0x33, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x1c, 0x7b,
	0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12,
	0x9d, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x2a, 0x19, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0xba, 0x3e, 0x3b,
	0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x24, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x22,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12,
	0x94, 0x01, 0x0a, 0x09, 0x43, 0x65, 0x72, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x73, 0xba,
	0x3e, 0x30, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x19, 0x7b, 0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x73,
	0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0x8e, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x5c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0xba, 0x3e, 0x38, 0x12, 0x13, 0x0a, 0x0f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01,
	0x42, 0x21, 0x7b, 0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0xba, 0x01, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x22, 0x78, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01,
	0x2a, 0x1a, 0x21, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x69, 0x64, 0x7d, 0xba, 0x3e, 0x49, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x32, 0x7b, 0x70,
	0x75, 0x74, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x69, 0x64, 0x7d, 0x22, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x22, 0x2a, 0x22, 0x7d,
	0x28, 0x00, 0x30, 0x00, 0x12, 0x8f, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x0a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x59, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a,
	0x01, 0x2a, 0x22, 0x11, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0xba, 0x3e, 0x3a, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x23, 0x7b,
	0x70, 0x6f, 0x73, 0x74, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x22, 0x2a,
	0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0x98, 0x01, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x5d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0xba, 0x3e, 0x3a, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x23, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x22, 0x7d, 0x28, 0x00, 0x30,
	0x00, 0x12, 0x86, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x0a,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x12, 0x18, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0xba, 0x3e, 0x37, 0x12,
	0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x10, 0x01, 0x42, 0x20, 0x7b, 0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x6e,
	0x61, 0x6d, 0x65, 0x7d, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0xab, 0x01, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x67, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0xba, 0x3e, 0x41, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x2a, 0x7b, 0x70, 0x6f, 0x73, 0x74, 0x3a,
	0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x3a,
	0x22, 0x2a, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0xad, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x2a,
	0x1f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0xba, 0x3e, 0x41, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x2a, 0x7b, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0xa2, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0x11, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x22,
	0x68, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x1f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0xba, 0x3e, 0x3e, 0x12, 0x13, 0x0a, 0x0f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01,
	0x42, 0x27, 0x7b, 0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0x85, 0x01,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0xba, 0x3e,
	0x30, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x19, 0x7b, 0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22,
	0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0xa1, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x12, 0x18, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72,
	0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0xba, 0x3e, 0x37, 0x12, 0x13,
	0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x10, 0x01, 0x42, 0x20, 0x7b, 0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x62, 0x69, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x12, 0xa2, 0x01, 0x0a, 0x0d, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0xba,
	0x3e, 0x38, 0x12, 0x13, 0x0a, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x10, 0x01, 0x42, 0x21, 0x7b, 0x67, 0x65, 0x74, 0x3a, 0x22, 0x2f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x7d, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x00,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x72, 0x61, 0x6c, 0x69, 0x63, 0x6b, 0x79, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2d, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_management_management_proto_rawDescOnce sync.Once
	file_pkg_management_management_proto_rawDescData = file_pkg_management_management_proto_rawDesc
)

func file_pkg_management_management_proto_rawDescGZIP() []byte {
	file_pkg_management_management_proto_rawDescOnce.Do(func() {
		file_pkg_management_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_management_management_proto_rawDescData)
	})
	return file_pkg_management_management_proto_rawDescData
}

var file_pkg_management_management_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_management_management_proto_goTypes = []interface{}{
	(*CreateBootstrapTokenRequest)(nil), // 0: management.CreateBootstrapTokenRequest
	(*CertsInfoResponse)(nil),           // 1: management.CertsInfoResponse
	(*ListClustersRequest)(nil),         // 2: management.ListClustersRequest
	(*EditClusterRequest)(nil),          // 3: management.EditClusterRequest
	nil,                                 // 4: management.EditClusterRequest.LabelsEntry
	(*durationpb.Duration)(nil),         // 5: google.protobuf.Duration
	(*core.CertInfo)(nil),               // 6: core.CertInfo
	(*core.LabelSelector)(nil),          // 7: core.LabelSelector
	(*core.Reference)(nil),              // 8: core.Reference
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
	(*core.Role)(nil),                   // 10: core.Role
	(*core.RoleBinding)(nil),            // 11: core.RoleBinding
	(*core.SubjectAccessRequest)(nil),   // 12: core.SubjectAccessRequest
	(*core.BootstrapToken)(nil),         // 13: core.BootstrapToken
	(*core.BootstrapTokenList)(nil),     // 14: core.BootstrapTokenList
	(*core.ClusterList)(nil),            // 15: core.ClusterList
	(*core.Cluster)(nil),                // 16: core.Cluster
	(*core.RoleList)(nil),               // 17: core.RoleList
	(*core.RoleBindingList)(nil),        // 18: core.RoleBindingList
	(*core.ReferenceList)(nil),          // 19: core.ReferenceList
}
var file_pkg_management_management_proto_depIdxs = []int32{
	5,  // 0: management.CreateBootstrapTokenRequest.ttl:type_name -> google.protobuf.Duration
	6,  // 1: management.CertsInfoResponse.chain:type_name -> core.CertInfo
	7,  // 2: management.ListClustersRequest.matchLabels:type_name -> core.LabelSelector
	8,  // 3: management.EditClusterRequest.cluster:type_name -> core.Reference
	4,  // 4: management.EditClusterRequest.labels:type_name -> management.EditClusterRequest.LabelsEntry
	0,  // 5: management.Management.CreateBootstrapToken:input_type -> management.CreateBootstrapTokenRequest
	8,  // 6: management.Management.RevokeBootstrapToken:input_type -> core.Reference
	9,  // 7: management.Management.ListBootstrapTokens:input_type -> google.protobuf.Empty
	2,  // 8: management.Management.ListClusters:input_type -> management.ListClustersRequest
	8,  // 9: management.Management.DeleteCluster:input_type -> core.Reference
	9,  // 10: management.Management.CertsInfo:input_type -> google.protobuf.Empty
	8,  // 11: management.Management.GetCluster:input_type -> core.Reference
	3,  // 12: management.Management.EditCluster:input_type -> management.EditClusterRequest
	10, // 13: management.Management.CreateRole:input_type -> core.Role
	8,  // 14: management.Management.DeleteRole:input_type -> core.Reference
	8,  // 15: management.Management.GetRole:input_type -> core.Reference
	11, // 16: management.Management.CreateRoleBinding:input_type -> core.RoleBinding
	8,  // 17: management.Management.DeleteRoleBinding:input_type -> core.Reference
	8,  // 18: management.Management.GetRoleBinding:input_type -> core.Reference
	9,  // 19: management.Management.ListRoles:input_type -> google.protobuf.Empty
	9,  // 20: management.Management.ListRoleBindings:input_type -> google.protobuf.Empty
	12, // 21: management.Management.SubjectAccess:input_type -> core.SubjectAccessRequest
	13, // 22: management.Management.CreateBootstrapToken:output_type -> core.BootstrapToken
	9,  // 23: management.Management.RevokeBootstrapToken:output_type -> google.protobuf.Empty
	14, // 24: management.Management.ListBootstrapTokens:output_type -> core.BootstrapTokenList
	15, // 25: management.Management.ListClusters:output_type -> core.ClusterList
	9,  // 26: management.Management.DeleteCluster:output_type -> google.protobuf.Empty
	1,  // 27: management.Management.CertsInfo:output_type -> management.CertsInfoResponse
	16, // 28: management.Management.GetCluster:output_type -> core.Cluster
	16, // 29: management.Management.EditCluster:output_type -> core.Cluster
	9,  // 30: management.Management.CreateRole:output_type -> google.protobuf.Empty
	9,  // 31: management.Management.DeleteRole:output_type -> google.protobuf.Empty
	10, // 32: management.Management.GetRole:output_type -> core.Role
	9,  // 33: management.Management.CreateRoleBinding:output_type -> google.protobuf.Empty
	9,  // 34: management.Management.DeleteRoleBinding:output_type -> google.protobuf.Empty
	11, // 35: management.Management.GetRoleBinding:output_type -> core.RoleBinding
	17, // 36: management.Management.ListRoles:output_type -> core.RoleList
	18, // 37: management.Management.ListRoleBindings:output_type -> core.RoleBindingList
	19, // 38: management.Management.SubjectAccess:output_type -> core.ReferenceList
	22, // [22:39] is the sub-list for method output_type
	5,  // [5:22] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_management_management_proto_init() }
func file_pkg_management_management_proto_init() {
	if File_pkg_management_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_management_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBootstrapTokenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_management_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertsInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_management_management_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListClustersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_management_management_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditClusterRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_management_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_management_management_proto_goTypes,
		DependencyIndexes: file_pkg_management_management_proto_depIdxs,
		MessageInfos:      file_pkg_management_management_proto_msgTypes,
	}.Build()
	File_pkg_management_management_proto = out.File
	file_pkg_management_management_proto_rawDesc = nil
	file_pkg_management_management_proto_goTypes = nil
	file_pkg_management_management_proto_depIdxs = nil
}
