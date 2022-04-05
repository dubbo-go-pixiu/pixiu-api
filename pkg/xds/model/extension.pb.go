// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: extension.proto

package model

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message map to dubbo-go.pixiu//v1/discovery:cluster
type PixiuExtensionClusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
}

func (x *PixiuExtensionClusters) Reset() {
	*x = PixiuExtensionClusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PixiuExtensionClusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PixiuExtensionClusters) ProtoMessage() {}

func (x *PixiuExtensionClusters) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PixiuExtensionClusters.ProtoReflect.Descriptor instead.
func (*PixiuExtensionClusters) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{0}
}

func (x *PixiuExtensionClusters) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

// Message map to dubbo-go.pixiu//v1/discovery:listener
type PixiuExtensionListeners struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listeners []*Listener `protobuf:"bytes,1,rep,name=listeners,proto3" json:"listeners,omitempty"`
}

func (x *PixiuExtensionListeners) Reset() {
	*x = PixiuExtensionListeners{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PixiuExtensionListeners) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PixiuExtensionListeners) ProtoMessage() {}

func (x *PixiuExtensionListeners) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PixiuExtensionListeners.ProtoReflect.Descriptor instead.
func (*PixiuExtensionListeners) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{1}
}

func (x *PixiuExtensionListeners) GetListeners() []*Listener {
	if x != nil {
		return x.Listeners
	}
	return nil
}

var File_extension_proto protoreflect.FileDescriptor

var file_extension_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x70, 0x69, 0x78, 0x69, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b,
	0x0a, 0x16, 0x50, 0x69, 0x78, 0x69, 0x75, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x69, 0x78,
	0x69, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x22, 0x4f, 0x0a, 0x17, 0x50,
	0x69, 0x78, 0x69, 0x75, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x69, 0x78, 0x69,
	0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extension_proto_rawDescOnce sync.Once
	file_extension_proto_rawDescData = file_extension_proto_rawDesc
)

func file_extension_proto_rawDescGZIP() []byte {
	file_extension_proto_rawDescOnce.Do(func() {
		file_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_extension_proto_rawDescData)
	})
	return file_extension_proto_rawDescData
}

var file_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_extension_proto_goTypes = []interface{}{
	(*PixiuExtensionClusters)(nil),  // 0: pixiu.api.v1.PixiuExtensionClusters
	(*PixiuExtensionListeners)(nil), // 1: pixiu.api.v1.PixiuExtensionListeners
	(*Cluster)(nil),                 // 2: pixiu.api.v1.Cluster
	(*Listener)(nil),                // 3: pixiu.api.v1.Listener
}
var file_extension_proto_depIdxs = []int32{
	2, // 0: pixiu.api.v1.PixiuExtensionClusters.clusters:type_name -> pixiu.api.v1.Cluster
	3, // 1: pixiu.api.v1.PixiuExtensionListeners.listeners:type_name -> pixiu.api.v1.Listener
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_extension_proto_init() }
func file_extension_proto_init() {
	if File_extension_proto != nil {
		return
	}
	file_cluster_proto_init()
	file_listener_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PixiuExtensionClusters); i {
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
		file_extension_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PixiuExtensionListeners); i {
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
			RawDescriptor: file_extension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extension_proto_goTypes,
		DependencyIndexes: file_extension_proto_depIdxs,
		MessageInfos:      file_extension_proto_msgTypes,
	}.Build()
	File_extension_proto = out.File
	file_extension_proto_rawDesc = nil
	file_extension_proto_goTypes = nil
	file_extension_proto_depIdxs = nil
}
