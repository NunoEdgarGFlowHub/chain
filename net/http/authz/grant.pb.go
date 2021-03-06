// Code generated by protoc-gen-go.
// source: grant.proto
// DO NOT EDIT!

/*
Package authz is a generated protocol buffer package.

It is generated from these files:
	grant.proto

It has these top-level messages:
	GrantList
	Grant
*/
package authz

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

type GrantList struct {
	Grants []*Grant `protobuf:"bytes,1,rep,name=grants" json:"grants,omitempty"`
}

func (m *GrantList) Reset()                    { *m = GrantList{} }
func (m *GrantList) String() string            { return proto.CompactTextString(m) }
func (*GrantList) ProtoMessage()               {}
func (*GrantList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GrantList) GetGrants() []*Grant {
	if m != nil {
		return m.Grants
	}
	return nil
}

type Grant struct {
	GuardType string `protobuf:"bytes,1,opt,name=guard_type,json=guardType" json:"guard_type,omitempty"`
	GuardData []byte `protobuf:"bytes,2,opt,name=guard_data,json=guardData,proto3" json:"guard_data,omitempty"`
	Policy    string `protobuf:"bytes,3,opt,name=policy" json:"policy,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	Protected bool   `protobuf:"varint,5,opt,name=protected" json:"protected,omitempty"`
}

func (m *Grant) Reset()                    { *m = Grant{} }
func (m *Grant) String() string            { return proto.CompactTextString(m) }
func (*Grant) ProtoMessage()               {}
func (*Grant) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Grant) GetGuardType() string {
	if m != nil {
		return m.GuardType
	}
	return ""
}

func (m *Grant) GetGuardData() []byte {
	if m != nil {
		return m.GuardData
	}
	return nil
}

func (m *Grant) GetPolicy() string {
	if m != nil {
		return m.Policy
	}
	return ""
}

func (m *Grant) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Grant) GetProtected() bool {
	if m != nil {
		return m.Protected
	}
	return false
}

func init() {
	proto.RegisterType((*GrantList)(nil), "authz.GrantList")
	proto.RegisterType((*Grant)(nil), "authz.Grant")
}

func init() { proto.RegisterFile("grant.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2f, 0x4a, 0xcc,
	0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x2c, 0x2d, 0xc9, 0xa8, 0x52, 0x32,
	0xe4, 0xe2, 0x74, 0x07, 0x89, 0xfa, 0x64, 0x16, 0x97, 0x08, 0xa9, 0x70, 0xb1, 0x81, 0x95, 0x14,
	0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xf1, 0xe8, 0x81, 0x15, 0xe9, 0x81, 0x55, 0x04, 0x41,
	0xe5, 0x94, 0x66, 0x31, 0x72, 0xb1, 0x82, 0x45, 0x84, 0x64, 0xb9, 0xb8, 0xd2, 0x4b, 0x13, 0x8b,
	0x52, 0xe2, 0x4b, 0x2a, 0x0b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38, 0xc1, 0x22,
	0x21, 0x95, 0x05, 0xa9, 0x08, 0xe9, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x1e,
	0xa8, 0xb4, 0x4b, 0x62, 0x49, 0xa2, 0x90, 0x18, 0x17, 0x5b, 0x41, 0x7e, 0x4e, 0x66, 0x72, 0xa5,
	0x04, 0x33, 0x58, 0x27, 0x94, 0x07, 0xd2, 0x96, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x9a, 0x12, 0x9f,
	0x58, 0x22, 0xc1, 0x02, 0x31, 0x15, 0x2a, 0xe2, 0x58, 0x22, 0x24, 0xc3, 0xc5, 0x09, 0xf2, 0x41,
	0x6a, 0x72, 0x49, 0x6a, 0x8a, 0x04, 0xab, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x42, 0x20, 0x89, 0x0d,
	0xec, 0x3b, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x08, 0xbf, 0xea, 0xec, 0x00, 0x00,
	0x00,
}
