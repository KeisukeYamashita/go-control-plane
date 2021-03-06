// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/zookeeper_proxy/v1alpha1/zookeeper_proxy.proto

package envoy_config_filter_network_zookeeper_proxy_v1alpha1

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ZooKeeperProxy struct {
	StatPrefix           string                `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	AccessLog            string                `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	MaxPacketBytes       *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_packet_bytes,json=maxPacketBytes,proto3" json:"max_packet_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ZooKeeperProxy) Reset()         { *m = ZooKeeperProxy{} }
func (m *ZooKeeperProxy) String() string { return proto.CompactTextString(m) }
func (*ZooKeeperProxy) ProtoMessage()    {}
func (*ZooKeeperProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_05247350458709ad, []int{0}
}

func (m *ZooKeeperProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZooKeeperProxy.Unmarshal(m, b)
}
func (m *ZooKeeperProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZooKeeperProxy.Marshal(b, m, deterministic)
}
func (m *ZooKeeperProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZooKeeperProxy.Merge(m, src)
}
func (m *ZooKeeperProxy) XXX_Size() int {
	return xxx_messageInfo_ZooKeeperProxy.Size(m)
}
func (m *ZooKeeperProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ZooKeeperProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ZooKeeperProxy proto.InternalMessageInfo

func (m *ZooKeeperProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ZooKeeperProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *ZooKeeperProxy) GetMaxPacketBytes() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPacketBytes
	}
	return nil
}

func init() {
	proto.RegisterType((*ZooKeeperProxy)(nil), "envoy.config.filter.network.zookeeper_proxy.v1alpha1.ZooKeeperProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/zookeeper_proxy/v1alpha1/zookeeper_proxy.proto", fileDescriptor_05247350458709ad)
}

var fileDescriptor_05247350458709ad = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbd, 0x8a, 0xdb, 0x40,
	0x14, 0x85, 0x19, 0x27, 0x71, 0xf0, 0x18, 0x8c, 0x51, 0x8a, 0x18, 0x13, 0x1b, 0x93, 0xca, 0xd5,
	0x0c, 0xb6, 0x92, 0x17, 0x50, 0x11, 0xc8, 0x4f, 0x21, 0x0c, 0x49, 0xe1, 0x46, 0x8c, 0xe5, 0x2b,
	0x45, 0x58, 0x9e, 0x3b, 0xcc, 0x8c, 0x6c, 0x39, 0x55, 0xde, 0x20, 0x6d, 0xaa, 0x40, 0x5e, 0x63,
	0x9f, 0x60, 0xdb, 0x7d, 0x95, 0x2d, 0xb7, 0x58, 0x16, 0xcd, 0x58, 0x8d, 0x97, 0x6d, 0xb6, 0x13,
	0xfa, 0x74, 0x3e, 0xee, 0x39, 0xa2, 0x5f, 0x40, 0x1e, 0xf0, 0xc4, 0x53, 0x94, 0x59, 0x91, 0xf3,
	0xac, 0x28, 0x2d, 0x68, 0x2e, 0xc1, 0x1e, 0x51, 0xef, 0xf8, 0x2f, 0xc4, 0x1d, 0x80, 0x02, 0x9d,
	0x28, 0x8d, 0xf5, 0x89, 0x1f, 0x16, 0xa2, 0x54, 0x3f, 0xc5, 0xe2, 0x12, 0x30, 0xa5, 0xd1, 0x62,
	0xf0, 0xc1, 0xb9, 0x98, 0x77, 0x31, 0xef, 0x62, 0x67, 0x17, 0xbb, 0x8c, 0xb4, 0xae, 0xf1, 0x34,
	0x47, 0xcc, 0x4b, 0xe0, 0xce, 0xb1, 0xa9, 0x32, 0x7e, 0xd4, 0x42, 0x29, 0xd0, 0xc6, 0x5b, 0xc7,
	0xd3, 0x6a, 0xab, 0x04, 0x17, 0x52, 0xa2, 0x15, 0xb6, 0x40, 0x69, 0xf8, 0xbe, 0xc8, 0xb5, 0xb0,
	0x70, 0xe6, 0x93, 0x47, 0xdc, 0x58, 0x61, 0xab, 0x36, 0xfe, 0xf6, 0x20, 0xca, 0x62, 0x2b, 0x2c,
	0xf0, 0xf6, 0xc1, 0x83, 0xf7, 0xff, 0x09, 0x1d, 0xac, 0x11, 0xbf, 0xba, 0xa3, 0xe2, 0xe6, 0xa6,
	0x60, 0x4e, 0xfb, 0x4d, 0x36, 0x51, 0x1a, 0xb2, 0xa2, 0x1e, 0x91, 0x19, 0x99, 0xf7, 0xa2, 0xd7,
	0x77, 0xd1, 0x4b, 0xdd, 0x99, 0x91, 0x15, 0x6d, 0x58, 0xec, 0x50, 0x30, 0xa1, 0x54, 0xa4, 0x29,
	0x18, 0x93, 0x94, 0x98, 0x8f, 0x3a, 0xcd, 0x87, 0xab, 0x9e, 0x7f, 0xf3, 0x0d, 0xf3, 0xe0, 0x13,
	0x1d, 0xee, 0x45, 0x9d, 0x28, 0x91, 0xee, 0xc0, 0x26, 0x9b, 0x93, 0x05, 0x33, 0x7a, 0x31, 0x23,
	0xf3, 0xfe, 0xf2, 0x1d, 0xf3, 0x75, 0x59, 0x5b, 0x97, 0x7d, 0xff, 0x2c, 0x6d, 0xb8, 0xfc, 0x21,
	0xca, 0x0a, 0x56, 0x83, 0xbd, 0xa8, 0x63, 0x17, 0x8a, 0x9a, 0x4c, 0xf4, 0x8f, 0xdc, 0xfe, 0xbd,
	0xff, 0xf3, 0xea, 0x63, 0x10, 0xfa, 0x69, 0xa1, 0xb6, 0x20, 0x4d, 0x53, 0xf2, 0x3c, 0xaf, 0x79,
	0x7a, 0xdf, 0xf0, 0xea, 0xf7, 0xf5, 0x4d, 0xb7, 0x33, 0x24, 0x34, 0x2a, 0x90, 0xb9, 0xbc, 0x27,
	0xcf, 0xf9, 0x4b, 0xd1, 0x9b, 0x75, 0x4b, 0xdc, 0x54, 0x71, 0x73, 0x7d, 0x4c, 0x36, 0x5d, 0x57,
	0x23, 0x7c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x7f, 0x88, 0x1f, 0x47, 0x02, 0x00, 0x00,
}
