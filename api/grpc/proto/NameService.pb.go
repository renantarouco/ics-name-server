// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NameService.proto

package proto

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

type HeartBeat struct {
	Header               *ConnectionHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *HeartBeat) Reset()         { *m = HeartBeat{} }
func (m *HeartBeat) String() string { return proto.CompactTextString(m) }
func (*HeartBeat) ProtoMessage()    {}
func (*HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_98ff7aff3b3ac208, []int{0}
}

func (m *HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartBeat.Unmarshal(m, b)
}
func (m *HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartBeat.Marshal(b, m, deterministic)
}
func (m *HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartBeat.Merge(m, src)
}
func (m *HeartBeat) XXX_Size() int {
	return xxx_messageInfo_HeartBeat.Size(m)
}
func (m *HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_HeartBeat proto.InternalMessageInfo

func (m *HeartBeat) GetHeader() *ConnectionHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ConnectionInfo struct {
	Header               *ConnectionHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ConnectionInfo) Reset()         { *m = ConnectionInfo{} }
func (m *ConnectionInfo) String() string { return proto.CompactTextString(m) }
func (*ConnectionInfo) ProtoMessage()    {}
func (*ConnectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_98ff7aff3b3ac208, []int{1}
}

func (m *ConnectionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionInfo.Unmarshal(m, b)
}
func (m *ConnectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionInfo.Marshal(b, m, deterministic)
}
func (m *ConnectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionInfo.Merge(m, src)
}
func (m *ConnectionInfo) XXX_Size() int {
	return xxx_messageInfo_ConnectionInfo.Size(m)
}
func (m *ConnectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionInfo proto.InternalMessageInfo

func (m *ConnectionInfo) GetHeader() *ConnectionHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ConnectionHeader struct {
	MessageServerId      string   `protobuf:"bytes,1,opt,name=message_server_id,json=messageServerId,proto3" json:"message_server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionHeader) Reset()         { *m = ConnectionHeader{} }
func (m *ConnectionHeader) String() string { return proto.CompactTextString(m) }
func (*ConnectionHeader) ProtoMessage()    {}
func (*ConnectionHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_98ff7aff3b3ac208, []int{2}
}

func (m *ConnectionHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionHeader.Unmarshal(m, b)
}
func (m *ConnectionHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionHeader.Marshal(b, m, deterministic)
}
func (m *ConnectionHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionHeader.Merge(m, src)
}
func (m *ConnectionHeader) XXX_Size() int {
	return xxx_messageInfo_ConnectionHeader.Size(m)
}
func (m *ConnectionHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionHeader.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionHeader proto.InternalMessageInfo

func (m *ConnectionHeader) GetMessageServerId() string {
	if m != nil {
		return m.MessageServerId
	}
	return ""
}

type CreateRoomRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomRequest) Reset()         { *m = CreateRoomRequest{} }
func (m *CreateRoomRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRoomRequest) ProtoMessage()    {}
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_98ff7aff3b3ac208, []int{3}
}

func (m *CreateRoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomRequest.Unmarshal(m, b)
}
func (m *CreateRoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomRequest.Marshal(b, m, deterministic)
}
func (m *CreateRoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomRequest.Merge(m, src)
}
func (m *CreateRoomRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRoomRequest.Size(m)
}
func (m *CreateRoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomRequest proto.InternalMessageInfo

func (m *CreateRoomRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateRoomResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRoomResponse) Reset()         { *m = CreateRoomResponse{} }
func (m *CreateRoomResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRoomResponse) ProtoMessage()    {}
func (*CreateRoomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_98ff7aff3b3ac208, []int{4}
}

func (m *CreateRoomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRoomResponse.Unmarshal(m, b)
}
func (m *CreateRoomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRoomResponse.Marshal(b, m, deterministic)
}
func (m *CreateRoomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRoomResponse.Merge(m, src)
}
func (m *CreateRoomResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRoomResponse.Size(m)
}
func (m *CreateRoomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRoomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRoomResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HeartBeat)(nil), "proto.HeartBeat")
	proto.RegisterType((*ConnectionInfo)(nil), "proto.ConnectionInfo")
	proto.RegisterType((*ConnectionHeader)(nil), "proto.ConnectionHeader")
	proto.RegisterType((*CreateRoomRequest)(nil), "proto.CreateRoomRequest")
	proto.RegisterType((*CreateRoomResponse)(nil), "proto.CreateRoomResponse")
}

func init() { proto.RegisterFile("NameService.proto", fileDescriptor_98ff7aff3b3ac208) }

var fileDescriptor_98ff7aff3b3ac208 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4f, 0x02, 0x31,
	0x10, 0x85, 0x29, 0x89, 0x24, 0x0c, 0x09, 0xb2, 0x13, 0x8c, 0xc8, 0xc9, 0xd4, 0x0b, 0xf1, 0x80,
	0x06, 0xaf, 0xc6, 0x44, 0x37, 0x26, 0x70, 0xd0, 0xc3, 0xf2, 0x03, 0x48, 0x65, 0x9f, 0xba, 0x87,
	0xed, 0x60, 0x5b, 0xf9, 0x2f, 0xfe, 0x5b, 0x63, 0x59, 0x61, 0x83, 0x5e, 0x38, 0x35, 0xe9, 0xf7,
	0x66, 0xe6, 0xbd, 0x47, 0xc9, 0xb3, 0x29, 0x31, 0x87, 0x5b, 0x17, 0x4b, 0x8c, 0x57, 0x4e, 0x82,
	0xf0, 0x51, 0x7c, 0xf4, 0x2d, 0xb5, 0xa7, 0x30, 0x2e, 0x3c, 0xc0, 0x04, 0xbe, 0xa2, 0xd6, 0x3b,
	0x4c, 0x0e, 0x37, 0x50, 0xe7, 0x6a, 0xd4, 0x99, 0x9c, 0x6e, 0xb4, 0xe3, 0x54, 0xac, 0xc5, 0x32,
	0x14, 0x62, 0xa7, 0x11, 0x67, 0x95, 0x4c, 0xdf, 0x53, 0x77, 0xc7, 0x66, 0xf6, 0x55, 0x0e, 0x5f,
	0x71, 0x47, 0xbd, 0x7d, 0xc6, 0x97, 0x94, 0x94, 0xf0, 0xde, 0xbc, 0x61, 0xe1, 0xe1, 0xd6, 0x70,
	0x8b, 0x22, 0x8f, 0xfb, 0xda, 0xd9, 0x71, 0x05, 0xe6, 0xf1, 0x7f, 0x96, 0xeb, 0x0b, 0x4a, 0x52,
	0x07, 0x13, 0x90, 0x89, 0x94, 0x19, 0x3e, 0x3e, 0xe1, 0x03, 0x77, 0xa9, 0xb9, 0x9d, 0x68, 0x16,
	0xb9, 0xee, 0x13, 0xd7, 0x45, 0x7e, 0x25, 0xd6, 0x63, 0xf2, 0xa5, 0xa8, 0x53, 0x2b, 0x86, 0x1f,
	0xa9, 0x5f, 0x59, 0x79, 0xaa, 0x1f, 0xe1, 0x5e, 0x95, 0x61, 0x5b, 0xd4, 0xf0, 0xe4, 0x4f, 0xaa,
	0x9f, 0xf0, 0xba, 0x31, 0x52, 0xd7, 0x8a, 0x53, 0xa2, 0xdd, 0x31, 0x1e, 0xfc, 0x4a, 0xf7, 0x4d,
	0x0e, 0xcf, 0xfe, 0x21, 0x1b, 0x67, 0xba, 0xf1, 0xd2, 0x8a, 0xec, 0xe6, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xfc, 0xd2, 0x02, 0xf2, 0xba, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NameServiceClient is the client API for NameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NameServiceClient interface {
	ConnectMessageServer(ctx context.Context, opts ...grpc.CallOption) (NameService_ConnectMessageServerClient, error)
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error)
}

type nameServiceClient struct {
	cc *grpc.ClientConn
}

func NewNameServiceClient(cc *grpc.ClientConn) NameServiceClient {
	return &nameServiceClient{cc}
}

func (c *nameServiceClient) ConnectMessageServer(ctx context.Context, opts ...grpc.CallOption) (NameService_ConnectMessageServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NameService_serviceDesc.Streams[0], "/proto.NameService/ConnectMessageServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &nameServiceConnectMessageServerClient{stream}
	return x, nil
}

type NameService_ConnectMessageServerClient interface {
	Send(*HeartBeat) error
	Recv() (*ConnectionInfo, error)
	grpc.ClientStream
}

type nameServiceConnectMessageServerClient struct {
	grpc.ClientStream
}

func (x *nameServiceConnectMessageServerClient) Send(m *HeartBeat) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nameServiceConnectMessageServerClient) Recv() (*ConnectionInfo, error) {
	m := new(ConnectionInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nameServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomResponse, error) {
	out := new(CreateRoomResponse)
	err := c.cc.Invoke(ctx, "/proto.NameService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NameServiceServer is the server API for NameService service.
type NameServiceServer interface {
	ConnectMessageServer(NameService_ConnectMessageServerServer) error
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomResponse, error)
}

// UnimplementedNameServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNameServiceServer struct {
}

func (*UnimplementedNameServiceServer) ConnectMessageServer(srv NameService_ConnectMessageServerServer) error {
	return status.Errorf(codes.Unimplemented, "method ConnectMessageServer not implemented")
}
func (*UnimplementedNameServiceServer) CreateRoom(ctx context.Context, req *CreateRoomRequest) (*CreateRoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}

func RegisterNameServiceServer(s *grpc.Server, srv NameServiceServer) {
	s.RegisterService(&_NameService_serviceDesc, srv)
}

func _NameService_ConnectMessageServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NameServiceServer).ConnectMessageServer(&nameServiceConnectMessageServerServer{stream})
}

type NameService_ConnectMessageServerServer interface {
	Send(*ConnectionInfo) error
	Recv() (*HeartBeat, error)
	grpc.ServerStream
}

type nameServiceConnectMessageServerServer struct {
	grpc.ServerStream
}

func (x *nameServiceConnectMessageServerServer) Send(m *ConnectionInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nameServiceConnectMessageServerServer) Recv() (*HeartBeat, error) {
	m := new(HeartBeat)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NameService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NameServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NameService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NameServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NameService",
	HandlerType: (*NameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _NameService_CreateRoom_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectMessageServer",
			Handler:       _NameService_ConnectMessageServer_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "NameService.proto",
}
