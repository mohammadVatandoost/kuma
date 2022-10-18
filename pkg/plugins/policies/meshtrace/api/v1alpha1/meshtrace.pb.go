// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.0
// source: pkg/plugins/policies/meshtrace/api/v1alpha1/meshtrace.proto

package v1alpha1

import (
	v1alpha1 "github.com/kumahq/kuma/api/common/v1alpha1"
	_ "github.com/kumahq/kuma/api/mesh"
	_ "github.com/kumahq/protoc-gen-kumadoc/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MeshTrace allows users to enable request tracing between services in the mesh
// and sending these traces to a third party storage.
type MeshTrace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TargetRef is a reference to the resource the policy takes an effect on.
	// The resource could be either a real store object or virtual resource
	// defined inplace.
	TargetRef *v1alpha1.TargetRef `protobuf:"bytes,1,opt,name=targetRef,proto3" json:"targetRef,omitempty"`
	// MeshTrace configuration.
	Default *MeshTrace_Conf `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *MeshTrace) Reset() {
	*x = MeshTrace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace) ProtoMessage() {}

func (x *MeshTrace) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace.ProtoReflect.Descriptor instead.
func (*MeshTrace) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0}
}

func (x *MeshTrace) GetTargetRef() *v1alpha1.TargetRef {
	if x != nil {
		return x.TargetRef
	}
	return nil
}

func (x *MeshTrace) GetDefault() *MeshTrace_Conf {
	if x != nil {
		return x.Default
	}
	return nil
}

// Datadog tracing backend configuration.
type MeshTrace_DatadogBackend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address of Datadog collector, only host and port are allowed (no paths,
	// fragments etc.)
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Determines if datadog service name should be split based on traffic
	// direction and destination. For example, with `splitService: true` and a
	// `backend` service that communicates with a couple of databases, you would
	// get service names like `backend_INBOUND`, `backend_OUTBOUND_db1`, and
	// `backend_OUTBOUND_db2` in Datadog. Default: false
	SplitService bool `protobuf:"varint,2,opt,name=splitService,proto3" json:"splitService,omitempty"`
}

func (x *MeshTrace_DatadogBackend) Reset() {
	*x = MeshTrace_DatadogBackend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace_DatadogBackend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace_DatadogBackend) ProtoMessage() {}

func (x *MeshTrace_DatadogBackend) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace_DatadogBackend.ProtoReflect.Descriptor instead.
func (*MeshTrace_DatadogBackend) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MeshTrace_DatadogBackend) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MeshTrace_DatadogBackend) GetSplitService() bool {
	if x != nil {
		return x.SplitService
	}
	return false
}

// Zipkin tracing backend configuration.
type MeshTrace_ZipkinBackend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address of Zipkin collector.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Generate 128bit traces. Default: false
	TraceId128Bit bool `protobuf:"varint,2,opt,name=traceId128bit,proto3" json:"traceId128bit,omitempty"`
	// Version of the API. values: httpJson, httpProto. Default:
	// httpJson see
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L66
	ApiVersion string `protobuf:"bytes,3,opt,name=apiVersion,proto3" json:"apiVersion,omitempty"`
	// Determines whether client and server spans will share the same span
	// context. Default: true.
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/trace/v3/zipkin.proto#L63
	SharedSpanContext *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=sharedSpanContext,proto3" json:"sharedSpanContext,omitempty"`
}

func (x *MeshTrace_ZipkinBackend) Reset() {
	*x = MeshTrace_ZipkinBackend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace_ZipkinBackend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace_ZipkinBackend) ProtoMessage() {}

func (x *MeshTrace_ZipkinBackend) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace_ZipkinBackend.ProtoReflect.Descriptor instead.
func (*MeshTrace_ZipkinBackend) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0, 1}
}

func (x *MeshTrace_ZipkinBackend) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MeshTrace_ZipkinBackend) GetTraceId128Bit() bool {
	if x != nil {
		return x.TraceId128Bit
	}
	return false
}

func (x *MeshTrace_ZipkinBackend) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *MeshTrace_ZipkinBackend) GetSharedSpanContext() *wrapperspb.BoolValue {
	if x != nil {
		return x.SharedSpanContext
	}
	return nil
}

