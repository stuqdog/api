// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: service/discovery/v1/discovery.proto

package v1

import (
	v1 "go.viam.com/api/app/v1"
	v11 "go.viam.com/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DiscoverResourcesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the discover service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Additional arguments to the method
	Extra *structpb.Struct `protobuf:"bytes,99,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *DiscoverResourcesRequest) Reset() {
	*x = DiscoverResourcesRequest{}
	mi := &file_service_discovery_v1_discovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverResourcesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverResourcesRequest) ProtoMessage() {}

func (x *DiscoverResourcesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_discovery_v1_discovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverResourcesRequest.ProtoReflect.Descriptor instead.
func (*DiscoverResourcesRequest) Descriptor() ([]byte, []int) {
	return file_service_discovery_v1_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *DiscoverResourcesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DiscoverResourcesRequest) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

type DiscoverResourcesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of ComponentConfigs that describe the components found by a discover service.
	Discoveries []*v1.ComponentConfig `protobuf:"bytes,1,rep,name=discoveries,proto3" json:"discoveries,omitempty"`
}

func (x *DiscoverResourcesResponse) Reset() {
	*x = DiscoverResourcesResponse{}
	mi := &file_service_discovery_v1_discovery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverResourcesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverResourcesResponse) ProtoMessage() {}

func (x *DiscoverResourcesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_discovery_v1_discovery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverResourcesResponse.ProtoReflect.Descriptor instead.
func (*DiscoverResourcesResponse) Descriptor() ([]byte, []int) {
	return file_service_discovery_v1_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoverResourcesResponse) GetDiscoveries() []*v1.ComponentConfig {
	if x != nil {
		return x.Discoveries
	}
	return nil
}

var File_service_discovery_v1_discovery_proto protoreflect.FileDescriptor

var file_service_discovery_v1_discovery_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x1a, 0x12, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x18, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x18, 0x63, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x5b, 0x0a, 0x19, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x32, 0xcf, 0x02, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xad, 0x01, 0x0a, 0x11, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x33, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x12, 0x8a, 0x01, 0x0a, 0x09, 0x44,
	0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x20, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x32, 0x22, 0x30, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x64, 0x6f, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x44, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x5a, 0x23, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x79, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_discovery_v1_discovery_proto_rawDescOnce sync.Once
	file_service_discovery_v1_discovery_proto_rawDescData = file_service_discovery_v1_discovery_proto_rawDesc
)

func file_service_discovery_v1_discovery_proto_rawDescGZIP() []byte {
	file_service_discovery_v1_discovery_proto_rawDescOnce.Do(func() {
		file_service_discovery_v1_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_discovery_v1_discovery_proto_rawDescData)
	})
	return file_service_discovery_v1_discovery_proto_rawDescData
}

var file_service_discovery_v1_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_service_discovery_v1_discovery_proto_goTypes = []any{
	(*DiscoverResourcesRequest)(nil),  // 0: viam.service.discovery.v1.DiscoverResourcesRequest
	(*DiscoverResourcesResponse)(nil), // 1: viam.service.discovery.v1.DiscoverResourcesResponse
	(*structpb.Struct)(nil),           // 2: google.protobuf.Struct
	(*v1.ComponentConfig)(nil),        // 3: viam.app.v1.ComponentConfig
	(*v11.DoCommandRequest)(nil),      // 4: viam.common.v1.DoCommandRequest
	(*v11.DoCommandResponse)(nil),     // 5: viam.common.v1.DoCommandResponse
}
var file_service_discovery_v1_discovery_proto_depIdxs = []int32{
	2, // 0: viam.service.discovery.v1.DiscoverResourcesRequest.extra:type_name -> google.protobuf.Struct
	3, // 1: viam.service.discovery.v1.DiscoverResourcesResponse.discoveries:type_name -> viam.app.v1.ComponentConfig
	0, // 2: viam.service.discovery.v1.DiscoveryService.DiscoverResources:input_type -> viam.service.discovery.v1.DiscoverResourcesRequest
	4, // 3: viam.service.discovery.v1.DiscoveryService.DoCommand:input_type -> viam.common.v1.DoCommandRequest
	1, // 4: viam.service.discovery.v1.DiscoveryService.DiscoverResources:output_type -> viam.service.discovery.v1.DiscoverResourcesResponse
	5, // 5: viam.service.discovery.v1.DiscoveryService.DoCommand:output_type -> viam.common.v1.DoCommandResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_discovery_v1_discovery_proto_init() }
func file_service_discovery_v1_discovery_proto_init() {
	if File_service_discovery_v1_discovery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_discovery_v1_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_discovery_v1_discovery_proto_goTypes,
		DependencyIndexes: file_service_discovery_v1_discovery_proto_depIdxs,
		MessageInfos:      file_service_discovery_v1_discovery_proto_msgTypes,
	}.Build()
	File_service_discovery_v1_discovery_proto = out.File
	file_service_discovery_v1_discovery_proto_rawDesc = nil
	file_service_discovery_v1_discovery_proto_goTypes = nil
	file_service_discovery_v1_discovery_proto_depIdxs = nil
}