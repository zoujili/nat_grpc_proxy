// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package chat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RestEvent struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestEvent) Reset()         { *m = RestEvent{} }
func (m *RestEvent) String() string { return proto.CompactTextString(m) }
func (*RestEvent) ProtoMessage()    {}
func (*RestEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *RestEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestEvent.Unmarshal(m, b)
}
func (m *RestEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestEvent.Marshal(b, m, deterministic)
}
func (m *RestEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestEvent.Merge(m, src)
}
func (m *RestEvent) XXX_Size() int {
	return xxx_messageInfo_RestEvent.Size(m)
}
func (m *RestEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RestEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RestEvent proto.InternalMessageInfo

func (m *RestEvent) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *RestEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PushEvent struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushEvent) Reset()         { *m = PushEvent{} }
func (m *PushEvent) String() string { return proto.CompactTextString(m) }
func (*PushEvent) ProtoMessage()    {}
func (*PushEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *PushEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushEvent.Unmarshal(m, b)
}
func (m *PushEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushEvent.Marshal(b, m, deterministic)
}
func (m *PushEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushEvent.Merge(m, src)
}
func (m *PushEvent) XXX_Size() int {
	return xxx_messageInfo_PushEvent.Size(m)
}
func (m *PushEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_PushEvent.DiscardUnknown(m)
}

var xxx_messageInfo_PushEvent proto.InternalMessageInfo

func (m *PushEvent) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *PushEvent) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RestEvent)(nil), "chat.RestEvent")
	proto.RegisterType((*PushEvent)(nil), "chat.PushEvent")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x1c, 0xb9, 0x38, 0x83, 0x52,
	0x8b, 0x4b, 0x5c, 0xcb, 0x52, 0xf3, 0x4a, 0x84, 0xa4, 0xb8, 0x38, 0x92, 0x73, 0x32, 0x53, 0xf3,
	0x4a, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x09, 0x2e, 0xf6,
	0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0xb0, 0x14, 0x8c, 0x0b, 0x32, 0x22, 0xa0,
	0xb4, 0x38, 0x83, 0x02, 0x23, 0x8c, 0x8c, 0xb9, 0x58, 0x9c, 0x33, 0x12, 0x4b, 0x84, 0xb4, 0xb9,
	0x98, 0x83, 0x13, 0x2b, 0x85, 0xf8, 0xf5, 0xc0, 0xee, 0x84, 0x3b, 0x4c, 0x0a, 0x2a, 0x00, 0xb7,
	0x46, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x0f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc2, 0xc7, 0x8b, 0x86, 0xd5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatClient interface {
	Say(ctx context.Context, opts ...grpc.CallOption) (Chat_SayClient, error)
}

type chatClient struct {
	cc *grpc.ClientConn
}

func NewChatClient(cc *grpc.ClientConn) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Say(ctx context.Context, opts ...grpc.CallOption) (Chat_SayClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/chat.Chat/Say", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatSayClient{stream}
	return x, nil
}

type Chat_SayClient interface {
	Send(*RestEvent) error
	Recv() (*PushEvent, error)
	grpc.ClientStream
}

type chatSayClient struct {
	grpc.ClientStream
}

func (x *chatSayClient) Send(m *RestEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatSayClient) Recv() (*PushEvent, error) {
	m := new(PushEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
type ChatServer interface {
	Say(Chat_SayServer) error
}

// UnimplementedChatServer can be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (*UnimplementedChatServer) Say(srv Chat_SayServer) error {
	return status.Errorf(codes.Unimplemented, "method Say not implemented")
}

func RegisterChatServer(s *grpc.Server, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_Say_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Say(&chatSayServer{stream})
}

type Chat_SayServer interface {
	Send(*PushEvent) error
	Recv() (*RestEvent, error)
	grpc.ServerStream
}

type chatSayServer struct {
	grpc.ServerStream
}

func (x *chatSayServer) Send(m *PushEvent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatSayServer) Recv() (*RestEvent, error) {
	m := new(RestEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Say",
			Handler:       _Chat_Say_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}