// Code generated by protoc-gen-go.
// source: isis_sh_te_adv.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_te_advertisements is a generated protocol buffer package.

It is generated from these files:
	isis_sh_te_adv.proto

It has these top-level messages:
	IsisShTeAdv_KEYS
	IsisShTeAdv
	IsisShTeAdvSubTlv
	IsisShTeAdvSubTlvP
	IsisShTeAdvEntry
	IsisShTeAdvEntryP
	IsisShTePceAdv
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_topologies_topology_topology_levels_topology_level_te_advertisements

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

// TE advertisements for an IS-IS level
type IsisShTeAdv_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	AfName       string `protobuf:"bytes,2,opt,name=af_name,json=afName" json:"af_name,omitempty"`
	SafName      string `protobuf:"bytes,3,opt,name=saf_name,json=safName" json:"saf_name,omitempty"`
	TopologyName string `protobuf:"bytes,4,opt,name=topology_name,json=topologyName" json:"topology_name,omitempty"`
	Level        string `protobuf:"bytes,5,opt,name=level" json:"level,omitempty"`
}

func (m *IsisShTeAdv_KEYS) Reset()                    { *m = IsisShTeAdv_KEYS{} }
func (m *IsisShTeAdv_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeAdv_KEYS) ProtoMessage()               {}
func (*IsisShTeAdv_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShTeAdv_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShTeAdv_KEYS) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IsisShTeAdv_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *IsisShTeAdv_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *IsisShTeAdv_KEYS) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

type IsisShTeAdv struct {
	// Indicates whether any TE advertisements are present
	TeAdvDataPresent bool `protobuf:"varint,50,opt,name=te_adv_data_present,json=teAdvDataPresent" json:"te_adv_data_present,omitempty"`
	// Local system ID
	TeSystemId string `protobuf:"bytes,51,opt,name=te_system_id,json=teSystemId" json:"te_system_id,omitempty"`
	// Local TE router ID
	TeLocalRouterId string `protobuf:"bytes,52,opt,name=te_local_router_id,json=teLocalRouterId" json:"te_local_router_id,omitempty"`
	// List of TE advertisement entries
	TeAdvList []*IsisShTeAdvEntryP `protobuf:"bytes,53,rep,name=te_adv_list,json=teAdvList" json:"te_adv_list,omitempty"`
	// TE PCE advertisements
	Tepceadv *IsisShTePceAdv `protobuf:"bytes,54,opt,name=tepceadv" json:"tepceadv,omitempty"`
}

func (m *IsisShTeAdv) Reset()                    { *m = IsisShTeAdv{} }
func (m *IsisShTeAdv) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeAdv) ProtoMessage()               {}
func (*IsisShTeAdv) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShTeAdv) GetTeAdvDataPresent() bool {
	if m != nil {
		return m.TeAdvDataPresent
	}
	return false
}

func (m *IsisShTeAdv) GetTeSystemId() string {
	if m != nil {
		return m.TeSystemId
	}
	return ""
}

func (m *IsisShTeAdv) GetTeLocalRouterId() string {
	if m != nil {
		return m.TeLocalRouterId
	}
	return ""
}

func (m *IsisShTeAdv) GetTeAdvList() []*IsisShTeAdvEntryP {
	if m != nil {
		return m.TeAdvList
	}
	return nil
}

func (m *IsisShTeAdv) GetTepceadv() *IsisShTePceAdv {
	if m != nil {
		return m.Tepceadv
	}
	return nil
}

type IsisShTeAdvSubTlv struct {
	TeSubTlvType   uint32 `protobuf:"varint,1,opt,name=te_sub_tlv_type,json=teSubTlvType" json:"te_sub_tlv_type,omitempty"`
	TeSubTlvLength uint32 `protobuf:"varint,2,opt,name=te_sub_tlv_length,json=teSubTlvLength" json:"te_sub_tlv_length,omitempty"`
	TeSubTlvValue  []byte `protobuf:"bytes,3,opt,name=te_sub_tlv_value,json=teSubTlvValue,proto3" json:"te_sub_tlv_value,omitempty"`
}

