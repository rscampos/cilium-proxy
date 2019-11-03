// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/transcoder/v2/transcoder.proto

package envoy_config_filter_http_transcoder_v2

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// [#next-free-field: 10]
type GrpcJsonTranscoder struct {
	// Types that are valid to be assigned to DescriptorSet:
	//	*GrpcJsonTranscoder_ProtoDescriptor
	//	*GrpcJsonTranscoder_ProtoDescriptorBin
	DescriptorSet isGrpcJsonTranscoder_DescriptorSet `protobuf_oneof:"descriptor_set"`
	// A list of strings that
	// supplies the fully qualified service names (i.e. "package_name.service_name") that
	// the transcoder will translate. If the service name doesn't exist in ``proto_descriptor``,
	// Envoy will fail at startup. The ``proto_descriptor`` may contain more services than
	// the service names specified here, but they won't be translated.
	Services []string `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty"`
	// Control options for response JSON. These options are passed directly to
	// `JsonPrintOptions <https://developers.google.com/protocol-buffers/docs/reference/cpp/
	// google.protobuf.util.json_util#JsonPrintOptions>`_.
	PrintOptions *GrpcJsonTranscoder_PrintOptions `protobuf:"bytes,3,opt,name=print_options,json=printOptions,proto3" json:"print_options,omitempty"`
	// Whether to keep the incoming request route after the outgoing headers have been transformed to
	// the match the upstream gRPC service. Note: This means that routes for gRPC services that are
	// not transcoded cannot be used in combination with *match_incoming_request_route*.
	MatchIncomingRequestRoute bool `protobuf:"varint,5,opt,name=match_incoming_request_route,json=matchIncomingRequestRoute,proto3" json:"match_incoming_request_route,omitempty"`
	// A list of query parameters to be ignored for transcoding method mapping.
	// By default, the transcoder filter will not transcode a request if there are any
	// unknown/invalid query parameters.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {
	//         option (google.api.http) = {
	//           get: "/shelves/{shelf}"
	//         };
	//       }
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The request ``/shelves/100?foo=bar`` will not be mapped to ``GetShelf``` because variable
	// binding for ``foo`` is not defined. Adding ``foo`` to ``ignored_query_parameters`` will allow
	// the same request to be mapped to ``GetShelf``.
	IgnoredQueryParameters []string `protobuf:"bytes,6,rep,name=ignored_query_parameters,json=ignoredQueryParameters,proto3" json:"ignored_query_parameters,omitempty"`
	// Whether to route methods without the ``google.api.http`` option.
	//
	// Example :
	//
	// .. code-block:: proto
	//
	//     package bookstore;
	//
	//     service Bookstore {
	//       rpc GetShelf(GetShelfRequest) returns (Shelf) {}
	//     }
	//
	//     message GetShelfRequest {
	//       int64 shelf = 1;
	//     }
	//
	//     message Shelf {}
	//
	// The client could ``post`` a json body ``{"shelf": 1234}`` with the path of
	// ``/bookstore.Bookstore/GetShelfRequest`` to call ``GetShelfRequest``.
	AutoMapping bool `protobuf:"varint,7,opt,name=auto_mapping,json=autoMapping,proto3" json:"auto_mapping,omitempty"`
	// Whether to ignore query parameters that cannot be mapped to a corresponding
	// protobuf field. Use this if you cannot control the query parameters and do
	// not know them beforehand. Otherwise use ``ignored_query_parameters``.
	// Defaults to false.
	IgnoreUnknownQueryParameters bool `protobuf:"varint,8,opt,name=ignore_unknown_query_parameters,json=ignoreUnknownQueryParameters,proto3" json:"ignore_unknown_query_parameters,omitempty"`
	// Whether to convert gRPC status headers to JSON.
	// When trailer indicates a gRPC error and there was no HTTP body, take ``google.rpc.Status``
	// from the ``grpc-status-details-bin`` header and use it as JSON body.
	// If there was no such header, make ``google.rpc.Status`` out of the ``grpc-status`` and
	// ``grpc-message`` headers.
	// The error details types must be present in the ``proto_descriptor``.
	//
	// For example, if an upstream server replies with headers:
	//
	// .. code-block:: none
	//
	//     grpc-status: 5
	//     grpc-status-details-bin:
	//         CAUaMwoqdHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUucnBjLlJlcXVlc3RJbmZvEgUKA3ItMQ
	//
	// The ``grpc-status-details-bin`` header contains a base64-encoded protobuf message
	// ``google.rpc.Status``. It will be transcoded into:
	//
	// .. code-block:: none
	//
	//     HTTP/1.1 404 Not Found
	//     content-type: application/json
	//
	//     {"code":5,"details":[{"@type":"type.googleapis.com/google.rpc.RequestInfo","requestId":"r-1"}]}
	//
	//  In order to transcode the message, the ``google.rpc.RequestInfo`` type from
	//  the ``google/rpc/error_details.proto`` should be included in the configured
	//  :ref:`proto descriptor set <config_grpc_json_generate_proto_descriptor_set>`.
	ConvertGrpcStatus    bool     `protobuf:"varint,9,opt,name=convert_grpc_status,json=convertGrpcStatus,proto3" json:"convert_grpc_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcJsonTranscoder) Reset()         { *m = GrpcJsonTranscoder{} }
func (m *GrpcJsonTranscoder) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder) ProtoMessage()    {}
func (*GrpcJsonTranscoder) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0}
}

