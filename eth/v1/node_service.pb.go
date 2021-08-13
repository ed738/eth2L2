// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: eth/v1/node_service.proto

package v1

import (
	context "context"
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/protoc-gen-go/descriptor"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_eth_v1_node_service_proto protoreflect.FileDescriptor

var file_eth_v1_node_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xde, 0x05, 0x0a, 0x0a, 0x42, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x67, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x66, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1d, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x12, 0x6c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50,
	0x65, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x68, 0x0a, 0x09, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x67, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x12, 0x64, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x20, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x65, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x58, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x7a, 0x0a, 0x13, 0x6f, 0x72, 0x67,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x42, 0x10, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x70,
	0x72, 0x79, 0x73, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0xaa, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74, 0x68,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45,
	0x74, 0x68, 0x5c, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_eth_v1_node_service_proto_goTypes = []interface{}{
	(*empty.Empty)(nil),       // 0: google.protobuf.Empty
	(*PeersRequest)(nil),      // 1: ethereum.eth.v1.PeersRequest
	(*PeerRequest)(nil),       // 2: ethereum.eth.v1.PeerRequest
	(*IdentityResponse)(nil),  // 3: ethereum.eth.v1.IdentityResponse
	(*PeersResponse)(nil),     // 4: ethereum.eth.v1.PeersResponse
	(*PeerResponse)(nil),      // 5: ethereum.eth.v1.PeerResponse
	(*PeerCountResponse)(nil), // 6: ethereum.eth.v1.PeerCountResponse
	(*SyncingResponse)(nil),   // 7: ethereum.eth.v1.SyncingResponse
	(*VersionResponse)(nil),   // 8: ethereum.eth.v1.VersionResponse
}
var file_eth_v1_node_service_proto_depIdxs = []int32{
	0, // 0: ethereum.eth.v1.BeaconNode.GetIdentity:input_type -> google.protobuf.Empty
	1, // 1: ethereum.eth.v1.BeaconNode.ListPeers:input_type -> ethereum.eth.v1.PeersRequest
	2, // 2: ethereum.eth.v1.BeaconNode.GetPeer:input_type -> ethereum.eth.v1.PeerRequest
	0, // 3: ethereum.eth.v1.BeaconNode.PeerCount:input_type -> google.protobuf.Empty
	0, // 4: ethereum.eth.v1.BeaconNode.GetSyncStatus:input_type -> google.protobuf.Empty
	0, // 5: ethereum.eth.v1.BeaconNode.GetVersion:input_type -> google.protobuf.Empty
	0, // 6: ethereum.eth.v1.BeaconNode.GetHealth:input_type -> google.protobuf.Empty
	3, // 7: ethereum.eth.v1.BeaconNode.GetIdentity:output_type -> ethereum.eth.v1.IdentityResponse
	4, // 8: ethereum.eth.v1.BeaconNode.ListPeers:output_type -> ethereum.eth.v1.PeersResponse
	5, // 9: ethereum.eth.v1.BeaconNode.GetPeer:output_type -> ethereum.eth.v1.PeerResponse
	6, // 10: ethereum.eth.v1.BeaconNode.PeerCount:output_type -> ethereum.eth.v1.PeerCountResponse
	7, // 11: ethereum.eth.v1.BeaconNode.GetSyncStatus:output_type -> ethereum.eth.v1.SyncingResponse
	8, // 12: ethereum.eth.v1.BeaconNode.GetVersion:output_type -> ethereum.eth.v1.VersionResponse
	0, // 13: ethereum.eth.v1.BeaconNode.GetHealth:output_type -> google.protobuf.Empty
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eth_v1_node_service_proto_init() }
func file_eth_v1_node_service_proto_init() {
	if File_eth_v1_node_service_proto != nil {
		return
	}
	file_eth_v1_node_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eth_v1_node_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eth_v1_node_service_proto_goTypes,
		DependencyIndexes: file_eth_v1_node_service_proto_depIdxs,
	}.Build()
	File_eth_v1_node_service_proto = out.File
	file_eth_v1_node_service_proto_rawDesc = nil
	file_eth_v1_node_service_proto_goTypes = nil
	file_eth_v1_node_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BeaconNodeClient is the client API for BeaconNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconNodeClient interface {
	GetIdentity(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IdentityResponse, error)
	ListPeers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersResponse, error)
	GetPeer(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerResponse, error)
	PeerCount(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PeerCountResponse, error)
	GetSyncStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SyncingResponse, error)
	GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type beaconNodeClient struct {
	cc grpc.ClientConnInterface
}