func (m *IsisShTeAdvSubTlv) Reset()                    { *m = IsisShTeAdvSubTlv{} }
func (m *IsisShTeAdvSubTlv) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeAdvSubTlv) ProtoMessage()               {}
func (*IsisShTeAdvSubTlv) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IsisShTeAdvSubTlv) GetTeSubTlvType() uint32 {
	if m != nil {
		return m.TeSubTlvType
	}
	return 0
}

func (m *IsisShTeAdvSubTlv) GetTeSubTlvLength() uint32 {
	if m != nil {
		return m.TeSubTlvLength
	}
	return 0
}

func (m *IsisShTeAdvSubTlv) GetTeSubTlvValue() []byte {
	if m != nil {
		return m.TeSubTlvValue
	}
	return nil
}

type IsisShTeAdvSubTlvP struct {
	Value *IsisShTeAdvSubTlv `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisShTeAdvSubTlvP) Reset()                    { *m = IsisShTeAdvSubTlvP{} }
func (m *IsisShTeAdvSubTlvP) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeAdvSubTlvP) ProtoMessage()               {}
func (*IsisShTeAdvSubTlvP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *IsisShTeAdvSubTlvP) GetValue() *IsisShTeAdvSubTlv {
	if m != nil {
		return m.Value
	}
	return nil
}

type IsisShTeAdvEntry struct {
	TeNeighborSystemId             string                `protobuf:"bytes,1,opt,name=te_neighbor_system_id,json=teNeighborSystemId" json:"te_neighbor_system_id,omitempty"`
	LinkType                       string                `protobuf:"bytes,2,opt,name=link_type,json=linkType" json:"link_type,omitempty"`
	LocalIpAddress                 string                `protobuf:"bytes,3,opt,name=local_ip_address,json=localIpAddress" json:"local_ip_address,omitempty"`
	TeNeighborIpAddress            string                `protobuf:"bytes,4,opt,name=te_neighbor_ip_address,json=teNeighborIpAddress" json:"te_neighbor_ip_address,omitempty"`
	TeMetric                       uint32                `protobuf:"varint,5,opt,name=te_metric,json=teMetric" json:"te_metric,omitempty"`
	TePhysicalLinkBandwidth        uint32                `protobuf:"varint,6,opt,name=te_physical_link_bandwidth,json=tePhysicalLinkBandwidth" json:"te_physical_link_bandwidth,omitempty"`
	TeReservedLinkBandwidth        uint32                `protobuf:"varint,7,opt,name=te_reserved_link_bandwidth,json=teReservedLinkBandwidth" json:"te_reserved_link_bandwidth,omitempty"`
	TeTransmittedBandwidth         []uint32              `protobuf:"varint,8,rep,packed,name=te_transmitted_bandwidth,json=teTransmittedBandwidth" json:"te_transmitted_bandwidth,omitempty"`
	TeSubpoolReservedLinkBandwidth uint32                `protobuf:"varint,9,opt,name=te_subpool_reserved_link_bandwidth,json=teSubpoolReservedLinkBandwidth" json:"te_subpool_reserved_link_bandwidth,omitempty"`
	TeSubpoolTransmittedBandwidth  []uint32              `protobuf:"varint,10,rep,packed,name=te_subpool_transmitted_bandwidth,json=teSubpoolTransmittedBandwidth" json:"te_subpool_transmitted_bandwidth,omitempty"`
	TeAffinity                     uint32                `protobuf:"varint,11,opt,name=te_affinity,json=teAffinity" json:"te_affinity,omitempty"`
	TeExtAdminNum                  uint32                `protobuf:"varint,12,opt,name=te_ext_admin_num,json=teExtAdminNum" json:"te_ext_admin_num,omitempty"`
	TeExtAdminSub                  []uint32              `protobuf:"varint,13,rep,packed,name=te_ext_admin_sub,json=teExtAdminSub" json:"te_ext_admin_sub,omitempty"`
	TeSubTlvDataPresent            bool                  `protobuf:"varint,14,opt,name=te_sub_tlv_data_present,json=teSubTlvDataPresent" json:"te_sub_tlv_data_present,omitempty"`
	TeSubTlvList                   []*IsisShTeAdvSubTlvP `protobuf:"bytes,15,rep,name=te_sub_tlv_list,json=teSubTlvList" json:"te_sub_tlv_list,omitempty"`
}

func (m *IsisShTeAdvEntry) Reset()                    { *m = IsisShTeAdvEntry{} }
func (m *IsisShTeAdvEntry) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeAdvEntry) ProtoMessage()               {}
func (*IsisShTeAdvEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IsisShTeAdvEntry) GetTeNeighborSystemId() string {
	if m != nil {
		return m.TeNeighborSystemId
	}
	return ""
}

func (m *IsisShTeAdvEntry) GetLinkType() string {
	if m != nil {
		return m.LinkType
	}
	return ""
}

func (m *IsisShTeAdvEntry) GetLocalIpAddress() string {
	if m != nil {
		return m.LocalIpAddress
	}
	return ""
}

func (m *IsisShTeAdvEntry) GetTeNeighborIpAddress() string {
	if m != nil {
		return m.TeNeighborIpAddress
	}
	return ""
}

func (m *IsisShTeAdvEntry) GetTeMetric() uint32 {
	if m != nil {
		return m.TeMetric
	}
	return 0
}

func (m *IsisShTeAdvEntry) GetTePhysicalLinkBandwidth() uint32 {
	if m != nil {
		return m.TePhysicalLinkBandwidth
	}
	return 0
}

func (m *IsisShTeAdvEntry) GetTeReservedLinkBandwidth() uint32 {
	if m != nil {
		return m.TeReservedLinkBandwidth
	}
	return 0
}

func (m *IsisShTeAdvEntry) GetTeTransmittedBandwidth() []uint32 {
	if m != nil {
		return m.TeTransmittedBandwidth
	}
	return nil
}

func (m *IsisShTeAdvEntry) GetTeSubpoolReservedLinkBandwidth() uint32 {
	if m != nil {
		return m.TeSubpoolReservedLinkBandwidth
	}
	return 0
}

func (m *IsisShTeAdvEntry) GetTeSubpoolTransmittedBandwidth() []uint32 {
	if m != nil {
		return m.TeSubpoolTransmittedBandwidth
	}
	return nil
}

func (m *IsisShTeAdvEntry) GetTeAffinity() uint32 {
	if m != nil {
		return m.TeAffinity
	}
	return 0
}

func (m *IsisShTeAdvEntry) GetTeExtAdminNum() uint32 {
	if m != nil {
		return m.TeExtAdminNum
	}
	return 0
}

func (m *IsisShTeAdvEntry) GetTeExtAdminSub() []uint32 {
	if m != nil {
		return m.TeExtAdminSub
	}
	return nil
}

func (m *IsisShTeAdvEntry) GetTeSubTlvDataPresent() bool {
	if m != nil {
		return m.TeSubTlvDataPresent
	}
	return false
}

func (m *IsisShTeAdvEntry) GetTeSubTlvList() []*IsisShTeAdvSubTlvP {
	if m != nil {
		return m.TeSubTlvList
	}
	return nil
}

type IsisShTeAdvEntryP struct {
	Value *IsisShTeAdvEntry `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *IsisShTeAdvEntryP) Reset()                    { *m = IsisShTeAdvEntryP{} }
