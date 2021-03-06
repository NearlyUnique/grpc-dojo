// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	SubType
	HelloRequest
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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SubType struct {
	Number  int32            `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Enabled bool             `protobuf:"varint,2,opt,name=enabled" json:"enabled,omitempty"`
	Lookup  map[string]int32 `protobuf:"bytes,3,rep,name=lookup" json:"lookup,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *SubType) Reset()                    { *m = SubType{} }
func (m *SubType) String() string            { return proto.CompactTextString(m) }
func (*SubType) ProtoMessage()               {}
func (*SubType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SubType) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *SubType) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *SubType) GetLookup() map[string]int32 {
	if m != nil {
		return m.Lookup
	}
	return nil
}

type HelloRequest struct {
	Name string     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	List []*SubType `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetList() []*SubType {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*SubType)(nil), "main.SubType")
	proto.RegisterType((*HelloRequest)(nil), "main.HelloRequest")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x4f, 0xbb, 0x4e, 0xc4, 0x30,
	0x10, 0x94, 0xf3, 0x3a, 0xd8, 0x3b, 0x24, 0xb4, 0x42, 0xc8, 0x50, 0x85, 0xab, 0x52, 0x59, 0x02,
	0x1a, 0xa0, 0x3f, 0x89, 0x82, 0xca, 0xf0, 0x03, 0xb6, 0xd8, 0xe2, 0x14, 0xc7, 0x0e, 0x8e, 0x1d,
	0x29, 0x9f, 0xc4, 0x5f, 0xa2, 0x18, 0x23, 0x71, 0xdd, 0xcc, 0xce, 0xee, 0xcc, 0x2c, 0xec, 0x26,
	0xf2, 0x33, 0x79, 0x31, 0x7a, 0x17, 0x1c, 0x56, 0x83, 0x3a, 0xda, 0xfd, 0x37, 0x83, 0xcd, 0x7b,
	0xd4, 0x1f, 0xcb, 0x48, 0x78, 0x0d, 0x8d, 0x8d, 0x83, 0x26, 0xcf, 0x59, 0xcb, 0xba, 0x5a, 0x66,
	0x86, 0x1c, 0x36, 0x64, 0x95, 0x36, 0xf4, 0xc9, 0x8b, 0x96, 0x75, 0x67, 0xf2, 0x8f, 0xe2, 0x3d,
	0x34, 0xc6, 0xb9, 0x3e, 0x8e, 0xbc, 0x6c, 0xcb, 0x6e, 0xfb, 0x70, 0x23, 0x56, 0x53, 0x91, 0x0d,
	0xc5, 0x5b, 0xd2, 0x0e, 0x36, 0xf8, 0x45, 0xe6, 0xc5, 0xdb, 0x67, 0xd8, 0xfe, 0x1b, 0xe3, 0x25,
	0x94, 0x3d, 0x2d, 0x29, 0xf0, 0x5c, 0xae, 0x10, 0xaf, 0xa0, 0x9e, 0x95, 0x89, 0x94, 0xb2, 0x6a,
	0xf9, 0x4b, 0x5e, 0x8a, 0x27, 0xb6, 0x3f, 0xc0, 0xee, 0x95, 0x8c, 0x71, 0x92, 0xbe, 0x22, 0x4d,
	0x01, 0x11, 0x2a, 0xab, 0x06, 0xca, 0xc7, 0x09, 0xe3, 0x1d, 0x54, 0xe6, 0x38, 0x05, 0x5e, 0xa4,
	0x3e, 0x17, 0x27, 0x7d, 0x64, 0x92, 0x74, 0x93, 0xfe, 0x7f, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff,
	0x32, 0xff, 0x4a, 0xc7, 0x0f, 0x01, 0x00, 0x00,
}
