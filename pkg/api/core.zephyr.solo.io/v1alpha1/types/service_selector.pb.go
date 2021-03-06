// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/service-mesh-hub/api/core/v1alpha1/service_selector.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
//Select Kubernetes services
//
//Only one of (labels + namespaces + cluster) or (resource refs) may be provided. If all four are provided, it will be
//considered an error, and the Status of the top level resource will be updated to reflect an IllegalSelection.
//
//Valid:
//1.
//selector:
//matcher:
//labels:
//foo: bar
//hello: world
//namespaces:
//- default
//cluster: "cluster-name"
//2.
//selector:
//matcher:
//refs:
//- name: foo
//namespace: bar
//
//Invalid:
//1.
//selector:
//matcher:
//labels:
//foo: bar
//hello: world
//namespaces:
//- default
//cluster: "cluster-name"
//refs:
//- name: foo
//namespace: bar
//
//By default labels will select across all namespaces, unless a list of namespaces is provided, in which case
//it will only select from those. An empty list is equal to AllNamespaces.
//
//If no labels are given, and only namespaces, all resources from the namespaces will be selected.
//
//The following selector will select all resources with the following labels in every namespace, in the local cluster:
//
//selector:
//matcher:
//labels:
//foo: bar
//hello: world
//
//Whereas the next selector will only select from the specified namespaces (foo, bar), in the local cluster:
//
//selector:
//matcher:
//labels:
//foo: bar
//hello: world
//namespaces
//- foo
//- bar
//
//This final selector will select all resources of a given type in the target namespace (foo), in the local cluster:
//
//selector
//matcher:
//namespaces
//- foo
//- bar
//labels:
//hello: world
//
//
type ServiceSelector struct {
	// If specified, select services using either a Matcher or direct reference. If not set, select all Services.
	//
	// Types that are valid to be assigned to ServiceSelectorType:
	//	*ServiceSelector_Matcher_
	//	*ServiceSelector_ServiceRefs_
	ServiceSelectorType  isServiceSelector_ServiceSelectorType `protobuf_oneof:"service_selector_type"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ServiceSelector) Reset()         { *m = ServiceSelector{} }
func (m *ServiceSelector) String() string { return proto.CompactTextString(m) }
func (*ServiceSelector) ProtoMessage()    {}
func (*ServiceSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f06c46809ababe4, []int{0}
}
func (m *ServiceSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSelector.Unmarshal(m, b)
}
func (m *ServiceSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSelector.Marshal(b, m, deterministic)
}
func (m *ServiceSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSelector.Merge(m, src)
}
func (m *ServiceSelector) XXX_Size() int {
	return xxx_messageInfo_ServiceSelector.Size(m)
}
func (m *ServiceSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSelector.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSelector proto.InternalMessageInfo

type isServiceSelector_ServiceSelectorType interface {
	isServiceSelector_ServiceSelectorType()
	Equal(interface{}) bool
}

type ServiceSelector_Matcher_ struct {
	Matcher *ServiceSelector_Matcher `protobuf:"bytes,1,opt,name=matcher,proto3,oneof" json:"matcher,omitempty"`
}
type ServiceSelector_ServiceRefs_ struct {
	ServiceRefs *ServiceSelector_ServiceRefs `protobuf:"bytes,2,opt,name=service_refs,json=serviceRefs,proto3,oneof" json:"service_refs,omitempty"`
}

func (*ServiceSelector_Matcher_) isServiceSelector_ServiceSelectorType()     {}
func (*ServiceSelector_ServiceRefs_) isServiceSelector_ServiceSelectorType() {}

func (m *ServiceSelector) GetServiceSelectorType() isServiceSelector_ServiceSelectorType {
	if m != nil {
		return m.ServiceSelectorType
	}
	return nil
}

func (m *ServiceSelector) GetMatcher() *ServiceSelector_Matcher {
	if x, ok := m.GetServiceSelectorType().(*ServiceSelector_Matcher_); ok {
		return x.Matcher
	}
	return nil
}

func (m *ServiceSelector) GetServiceRefs() *ServiceSelector_ServiceRefs {
	if x, ok := m.GetServiceSelectorType().(*ServiceSelector_ServiceRefs_); ok {
		return x.ServiceRefs
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ServiceSelector) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ServiceSelector_Matcher_)(nil),
		(*ServiceSelector_ServiceRefs_)(nil),
	}
}

type ServiceSelector_Matcher struct {
	// If specified, all labels must exist on k8s Service, else match on any labels.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If specified, match k8s Services if they exist in one of the specified namespaces. If not specified, match on any namespace.
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// If specified, match k8s Services if they exist in one of the specified clusters. If not specified, match on any cluster.
	Clusters             []string `protobuf:"bytes,3,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceSelector_Matcher) Reset()         { *m = ServiceSelector_Matcher{} }
func (m *ServiceSelector_Matcher) String() string { return proto.CompactTextString(m) }
func (*ServiceSelector_Matcher) ProtoMessage()    {}
func (*ServiceSelector_Matcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f06c46809ababe4, []int{0, 0}
}
func (m *ServiceSelector_Matcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSelector_Matcher.Unmarshal(m, b)
}
func (m *ServiceSelector_Matcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSelector_Matcher.Marshal(b, m, deterministic)
}
func (m *ServiceSelector_Matcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSelector_Matcher.Merge(m, src)
}
func (m *ServiceSelector_Matcher) XXX_Size() int {
	return xxx_messageInfo_ServiceSelector_Matcher.Size(m)
}
func (m *ServiceSelector_Matcher) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSelector_Matcher.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSelector_Matcher proto.InternalMessageInfo

func (m *ServiceSelector_Matcher) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ServiceSelector_Matcher) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func (m *ServiceSelector_Matcher) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type ServiceSelector_ServiceRefs struct {
	// Match k8s Services by direct reference.
	Services             []*ResourceRef `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ServiceSelector_ServiceRefs) Reset()         { *m = ServiceSelector_ServiceRefs{} }
func (m *ServiceSelector_ServiceRefs) String() string { return proto.CompactTextString(m) }
func (*ServiceSelector_ServiceRefs) ProtoMessage()    {}
func (*ServiceSelector_ServiceRefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f06c46809ababe4, []int{0, 1}
}
func (m *ServiceSelector_ServiceRefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSelector_ServiceRefs.Unmarshal(m, b)
}
func (m *ServiceSelector_ServiceRefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSelector_ServiceRefs.Marshal(b, m, deterministic)
}
func (m *ServiceSelector_ServiceRefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSelector_ServiceRefs.Merge(m, src)
}
func (m *ServiceSelector_ServiceRefs) XXX_Size() int {
	return xxx_messageInfo_ServiceSelector_ServiceRefs.Size(m)
}
func (m *ServiceSelector_ServiceRefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSelector_ServiceRefs.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSelector_ServiceRefs proto.InternalMessageInfo

func (m *ServiceSelector_ServiceRefs) GetServices() []*ResourceRef {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*ServiceSelector)(nil), "core.zephyr.solo.io.ServiceSelector")
	proto.RegisterType((*ServiceSelector_Matcher)(nil), "core.zephyr.solo.io.ServiceSelector.Matcher")
	proto.RegisterMapType((map[string]string)(nil), "core.zephyr.solo.io.ServiceSelector.Matcher.LabelsEntry")
	proto.RegisterType((*ServiceSelector_ServiceRefs)(nil), "core.zephyr.solo.io.ServiceSelector.ServiceRefs")
}

func init() {
	proto.RegisterFile("github.com/solo-io/service-mesh-hub/api/core/v1alpha1/service_selector.proto", fileDescriptor_4f06c46809ababe4)
}

var fileDescriptor_4f06c46809ababe4 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x7d, 0x79, 0xc1, 0xf7, 0x5e, 0x27, 0x82, 0x32, 0x56, 0x0c, 0x59, 0x94, 0xe2, 0xaa, 0x8b,
	0x76, 0xc6, 0xd6, 0x4d, 0x15, 0x57, 0x05, 0xa1, 0x60, 0x15, 0x99, 0xea, 0xc6, 0x4d, 0x99, 0x84,
	0x9b, 0x0f, 0x3a, 0xe9, 0x0c, 0x33, 0x49, 0x25, 0xfe, 0x22, 0x7f, 0x8f, 0x3f, 0xc1, 0xbd, 0xff,
	0x41, 0x32, 0x49, 0xda, 0x52, 0xba, 0xe8, 0xdb, 0xdd, 0x3b, 0x73, 0xcf, 0xb9, 0xe7, 0x5c, 0x0e,
	0x5a, 0x25, 0x59, 0x91, 0x96, 0x21, 0x89, 0x64, 0x4e, 0x8d, 0x14, 0x72, 0x92, 0x49, 0x6a, 0x40,
	0xef, 0xb3, 0x08, 0x26, 0x39, 0x98, 0x74, 0x92, 0x96, 0x21, 0xe5, 0x2a, 0xa3, 0x91, 0xd4, 0x40,
	0xf7, 0x53, 0x2e, 0x54, 0xca, 0xa7, 0xdd, 0xc8, 0xc6, 0x80, 0x80, 0xa8, 0x90, 0x9a, 0x28, 0x2d,
	0x0b, 0x89, 0x5f, 0xd4, 0x53, 0xe4, 0x17, 0xa8, 0xb4, 0xd2, 0xa4, 0xa6, 0x23, 0x99, 0x0c, 0x06,
	0x89, 0x94, 0x89, 0x00, 0x6a, 0x47, 0xc2, 0x32, 0xa6, 0x3f, 0x35, 0x57, 0x0a, 0xb4, 0x69, 0x40,
	0xc1, 0xf8, 0x8a, 0x7d, 0x1a, 0xe2, 0x76, 0xba, 0x9f, 0xc8, 0x44, 0xda, 0x92, 0xd6, 0x55, 0xf3,
	0xfa, 0xfa, 0x9f, 0x8b, 0x9e, 0xad, 0x1b, 0x9a, 0x75, 0x2b, 0x09, 0x2f, 0xd1, 0x7d, 0xce, 0x8b,
	0x28, 0x05, 0xed, 0x3b, 0x43, 0x67, 0xe4, 0xcd, 0xc6, 0xe4, 0x82, 0x3c, 0x72, 0x06, 0x23, 0x9f,
	0x1b, 0xcc, 0xf2, 0x86, 0x75, 0x70, 0xfc, 0x1d, 0x3d, 0xed, 0x0c, 0x6b, 0x88, 0x8d, 0x7f, 0x6b,
	0xe9, 0xde, 0x5c, 0x45, 0xd7, 0xf6, 0x0c, 0x62, 0xb3, 0xbc, 0x61, 0x9e, 0x39, 0xb6, 0xc1, 0x1f,
	0x07, 0xdd, 0xb7, 0xdb, 0xf0, 0x57, 0x74, 0x27, 0x78, 0x08, 0xc2, 0xf8, 0xce, 0xd0, 0x1d, 0x79,
	0xb3, 0xf9, 0x63, 0xb4, 0x92, 0x95, 0x85, 0x7e, 0xdc, 0x15, 0xba, 0x62, 0x2d, 0x0f, 0x1e, 0x20,
	0xb4, 0xe3, 0x39, 0x18, 0xc5, 0x23, 0xa8, 0x25, 0xbb, 0xa3, 0x1e, 0x3b, 0x79, 0xc1, 0x01, 0x7a,
	0x88, 0x44, 0x69, 0x0a, 0xd0, 0xc6, 0x77, 0xed, 0xef, 0xa1, 0x0f, 0xde, 0x21, 0xef, 0x84, 0x12,
	0x3f, 0x47, 0xee, 0x16, 0x2a, 0x7b, 0xc5, 0x1e, 0xab, 0x4b, 0xdc, 0x47, 0x4f, 0xf6, 0x5c, 0x94,
	0x60, 0x4f, 0xd1, 0x63, 0x4d, 0xf3, 0xfe, 0x76, 0xee, 0x04, 0x9f, 0x90, 0x77, 0x62, 0x19, 0x7f,
	0x40, 0x0f, 0xad, 0xe5, 0xce, 0xd9, 0xf0, 0xa2, 0x33, 0x06, 0x46, 0x96, 0xda, 0x82, 0xd8, 0x01,
	0xb1, 0x78, 0x85, 0x5e, 0x9e, 0x27, 0x6d, 0x53, 0x54, 0x0a, 0x16, 0xdf, 0x7e, 0xff, 0x1d, 0x38,
	0x3f, 0xbe, 0x5c, 0x13, 0x5e, 0xb5, 0x4d, 0x0e, 0x81, 0x3a, 0xdb, 0x7a, 0xcc, 0x57, 0x4d, 0x6a,
	0xc2, 0x3b, 0x1b, 0xa6, 0xb7, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x59, 0x0e, 0xc6, 0x95, 0x15,
	0x03, 0x00, 0x00,
}

func (this *ServiceSelector) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSelector)
	if !ok {
		that2, ok := that.(ServiceSelector)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.ServiceSelectorType == nil {
		if this.ServiceSelectorType != nil {
			return false
		}
	} else if this.ServiceSelectorType == nil {
		return false
	} else if !this.ServiceSelectorType.Equal(that1.ServiceSelectorType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSelector_Matcher_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSelector_Matcher_)
	if !ok {
		that2, ok := that.(ServiceSelector_Matcher_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Matcher.Equal(that1.Matcher) {
		return false
	}
	return true
}
func (this *ServiceSelector_ServiceRefs_) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSelector_ServiceRefs_)
	if !ok {
		that2, ok := that.(ServiceSelector_ServiceRefs_)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ServiceRefs.Equal(that1.ServiceRefs) {
		return false
	}
	return true
}
func (this *ServiceSelector_Matcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSelector_Matcher)
	if !ok {
		that2, ok := that.(ServiceSelector_Matcher)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if this.Clusters[i] != that1.Clusters[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ServiceSelector_ServiceRefs) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceSelector_ServiceRefs)
	if !ok {
		that2, ok := that.(ServiceSelector_ServiceRefs)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Services) != len(that1.Services) {
		return false
	}
	for i := range this.Services {
		if !this.Services[i].Equal(that1.Services[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