func NewBeaconNodeClient(cc grpc.ClientConnInterface) BeaconNodeClient {
	return &beaconNodeClient{cc}
}

func (c *beaconNodeClient) GetIdentity(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IdentityResponse, error) {
	out := new(IdentityResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconNode/GetIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeClient) ListPeers(ctx context.Context, in *PeersRequest, opts ...grpc.CallOption) (*PeersResponse, error) {
	out := new(PeersResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconNode/ListPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeClient) GetPeer(ctx context.Context, in *PeerRequest, opts ...grpc.CallOption) (*PeerResponse, error) {
	out := new(PeerResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconNode/GetPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeClient) PeerCount(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PeerCountResponse, error) {
	out := new(PeerCountResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconNode/PeerCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeClient) GetSyncStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SyncingResponse, error) {
	out := new(SyncingResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconNode/GetSyncStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeClient) GetVersion(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconNode/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconNodeClient) GetHealth(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1.BeaconNode/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BeaconNodeServer is the server API for BeaconNode service.
type BeaconNodeServer interface {
	GetIdentity(context.Context, *empty.Empty) (*IdentityResponse, error)
	ListPeers(context.Context, *PeersRequest) (*PeersResponse, error)
	GetPeer(context.Context, *PeerRequest) (*PeerResponse, error)
	PeerCount(context.Context, *empty.Empty) (*PeerCountResponse, error)
	GetSyncStatus(context.Context, *empty.Empty) (*SyncingResponse, error)
	GetVersion(context.Context, *empty.Empty) (*VersionResponse, error)
	GetHealth(context.Context, *empty.Empty) (*empty.Empty, error)
}

// UnimplementedBeaconNodeServer can be embedded to have forward compatible implementations.
type UnimplementedBeaconNodeServer struct {
}

func (*UnimplementedBeaconNodeServer) GetIdentity(context.Context, *empty.Empty) (*IdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdentity not implemented")
}
func (*UnimplementedBeaconNodeServer) ListPeers(context.Context, *PeersRequest) (*PeersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPeers not implemented")
}
func (*UnimplementedBeaconNodeServer) GetPeer(context.Context, *PeerRequest) (*PeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeer not implemented")
}
func (*UnimplementedBeaconNodeServer) PeerCount(context.Context, *empty.Empty) (*PeerCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeerCount not implemented")
}
func (*UnimplementedBeaconNodeServer) GetSyncStatus(context.Context, *empty.Empty) (*SyncingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSyncStatus not implemented")
}
func (*UnimplementedBeaconNodeServer) GetVersion(context.Context, *empty.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (*UnimplementedBeaconNodeServer) GetHealth(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}

func RegisterBeaconNodeServer(s *grpc.Server, srv BeaconNodeServer) {
	s.RegisterService(&_BeaconNode_serviceDesc, srv)
}

func _BeaconNode_GetIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeServer).GetIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconNode/GetIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeServer).GetIdentity(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNode_ListPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeServer).ListPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconNode/ListPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeServer).ListPeers(ctx, req.(*PeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNode_GetPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeServer).GetPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconNode/GetPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeServer).GetPeer(ctx, req.(*PeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNode_PeerCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeServer).PeerCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconNode/PeerCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeServer).PeerCount(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNode_GetSyncStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeServer).GetSyncStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconNode/GetSyncStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeServer).GetSyncStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNode_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconNode/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeServer).GetVersion(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BeaconNode_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconNodeServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1.BeaconNode/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconNodeServer).GetHealth(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _BeaconNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1.BeaconNode",
	HandlerType: (*BeaconNodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetIdentity",
			Handler:    _BeaconNode_GetIdentity_Handler,
		},
		{
			MethodName: "ListPeers",
			Handler:    _BeaconNode_ListPeers_Handler,
		},
		{
			MethodName: "GetPeer",
			Handler:    _BeaconNode_GetPeer_Handler,
		},
		{
			MethodName: "PeerCount",
			Handler:    _BeaconNode_PeerCount_Handler,
		},
		{
			MethodName: "GetSyncStatus",
			Handler:    _BeaconNode_GetSyncStatus_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _BeaconNode_GetVersion_Handler,
		},
		{
			MethodName: "GetHealth",
			Handler:    _BeaconNode_GetHealth_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eth/v1/node_service.proto",
}
