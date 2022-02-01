// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_gcp_pubsub.proto

package args

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

type GCPPubSubConn struct {
	// @gotags: kong:"help='Project ID',env='PLUMBER_RELAY_GCP_PROJECT_ID',required"
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" kong:"help='Project ID',env='PLUMBER_RELAY_GCP_PROJECT_ID',required"`
	// @gotags: kong:"help='GCP Credentials in JSON format',env='PLUMBER_RELAY_GCP_CREDENTIALS'"
	CredentialsJson string `protobuf:"bytes,2,opt,name=credentials_json,json=credentialsJson,proto3" json:"credentials_json,omitempty" kong:"help='GCP Credentials in JSON format',env='PLUMBER_RELAY_GCP_CREDENTIALS'"`
	// @gotags: kong:"help='Path to GCP credentials JSON file',env='GOOGLE_APPLICATION_CREDENTIALS'"
	CredentialsFile      string   `protobuf:"bytes,3,opt,name=credentials_file,json=credentialsFile,proto3" json:"credentials_file,omitempty" kong:"help='Path to GCP credentials JSON file',env='GOOGLE_APPLICATION_CREDENTIALS'"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCPPubSubConn) Reset()         { *m = GCPPubSubConn{} }
func (m *GCPPubSubConn) String() string { return proto.CompactTextString(m) }
func (*GCPPubSubConn) ProtoMessage()    {}
func (*GCPPubSubConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_e51274dec22f496a, []int{0}
}

func (m *GCPPubSubConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCPPubSubConn.Unmarshal(m, b)
}
func (m *GCPPubSubConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCPPubSubConn.Marshal(b, m, deterministic)
}
func (m *GCPPubSubConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCPPubSubConn.Merge(m, src)
}
func (m *GCPPubSubConn) XXX_Size() int {
	return xxx_messageInfo_GCPPubSubConn.Size(m)
}
func (m *GCPPubSubConn) XXX_DiscardUnknown() {
	xxx_messageInfo_GCPPubSubConn.DiscardUnknown(m)
}

var xxx_messageInfo_GCPPubSubConn proto.InternalMessageInfo

func (m *GCPPubSubConn) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *GCPPubSubConn) GetCredentialsJson() string {
	if m != nil {
		return m.CredentialsJson
	}
	return ""
}

func (m *GCPPubSubConn) GetCredentialsFile() string {
	if m != nil {
		return m.CredentialsFile
	}
	return ""
}

type GCPPubSubReadArgs struct {
	// @gotags: kong:"help='Subscription ID',env='PLUMBER_RELAY_GCP_SUBSCRIPTION_ID',required"
	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty" kong:"help='Subscription ID',env='PLUMBER_RELAY_GCP_SUBSCRIPTION_ID',required"`
	// @gotags: kong:"help='Whether to acknowledge message receive',env='PLUMBER_RELAY_GCP_ACK_MESSAGE',default=true"
	AckMessages          bool     `protobuf:"varint,2,opt,name=ack_messages,json=ackMessages,proto3" json:"ack_messages,omitempty" kong:"help='Whether to acknowledge message receive',env='PLUMBER_RELAY_GCP_ACK_MESSAGE',default=true"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCPPubSubReadArgs) Reset()         { *m = GCPPubSubReadArgs{} }
func (m *GCPPubSubReadArgs) String() string { return proto.CompactTextString(m) }
func (*GCPPubSubReadArgs) ProtoMessage()    {}
func (*GCPPubSubReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_e51274dec22f496a, []int{1}
}

func (m *GCPPubSubReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCPPubSubReadArgs.Unmarshal(m, b)
}
func (m *GCPPubSubReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCPPubSubReadArgs.Marshal(b, m, deterministic)
}
func (m *GCPPubSubReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCPPubSubReadArgs.Merge(m, src)
}
func (m *GCPPubSubReadArgs) XXX_Size() int {
	return xxx_messageInfo_GCPPubSubReadArgs.Size(m)
}
func (m *GCPPubSubReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GCPPubSubReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GCPPubSubReadArgs proto.InternalMessageInfo

func (m *GCPPubSubReadArgs) GetSubscriptionId() string {
	if m != nil {
		return m.SubscriptionId
	}
	return ""
}

func (m *GCPPubSubReadArgs) GetAckMessages() bool {
	if m != nil {
		return m.AckMessages
	}
	return false
}

type GCPPubSubWriteArgs struct {
	// @gotags: kong:"help='Topic ID to publish message(s) to',required"
	TopicId              string   `protobuf:"bytes,1,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty" kong:"help='Topic ID to publish message(s) to',required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GCPPubSubWriteArgs) Reset()         { *m = GCPPubSubWriteArgs{} }
func (m *GCPPubSubWriteArgs) String() string { return proto.CompactTextString(m) }
func (*GCPPubSubWriteArgs) ProtoMessage()    {}
func (*GCPPubSubWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_e51274dec22f496a, []int{2}
}

func (m *GCPPubSubWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCPPubSubWriteArgs.Unmarshal(m, b)
}
func (m *GCPPubSubWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCPPubSubWriteArgs.Marshal(b, m, deterministic)
}
func (m *GCPPubSubWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCPPubSubWriteArgs.Merge(m, src)
}
func (m *GCPPubSubWriteArgs) XXX_Size() int {
	return xxx_messageInfo_GCPPubSubWriteArgs.Size(m)
}
func (m *GCPPubSubWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_GCPPubSubWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_GCPPubSubWriteArgs proto.InternalMessageInfo

func (m *GCPPubSubWriteArgs) GetTopicId() string {
	if m != nil {
		return m.TopicId
	}
	return ""
}

func init() {
	proto.RegisterType((*GCPPubSubConn)(nil), "protos.args.GCPPubSubConn")
	proto.RegisterType((*GCPPubSubReadArgs)(nil), "protos.args.GCPPubSubReadArgs")
	proto.RegisterType((*GCPPubSubWriteArgs)(nil), "protos.args.GCPPubSubWriteArgs")
}

func init() { proto.RegisterFile("ps_args_gcp_pubsub.proto", fileDescriptor_e51274dec22f496a) }

var fileDescriptor_e51274dec22f496a = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0x33, 0x31,
	0x18, 0xc4, 0xd9, 0xf7, 0x05, 0x6d, 0x53, 0xb5, 0x9a, 0xd3, 0x7a, 0x10, 0x74, 0x2f, 0xea, 0xc1,
	0xcd, 0xc1, 0x93, 0x78, 0xd2, 0x82, 0xb2, 0x82, 0x50, 0xd6, 0x83, 0xe0, 0x25, 0xe4, 0x9f, 0xd9,
	0xb4, 0xbb, 0x49, 0xc8, 0x93, 0x7c, 0x03, 0x3f, 0xb8, 0x34, 0x96, 0xb6, 0xe8, 0x29, 0xcc, 0x2f,
	0xc3, 0x33, 0xc3, 0xa0, 0xd2, 0x03, 0x65, 0x41, 0x03, 0xd5, 0xc2, 0x53, 0x9f, 0x38, 0x24, 0x5e,
	0xfb, 0xe0, 0xa2, 0xc3, 0x93, 0xfc, 0x40, 0xbd, 0xfa, 0xad, 0xbe, 0x0a, 0x74, 0xf8, 0x3c, 0x9b,
	0xcf, 0x13, 0x7f, 0x4b, 0x7c, 0xe6, 0xac, 0xc5, 0x67, 0x08, 0xf9, 0xe0, 0x16, 0x4a, 0x44, 0x6a,
	0x64, 0x59, 0x9c, 0x17, 0x57, 0xe3, 0x76, 0xbc, 0x26, 0x8d, 0xc4, 0xd7, 0xe8, 0x58, 0x04, 0x25,
	0x95, 0x8d, 0x86, 0xf5, 0x40, 0x17, 0xe0, 0x6c, 0xf9, 0x2f, 0x9b, 0xa6, 0x3b, 0xfc, 0x05, 0x9c,
	0xfd, 0x6d, 0xfd, 0x34, 0xbd, 0x2a, 0xff, 0xff, 0xb1, 0x3e, 0x99, 0x5e, 0x55, 0x14, 0x9d, 0x6c,
	0x5a, 0xb4, 0x8a, 0xc9, 0x87, 0xa0, 0x01, 0x5f, 0xa2, 0x29, 0x24, 0x0e, 0x22, 0x18, 0x1f, 0x8d,
	0xb3, 0xdb, 0x3a, 0x47, 0xbb, 0xb8, 0x91, 0xf8, 0x02, 0x1d, 0x30, 0xb1, 0xa4, 0x83, 0x02, 0x60,
	0x5a, 0x41, 0xee, 0x33, 0x6a, 0x27, 0x4c, 0x2c, 0x5f, 0xd7, 0xa8, 0x22, 0x08, 0x6f, 0x02, 0xde,
	0x83, 0x89, 0x2a, 0x27, 0x9c, 0xa2, 0x51, 0x74, 0xde, 0x88, 0xed, 0xe9, 0xfd, 0xac, 0x1b, 0xf9,
	0x78, 0xff, 0x71, 0xa7, 0x4d, 0xec, 0x12, 0xaf, 0x85, 0x1b, 0x08, 0x67, 0x51, 0x74, 0xc2, 0x05,
	0x4f, 0x7c, 0x9f, 0x06, 0xae, 0xc2, 0x0d, 0x88, 0x4e, 0x0d, 0x0c, 0x08, 0x4f, 0xa6, 0x97, 0x44,
	0x3b, 0xf2, 0xb3, 0x2a, 0x59, 0xad, 0xca, 0xf7, 0xb2, 0xb8, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xac, 0xe8, 0x15, 0xc6, 0x85, 0x01, 0x00, 0x00,
}