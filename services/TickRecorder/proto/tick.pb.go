// Code generated by protoc-gen-go.
// source: open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto/tick.proto
// DO NOT EDIT!

/*
Package go_srv_TickRecorder is a generated protocol buffer package.

It is generated from these files:
	open-algot.servebeer.com/open-algot/open-algot-platform/services/TickRecorder/proto/tick.proto

It has these top-level messages:
	Tick
	Trade
*/
package go_srv_TickRecorder

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Tick struct {
	Time   int64   `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Bid    float64 `protobuf:"fixed64,2,opt,name=bid" json:"bid,omitempty"`
	Ask    float64 `protobuf:"fixed64,3,opt,name=ask" json:"ask,omitempty"`
	Last   float64 `protobuf:"fixed64,4,opt,name=last" json:"last,omitempty"`
	Pair   string  `protobuf:"bytes,5,opt,name=pair" json:"pair,omitempty"`
	Broker string  `protobuf:"bytes,6,opt,name=broker" json:"broker,omitempty"`
}

func (m *Tick) Reset()                    { *m = Tick{} }
func (m *Tick) String() string            { return proto.CompactTextString(m) }
func (*Tick) ProtoMessage()               {}
func (*Tick) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Trade struct {
	Time   int64   `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Price  float64 `protobuf:"fixed64,2,opt,name=price" json:"price,omitempty"`
	Amount float64 `protobuf:"fixed64,3,opt,name=amount" json:"amount,omitempty"`
	Type   int32   `protobuf:"varint,4,opt,name=type" json:"type,omitempty"`
	Broker string  `protobuf:"bytes,6,opt,name=broker" json:"broker,omitempty"`
	Pair   string  `protobuf:"bytes,7,opt,name=pair" json:"pair,omitempty"`
}

func (m *Trade) Reset()                    { *m = Trade{} }
func (m *Trade) String() string            { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()               {}
func (*Trade) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Tick)(nil), "go.srv.TickRecorder.Tick")
	proto.RegisterType((*Trade)(nil), "go.srv.TickRecorder.Trade")
}

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x8f, 0xbd, 0x4e, 0x03, 0x31,
	0x10, 0x84, 0x75, 0xdc, 0x0f, 0xc2, 0xfc, 0x14, 0xa6, 0x71, 0x89, 0x52, 0xd1, 0xc4, 0x2e, 0x78,
	0x13, 0x94, 0x8a, 0x02, 0x61, 0xfb, 0x96, 0xc8, 0xba, 0x73, 0xd6, 0x5a, 0x2f, 0x91, 0x78, 0x7b,
	0x6c, 0x0b, 0x48, 0x94, 0x6e, 0x3e, 0x7b, 0x76, 0x66, 0x57, 0xbc, 0x63, 0x82, 0xc3, 0xd6, 0xae,
	0x7b, 0x64, 0x9d, 0x81, 0x8e, 0xe0, 0x00, 0x48, 0x7b, 0x8c, 0xe6, 0xf4, 0x71, 0x26, 0xb7, 0x69,
	0xb5, 0xfc, 0x89, 0x14, 0x4d, 0x35, 0x07, 0x0f, 0xd9, 0xec, 0x82, 0x5f, 0x5e, 0xc1, 0x23, 0xcd,
	0x40, 0x26, 0x11, 0x32, 0x1a, 0x2e, 0x4f, 0xba, 0x49, 0xf9, 0xb8, 0x47, 0x9d, 0xe9, 0xa8, 0xcf,
	0x5d, 0x9b, 0x37, 0x31, 0x54, 0x96, 0x77, 0x62, 0xe0, 0x10, 0x41, 0x75, 0x4f, 0xdd, 0x73, 0x2f,
	0x6f, 0x45, 0xef, 0xc2, 0xac, 0xae, 0x0a, 0x74, 0x15, 0x6c, 0x5e, 0x54, 0xdf, 0xa0, 0xf8, 0x56,
	0x9b, 0x59, 0x0d, 0x7f, 0x94, 0x6c, 0x20, 0x35, 0x16, 0xba, 0x91, 0x0f, 0x62, 0x72, 0x84, 0x0b,
	0x90, 0x9a, 0x2a, 0x6f, 0x3e, 0xc4, 0xb8, 0x23, 0x3b, 0xc3, 0x45, 0xf8, 0xbd, 0x18, 0x13, 0x95,
	0x7d, 0x7f, 0xe3, 0xcb, 0x94, 0x8d, 0xf8, 0x75, 0xe0, 0x53, 0x03, 0x7f, 0x27, 0x68, 0x0d, 0xe3,
	0x65, 0xe6, 0x7f, 0xe3, 0x75, 0x25, 0x37, 0xb5, 0xcb, 0x5e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x89, 0xed, 0xa5, 0x26, 0x3b, 0x01, 0x00, 0x00,
}
