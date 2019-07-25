// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package proto

import (
	fmt "fmt"
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

type TestInfo struct {
	Authorization        string   `protobuf:"bytes,1,opt,name=Authorization,proto3" json:"Authorization,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestInfo) Reset()         { *m = TestInfo{} }
func (m *TestInfo) String() string { return proto.CompactTextString(m) }
func (*TestInfo) ProtoMessage()    {}
func (*TestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestInfo.Unmarshal(m, b)
}
func (m *TestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestInfo.Marshal(b, m, deterministic)
}
func (m *TestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestInfo.Merge(m, src)
}
func (m *TestInfo) XXX_Size() int {
	return xxx_messageInfo_TestInfo.Size(m)
}
func (m *TestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TestInfo proto.InternalMessageInfo

func (m *TestInfo) GetAuthorization() string {
	if m != nil {
		return m.Authorization
	}
	return ""
}

func init() {
	proto.RegisterType((*TestInfo)(nil), "proto.TestInfo")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 83 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x06, 0x5c, 0x1c, 0x21, 0xa9,
	0xc5, 0x25, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x2a, 0x5c, 0xbc, 0x8e, 0xa5, 0x25, 0x19, 0xf9, 0x45,
	0x99, 0x55, 0x89, 0x25, 0x99, 0xf9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xa8, 0x82,
	0x49, 0x6c, 0x60, 0x8d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0xab, 0x62, 0x27, 0x4d,
	0x00, 0x00, 0x00,
}
