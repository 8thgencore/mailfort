// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mail/mail.proto

package mailv1

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

type SendRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	OtpCode              string   `protobuf:"bytes,2,opt,name=otp_code,json=otpCode,proto3" json:"otp_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ac4131798133d9, []int{0}
}

func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendRequest.Unmarshal(m, b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return xxx_messageInfo_SendRequest.Size(m)
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SendRequest) GetOtpCode() string {
	if m != nil {
		return m.OtpCode
	}
	return ""
}

type SendResponse struct {
	IsSuccess            bool     `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ac4131798133d9, []int{1}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "mail.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "mail.SendResponse")
}

func init() {
	proto.RegisterFile("mail/mail.proto", fileDescriptor_d0ac4131798133d9)
}

var fileDescriptor_d0ac4131798133d9 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0xcc, 0xcc,
	0xd1, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x1d, 0x17,
	0x77, 0x70, 0x6a, 0x5e, 0x4a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08, 0x17, 0x6b,
	0x2a, 0x48, 0x5c, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x11, 0x92, 0xe4, 0xe2, 0xc8,
	0x2f, 0x29, 0x88, 0x4f, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x02, 0x4b, 0xb0, 0xe7, 0x97, 0x14, 0x38,
	0xe7, 0xa7, 0xa4, 0x2a, 0xe9, 0x72, 0xf1, 0x40, 0xf4, 0x17, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a,
	0xc9, 0x72, 0x71, 0x65, 0x16, 0xc7, 0x17, 0x97, 0x26, 0x27, 0xa7, 0x16, 0x17, 0x83, 0x4d, 0xe1,
	0x08, 0xe2, 0xcc, 0x2c, 0x0e, 0x86, 0x08, 0x18, 0x35, 0x31, 0x72, 0xb1, 0xf8, 0x82, 0x8c, 0xb4,
	0xe3, 0x12, 0x05, 0xe9, 0x73, 0xce, 0xcf, 0x4b, 0xcb, 0x2c, 0xca, 0x4d, 0x2c, 0xc9, 0xcc, 0xcf,
	0x73, 0x05, 0xdb, 0x25, 0xa8, 0x07, 0x76, 0x23, 0x92, 0xa3, 0xa4, 0x84, 0x90, 0x85, 0xa0, 0xf6,
	0x58, 0x71, 0x09, 0x82, 0xf8, 0x01, 0x89, 0xc5, 0xc5, 0xe5, 0xf9, 0x45, 0x20, 0xf1, 0xd4, 0x12,
	0x22, 0xf5, 0x3a, 0x09, 0x44, 0xf1, 0x81, 0x05, 0xcb, 0x0c, 0xad, 0x41, 0x74, 0x99, 0x61, 0x12,
	0x1b, 0x38, 0x48, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x0c, 0x2f, 0xb7, 0x25, 0x01,
	0x00, 0x00,
}
