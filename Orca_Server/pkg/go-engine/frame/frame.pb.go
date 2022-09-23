// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frame.proto

package frame

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

type FrameData_TYPE int32

const (
	FrameData_USER_DATA FrameData_TYPE = 0
	FrameData_CONN      FrameData_TYPE = 1
	FrameData_CONNRSP   FrameData_TYPE = 2
	FrameData_CLOSE     FrameData_TYPE = 3
	FrameData_HB        FrameData_TYPE = 4
)

var FrameData_TYPE_name = map[int32]string{
	0: "USER_DATA",
	1: "CONN",
	2: "CONNRSP",
	3: "CLOSE",
	4: "HB",
}

var FrameData_TYPE_value = map[string]int32{
	"USER_DATA": 0,
	"CONN":      1,
	"CONNRSP":   2,
	"CLOSE":     3,
	"HB":        4,
}

func (x FrameData_TYPE) String() string {
	return proto.EnumName(FrameData_TYPE_name, int32(x))
}

func (FrameData_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{0, 0}
}

type Frame_TYPE int32

const (
	Frame_DATA Frame_TYPE = 0
	Frame_REQ  Frame_TYPE = 1
	Frame_ACK  Frame_TYPE = 2
	Frame_PING Frame_TYPE = 3
	Frame_PONG Frame_TYPE = 4
)

var Frame_TYPE_name = map[int32]string{
	0: "DATA",
	1: "REQ",
	2: "ACK",
	3: "PING",
	4: "PONG",
}

var Frame_TYPE_value = map[string]int32{
	"DATA": 0,
	"REQ":  1,
	"ACK":  2,
	"PING": 3,
	"PONG": 4,
}

func (x Frame_TYPE) String() string {
	return proto.EnumName(Frame_TYPE_name, int32(x))
}

func (Frame_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{1, 0}
}