func (m *IsisShTeAdvEntryP) String() string            { return proto.CompactTextString(m) }
func (*IsisShTeAdvEntryP) ProtoMessage()               {}
func (*IsisShTeAdvEntryP) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IsisShTeAdvEntryP) GetValue() *IsisShTeAdvEntry {
	if m != nil {
		return m.Value
	}
	return nil
}

// TE PCE advertisements for an IS-IS level
type IsisShTePceAdv struct {
	// Indicates whether any PCE advertisements are present
	PceAdvDataPresent bool `protobuf:"varint,1,opt,name=pce_adv_data_present,json=pceAdvDataPresent" json:"pce_adv_data_present,omitempty"`
	// Flooding scope for PCE advertisement
	PceFloodingScope string `protobuf:"bytes,2,opt,name=pce_flooding_scope,json=pceFloodingScope" json:"pce_flooding_scope,omitempty"`
	// IPv4 PCE address
	PceAddressIpv4 string `protobuf:"bytes,3,opt,name=pce_address_ipv4,json=pceAddressIpv4" json:"pce_address_ipv4,omitempty"`
	// PCE path scope bits
	PcePathScopeBits uint32 `protobuf:"varint,4,opt,name=pce_path_scope_bits,json=pcePathScopeBits" json:"pce_path_scope_bits,omitempty"`
	// PCE path scope preferences
	PcePathScopePrefs uint32 `protobuf:"varint,5,opt,name=pce_path_scope_prefs,json=pcePathScopePrefs" json:"pce_path_scope_prefs,omitempty"`
}

