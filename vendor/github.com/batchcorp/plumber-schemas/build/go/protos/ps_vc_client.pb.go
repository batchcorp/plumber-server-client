// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_vc_client.proto

package protos

import (
	fmt "fmt"
	common "github.com/batchcorp/plumber-schemas/build/go/protos/common"
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

// GetVCEventsRequest is used to call GetVCEvents() which returns a stream of events from ps_ghserver.proto's server
type GetVCEventsRequest struct {
	// Every gRPC request must have a valid auth config
	Auth                 *common.Auth `protobuf:"bytes,9999,opt,name=auth,proto3" json:"auth,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetVCEventsRequest) Reset()         { *m = GetVCEventsRequest{} }
func (m *GetVCEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetVCEventsRequest) ProtoMessage()    {}
func (*GetVCEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6afb08e90b4c9a32, []int{0}
}

func (m *GetVCEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetVCEventsRequest.Unmarshal(m, b)
}
func (m *GetVCEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetVCEventsRequest.Marshal(b, m, deterministic)
}
func (m *GetVCEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetVCEventsRequest.Merge(m, src)
}
func (m *GetVCEventsRequest) XXX_Size() int {
	return xxx_messageInfo_GetVCEventsRequest.Size(m)
}
func (m *GetVCEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetVCEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetVCEventsRequest proto.InternalMessageInfo

func (m *GetVCEventsRequest) GetAuth() *common.Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func init() {
	proto.RegisterType((*GetVCEventsRequest)(nil), "protos.GetVCEventsRequest")
}

func init() { proto.RegisterFile("ps_vc_client.proto", fileDescriptor_6afb08e90b4c9a32) }

var fileDescriptor_6afb08e90b4c9a32 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0x8e, 0x2f,
	0x4b, 0x8e, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x03, 0x53, 0xc5, 0x52, 0xd2, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0xfa, 0x05, 0xc5, 0xf1, 0x10,
	0x56, 0x7c, 0x62, 0x69, 0x49, 0x06, 0x44, 0x91, 0x92, 0x1d, 0x97, 0x90, 0x7b, 0x6a, 0x49, 0x98,
	0xb3, 0x6b, 0x59, 0x6a, 0x5e, 0x49, 0x71, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x06,
	0x17, 0x0b, 0x48, 0x8d, 0x44, 0xbf, 0x9f, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0x30, 0x44, 0x71, 0xb1,
	0x1e, 0x44, 0xbf, 0x9e, 0x63, 0x69, 0x49, 0x46, 0x10, 0x58, 0x85, 0x93, 0x59, 0x94, 0x49, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x12, 0x48, 0x4e, 0x3f, 0x29, 0xb1, 0x24, 0x39, 0x23, 0x39, 0xbf, 0xa8,
	0x40, 0xbf, 0x20, 0xa7, 0x34, 0x37, 0x29, 0xb5, 0x48, 0xb7, 0x38, 0x39, 0x23, 0x35, 0x37, 0xb1,
	0x58, 0x3f, 0xa9, 0x34, 0x33, 0x27, 0x45, 0x3f, 0x3d, 0x5f, 0x1f, 0x62, 0x52, 0x12, 0xc4, 0x71,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x43, 0x7d, 0x44, 0xb9, 0x00, 0x00, 0x00,
}
