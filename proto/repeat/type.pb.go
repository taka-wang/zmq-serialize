// Code generated by protoc-gen-go.
// source: type.proto
// DO NOT EDIT!

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	type.proto

It has these top-level messages:
	CmdHeader
	MbTcpHeader
	MbWriteRequest
	MbTcpSingleWriteReq
	MbTcpMultipleWriteReq
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

type CmdHeader struct {
	Receiver string `protobuf:"bytes,1,opt,name=receiver" json:"receiver,omitempty"`
	Sender   string `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Version  string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Tid      uint64 `protobuf:"varint,4,opt,name=tid" json:"tid,omitempty"`
	Method   string `protobuf:"bytes,5,opt,name=method" json:"method,omitempty"`
}

func (m *CmdHeader) Reset()                    { *m = CmdHeader{} }
func (m *CmdHeader) String() string            { return proto.CompactTextString(m) }
func (*CmdHeader) ProtoMessage()               {}
func (*CmdHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MbTcpHeader struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port uint32 `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Id   uint32 `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
}

func (m *MbTcpHeader) Reset()                    { *m = MbTcpHeader{} }
func (m *MbTcpHeader) String() string            { return proto.CompactTextString(m) }
func (*MbTcpHeader) ProtoMessage()               {}
func (*MbTcpHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type MbWriteRequest struct {
	Code     uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Register uint32 `protobuf:"varint,2,opt,name=register" json:"register,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	Type     string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Alias    string `protobuf:"bytes,5,opt,name=alias" json:"alias,omitempty"`
}

