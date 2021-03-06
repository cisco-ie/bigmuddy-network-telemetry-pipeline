// Code generated by protoc-gen-go.
// source: ipv4_acl_edm_oor_detail.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_acl_oper_ipv4_acl_and_prefix_list_oor_access_list_summary_details is a generated protocol buffer package.

It is generated from these files:
	ipv4_acl_edm_oor_detail.proto

It has these top-level messages:
	Ipv4AclEdmOorDetail_KEYS
	Ipv4AclEdmOorDetail
*/
package cisco_ios_xr_ipv4_acl_oper_ipv4_acl_and_prefix_list_oor_access_list_summary_details

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Oor deatil config BAG
type Ipv4AclEdmOorDetail_KEYS struct {
}

func (m *Ipv4AclEdmOorDetail_KEYS) Reset()                    { *m = Ipv4AclEdmOorDetail_KEYS{} }
func (m *Ipv4AclEdmOorDetail_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ipv4AclEdmOorDetail_KEYS) ProtoMessage()               {}
func (*Ipv4AclEdmOorDetail_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Ipv4AclEdmOorDetail struct {
	// default max configurable acls
	DefaultMaxAcLs uint32 `protobuf:"varint,50,opt,name=default_max_ac_ls,json=defaultMaxAcLs" json:"default_max_ac_ls,omitempty"`
	// default max configurable aces
	DefaultMaxAcEs uint32 `protobuf:"varint,51,opt,name=default_max_ac_es,json=defaultMaxAcEs" json:"default_max_ac_es,omitempty"`
	// Current configured acls
	CurrentConfiguredAcLs uint32 `protobuf:"varint,52,opt,name=current_configured_ac_ls,json=currentConfiguredAcLs" json:"current_configured_ac_ls,omitempty"`
	// Current configured aces
	CurrentConfiguredAcEs uint32 `protobuf:"varint,53,opt,name=current_configured_ac_es,json=currentConfiguredAcEs" json:"current_configured_ac_es,omitempty"`
	// Current max configurable acls
	CurrentMaxConfigurableAcLs uint32 `protobuf:"varint,54,opt,name=current_max_configurable_ac_ls,json=currentMaxConfigurableAcLs" json:"current_max_configurable_ac_ls,omitempty"`
	// Current max configurable aces
	CurrentMaxConfigurableAcEs uint32 `protobuf:"varint,55,opt,name=current_max_configurable_ac_es,json=currentMaxConfigurableAcEs" json:"current_max_configurable_ac_es,omitempty"`
	// max configurable acls
	MaxConfigurableAcLs uint32 `protobuf:"varint,56,opt,name=max_configurable_ac_ls,json=maxConfigurableAcLs" json:"max_configurable_ac_ls,omitempty"`
	// max configurable aces
	MaxConfigurableAcEs uint32 `protobuf:"varint,57,opt,name=max_configurable_ac_es,json=maxConfigurableAcEs" json:"max_configurable_ac_es,omitempty"`
}

func (m *Ipv4AclEdmOorDetail) Reset()                    { *m = Ipv4AclEdmOorDetail{} }
func (m *Ipv4AclEdmOorDetail) String() string            { return proto.CompactTextString(m) }
func (*Ipv4AclEdmOorDetail) ProtoMessage()               {}
func (*Ipv4AclEdmOorDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ipv4AclEdmOorDetail) GetDefaultMaxAcLs() uint32 {
	if m != nil {
		return m.DefaultMaxAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetDefaultMaxAcEs() uint32 {
	if m != nil {
		return m.DefaultMaxAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentConfiguredAcLs() uint32 {
	if m != nil {
		return m.CurrentConfiguredAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentConfiguredAcEs() uint32 {
	if m != nil {
		return m.CurrentConfiguredAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentMaxConfigurableAcLs() uint32 {
	if m != nil {
		return m.CurrentMaxConfigurableAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetCurrentMaxConfigurableAcEs() uint32 {
	if m != nil {
		return m.CurrentMaxConfigurableAcEs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetMaxConfigurableAcLs() uint32 {
	if m != nil {
		return m.MaxConfigurableAcLs
	}
	return 0
}

func (m *Ipv4AclEdmOorDetail) GetMaxConfigurableAcEs() uint32 {
	if m != nil {
		return m.MaxConfigurableAcEs
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv4AclEdmOorDetail_KEYS)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.oor.access_list_summary.details.ipv4_acl_edm_oor_detail_KEYS")
	proto.RegisterType((*Ipv4AclEdmOorDetail)(nil), "cisco_ios_xr_ipv4_acl_oper.ipv4_acl_and_prefix_list.oor.access_list_summary.details.ipv4_acl_edm_oor_detail")
}

func init() { proto.RegisterFile("ipv4_acl_edm_oor_detail.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x11, 0xc1, 0x62, 0x41, 0xc1, 0x88, 0xba, 0x88, 0x1e, 0x72, 0x95, 0x36, 0x29, 0xbc,
	0xd3, 0xd3, 0x52, 0x8f, 0x54, 0x7a, 0x8d, 0x57, 0x59, 0x0d, 0x73, 0x9b, 0x89, 0x2c, 0x6c, 0xb2,
	0x61, 0x26, 0x91, 0xf8, 0x3b, 0xfd, 0x43, 0xe2, 0x9a, 0x0b, 0x82, 0xb7, 0x5a, 0x0e, 0xef, 0x7d,
	0xef, 0xdb, 0x62, 0xd5, 0x99, 0xad, 0xdf, 0xa6, 0x80, 0xc6, 0x01, 0xe5, 0x25, 0x78, 0xcf, 0x90,
	0x53, 0x83, 0xd6, 0xa5, 0x35, 0xfb, 0xc6, 0x27, 0x4b, 0x63, 0xc5, 0x78, 0xb0, 0x5e, 0xa0, 0x63,
	0x18, 0xba, 0xbe, 0x26, 0x4e, 0x87, 0x0b, 0xab, 0x1c, 0x6a, 0xa6, 0xc2, 0x76, 0xe0, 0xac, 0x34,
	0xa9, 0xf7, 0x9c, 0xa2, 0x31, 0x24, 0x12, 0x6e, 0x90, 0xb6, 0x2c, 0x91, 0xdf, 0xd3, 0xef, 0x65,
	0x19, 0x8f, 0xd4, 0x69, 0xc4, 0x0a, 0x8f, 0xd9, 0xcb, 0x72, 0xfc, 0xb1, 0xad, 0x8e, 0x23, 0x85,
	0xe4, 0x52, 0xed, 0xe7, 0x54, 0x60, 0xeb, 0x1a, 0x28, 0xb1, 0x03, 0x34, 0xe0, 0x44, 0x5f, 0x9d,
	0x6f, 0x5d, 0xec, 0x3e, 0xef, 0xf5, 0xc1, 0x02, 0xbb, 0x7b, 0xf3, 0x24, 0x1b, 0xaa, 0x24, 0x7a,
	0xf2, 0xbb, 0x9a, 0x49, 0x32, 0x53, 0xda, 0xb4, 0xcc, 0x54, 0x35, 0x60, 0x7c, 0x55, 0xd8, 0xd7,
	0x96, 0x29, 0xef, 0xc7, 0xa7, 0x81, 0x38, 0xec, 0xf3, 0xf9, 0x10, 0x07, 0x47, 0x14, 0x24, 0xd1,
	0xd7, 0x51, 0x30, 0x93, 0xe4, 0x41, 0x8d, 0xd6, 0xe0, 0xd7, 0xe3, 0xd6, 0x30, 0xae, 0x1c, 0xf5,
	0xde, 0x9b, 0x80, 0x9f, 0xf4, 0xad, 0x05, 0x76, 0xf3, 0x1f, 0x9d, 0x20, 0xff, 0x67, 0x83, 0x44,
	0xcf, 0xfe, 0xde, 0xc8, 0x24, 0x99, 0xa8, 0xa3, 0x88, 0xff, 0x36, 0xb0, 0x07, 0xe5, 0x06, 0x71,
	0x04, 0x22, 0xd1, 0x77, 0x11, 0x28, 0x93, 0xd5, 0x4e, 0xf8, 0x51, 0x93, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x6f, 0x1d, 0x63, 0x20, 0x72, 0x02, 0x00, 0x00,
}
