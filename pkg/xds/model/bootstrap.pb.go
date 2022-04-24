// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: bootstrap.proto

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

// HeaderValue
type HeaderValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *HeaderValue) Reset() {
	*x = HeaderValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderValue) ProtoMessage() {}

func (x *HeaderValue) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderValue.ProtoReflect.Descriptor instead.
func (*HeaderValue) Descriptor() ([]byte, []int) {
	return file_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *HeaderValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HeaderValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// Metric config for otel metric.
type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable         bool  `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	PrometheusPort int64 `protobuf:"varint,2,opt,name=prometheus_port,json=prometheusPort,proto3" json:"prometheus_port,omitempty"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_bootstrap_proto_rawDescGZIP(), []int{1}
}

func (x *Metric) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *Metric) GetPrometheusPort() int64 {
	if x != nil {
		return x.PrometheusPort
	}
	return 0
}

// PprofConf config how to connect to golang pprof
type PprofConf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable  bool     `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	Address *Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *PprofConf) Reset() {
	*x = PprofConf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PprofConf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PprofConf) ProtoMessage() {}

func (x *PprofConf) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PprofConf.ProtoReflect.Descriptor instead.
func (*PprofConf) Descriptor() ([]byte, []int) {
	return file_bootstrap_proto_rawDescGZIP(), []int{2}
}

func (x *PprofConf) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *PprofConf) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

// ShutdownConfig how to shutdown server.
type ShutdownConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout      string `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	StepTimeout  string `protobuf:"bytes,2,opt,name=step_timeout,json=stepTimeout,proto3" json:"step_timeout,omitempty"`
	RejectPolicy string `protobuf:"bytes,3,opt,name=reject_policy,json=rejectPolicy,proto3" json:"reject_policy,omitempty"`
}

func (x *ShutdownConfig) Reset() {
	*x = ShutdownConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShutdownConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShutdownConfig) ProtoMessage() {}

func (x *ShutdownConfig) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShutdownConfig.ProtoReflect.Descriptor instead.
func (*ShutdownConfig) Descriptor() ([]byte, []int) {
	return file_bootstrap_proto_rawDescGZIP(), []int{3}
}

func (x *ShutdownConfig) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

func (x *ShutdownConfig) GetStepTimeout() string {
	if x != nil {
		return x.StepTimeout
	}
	return ""
}

func (x *ShutdownConfig) GetRejectPolicy() string {
	if x != nil {
		return x.RejectPolicy
	}
	return ""
}