// Only one of zipkin or datadog can be used.
type MeshTrace_Backend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Zipkin backend configuration.
	Zipkin *MeshTrace_ZipkinBackend `protobuf:"bytes,1,opt,name=zipkin,proto3" json:"zipkin,omitempty"`
	// Datadog backend configuration.
	Datadog *MeshTrace_DatadogBackend `protobuf:"bytes,2,opt,name=datadog,proto3" json:"datadog,omitempty"`
}

func (x *MeshTrace_Backend) Reset() {
	*x = MeshTrace_Backend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace_Backend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace_Backend) ProtoMessage() {}

func (x *MeshTrace_Backend) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace_Backend.ProtoReflect.Descriptor instead.
func (*MeshTrace_Backend) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0, 2}
}

func (x *MeshTrace_Backend) GetZipkin() *MeshTrace_ZipkinBackend {
	if x != nil {
		return x.Zipkin
	}
	return nil
}

func (x *MeshTrace_Backend) GetDatadog() *MeshTrace_DatadogBackend {
	if x != nil {
		return x.Datadog
	}
	return nil
}

// Wrapper type.
type MeshTrace_UInt32Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The uint32 value.
	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MeshTrace_UInt32Value) Reset() {
	*x = MeshTrace_UInt32Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace_UInt32Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace_UInt32Value) ProtoMessage() {}

func (x *MeshTrace_UInt32Value) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace_UInt32Value.ProtoReflect.Descriptor instead.
func (*MeshTrace_UInt32Value) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0, 3}
}

func (x *MeshTrace_UInt32Value) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Sampling configuration.
type MeshTrace_Sampling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target percentage of requests will be traced
	// after all other sampling checks have been applied (client, force tracing,
	// random sampling). This field functions as an upper limit on the total
	// configured sampling rate. For instance, setting client_sampling to 100%
	// but overall_sampling to 1% will result in only 1% of client requests with
	// the appropriate headers to be force traced. Default: 100% Mirror of
	// overall_sampling in Envoy
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L142-L150
	Overall *MeshTrace_UInt32Value `protobuf:"bytes,1,opt,name=overall,proto3" json:"overall,omitempty"`
	// Target percentage of requests that will be force traced if the
	// 'x-client-trace-id' header is set. Default: 100% Mirror of
	// client_sampling in Envoy
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L127-L133
	Client *MeshTrace_UInt32Value `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	// Target percentage of requests that will be randomly selected for trace
	// generation, if not requested by the client or not forced. Default: 100%
	// Mirror of random_sampling in Envoy
	// https://github.com/envoyproxy/envoy/blob/v1.22.0/api/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.proto#L135-L140
	Random *MeshTrace_UInt32Value `protobuf:"bytes,3,opt,name=random,proto3" json:"random,omitempty"`
}

func (x *MeshTrace_Sampling) Reset() {
	*x = MeshTrace_Sampling{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace_Sampling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace_Sampling) ProtoMessage() {}

func (x *MeshTrace_Sampling) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace_Sampling.ProtoReflect.Descriptor instead.
func (*MeshTrace_Sampling) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0, 4}
}

func (x *MeshTrace_Sampling) GetOverall() *MeshTrace_UInt32Value {
	if x != nil {
		return x.Overall
	}
	return nil
}

func (x *MeshTrace_Sampling) GetClient() *MeshTrace_UInt32Value {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *MeshTrace_Sampling) GetRandom() *MeshTrace_UInt32Value {
	if x != nil {
		return x.Random
	}
	return nil
}

// Tag taken from a header configuration.
type MeshTrace_HeaderTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the header.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Default value to use if header is missing.
	// If the default is missing and there is no value the tag will not be
	// included.
	Default string `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
}

func (x *MeshTrace_HeaderTag) Reset() {
	*x = MeshTrace_HeaderTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace_HeaderTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace_HeaderTag) ProtoMessage() {}

func (x *MeshTrace_HeaderTag) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace_HeaderTag.ProtoReflect.Descriptor instead.
func (*MeshTrace_HeaderTag) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0, 5}
}

