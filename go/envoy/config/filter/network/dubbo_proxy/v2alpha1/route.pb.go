// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto

package envoy_config_filter_network_dubbo_proxy_v2alpha1

import (
	fmt "fmt"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	_type "github.com/cilium/proxy/go/envoy/type"
	matcher "github.com/cilium/proxy/go/envoy/type/matcher"
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

// [#next-free-field: 6]
type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The interface name of the service.
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	// Which group does the interface belong to.
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	// The version number of the interface.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes               []*Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RouteConfiguration) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RouteConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Method level routing matching.
	Method *MethodMatch `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config).
	Headers              []*route.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

func (m *RouteMatch) GetMethod() *MethodMatch {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *route.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *route.WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

type MethodMatch struct {
	// The name of the method.
	Name *matcher.StringMatcher `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Method parameter definition.
	// The key is the parameter index, starting from 0.
	// The value is the parameter matching type.
	ParamsMatch          map[uint32]*MethodMatch_ParameterMatchSpecifier `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch,proto3" json:"params_match,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *MethodMatch) Reset()         { *m = MethodMatch{} }
func (m *MethodMatch) String() string { return proto.CompactTextString(m) }
func (*MethodMatch) ProtoMessage()    {}
func (*MethodMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{4}
}

func (m *MethodMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch.Unmarshal(m, b)
}
func (m *MethodMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch.Marshal(b, m, deterministic)
}
func (m *MethodMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch.Merge(m, src)
}
func (m *MethodMatch) XXX_Size() int {
	return xxx_messageInfo_MethodMatch.Size(m)
}
func (m *MethodMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch proto.InternalMessageInfo

func (m *MethodMatch) GetName() *matcher.StringMatcher {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodMatch) GetParamsMatch() map[uint32]*MethodMatch_ParameterMatchSpecifier {
	if m != nil {
		return m.ParamsMatch
	}
	return nil
}

// The parameter matching type.
type MethodMatch_ParameterMatchSpecifier struct {
	// Types that are valid to be assigned to ParameterMatchSpecifier:
	//	*MethodMatch_ParameterMatchSpecifier_ExactMatch
	//	*MethodMatch_ParameterMatchSpecifier_RangeMatch
	ParameterMatchSpecifier isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier `protobuf_oneof:"parameter_match_specifier"`
	XXX_NoUnkeyedLiteral    struct{}                                                      `json:"-"`
	XXX_unrecognized        []byte                                                        `json:"-"`
	XXX_sizecache           int32                                                         `json:"-"`
}

func (m *MethodMatch_ParameterMatchSpecifier) Reset()         { *m = MethodMatch_ParameterMatchSpecifier{} }
func (m *MethodMatch_ParameterMatchSpecifier) String() string { return proto.CompactTextString(m) }
func (*MethodMatch_ParameterMatchSpecifier) ProtoMessage()    {}
func (*MethodMatch_ParameterMatchSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{4, 0}
}

func (m *MethodMatch_ParameterMatchSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Unmarshal(m, b)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Marshal(b, m, deterministic)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Merge(m, src)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Size() int {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Size(m)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch_ParameterMatchSpecifier proto.InternalMessageInfo

type isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier interface {
	isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier()
}

type MethodMatch_ParameterMatchSpecifier_ExactMatch struct {
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type MethodMatch_ParameterMatchSpecifier_RangeMatch struct {
	RangeMatch *_type.Int64Range `protobuf:"bytes,4,opt,name=range_match,json=rangeMatch,proto3,oneof"`
}

func (*MethodMatch_ParameterMatchSpecifier_ExactMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (*MethodMatch_ParameterMatchSpecifier_RangeMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (m *MethodMatch_ParameterMatchSpecifier) GetParameterMatchSpecifier() isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier {
	if m != nil {
		return m.ParameterMatchSpecifier
	}
	return nil
}

func (m *MethodMatch_ParameterMatchSpecifier) GetExactMatch() string {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (m *MethodMatch_ParameterMatchSpecifier) GetRangeMatch() *_type.Int64Range {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		return x.RangeMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MethodMatch_ParameterMatchSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MethodMatch_ParameterMatchSpecifier_ExactMatch)(nil),
		(*MethodMatch_ParameterMatchSpecifier_RangeMatch)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteAction")
	proto.RegisterType((*MethodMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch")
	proto.RegisterMapType((map[uint32]*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch.ParamsMatchEntry")
	proto.RegisterType((*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch.ParameterMatchSpecifier")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto", fileDescriptor_74a572433a3292e0)
}

var fileDescriptor_74a572433a3292e0 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0x93, 0xba, 0xfd, 0x7a, 0xfd, 0x21, 0xa5, 0x23, 0xd4, 0x9a, 0x80, 0xa0, 0x0d, 0x9b,
	0xae, 0xc6, 0x10, 0xfe, 0xa1, 0x20, 0x91, 0x0a, 0xa9, 0x2c, 0x5a, 0xaa, 0xa9, 0x2a, 0x36, 0xa0,
	0x68, 0xea, 0x4c, 0x92, 0x51, 0x13, 0x8f, 0x19, 0x4f, 0xdc, 0xe6, 0x15, 0x58, 0x83, 0xd8, 0xf2,
	0x30, 0x48, 0xbc, 0x09, 0xcf, 0x80, 0x58, 0x21, 0xdf, 0x19, 0x13, 0xf3, 0xb3, 0x69, 0x60, 0x37,
	0xbe, 0x73, 0xcf, 0x39, 0x73, 0x8e, 0xee, 0x35, 0x6c, 0x8b, 0x24, 0x57, 0xd3, 0x28, 0x56, 0x49,
	0x5f, 0x0e, 0xa2, 0xbe, 0x1c, 0x19, 0xa1, 0xa3, 0x44, 0x98, 0x53, 0xa5, 0x4f, 0xa2, 0xde, 0xe4,
	0xf8, 0x58, 0x75, 0x53, 0xad, 0xce, 0xa6, 0x51, 0xde, 0xe6, 0xa3, 0x74, 0xc8, 0x6f, 0x46, 0x5a,
	0x4d, 0x8c, 0xa0, 0xa9, 0x56, 0x46, 0x91, 0x1b, 0x88, 0xa6, 0x16, 0x4d, 0x2d, 0x9a, 0x3a, 0x34,
	0xad, 0xa0, 0x69, 0x89, 0x6e, 0x5e, 0xb5, 0x7a, 0x3c, 0x95, 0x51, 0xde, 0xb6, 0x5c, 0x55, 0xc6,
	0xe6, 0x35, 0x7b, 0x6f, 0xa6, 0xa9, 0x88, 0xc6, 0xdc, 0xc4, 0x43, 0xa1, 0xa3, 0xcc, 0x68, 0x99,
	0x0c, 0x5c, 0xc3, 0x5a, 0xa5, 0x41, 0xf3, 0x64, 0x50, 0x02, 0xd7, 0x73, 0x3e, 0x92, 0x3d, 0x6e,
	0x44, 0x54, 0x1e, 0xec, 0x45, 0xeb, 0xb3, 0x07, 0x84, 0x15, 0x0a, 0x3b, 0xf8, 0xca, 0x89, 0xe6,
	0x46, 0xaa, 0x84, 0x10, 0x58, 0x4c, 0xf8, 0x58, 0x84, 0xde, 0x86, 0xb7, 0xb5, 0xc2, 0xf0, 0x4c,
	0xae, 0xc0, 0x8a, 0x4c, 0x8c, 0xd0, 0x7d, 0x1e, 0x8b, 0xb0, 0x86, 0x17, 0xb3, 0x02, 0xb9, 0x08,
	0xfe, 0x40, 0xab, 0x49, 0x1a, 0xd6, 0xf1, 0xc6, 0x7e, 0x90, 0x10, 0x96, 0x73, 0xa1, 0x33, 0xa9,
	0x92, 0x70, 0x11, 0xeb, 0xe5, 0x27, 0x79, 0x01, 0x4b, 0xe8, 0x2c, 0x0b, 0xfd, 0x8d, 0xfa, 0x56,
	0xd0, 0xbe, 0x47, 0xcf, 0x9b, 0x16, 0xc5, 0x77, 0x33, 0x47, 0xd3, 0xfa, 0xe4, 0x81, 0x8f, 0x15,
	0xf2, 0x0a, 0x7c, 0x0c, 0x07, 0x5f, 0x1f, 0xb4, 0xb7, 0xe7, 0x64, 0xde, 0x2b, 0x38, 0x3a, 0xff,
	0x7d, 0xeb, 0xf8, 0x6f, 0xbd, 0x5a, 0xc3, 0x63, 0x96, 0x94, 0xbc, 0x06, 0x1f, 0x15, 0x31, 0x82,
	0xa0, 0xfd, 0x78, 0x4e, 0xf6, 0xa7, 0x71, 0x11, 0x74, 0x95, 0x1e, 0x59, 0x5b, 0x1f, 0x3d, 0x80,
	0x99, 0x3c, 0x39, 0x82, 0xa5, 0xb1, 0x30, 0x43, 0xd5, 0x73, 0x66, 0xe6, 0x90, 0xdb, 0x43, 0x3c,
	0xd2, 0x31, 0x47, 0x46, 0x1e, 0xc1, 0xf2, 0x50, 0xf0, 0x9e, 0xd0, 0x59, 0x58, 0xc3, 0xf8, 0x37,
	0x1d, 0x2f, 0x4f, 0x25, 0xcd, 0xdb, 0xd4, 0x0e, 0xdd, 0x2e, 0xb6, 0xec, 0xd9, 0x41, 0x63, 0x25,
	0xa2, 0xf5, 0xc1, 0x83, 0xa0, 0xe2, 0x81, 0x34, 0x61, 0x39, 0x1e, 0x4d, 0x32, 0x23, 0xb4, 0x9d,
	0x97, 0xdd, 0x05, 0x56, 0x16, 0x08, 0x83, 0xd5, 0x53, 0x21, 0x07, 0x43, 0x23, 0x7a, 0x5d, 0x57,
	0xcb, 0x5c, 0x72, 0xd7, 0xff, 0x24, 0xf9, 0xd2, 0x35, 0xef, 0xd8, 0xde, 0xdd, 0x05, 0xd6, 0x38,
	0xfd, 0xb9, 0x94, 0x75, 0x42, 0x58, 0x75, 0x54, 0xdd, 0x2c, 0x15, 0xb1, 0xec, 0x4b, 0xa1, 0x49,
	0xfd, 0x6b, 0xc7, 0x6b, 0x7d, 0xa9, 0x43, 0x50, 0xb1, 0x4b, 0xee, 0x54, 0xc6, 0x78, 0xe6, 0xb1,
	0xd8, 0x0e, 0xea, 0xd6, 0x87, 0x1e, 0xe2, 0xfa, 0x94, 0x1e, 0xed, 0xa4, 0xbf, 0x81, 0xff, 0x53,
	0xae, 0xf9, 0x38, 0xeb, 0xda, 0x39, 0xb2, 0x11, 0xed, 0xff, 0x55, 0xf4, 0xf4, 0x00, 0x19, 0xf1,
	0xfc, 0x2c, 0x31, 0x7a, 0xca, 0x82, 0x74, 0x56, 0x69, 0xbe, 0xf3, 0x60, 0x1d, 0x3b, 0x84, 0x71,
	0x89, 0x1f, 0xfe, 0xb0, 0xb6, 0x09, 0x81, 0x38, 0xe3, 0xb1, 0x71, 0xaf, 0xa9, 0xbb, 0x8c, 0x01,
	0x8b, 0xd6, 0xe8, 0x03, 0x08, 0x70, 0xdd, 0x5d, 0xcb, 0x22, 0xfa, 0x5d, 0xab, 0xfa, 0x7d, 0x9e,
	0x98, 0xbb, 0xb7, 0x59, 0xd1, 0x53, 0x40, 0xb1, 0xd9, 0x0e, 0xf8, 0x65, 0xb8, 0x94, 0x96, 0xc2,
	0x16, 0x3e, 0x4b, 0xb5, 0xf9, 0xde, 0x83, 0xc6, 0xaf, 0x0f, 0x27, 0x0d, 0xa8, 0x9f, 0x88, 0x29,
	0x86, 0x7a, 0x81, 0x15, 0x47, 0x72, 0x02, 0x7e, 0xce, 0x47, 0x93, 0x72, 0x27, 0x8e, 0xfe, 0x41,
	0x52, 0xbf, 0xe7, 0xc0, 0xac, 0xc6, 0xc3, 0xda, 0x7d, 0xaf, 0xb3, 0x0f, 0x4f, 0xa4, 0xb2, 0x2a,
	0x96, 0xe8, 0xbc, 0x82, 0x1d, 0xbb, 0x64, 0x07, 0xc5, 0x4f, 0xf0, 0xc0, 0x3b, 0x5e, 0xc2, 0xbf,
	0xe1, 0xad, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x4c, 0xcf, 0x0d, 0xf1, 0x05, 0x00, 0x00,
}