type FrameData struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Compress             bool     `protobuf:"varint,3,opt,name=compress,proto3" json:"compress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FrameData) Reset()         { *m = FrameData{} }
func (m *FrameData) String() string { return proto.CompactTextString(m) }
func (*FrameData) ProtoMessage()    {}
func (*FrameData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{0}
}

func (m *FrameData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrameData.Unmarshal(m, b)
}
func (m *FrameData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrameData.Marshal(b, m, deterministic)
}
func (m *FrameData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrameData.Merge(m, src)
}
func (m *FrameData) XXX_Size() int {
	return xxx_messageInfo_FrameData.Size(m)
}
func (m *FrameData) XXX_DiscardUnknown() {
	xxx_messageInfo_FrameData.DiscardUnknown(m)
}

var xxx_messageInfo_FrameData proto.InternalMessageInfo

func (m *FrameData) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *FrameData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *FrameData) GetCompress() bool {
	if m != nil {
		return m.Compress
	}
	return false
}

type Frame struct {
	Type                 int32      `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Resend               bool       `protobuf:"varint,2,opt,name=resend,proto3" json:"resend,omitempty"`
	Sendtime             int64      `protobuf:"varint,3,opt,name=sendtime,proto3" json:"sendtime,omitempty"`
	Id                   int32      `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	Data                 *FrameData `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Dataid               []int32    `protobuf:"varint,6,rep,packed,name=dataid,proto3" json:"dataid,omitempty"`
	Acked                bool       `protobuf:"varint,7,opt,name=acked,proto3" json:"acked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{1}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Frame) GetResend() bool {
	if m != nil {
		return m.Resend
	}
	return false
}

func (m *Frame) GetSendtime() int64 {
	if m != nil {
		return m.Sendtime
	}
	return 0
}

func (m *Frame) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Frame) GetData() *FrameData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Frame) GetDataid() []int32 {
	if m != nil {
		return m.Dataid
	}
	return nil
}

func (m *Frame) GetAcked() bool {
	if m != nil {
		return m.Acked
	}
	return false
}

func init() {
	proto.RegisterEnum("FrameData_TYPE", FrameData_TYPE_name, FrameData_TYPE_value)
	proto.RegisterEnum("Frame_TYPE", Frame_TYPE_name, Frame_TYPE_value)
	proto.RegisterType((*FrameData)(nil), "FrameData")
	proto.RegisterType((*Frame)(nil), "Frame")
}

func init() { proto.RegisterFile("frame.proto", fileDescriptor_5379e2b825e15002) }

var fileDescriptor_5379e2b825e15002 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4a, 0xf4, 0x30,
	0x14, 0x86, 0xbf, 0xfc, 0xb5, 0x9d, 0x33, 0x9f, 0x12, 0x82, 0x0c, 0xc1, 0x85, 0x94, 0xae, 0xba,
	0x9a, 0x85, 0x82, 0x5b, 0x99, 0x9f, 0x3a, 0x8a, 0xd2, 0x19, 0x33, 0xe3, 0x42, 0x37, 0x12, 0x27,
	0x11, 0x8a, 0x8c, 0x2d, 0x6d, 0x37, 0xde, 0x85, 0x37, 0xe9, 0x7d, 0x48, 0x62, 0xe8, 0xca, 0xd5,
	0x79, 0x9f, 0x16, 0x92, 0xe7, 0xcd, 0x81, 0xf1, 0x5b, 0xab, 0x0f, 0x76, 0xda, 0xb4, 0x75, 0x5f,
	0x67, 0x5f, 0x08, 0x46, 0xd7, 0x8e, 0x97, 0xba, 0xd7, 0x42, 0x00, 0xed, 0x3f, 0x1b, 0x2b, 0x51,
	0x8a, 0x72, 0xa6, 0x7c, 0x76, 0xdf, 0x8c, 0xee, 0xb5, 0xc4, 0x29, 0xca, 0xff, 0x2b, 0x9f, 0xc5,
	0x29, 0x24, 0xfb, 0xfa, 0xd0, 0xb4, 0xb6, 0xeb, 0x24, 0x49, 0x51, 0x9e, 0xa8, 0x81, 0xb3, 0x2b,
	0xa0, 0xbb, 0xa7, 0x4d, 0x21, 0x8e, 0x60, 0xf4, 0xb8, 0x2d, 0xd4, 0xcb, 0x72, 0xb6, 0x9b, 0xf1,
	0x7f, 0x22, 0x01, 0xba, 0x58, 0x97, 0x25, 0x47, 0x62, 0x0c, 0xb1, 0x4b, 0x6a, 0xbb, 0xe1, 0x58,
	0x8c, 0x80, 0x2d, 0xee, 0xd7, 0xdb, 0x82, 0x13, 0x11, 0x01, 0xbe, 0x99, 0x73, 0x9a, 0x7d, 0x23,
	0x60, 0x5e, 0xe9, 0x4f, 0x9d, 0x09, 0x44, 0xad, 0xed, 0xec, 0x87, 0xf1, 0x42, 0x89, 0x0a, 0xe4,
	0x94, 0xdc, 0xec, 0xab, 0x83, 0xf5, 0x4a, 0x44, 0x0d, 0x2c, 0x8e, 0x01, 0x57, 0x46, 0x52, 0x7f,
	0x0a, 0xae, 0x8c, 0x38, 0x0b, 0x95, 0x58, 0x8a, 0xf2, 0xf1, 0x39, 0x4c, 0x87, 0x07, 0x08, 0xf5,
	0x26, 0x10, 0xb9, 0x59, 0x19, 0x19, 0xa5, 0x24, 0x67, 0x2a, 0x90, 0x38, 0x01, 0xa6, 0xf7, 0xef,
	0xd6, 0xc8, 0xd8, 0x5f, 0xfd, 0x0b, 0xd9, 0x65, 0x28, 0x9c, 0x00, 0x0d, 0x5d, 0x63, 0x20, 0xaa,
	0x78, 0xe0, 0xc8, 0x85, 0xd9, 0xe2, 0x8e, 0x63, 0xf7, 0x6f, 0x73, 0x5b, 0xae, 0x38, 0xf1, 0x69,
	0x5d, 0xae, 0x38, 0x9d, 0xc7, 0xcf, 0xcc, 0x6f, 0xe2, 0x35, 0xf2, 0xab, 0xb8, 0xf8, 0x09, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0x07, 0xe4, 0x02, 0x99, 0x01, 0x00, 0x00,
}