func (m *GrpcJsonTranscoder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder.Merge(m, src)
}
func (m *GrpcJsonTranscoder) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder.Size(m)
}
func (m *GrpcJsonTranscoder) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder proto.InternalMessageInfo

type isGrpcJsonTranscoder_DescriptorSet interface {
	isGrpcJsonTranscoder_DescriptorSet()
}

type GrpcJsonTranscoder_ProtoDescriptor struct {
	ProtoDescriptor string `protobuf:"bytes,1,opt,name=proto_descriptor,json=protoDescriptor,proto3,oneof"`
}

type GrpcJsonTranscoder_ProtoDescriptorBin struct {
	ProtoDescriptorBin []byte `protobuf:"bytes,4,opt,name=proto_descriptor_bin,json=protoDescriptorBin,proto3,oneof"`
}

func (*GrpcJsonTranscoder_ProtoDescriptor) isGrpcJsonTranscoder_DescriptorSet() {}

func (*GrpcJsonTranscoder_ProtoDescriptorBin) isGrpcJsonTranscoder_DescriptorSet() {}

func (m *GrpcJsonTranscoder) GetDescriptorSet() isGrpcJsonTranscoder_DescriptorSet {
	if m != nil {
		return m.DescriptorSet
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetProtoDescriptor() string {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptor); ok {
		return x.ProtoDescriptor
	}
	return ""
}

func (m *GrpcJsonTranscoder) GetProtoDescriptorBin() []byte {
	if x, ok := m.GetDescriptorSet().(*GrpcJsonTranscoder_ProtoDescriptorBin); ok {
		return x.ProtoDescriptorBin
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetPrintOptions() *GrpcJsonTranscoder_PrintOptions {
	if m != nil {
		return m.PrintOptions
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetMatchIncomingRequestRoute() bool {
	if m != nil {
		return m.MatchIncomingRequestRoute
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoredQueryParameters() []string {
	if m != nil {
		return m.IgnoredQueryParameters
	}
	return nil
}

func (m *GrpcJsonTranscoder) GetAutoMapping() bool {
	if m != nil {
		return m.AutoMapping
	}
	return false
}

func (m *GrpcJsonTranscoder) GetIgnoreUnknownQueryParameters() bool {
	if m != nil {
		return m.IgnoreUnknownQueryParameters
	}
	return false
}

func (m *GrpcJsonTranscoder) GetConvertGrpcStatus() bool {
	if m != nil {
		return m.ConvertGrpcStatus
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GrpcJsonTranscoder) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GrpcJsonTranscoder_ProtoDescriptor)(nil),
		(*GrpcJsonTranscoder_ProtoDescriptorBin)(nil),
	}
}

type GrpcJsonTranscoder_PrintOptions struct {
	// Whether to add spaces, line breaks and indentation to make the JSON
	// output easy to read. Defaults to false.
	AddWhitespace bool `protobuf:"varint,1,opt,name=add_whitespace,json=addWhitespace,proto3" json:"add_whitespace,omitempty"`
	// Whether to always print primitive fields. By default primitive
	// fields with default values will be omitted in JSON output. For
	// example, an int32 field set to 0 will be omitted. Setting this flag to
	// true will override the default behavior and print primitive fields
	// regardless of their values. Defaults to false.
	AlwaysPrintPrimitiveFields bool `protobuf:"varint,2,opt,name=always_print_primitive_fields,json=alwaysPrintPrimitiveFields,proto3" json:"always_print_primitive_fields,omitempty"`
	// Whether to always print enums as ints. By default they are rendered
	// as strings. Defaults to false.
	AlwaysPrintEnumsAsInts bool `protobuf:"varint,3,opt,name=always_print_enums_as_ints,json=alwaysPrintEnumsAsInts,proto3" json:"always_print_enums_as_ints,omitempty"`
	// Whether to preserve proto field names. By default protobuf will
	// generate JSON field names using the ``json_name`` option, or lower camel case,
	// in that order. Setting this flag will preserve the original field names. Defaults to false.
	PreserveProtoFieldNames bool     `protobuf:"varint,4,opt,name=preserve_proto_field_names,json=preserveProtoFieldNames,proto3" json:"preserve_proto_field_names,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *GrpcJsonTranscoder_PrintOptions) Reset()         { *m = GrpcJsonTranscoder_PrintOptions{} }
func (m *GrpcJsonTranscoder_PrintOptions) String() string { return proto.CompactTextString(m) }
func (*GrpcJsonTranscoder_PrintOptions) ProtoMessage()    {}
func (*GrpcJsonTranscoder_PrintOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_540f2f6de8b0585c, []int{0, 0}
}

func (m *GrpcJsonTranscoder_PrintOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Unmarshal(m, b)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Marshal(b, m, deterministic)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Merge(m, src)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_Size() int {
	return xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.Size(m)
}
func (m *GrpcJsonTranscoder_PrintOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcJsonTranscoder_PrintOptions.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcJsonTranscoder_PrintOptions proto.InternalMessageInfo

func (m *GrpcJsonTranscoder_PrintOptions) GetAddWhitespace() bool {
	if m != nil {
		return m.AddWhitespace
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintPrimitiveFields() bool {
	if m != nil {
		return m.AlwaysPrintPrimitiveFields
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetAlwaysPrintEnumsAsInts() bool {
	if m != nil {
		return m.AlwaysPrintEnumsAsInts
	}
	return false
}

func (m *GrpcJsonTranscoder_PrintOptions) GetPreserveProtoFieldNames() bool {
	if m != nil {
		return m.PreserveProtoFieldNames
	}
	return false
}

func init() {
	proto.RegisterType((*GrpcJsonTranscoder)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder")
	proto.RegisterType((*GrpcJsonTranscoder_PrintOptions)(nil), "envoy.config.filter.http.transcoder.v2.GrpcJsonTranscoder.PrintOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/transcoder/v2/transcoder.proto", fileDescriptor_540f2f6de8b0585c)
}

var fileDescriptor_540f2f6de8b0585c = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x1c, 0xc5, 0x97, 0x8d, 0x8d, 0xcc, 0xeb, 0x36, 0x30, 0xb0, 0x85, 0x6a, 0x88, 0x82, 0x00, 0x55,
	0x42, 0x4a, 0xa4, 0x82, 0x04, 0x82, 0x03, 0x5a, 0xc4, 0xd8, 0x86, 0x04, 0x94, 0x00, 0xe2, 0x68,
	0x79, 0xc9, 0xbf, 0xad, 0x45, 0x63, 0x7b, 0xb6, 0x93, 0xd2, 0xaf, 0xc1, 0xb7, 0xe0, 0xe3, 0x71,
	0x43, 0x9c, 0x90, 0xed, 0xb6, 0x84, 0xee, 0xb2, 0x5b, 0xec, 0xf7, 0x7b, 0xff, 0x17, 0xdb, 0x0f,
	0x3d, 0x03, 0x5e, 0x8b, 0x69, 0x92, 0x0b, 0x3e, 0x60, 0xc3, 0x64, 0xc0, 0xc6, 0x06, 0x54, 0x32,
	0x32, 0x46, 0x26, 0x46, 0x51, 0xae, 0x73, 0x51, 0x80, 0x4a, 0xea, 0x5e, 0x63, 0x15, 0x4b, 0x25,
	0x8c, 0xc0, 0x8f, 0x9c, 0x31, 0xf6, 0xc6, 0xd8, 0x1b, 0x63, 0x6b, 0x8c, 0x1b, 0x68, 0xdd, 0x6b,
	0xef, 0xd7, 0x74, 0xcc, 0x0a, 0x6a, 0x20, 0x99, 0x7f, 0xf8, 0x01, 0xf7, 0x7f, 0x6e, 0x20, 0x7c,
	0xac, 0x64, 0xfe, 0x56, 0x0b, 0xfe, 0x79, 0x61, 0xc1, 0x8f, 0xd1, 0x35, 0xa7, 0x93, 0x02, 0x74,
	0xae, 0x98, 0x34, 0x42, 0x45, 0x41, 0x27, 0xe8, 0x6e, 0x9e, 0xac, 0x64, 0xbb, 0x4e, 0x79, 0xbd,
	0x10, 0x70, 0x0f, 0xdd, 0x5c, 0x86, 0xc9, 0x19, 0xe3, 0xd1, 0x95, 0x4e, 0xd0, 0x6d, 0x9d, 0xac,
	0x64, 0x78, 0xc9, 0x90, 0x32, 0x8e, 0x1f, 0xa0, 0x50, 0x83, 0xaa, 0x59, 0x0e, 0x3a, 0x5a, 0xed,
	0xac, 0x75, 0x37, 0xd3, 0xf0, 0x4f, 0xba, 0xfe, 0x23, 0x58, 0x0d, 0x83, 0x6c, 0xa1, 0xe0, 0x31,
	0xda, 0x96, 0x8a, 0x71, 0x43, 0x84, 0x34, 0x4c, 0x70, 0x1d, 0xad, 0x75, 0x82, 0xee, 0x56, 0xef,
	0x38, 0xbe, 0xdc, 0xb1, 0xe3, 0x8b, 0x27, 0x8b, 0xfb, 0x76, 0xde, 0x07, 0x3f, 0x2e, 0x6b, 0xc9,
	0xc6, 0x0a, 0xbf, 0x42, 0x07, 0x25, 0x35, 0xf9, 0x88, 0x30, 0x9e, 0x8b, 0x92, 0xf1, 0x21, 0x51,
	0x70, 0x5e, 0x81, 0x36, 0x44, 0x89, 0xca, 0x40, 0xb4, 0xde, 0x09, 0xba, 0x61, 0x76, 0xdb, 0x31,
	0xa7, 0x33, 0x24, 0xf3, 0x44, 0x66, 0x01, 0xfc, 0x1c, 0x45, 0x6c, 0xc8, 0x85, 0x82, 0x82, 0x9c,
	0x57, 0xa0, 0xa6, 0x44, 0x52, 0x45, 0x4b, 0x30, 0xa0, 0x74, 0xb4, 0x61, 0x0f, 0x99, 0xed, 0xcd,
	0xf4, 0x8f, 0x56, 0xee, 0x2f, 0x54, 0x7c, 0x0f, 0xb5, 0x68, 0x65, 0x04, 0x29, 0xa9, 0x94, 0x8c,
	0x0f, 0xa3, 0xab, 0x2e, 0x6a, 0xcb, 0xee, 0xbd, 0xf3, 0x5b, 0xf8, 0x08, 0xdd, 0xf5, 0x66, 0x52,
	0xf1, 0x6f, 0x5c, 0x4c, 0xf8, 0xc5, 0x8c, 0xd0, 0xb9, 0x0e, 0x3c, 0xf6, 0xc5, 0x53, 0xcb, 0x49,
	0x31, 0xba, 0x91, 0x0b, 0x5e, 0x83, 0x32, 0x64, 0xa8, 0x64, 0x4e, 0xb4, 0xa1, 0xa6, 0xd2, 0xd1,
	0xa6, 0xb3, 0x5e, 0x9f, 0x49, 0xf6, 0xde, 0x3e, 0x39, 0xa1, 0xfd, 0x2b, 0x40, 0xad, 0xe6, 0x9d,
	0xe1, 0x87, 0x68, 0x87, 0x16, 0x05, 0x99, 0x8c, 0x98, 0x01, 0x2d, 0x69, 0x0e, 0xae, 0x18, 0x61,
	0xb6, 0x4d, 0x8b, 0xe2, 0xeb, 0x62, 0x13, 0x1f, 0xa2, 0x3b, 0x74, 0x3c, 0xa1, 0x53, 0x4d, 0xfc,
	0x0b, 0x4a, 0xc5, 0x4a, 0x66, 0x58, 0x0d, 0x64, 0xc0, 0x60, 0x5c, 0xd8, 0x57, 0xb7, 0xae, 0xb6,
	0x87, 0x5c, 0x42, 0x7f, 0x8e, 0xbc, 0x71, 0x04, 0x7e, 0x81, 0xda, 0xff, 0x8d, 0x00, 0x5e, 0x95,
	0x9a, 0x50, 0x4d, 0x18, 0x37, 0xbe, 0x0a, 0x61, 0xb6, 0xd7, 0xf0, 0x1f, 0x59, 0xfd, 0x50, 0x9f,
	0x72, 0xa3, 0xf1, 0x4b, 0xd4, 0x96, 0x0a, 0x6c, 0x91, 0x80, 0xf8, 0x72, 0xba, 0x58, 0xc2, 0x69,
	0x09, 0xda, 0x35, 0x33, 0xcc, 0xf6, 0xe7, 0x44, 0xdf, 0x02, 0x2e, 0xf4, 0xbd, 0x95, 0xd3, 0x5b,
	0x68, 0xa7, 0x51, 0x65, 0x0d, 0x06, 0xaf, 0xfd, 0x4e, 0x83, 0xf4, 0x14, 0x3d, 0x65, 0xc2, 0x57,
	0x4f, 0x2a, 0xf1, 0x7d, 0x7a, 0xc9, 0x16, 0xa6, 0xbb, 0xff, 0xea, 0xe7, 0x92, 0xfa, 0xc1, 0xd9,
	0x86, 0xfb, 0xa7, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x14, 0x4c, 0x05, 0x21, 0xf9, 0x03,
	0x00, 0x00,
}