func (x *MeshTrace_HeaderTag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MeshTrace_HeaderTag) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

// Custom tags configuration.
// Only one of literal or header can be used.
type MeshTrace_Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the tag.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Tag taken from literal value.
	Literal string `protobuf:"bytes,2,opt,name=literal,proto3" json:"literal,omitempty"`
	// Tag taken from a header.
	Header *MeshTrace_HeaderTag `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *MeshTrace_Tag) Reset() {
	*x = MeshTrace_Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace_Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace_Tag) ProtoMessage() {}

func (x *MeshTrace_Tag) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace_Tag.ProtoReflect.Descriptor instead.
func (*MeshTrace_Tag) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0, 6}
}

func (x *MeshTrace_Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MeshTrace_Tag) GetLiteral() string {
	if x != nil {
		return x.Literal
	}
	return ""
}

func (x *MeshTrace_Tag) GetHeader() *MeshTrace_HeaderTag {
	if x != nil {
		return x.Header
	}
	return nil
}

type MeshTrace_Conf struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A one element array of backend definition.
	// Envoy allows configuring only 1 backend, so the natural way of
	// representing that would be just one object. Unfortunately due to the
	// reasons explained in MADR 009-tracing-policy this has to be a one element
	// array for now.
	// +optional
	// +nullable
	Backends []*MeshTrace_Backend `protobuf:"bytes,1,rep,name=backends,proto3" json:"backends"`
	// Sampling configuration.
	// Sampling is the process by which a decision is made on whether to
	// process/export a span or not.
	Sampling *MeshTrace_Sampling `protobuf:"bytes,2,opt,name=sampling,proto3" json:"sampling,omitempty"`
	// Custom tags configuration. You can add custom tags to traces based on
	// headers or literal values.
	// +optional
	// +nullable
	Tags []*MeshTrace_Tag `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags"`
}

func (x *MeshTrace_Conf) Reset() {
	*x = MeshTrace_Conf{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeshTrace_Conf) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeshTrace_Conf) ProtoMessage() {}

func (x *MeshTrace_Conf) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeshTrace_Conf.ProtoReflect.Descriptor instead.
func (*MeshTrace_Conf) Descriptor() ([]byte, []int) {
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP(), []int{0, 7}
}

func (x *MeshTrace_Conf) GetBackends() []*MeshTrace_Backend {
	if x != nil {
		return x.Backends
	}
	return nil
}

func (x *MeshTrace_Conf) GetSampling() *MeshTrace_Sampling {
	if x != nil {
		return x.Sampling
	}
	return nil
}

func (x *MeshTrace_Conf) GetTags() []*MeshTrace_Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto protoreflect.FileDescriptor

var file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x6b,
	0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6b, 0x75,
	0x6d, 0x61, 0x2d, 0x64, 0x6f, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x95, 0x0b, 0x0a, 0x09, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x52, 0x65, 0x66, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x66,
	0x12, 0x52, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x38, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73,
	0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x52, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x1a, 0x4c, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x42,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0xbd, 0x01, 0x0a, 0x0d, 0x5a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x42, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x31, 0x32, 0x38, 0x62, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x31, 0x32, 0x38, 0x62,
	0x69, 0x74, 0x12, 0x24, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x0a, 0x61, 0x70,
	0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x11, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x53, 0x70, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x53, 0x70, 0x61, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x1a, 0xc2, 0x01, 0x0a, 0x07, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x12, 0x59,
	0x0a, 0x06, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41,
	0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x5a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x52, 0x06, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x12, 0x5c, 0x0a, 0x07, 0x64, 0x61, 0x74,
	0x61, 0x64, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x6b, 0x75, 0x6d,
	0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x42, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x52, 0x07,
	0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x1a, 0x23, 0x0a, 0x0b, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x97, 0x02, 0x0a,
	0x08, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x59, 0x0a, 0x07, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x6b, 0x75, 0x6d,
	0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x6c, 0x6c, 0x12, 0x57, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x57, 0x0a,
	0x06, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e,
	0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x72, 0x61,
	0x63, 0x65, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x1a, 0x3f, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x54, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0x88, 0xb5, 0x18, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x1a, 0x90, 0x01, 0x0a, 0x03, 0x54, 0x61, 0x67, 0x12,
	0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0x88,
	0xb5, 0x18, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x74,
	0x65, 0x72, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x74, 0x65,
	0x72, 0x61, 0x6c, 0x12, 0x55, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x54,
	0x61, 0x67, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x86, 0x02, 0x0a, 0x04, 0x43,
	0x6f, 0x6e, 0x66, 0x12, 0x57, 0x0a, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x52, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x58, 0x0a, 0x08,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x4b, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x6b, 0x75, 0x6d, 0x61, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x73, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x4d, 0x65, 0x73, 0x68, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x3a, 0x06, 0xb2, 0x8c, 0x89, 0xa6, 0x01, 0x00, 0x42, 0x62, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x68, 0x71,
	0x2f, 0x6b, 0x75, 0x6d, 0x61, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x8a, 0xb5, 0x18, 0x1a, 0x50, 0x01, 0xa2, 0x01, 0x09, 0x4d, 0x65, 0x73, 0x68, 0x54, 0x72,
	0x61, 0x63, 0x65, 0xf2, 0x01, 0x09, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x72, 0x61, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescOnce sync.Once
	file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescData = file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDesc
)

