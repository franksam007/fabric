// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dah.proto

/*
Package util is a generated protocol buffer package.

It is generated from these files:
	dah.proto

It has these top-level messages:
	TX
*/
package util

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

type TX struct {
	Version  uint32      `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	LockTime uint32      `protobuf:"varint,2,opt,name=lockTime" json:"lockTime,omitempty"`
	Txin     []*TX_TXIN  `protobuf:"bytes,3,rep,name=txin" json:"txin,omitempty"`
	Txout    []*TX_TXOUT `protobuf:"bytes,4,rep,name=txout" json:"txout,omitempty"`
	Blocks   [][]byte    `protobuf:"bytes,5,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Fee      uint64      `protobuf:"varint,6,opt,name=fee" json:"fee,omitempty"`
}

func (m *TX) Reset()                    { *m = TX{} }
func (m *TX) String() string            { return proto.CompactTextString(m) }
func (*TX) ProtoMessage()               {}
func (*TX) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TX) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TX) GetLockTime() uint32 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

func (m *TX) GetTxin() []*TX_TXIN {
	if m != nil {
		return m.Txin
	}
	return nil
}

func (m *TX) GetTxout() []*TX_TXOUT {
	if m != nil {
		return m.Txout
	}
	return nil
}

func (m *TX) GetBlocks() [][]byte {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *TX) GetFee() uint64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type TX_TXIN struct {
	Ix         uint32 `protobuf:"varint,1,opt,name=ix" json:"ix,omitempty"`
	SourceHash []byte `protobuf:"bytes,2,opt,name=sourceHash,proto3" json:"sourceHash,omitempty"`
	Script     []byte `protobuf:"bytes,3,opt,name=script,proto3" json:"script,omitempty"`
	Sequence   uint32 `protobuf:"varint,4,opt,name=sequence" json:"sequence,omitempty"`
}

func (m *TX_TXIN) Reset()                    { *m = TX_TXIN{} }
func (m *TX_TXIN) String() string            { return proto.CompactTextString(m) }
func (*TX_TXIN) ProtoMessage()               {}
func (*TX_TXIN) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *TX_TXIN) GetIx() uint32 {
	if m != nil {
		return m.Ix
	}
	return 0
}

func (m *TX_TXIN) GetSourceHash() []byte {
	if m != nil {
		return m.SourceHash
	}
	return nil
}

func (m *TX_TXIN) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *TX_TXIN) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

type TX_TXOUT struct {
	Value    uint64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Script   []byte `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Color    []byte `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Quantity uint64 `protobuf:"varint,4,opt,name=quantity" json:"quantity,omitempty"`
}

func (m *TX_TXOUT) Reset()                    { *m = TX_TXOUT{} }
func (m *TX_TXOUT) String() string            { return proto.CompactTextString(m) }
func (*TX_TXOUT) ProtoMessage()               {}
func (*TX_TXOUT) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func (m *TX_TXOUT) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *TX_TXOUT) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *TX_TXOUT) GetColor() []byte {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *TX_TXOUT) GetQuantity() uint64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func init() {
	proto.RegisterType((*TX)(nil), "util.TX")
	proto.RegisterType((*TX_TXIN)(nil), "util.TX.TXIN")
	proto.RegisterType((*TX_TXOUT)(nil), "util.TX.TXOUT")
}

func init() { proto.RegisterFile("dah.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x18, 0xc5, 0x59, 0x9b, 0x4e, 0x8d, 0xdb, 0x90, 0x30, 0x24, 0xec, 0x20, 0x55, 0x3c, 0xf4, 0xd4,
	0x82, 0x9e, 0x45, 0xf0, 0xa4, 0x17, 0x85, 0x50, 0x61, 0x78, 0x6b, 0xb3, 0x6f, 0x6d, 0xb4, 0x6b,
	0xba, 0x34, 0x19, 0xdd, 0xd5, 0xbf, 0x5c, 0x92, 0x76, 0x63, 0x5e, 0x42, 0x7e, 0x79, 0x1f, 0xef,
	0xbd, 0x8f, 0xe0, 0x8b, 0x55, 0x56, 0xc6, 0x8d, 0x92, 0x5a, 0x12, 0x64, 0xb4, 0xa8, 0xee, 0x7e,
	0x7d, 0xec, 0xa5, 0x4b, 0x42, 0xf1, 0xd9, 0x0e, 0x54, 0x2b, 0x64, 0x4d, 0x47, 0xe1, 0x28, 0x9a,
	0xb2, 0x03, 0x92, 0x05, 0x3e, 0xaf, 0x24, 0xff, 0x49, 0xc5, 0x06, 0xa8, 0xe7, 0xa4, 0x23, 0x93,
	0x5b, 0x8c, 0x74, 0x27, 0x6a, 0xea, 0x87, 0x7e, 0x74, 0xf9, 0x30, 0x8d, 0xad, 0x63, 0x9c, 0x2e,
	0xe3, 0x74, 0xf9, 0xf6, 0xce, 0x9c, 0x44, 0xee, 0x71, 0xa0, 0x3b, 0x69, 0x34, 0x45, 0x6e, 0x66,
	0x76, 0x32, 0xf3, 0xf1, 0x99, 0xb2, 0x5e, 0x24, 0xd7, 0x78, 0x9c, 0x5b, 0xd7, 0x96, 0x06, 0xa1,
	0x1f, 0x4d, 0xd8, 0x40, 0xe4, 0x0a, 0xfb, 0x6b, 0x00, 0x3a, 0x0e, 0x47, 0x11, 0x62, 0xf6, 0xba,
	0xf8, 0xc6, 0xc8, 0xba, 0x93, 0x19, 0xf6, 0x44, 0x37, 0x74, 0xf5, 0x44, 0x47, 0x6e, 0x30, 0x6e,
	0xa5, 0x51, 0x1c, 0x5e, 0xb3, 0xb6, 0x74, 0x45, 0x27, 0xec, 0xe4, 0xc5, 0x26, 0xb4, 0x5c, 0x89,
	0x46, 0x53, 0xdf, 0x69, 0x03, 0xd9, 0xf5, 0x5a, 0xd8, 0x1a, 0xa8, 0x39, 0x50, 0xd4, 0xaf, 0x77,
	0xe0, 0x45, 0x81, 0x03, 0xd7, 0x92, 0xcc, 0x71, 0xb0, 0xcb, 0x2a, 0x03, 0x2e, 0x0f, 0xb1, 0x1e,
	0x4e, 0x2c, 0xbd, 0x7f, 0x96, 0x73, 0x1c, 0x70, 0x59, 0x49, 0x35, 0x24, 0xf5, 0x60, 0x83, 0xb6,
	0x26, 0xab, 0xb5, 0xd0, 0x7b, 0x17, 0x84, 0xd8, 0x91, 0x5f, 0x9e, 0xbf, 0x9e, 0x0a, 0xa1, 0x4b,
	0x93, 0xc7, 0x5c, 0x6e, 0x92, 0x72, 0xdf, 0x80, 0xaa, 0x60, 0x55, 0x80, 0x4a, 0xd6, 0x59, 0xae,
	0x04, 0x4f, 0xa0, 0xcb, 0x36, 0x4d, 0x05, 0x6d, 0xc2, 0xcb, 0x4c, 0xd4, 0x5c, 0xae, 0x20, 0x29,
	0x64, 0x62, 0x74, 0x67, 0x0f, 0x51, 0xe5, 0x63, 0xf7, 0xa5, 0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0x8b, 0x8c, 0x33, 0xdf, 0x01, 0x00, 0x00,
}