func (m *MbWriteRequest) Reset()                    { *m = MbWriteRequest{} }
func (m *MbWriteRequest) String() string            { return proto.CompactTextString(m) }
func (*MbWriteRequest) ProtoMessage()               {}
func (*MbWriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type MbTcpSingleWriteReq struct {
	CmdHeader      *CmdHeader      `protobuf:"bytes,1,opt,name=cmd_header,json=cmdHeader" json:"cmd_header,omitempty"`
	MbTcpHeader    *MbTcpHeader    `protobuf:"bytes,2,opt,name=mb_tcp_header,json=mbTcpHeader" json:"mb_tcp_header,omitempty"`
	MbWriteRequest *MbWriteRequest `protobuf:"bytes,3,opt,name=mb_write_request,json=mbWriteRequest" json:"mb_write_request,omitempty"`
}

func (m *MbTcpSingleWriteReq) Reset()                    { *m = MbTcpSingleWriteReq{} }
func (m *MbTcpSingleWriteReq) String() string            { return proto.CompactTextString(m) }
func (*MbTcpSingleWriteReq) ProtoMessage()               {}
func (*MbTcpSingleWriteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *MbTcpSingleWriteReq) GetCmdHeader() *CmdHeader {
	if m != nil {
		return m.CmdHeader
	}
	return nil
}

func (m *MbTcpSingleWriteReq) GetMbTcpHeader() *MbTcpHeader {
	if m != nil {
		return m.MbTcpHeader
	}
	return nil
}

func (m *MbTcpSingleWriteReq) GetMbWriteRequest() *MbWriteRequest {
	if m != nil {
		return m.MbWriteRequest
	}
	return nil
}

type MbTcpMultipleWriteReq struct {
	CmdHeader   *CmdHeader        `protobuf:"bytes,1,opt,name=cmd_header,json=cmdHeader" json:"cmd_header,omitempty"`
	MbTcpHeader *MbTcpHeader      `protobuf:"bytes,2,opt,name=mb_tcp_header,json=mbTcpHeader" json:"mb_tcp_header,omitempty"`
	Requests    []*MbWriteRequest `protobuf:"bytes,3,rep,name=requests" json:"requests,omitempty"`
}

func (m *MbTcpMultipleWriteReq) Reset()                    { *m = MbTcpMultipleWriteReq{} }
func (m *MbTcpMultipleWriteReq) String() string            { return proto.CompactTextString(m) }
func (*MbTcpMultipleWriteReq) ProtoMessage()               {}
func (*MbTcpMultipleWriteReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *MbTcpMultipleWriteReq) GetCmdHeader() *CmdHeader {
	if m != nil {
		return m.CmdHeader
	}
	return nil
}

func (m *MbTcpMultipleWriteReq) GetMbTcpHeader() *MbTcpHeader {
	if m != nil {
		return m.MbTcpHeader
	}
	return nil
}

func (m *MbTcpMultipleWriteReq) GetRequests() []*MbWriteRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

func init() {
	proto.RegisterType((*CmdHeader)(nil), "main.CmdHeader")
	proto.RegisterType((*MbTcpHeader)(nil), "main.MbTcpHeader")
	proto.RegisterType((*MbWriteRequest)(nil), "main.MbWriteRequest")
	proto.RegisterType((*MbTcpSingleWriteReq)(nil), "main.MbTcpSingleWriteReq")
	proto.RegisterType((*MbTcpMultipleWriteReq)(nil), "main.MbTcpMultipleWriteReq")
}

var fileDescriptor0 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x52, 0xbd, 0x4e, 0xf3, 0x40,
	0x10, 0x94, 0x13, 0x27, 0xdf, 0xe7, 0xb5, 0x1c, 0xc2, 0x11, 0x90, 0x45, 0x85, 0x5c, 0x51, 0x59,
	0x28, 0x88, 0x16, 0x09, 0xd1, 0xd0, 0xa4, 0x39, 0x90, 0x28, 0x23, 0xdb, 0xb7, 0x4a, 0x4e, 0xf2,
	0x1f, 0xf6, 0x25, 0x88, 0x0e, 0x89, 0x17, 0xe2, 0x01, 0x78, 0x38, 0xee, 0xf6, 0x6c, 0x2b, 0x14,
	0xb4, 0x74, 0x3b, 0xde, 0x9d, 0xdd, 0x99, 0x39, 0x03, 0xa8, 0xb7, 0x1a, 0xe3, 0xba, 0xa9, 0x54,
	0xc5, 0xdc, 0x22, 0x91, 0x65, 0xf4, 0xe1, 0x80, 0x77, 0x5f, 0x88, 0x07, 0x4c, 0x04, 0x36, 0xec,
	0x1c, 0xfe, 0x37, 0x98, 0xa1, 0xdc, 0x63, 0x13, 0x3a, 0x17, 0xce, 0xa5, 0xc7, 0x07, 0xcc, 0xce,
	0x60, 0xda, 0x62, 0xa9, 0xa7, 0xc2, 0x11, 0x75, 0x3a, 0xc4, 0x42, 0xf8, 0xa7, 0xdb, 0xad, 0xac,
	0xca, 0x70, 0x4c, 0x8d, 0x1e, 0xb2, 0x39, 0x8c, 0x95, 0x14, 0xa1, 0xab, 0xbf, 0xba, 0xdc, 0x94,
	0x66, 0x47, 0x81, 0x6a, 0x5b, 0x89, 0x70, 0x62, 0x77, 0x58, 0x14, 0xdd, 0x81, 0xbf, 0x4a, 0x9f,
	0xb2, 0xba, 0x93, 0x31, 0x83, 0x91, 0xac, 0x3b, 0x01, 0xba, 0x62, 0x0c, 0xdc, 0xba, 0x6a, 0x14,
	0x1d, 0x0e, 0x38, 0xd5, 0x34, 0x23, 0xe8, 0x62, 0xa0, 0x67, 0x44, 0xf4, 0xee, 0xc0, 0x6c, 0x95,
	0x3e, 0x37, 0x52, 0x21, 0xc7, 0x97, 0x1d, 0xb6, 0xca, 0xd0, 0xb2, 0x4a, 0x20, 0x2d, 0xd2, 0x34,
	0x53, 0x5b, 0x87, 0x1b, 0xd9, 0xaa, 0xce, 0x47, 0xc0, 0x07, 0xcc, 0x16, 0x30, 0xd9, 0x27, 0xf9,
	0x0e, 0x3b, 0x1f, 0x16, 0x98, 0x2d, 0x26, 0x35, 0xb2, 0xe1, 0x71, 0xaa, 0xcd, 0x64, 0x92, 0xcb,
	0xa4, 0xed, 0x6c, 0x58, 0x10, 0x7d, 0x39, 0x70, 0x42, 0x36, 0x1e, 0x65, 0xb9, 0xc9, 0xb1, 0xd7,
	0xc2, 0x62, 0x80, 0xac, 0x10, 0xeb, 0x2d, 0x99, 0x23, 0x35, 0xfe, 0xf2, 0x28, 0x36, 0xf1, 0xc7,
	0x43, 0xf4, 0xdc, 0xcb, 0x86, 0x57, 0xb8, 0x81, 0xa0, 0x48, 0xd7, 0x2a, 0xab, 0x7b, 0xca, 0x88,
	0x28, 0xc7, 0x96, 0x72, 0x10, 0x14, 0xf7, 0x8b, 0x83, 0xd4, 0x6e, 0x61, 0xae, 0x69, 0xaf, 0xe6,
	0xea, 0xba, 0xb1, 0x11, 0x90, 0x13, 0x7f, 0xb9, 0xe8, 0x99, 0x87, 0xf1, 0xf0, 0x59, 0xf1, 0x03,
	0x47, 0x9f, 0x0e, 0x9c, 0xd2, 0xf2, 0xd5, 0x2e, 0x57, 0xb2, 0xfe, 0x7b, 0x03, 0x57, 0xe6, 0x6d,
	0x48, 0x4b, 0xab, 0x85, 0x8f, 0x7f, 0x15, 0x3e, 0x4c, 0xa5, 0x53, 0xfa, 0x95, 0xaf, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x7e, 0x27, 0x50, 0x48, 0xd8, 0x02, 0x00, 0x00,
}