func file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescGZIP() []byte {
	file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescOnce.Do(func() {
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescData)
	})
	return file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDescData
}

var file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_goTypes = []interface{}{
	(*MeshTrace)(nil),                // 0: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace
	(*MeshTrace_DatadogBackend)(nil), // 1: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.DatadogBackend
	(*MeshTrace_ZipkinBackend)(nil),  // 2: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.ZipkinBackend
	(*MeshTrace_Backend)(nil),        // 3: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Backend
	(*MeshTrace_UInt32Value)(nil),    // 4: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.UInt32Value
	(*MeshTrace_Sampling)(nil),       // 5: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Sampling
	(*MeshTrace_HeaderTag)(nil),      // 6: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.HeaderTag
	(*MeshTrace_Tag)(nil),            // 7: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Tag
	(*MeshTrace_Conf)(nil),           // 8: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Conf
	(*v1alpha1.TargetRef)(nil),       // 9: kuma.common.v1alpha1.TargetRef
	(*wrapperspb.BoolValue)(nil),     // 10: google.protobuf.BoolValue
}
var file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_depIdxs = []int32{
	9,  // 0: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.targetRef:type_name -> kuma.common.v1alpha1.TargetRef
	8,  // 1: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.default:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Conf
	10, // 2: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.ZipkinBackend.sharedSpanContext:type_name -> google.protobuf.BoolValue
	2,  // 3: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Backend.zipkin:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.ZipkinBackend
	1,  // 4: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Backend.datadog:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.DatadogBackend
	4,  // 5: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Sampling.overall:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.UInt32Value
	4,  // 6: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Sampling.client:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.UInt32Value
	4,  // 7: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Sampling.random:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.UInt32Value
	6,  // 8: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Tag.header:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.HeaderTag
	3,  // 9: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Conf.backends:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Backend
	5,  // 10: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Conf.sampling:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Sampling
	7,  // 11: kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Conf.tags:type_name -> kuma.plugins.policies.meshtrace.v1alpha1.MeshTrace.Tag
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_init() }
func file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_init() {
	if File_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace); i {
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
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace_DatadogBackend); i {
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
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace_ZipkinBackend); i {
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
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace_Backend); i {
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
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace_UInt32Value); i {
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
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace_Sampling); i {
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
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace_HeaderTag); i {
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
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace_Tag); i {
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
		file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeshTrace_Conf); i {
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
			RawDescriptor: file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_goTypes,
		DependencyIndexes: file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_depIdxs,
		MessageInfos:      file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_msgTypes,
	}.Build()
	File_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto = out.File
	file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_rawDesc = nil
	file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_goTypes = nil
	file_pkg_plugins_policies_meshtrace_api_v1alpha1_meshtrace_proto_depIdxs = nil
}
