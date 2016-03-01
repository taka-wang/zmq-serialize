// Code generated by protoc-gen-go.
// source: type.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	type.proto

It has these top-level messages:
	MbTcpHeader
*/
package main

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

type MbTcpHeader struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Id   uint32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
}

func (m *MbTcpHeader) Reset()                    { *m = MbTcpHeader{} }
func (m *MbTcpHeader) String() string            { return proto.CompactTextString(m) }
func (*MbTcpHeader) ProtoMessage()               {}
func (*MbTcpHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*MbTcpHeader)(nil), "main.MbTcpHeader")
}

var fileDescriptor0 = []byte{
	// 101 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xa9, 0x2c, 0x48,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x72, 0xe4, 0xe2,
	0xf6, 0x4d, 0x0a, 0x49, 0x2e, 0xf0, 0x48, 0x4d, 0x4c, 0x49, 0x2d, 0x12, 0xe2, 0xe3, 0x62, 0xca,
	0x2c, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb2, 0x84, 0x84, 0xb8, 0x58, 0x0a, 0xf2,
	0x8b, 0x4a, 0x24, 0x98, 0x80, 0x22, 0xbc, 0x41, 0x60, 0x36, 0x58, 0x4d, 0x8a, 0x04, 0x33, 0x58,
	0x04, 0xc8, 0x4a, 0x62, 0x03, 0x9b, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xf7, 0xcf,
	0x94, 0x5d, 0x00, 0x00, 0x00,
}