func (m *IsisShTePceAdv) Reset()                    { *m = IsisShTePceAdv{} }
func (m *IsisShTePceAdv) String() string            { return proto.CompactTextString(m) }
func (*IsisShTePceAdv) ProtoMessage()               {}
func (*IsisShTePceAdv) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *IsisShTePceAdv) GetPceAdvDataPresent() bool {
	if m != nil {
		return m.PceAdvDataPresent
	}
	return false
}

func (m *IsisShTePceAdv) GetPceFloodingScope() string {
	if m != nil {
		return m.PceFloodingScope
	}
	return ""
}

func (m *IsisShTePceAdv) GetPceAddressIpv4() string {
	if m != nil {
		return m.PceAddressIpv4
	}
	return ""
}

func (m *IsisShTePceAdv) GetPcePathScopeBits() uint32 {
	if m != nil {
		return m.PcePathScopeBits
	}
	return 0
}

func (m *IsisShTePceAdv) GetPcePathScopePrefs() uint32 {
	if m != nil {
		return m.PcePathScopePrefs
	}
	return 0
}

func init() {
	proto.RegisterType((*IsisShTeAdv_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_advertisements.isis_sh_te_adv_KEYS")
	proto.RegisterType((*IsisShTeAdv)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_advertisements.isis_sh_te_adv")
	proto.RegisterType((*IsisShTeAdvSubTlv)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_advertisements.isis_sh_te_adv_sub_tlv")
	proto.RegisterType((*IsisShTeAdvSubTlvP)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_advertisements.isis_sh_te_adv_sub_tlv_p")
	proto.RegisterType((*IsisShTeAdvEntry)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_advertisements.isis_sh_te_adv_entry")
	proto.RegisterType((*IsisShTeAdvEntryP)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_advertisements.isis_sh_te_adv_entry_p")
	proto.RegisterType((*IsisShTePceAdv)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.topologies.topology.topology_levels.topology_level.te_advertisements.isis_sh_te_pce_adv")
}

func init() { proto.RegisterFile("isis_sh_te_adv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 904 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xcb, 0x6e, 0x2b, 0x45,
	0x10, 0xd5, 0x10, 0x92, 0xd8, 0x1d, 0x4f, 0x1e, 0x9d, 0x90, 0x0c, 0x5c, 0x01, 0x96, 0xaf, 0x10,
	0x46, 0x80, 0x11, 0x49, 0x40, 0x48, 0xac, 0x72, 0xc5, 0x05, 0x05, 0x42, 0x14, 0x8d, 0x23, 0x24,
	0x56, 0xad, 0xf6, 0x4c, 0x39, 0x6e, 0xdd, 0x79, 0xb4, 0xa6, 0xcb, 0x43, 0xfc, 0x0f, 0x08, 0xf1,
	0xf8, 0x06, 0x84, 0xb2, 0x64, 0xcb, 0x8a, 0x05, 0xdf, 0xc3, 0x1f, 0x20, 0xa1, 0xae, 0x9e, 0xb1,
	0xc7, 0x4e, 0x58, 0x93, 0x5d, 0x77, 0x9d, 0x53, 0x55, 0xc7, 0xee, 0xaa, 0x63, 0xb3, 0x03, 0x65,
	0x94, 0x11, 0x66, 0x22, 0x10, 0x84, 0x8c, 0xcb, 0x81, 0x2e, 0x72, 0xcc, 0xf9, 0x6d, 0xa4, 0x4c,
	0x94, 0x0b, 0x95, 0x1b, 0x71, 0x5b, 0x88, 0x28, 0xc9, 0x8c, 0x20, 0x5e, 0xae, 0xa1, 0x18, 0xd8,
	0xd3, 0x40, 0x65, 0x06, 0x65, 0x16, 0xc1, 0xe2, 0x34, 0xc0, 0x5c, 0xe7, 0x49, 0x7e, 0xa3, 0xc0,
	0xd4, 0xc7, 0xd9, 0xfc, 0x20, 0x12, 0x28, 0x21, 0x31, 0x2b, 0xf7, 0x81, 0x6b, 0x0a, 0x05, 0x2a,
	0x03, 0x29, 0x64, 0x68, 0x7a, 0x77, 0x1e, 0xdb, 0x5f, 0x96, 0x24, 0xbe, 0x7a, 0xfe, 0xed, 0x90,
	0x3f, 0x65, 0x7e, 0xdd, 0x48, 0x64, 0x32, 0x85, 0xc0, 0xeb, 0x7a, 0xfd, 0x76, 0xd8, 0xa9, 0x83,
	0x97, 0x32, 0x05, 0x7e, 0xc4, 0x36, 0xe5, 0xd8, 0xc1, 0x2f, 0x11, 0xbc, 0x21, 0xc7, 0x04, 0xbc,
	0xca, 0x5a, 0xa6, 0x46, 0xd6, 0x08, 0xd9, 0x34, 0x15, 0xf4, 0x94, 0xf9, 0x73, 0x51, 0x84, 0xbf,
	0xec, 0x0a, 0xd7, 0x41, 0x22, 0x1d, 0xb0, 0x75, 0x12, 0x1c, 0xac, 0x13, 0xe8, 0x2e, 0xbd, 0xbf,
	0xd7, 0xd8, 0xf6, 0xb2, 0x56, 0xfe, 0x3e, 0xdb, 0xaf, 0x54, 0xc7, 0x12, 0xa5, 0xd0, 0x05, 0x18,
	0xc8, 0x30, 0x38, 0xee, 0x7a, 0xfd, 0x56, 0xb8, 0x8b, 0x70, 0x16, 0x97, 0x9f, 0x49, 0x94, 0x57,
	0x2e, 0xce, 0xbb, 0xac, 0x83, 0x20, 0xcc, 0xcc, 0x20, 0xa4, 0x42, 0xc5, 0xc1, 0x09, 0x95, 0x67,
	0x08, 0x43, 0x0a, 0x9d, 0xc7, 0xfc, 0x5d, 0xc6, 0x11, 0x44, 0x92, 0x47, 0x32, 0x11, 0x45, 0x3e,
	0x45, 0x28, 0x2c, 0xef, 0x94, 0x78, 0x3b, 0x08, 0x17, 0x16, 0x08, 0x29, 0x7e, 0x1e, 0xf3, 0xdf,
	0x3d, 0xb6, 0x55, 0xb5, 0x4f, 0x94, 0xc1, 0xe0, 0xa3, 0xee, 0x5a, 0x7f, 0xeb, 0xf8, 0x47, 0x6f,
	0xf0, 0x7f, 0x3d, 0xe7, 0x60, 0xe5, 0x29, 0x21, 0xc3, 0x62, 0x26, 0x74, 0xd8, 0xa6, 0x6f, 0xe2,
	0x42, 0x19, 0xe4, 0x77, 0x1e, 0x6b, 0x21, 0xe8, 0x08, 0x64, 0x5c, 0x06, 0x1f, 0x77, 0xbd, 0xfe,
	0xd6, 0xf1, 0xf7, 0x8f, 0x43, 0xb0, 0x8e, 0x08, 0x0d, 0xe7, 0xf2, 0x7a, 0xbf, 0x78, 0xec, 0x70,
	0xe5, 0x13, 0x99, 0xe9, 0x48, 0x60, 0x52, 0xf2, 0xb7, 0xd8, 0x8e, 0x7d, 0x49, 0x77, 0x13, 0x38,
	0xd3, 0x6e, 0x42, 0xfd, 0xb0, 0x83, 0x30, 0x9c, 0x8e, 0xae, 0x93, 0xf2, 0x7a, 0xa6, 0x81, 0xbf,
	0xc3, 0xf6, 0x1a, 0xb4, 0x04, 0xb2, 0x1b, 0x9c, 0xd0, 0xac, 0xfa, 0xe1, 0x76, 0x4d, 0xbc, 0xa0,
	0x28, 0x7f, 0x9b, 0xed, 0x36, 0xa8, 0xa5, 0x4c, 0xa6, 0x6e, 0x76, 0x3b, 0xa1, 0x5f, 0x33, 0xbf,
	0xb1, 0xc1, 0xde, 0x5f, 0x1e, 0x0b, 0x1e, 0x56, 0x25, 0x34, 0xff, 0xcd, 0x63, 0xeb, 0x2e, 0xd7,
	0xa3, 0xef, 0xf6, 0x11, 0x0d, 0x43, 0x25, 0x32, 0x74, 0xfa, 0x7a, 0x3f, 0x6d, 0xae, 0x9a, 0x91,
	0x1b, 0x17, 0xfe, 0x21, 0x7b, 0x05, 0x41, 0x64, 0xa0, 0x6e, 0x26, 0xa3, 0xbc, 0x68, 0x6c, 0x8b,
	0xb3, 0x00, 0x8e, 0x70, 0x59, 0x61, 0xf3, 0xad, 0x79, 0xc2, 0xda, 0x89, 0xca, 0x5e, 0xb8, 0x77,
	0x70, 0x56, 0xd0, 0xb2, 0x01, 0x7a, 0x83, 0x3e, 0xdb, 0x75, 0xfb, 0xa4, 0xb4, 0x90, 0x71, 0x5c,
	0x80, 0x31, 0x95, 0x29, 0x6c, 0x53, 0xfc, 0x5c, 0x9f, 0xb9, 0x28, 0x3f, 0x61, 0x87, 0xcd, 0xce,
	0x0d, 0xbe, 0x33, 0x89, 0xfd, 0x45, 0xeb, 0x45, 0xd2, 0x13, 0xd6, 0x46, 0x10, 0x29, 0x60, 0xa1,
	0x22, 0xf2, 0x0b, 0xdf, 0x4e, 0xd0, 0xd7, 0x74, 0xe7, 0x9f, 0xb2, 0xd7, 0xec, 0x64, 0x4d, 0x66,
	0x46, 0x59, 0x05, 0x24, 0x72, 0x24, 0xb3, 0xf8, 0x3b, 0x15, 0xe3, 0x24, 0xd8, 0x20, 0xf6, 0x11,
	0xc2, 0x55, 0x45, 0xb8, 0x50, 0xd9, 0x8b, 0x67, 0x35, 0x5c, 0x25, 0x5b, 0xeb, 0x28, 0x4a, 0x88,
	0x57, 0x93, 0x37, 0xeb, 0xe4, 0xb0, 0x22, 0x2c, 0x27, 0x7f, 0xc2, 0x02, 0x04, 0x81, 0x85, 0xcc,
	0x4c, 0xaa, 0x10, 0x21, 0x6e, 0xa4, 0xb6, 0xba, 0x6b, 0x7d, 0x3f, 0x3c, 0x44, 0xb8, 0x5e, 0xc0,
	0x8b, 0xcc, 0x2f, 0x59, 0xcf, 0x0d, 0xa2, 0xce, 0xf3, 0xe4, 0x3f, 0xdb, 0xb7, 0xa9, 0xfd, 0x1b,
	0x34, 0x9a, 0x96, 0xf8, 0xb0, 0x8a, 0x2f, 0x58, 0xb7, 0x51, 0xeb, 0x61, 0x35, 0x8c, 0xd4, 0xbc,
	0x3e, 0xaf, 0xf4, 0xa0, 0xa8, 0x37, 0x9d, 0xd3, 0x8d, 0xc7, 0x2a, 0x53, 0x38, 0x0b, 0xb6, 0xa8,
	0x3b, 0x43, 0x38, 0xab, 0x22, 0xd5, 0xfa, 0xc0, 0x2d, 0x0a, 0x19, 0xa7, 0x2a, 0x13, 0xd9, 0x34,
	0x0d, 0x3a, 0xc4, 0xf2, 0x11, 0x9e, 0xdf, 0xe2, 0x99, 0x8d, 0x5e, 0x4e, 0xd3, 0x7b, 0x44, 0x33,
	0x1d, 0x05, 0x3e, 0x49, 0x68, 0x10, 0x87, 0xd3, 0x11, 0x3f, 0x65, 0x47, 0x8d, 0x85, 0x5c, 0xf2,
	0xf7, 0x6d, 0xf2, 0xf7, 0xfd, 0x7a, 0x2f, 0x9b, 0x16, 0xff, 0x87, 0xb7, 0xe4, 0x0c, 0xe4, 0xcb,
	0x3b, 0xe4, 0xcb, 0x3f, 0x3f, 0xba, 0x55, 0x14, 0x7a, 0x61, 0x57, 0xd6, 0x9c, 0x7b, 0x7f, 0xde,
	0x37, 0xbc, 0xca, 0xc2, 0xf9, 0xaf, 0x2b, 0xc6, 0xf2, 0xc3, 0x23, 0xfb, 0x95, 0xa9, 0x6d, 0xe5,
	0x1f, 0x8f, 0xf1, 0xfb, 0xa6, 0xce, 0x3f, 0x60, 0x07, 0xd5, 0x71, 0xf9, 0x25, 0x3d, 0x7a, 0xc9,
	0x3d, 0x1d, 0xad, 0xfe, 0x54, 0xbf, 0xc7, 0xb8, 0x4d, 0x18, 0x27, 0x79, 0x1e, 0xab, 0xec, 0x46,
	0x98, 0x28, 0x9f, 0x7b, 0xcb, 0xae, 0x8e, 0xe0, 0xf3, 0x0a, 0x18, 0xda, 0xb8, 0xf5, 0x18, 0x57,
	0x9e, 0x3c, 0x41, 0x28, 0x5d, 0x9e, 0xd6, 0x1e, 0x43, 0xa5, 0x29, 0x7c, 0xae, 0xcb, 0x53, 0xfb,
	0x8f, 0xc1, 0x32, 0xb5, 0xc4, 0x89, 0xab, 0x29, 0x46, 0x0a, 0x9d, 0xc1, 0xf8, 0x54, 0xf8, 0x4a,
	0xe2, 0x84, 0x8a, 0x3e, 0x53, 0x68, 0x6a, 0xdd, 0x0d, 0xba, 0x2e, 0x60, 0x6c, 0x2a, 0xa3, 0xd9,
	0x6b, 0xf2, 0xaf, 0x2c, 0x30, 0xda, 0xa0, 0x7f, 0x74, 0x27, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x85, 0xbc, 0xfb, 0xc6, 0xe9, 0x09, 0x00, 0x00,
}