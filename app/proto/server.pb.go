// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/server.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetDataRequest struct {
	Param                string   `protobuf:"bytes,1,opt,name=param" json:"param,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataRequest) Reset()         { *m = GetDataRequest{} }
func (m *GetDataRequest) String() string { return proto.CompactTextString(m) }
func (*GetDataRequest) ProtoMessage()    {}
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_659654d932168da8, []int{0}
}
func (m *GetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataRequest.Unmarshal(m, b)
}
func (m *GetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataRequest.Marshal(b, m, deterministic)
}
func (dst *GetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataRequest.Merge(dst, src)
}
func (m *GetDataRequest) XXX_Size() int {
	return xxx_messageInfo_GetDataRequest.Size(m)
}
func (m *GetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataRequest proto.InternalMessageInfo

func (m *GetDataRequest) GetParam() string {
	if m != nil {
		return m.Param
	}
	return ""
}

type GetDataResponse struct {
	Str                  string   `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
	Num                  uint32   `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDataResponse) Reset()         { *m = GetDataResponse{} }
func (m *GetDataResponse) String() string { return proto.CompactTextString(m) }
func (*GetDataResponse) ProtoMessage()    {}
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_659654d932168da8, []int{1}
}
func (m *GetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDataResponse.Unmarshal(m, b)
}
func (m *GetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDataResponse.Marshal(b, m, deterministic)
}
func (dst *GetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDataResponse.Merge(dst, src)
}
func (m *GetDataResponse) XXX_Size() int {
	return xxx_messageInfo_GetDataResponse.Size(m)
}
func (m *GetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDataResponse proto.InternalMessageInfo

func (m *GetDataResponse) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *GetDataResponse) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*GetDataRequest)(nil), "proto.GetDataRequest")
	proto.RegisterType((*GetDataResponse)(nil), "proto.GetDataResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Server service

type ServerClient interface {
	GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error)
}

type serverClient struct {
	cc *grpc.ClientConn
}

func NewServerClient(cc *grpc.ClientConn) ServerClient {
	return &serverClient{cc}
}

func (c *serverClient) GetData(ctx context.Context, in *GetDataRequest, opts ...grpc.CallOption) (*GetDataResponse, error) {
	out := new(GetDataResponse)
	err := grpc.Invoke(ctx, "/proto.Server/GetData", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Server service

type ServerServer interface {
	GetData(context.Context, *GetDataRequest) (*GetDataResponse, error)
}

func RegisterServerServer(s *grpc.Server, srv ServerServer) {
	s.RegisterService(&_Server_serviceDesc, srv)
}

func _Server_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Server/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerServer).GetData(ctx, req.(*GetDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Server_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Server",
	HandlerType: (*ServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _Server_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server.proto",
}

func init() { proto.RegisterFile("proto/server.proto", fileDescriptor_server_659654d932168da8) }

var fileDescriptor_server_659654d932168da8 = []byte{
	// 190 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x94,
	0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e,
	0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44, 0x91, 0x92, 0x1a, 0x17, 0x9f, 0x7b, 0x6a, 0x89,
	0x4b, 0x62, 0x49, 0x62, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x08, 0x17, 0x6b, 0x41,
	0x62, 0x51, 0x62, 0xae, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0xa3, 0x64, 0xca, 0xc5,
	0x0f, 0x57, 0x57, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc0, 0xc5, 0x5c, 0x5c, 0x52, 0x04,
	0x55, 0x06, 0x62, 0x82, 0x44, 0xf2, 0x4a, 0x73, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x83, 0x40,
	0x4c, 0xa3, 0x40, 0x2e, 0xb6, 0x60, 0xb0, 0x9b, 0x84, 0xdc, 0xb9, 0xd8, 0xa1, 0x06, 0x08, 0x89,
	0x42, 0xec, 0xd6, 0x43, 0xb5, 0x58, 0x4a, 0x0c, 0x5d, 0x18, 0x62, 0x8f, 0x12, 0x6f, 0xd3, 0xe5,
	0x27, 0x93, 0x99, 0xd8, 0x85, 0x58, 0xf5, 0x53, 0x12, 0x4b, 0x12, 0x93, 0xd8, 0xc0, 0xaa, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x73, 0x5e, 0x20, 0xf3, 0x00, 0x00, 0x00,
}