// StaticResources
type StaticResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Listeners      []*Listener     `protobuf:"bytes,1,rep,name=listeners,proto3" json:"listeners,omitempty"`
	Clusters       []*Cluster      `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	Adapters       []*Adapter      `protobuf:"bytes,3,rep,name=adapters,proto3" json:"adapters,omitempty"`
	ShutdownConfig *ShutdownConfig `protobuf:"bytes,4,opt,name=shutdown_config,json=shutdownConfig,proto3" json:"shutdown_config,omitempty"`
	PprofConf      *PprofConf      `protobuf:"bytes,5,opt,name=pprof_conf,json=pprofConf,proto3" json:"pprof_conf,omitempty"`
}

func (x *StaticResources) Reset() {
	*x = StaticResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootstrap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticResources) ProtoMessage() {}

func (x *StaticResources) ProtoReflect() protoreflect.Message {
	mi := &file_bootstrap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticResources.ProtoReflect.Descriptor instead.
func (*StaticResources) Descriptor() ([]byte, []int) {
	return file_bootstrap_proto_rawDescGZIP(), []int{4}
}

func (x *StaticResources) GetListeners() []*Listener {
	if x != nil {
		return x.Listeners
	}
	return nil
}

func (x *StaticResources) GetClusters() []*Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *StaticResources) GetAdapters() []*Adapter {
	if x != nil {
		return x.Adapters
	}
	return nil
}

func (x *StaticResources) GetShutdownConfig() *ShutdownConfig {
	if x != nil {
		return x.ShutdownConfig
	}
	return nil
}

func (x *StaticResources) GetPprofConf() *PprofConf {
	if x != nil {
		return x.PprofConf
	}
	return nil
}

var File_bootstrap_proto protoreflect.FileDescriptor

var file_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x70, 0x69, 0x78, 0x69, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a,
	0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x65, 0x75, 0x73, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x70, 0x72, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x65, 0x75, 0x73, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x54,
	0x0a, 0x09, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x69, 0x78, 0x69, 0x75, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x72, 0x0a, 0x0e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x65, 0x70, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x65, 0x70, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0xac, 0x02, 0x0a, 0x0f, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x69, 0x78, 0x69, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x69, 0x78, 0x69, 0x75, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x08, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x69, 0x78, 0x69, 0x75, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x08,
	0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x45, 0x0a, 0x0f, 0x73, 0x68, 0x75, 0x74,
	0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x69, 0x78, 0x69, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0e, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x36, 0x0a, 0x0a, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x69, 0x78, 0x69, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x70, 0x72, 0x6f, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x09, 0x70, 0x70,
	0x72, 0x6f, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x42, 0x07, 0x5a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bootstrap_proto_rawDescOnce sync.Once
	file_bootstrap_proto_rawDescData = file_bootstrap_proto_rawDesc
)

func file_bootstrap_proto_rawDescGZIP() []byte {
	file_bootstrap_proto_rawDescOnce.Do(func() {
		file_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_bootstrap_proto_rawDescData)
	})
	return file_bootstrap_proto_rawDescData
}

var file_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bootstrap_proto_goTypes = []interface{}{
	(*HeaderValue)(nil),     // 0: pixiu.api.v1.HeaderValue
	(*Metric)(nil),          // 1: pixiu.api.v1.Metric
	(*PprofConf)(nil),       // 2: pixiu.api.v1.PprofConf
	(*ShutdownConfig)(nil),  // 3: pixiu.api.v1.ShutdownConfig
	(*StaticResources)(nil), // 4: pixiu.api.v1.StaticResources
	(*Address)(nil),         // 5: pixiu.api.v1.Address
	(*Listener)(nil),        // 6: pixiu.api.v1.Listener
	(*Cluster)(nil),         // 7: pixiu.api.v1.Cluster
	(*Adapter)(nil),         // 8: pixiu.api.v1.Adapter
}
var file_bootstrap_proto_depIdxs = []int32{
	5, // 0: pixiu.api.v1.PprofConf.address:type_name -> pixiu.api.v1.Address
	6, // 1: pixiu.api.v1.StaticResources.listeners:type_name -> pixiu.api.v1.Listener
	7, // 2: pixiu.api.v1.StaticResources.clusters:type_name -> pixiu.api.v1.Cluster
	8, // 3: pixiu.api.v1.StaticResources.adapters:type_name -> pixiu.api.v1.Adapter
	3, // 4: pixiu.api.v1.StaticResources.shutdown_config:type_name -> pixiu.api.v1.ShutdownConfig
	2, // 5: pixiu.api.v1.StaticResources.pprof_conf:type_name -> pixiu.api.v1.PprofConf
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_bootstrap_proto_init() }
func file_bootstrap_proto_init() {
	if File_bootstrap_proto != nil {
		return
	}
	file_address_proto_init()
	file_cluster_proto_init()
	file_adapter_proto_init()
	file_listener_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bootstrap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderValue); i {
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
		file_bootstrap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metric); i {
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
		file_bootstrap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PprofConf); i {
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
		file_bootstrap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShutdownConfig); i {
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
		file_bootstrap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticResources); i {
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
			RawDescriptor: file_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bootstrap_proto_goTypes,
		DependencyIndexes: file_bootstrap_proto_depIdxs,
		MessageInfos:      file_bootstrap_proto_msgTypes,
	}.Build()
	File_bootstrap_proto = out.File
	file_bootstrap_proto_rawDesc = nil
	file_bootstrap_proto_goTypes = nil
	file_bootstrap_proto_depIdxs = nil
}
