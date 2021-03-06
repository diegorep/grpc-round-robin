// Code generated by protoc-gen-go. DO NOT EDIT.
// source: threes.proto

package threes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Three struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Three) Reset()         { *m = Three{} }
func (m *Three) String() string { return proto.CompactTextString(m) }
func (*Three) ProtoMessage()    {}
func (*Three) Descriptor() ([]byte, []int) {
	return fileDescriptor_threes_1033e26727525c11, []int{0}
}
func (m *Three) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Three.Unmarshal(m, b)
}
func (m *Three) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Three.Marshal(b, m, deterministic)
}
func (dst *Three) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Three.Merge(dst, src)
}
func (m *Three) XXX_Size() int {
	return xxx_messageInfo_Three.Size(m)
}
func (m *Three) XXX_DiscardUnknown() {
	xxx_messageInfo_Three.DiscardUnknown(m)
}

var xxx_messageInfo_Three proto.InternalMessageInfo

func (m *Three) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GetThreeRequest struct {
	Value                *Three   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetThreeRequest) Reset()         { *m = GetThreeRequest{} }
func (m *GetThreeRequest) String() string { return proto.CompactTextString(m) }
func (*GetThreeRequest) ProtoMessage()    {}
func (*GetThreeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_threes_1033e26727525c11, []int{1}
}
func (m *GetThreeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThreeRequest.Unmarshal(m, b)
}
func (m *GetThreeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThreeRequest.Marshal(b, m, deterministic)
}
func (dst *GetThreeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThreeRequest.Merge(dst, src)
}
func (m *GetThreeRequest) XXX_Size() int {
	return xxx_messageInfo_GetThreeRequest.Size(m)
}
func (m *GetThreeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThreeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetThreeRequest proto.InternalMessageInfo

func (m *GetThreeRequest) GetValue() *Three {
	if m != nil {
		return m.Value
	}
	return nil
}

type GetThreeResponse struct {
	Value                *Three   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetThreeResponse) Reset()         { *m = GetThreeResponse{} }
func (m *GetThreeResponse) String() string { return proto.CompactTextString(m) }
func (*GetThreeResponse) ProtoMessage()    {}
func (*GetThreeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_threes_1033e26727525c11, []int{2}
}
func (m *GetThreeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetThreeResponse.Unmarshal(m, b)
}
func (m *GetThreeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetThreeResponse.Marshal(b, m, deterministic)
}
func (dst *GetThreeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetThreeResponse.Merge(dst, src)
}
func (m *GetThreeResponse) XXX_Size() int {
	return xxx_messageInfo_GetThreeResponse.Size(m)
}
func (m *GetThreeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetThreeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetThreeResponse proto.InternalMessageInfo

func (m *GetThreeResponse) GetValue() *Three {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Three)(nil), "Three")
	proto.RegisterType((*GetThreeRequest)(nil), "GetThreeRequest")
	proto.RegisterType((*GetThreeResponse)(nil), "GetThreeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ThreesClient is the client API for Threes service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThreesClient interface {
	EchoThree(ctx context.Context, in *GetThreeRequest, opts ...grpc.CallOption) (*GetThreeResponse, error)
}

type threesClient struct {
	cc *grpc.ClientConn
}

func NewThreesClient(cc *grpc.ClientConn) ThreesClient {
	return &threesClient{cc}
}

func (c *threesClient) EchoThree(ctx context.Context, in *GetThreeRequest, opts ...grpc.CallOption) (*GetThreeResponse, error) {
	out := new(GetThreeResponse)
	err := c.cc.Invoke(ctx, "/Threes/EchoThree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThreesServer is the server API for Threes service.
type ThreesServer interface {
	EchoThree(context.Context, *GetThreeRequest) (*GetThreeResponse, error)
}

func RegisterThreesServer(s *grpc.Server, srv ThreesServer) {
	s.RegisterService(&_Threes_serviceDesc, srv)
}

func _Threes_EchoThree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetThreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreesServer).EchoThree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Threes/EchoThree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreesServer).EchoThree(ctx, req.(*GetThreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Threes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Threes",
	HandlerType: (*ThreesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoThree",
			Handler:    _Threes_EchoThree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "threes.proto",
}

func init() { proto.RegisterFile("threes.proto", fileDescriptor_threes_1033e26727525c11) }

var fileDescriptor_threes_1033e26727525c11 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc9, 0x28, 0x4a,
	0x4d, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x92, 0xe5, 0x62, 0x0d, 0x01, 0xf1, 0x85,
	0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x98, 0x82, 0x20,
	0x1c, 0x25, 0x7d, 0x2e, 0x7e, 0xf7, 0xd4, 0x12, 0xb0, 0x8a, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x19, 0x64, 0x85, 0xdc, 0x46, 0x6c, 0x7a, 0x10, 0x59, 0xa8, 0x06, 0x03, 0x2e, 0x01,
	0x84, 0x86, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xfc, 0x3a, 0x8c, 0x6c, 0xb8, 0xd8, 0xc0, 0xfc,
	0x62, 0x21, 0x23, 0x2e, 0x4e, 0xd7, 0xe4, 0x8c, 0x7c, 0x88, 0x7b, 0x04, 0xf4, 0xd0, 0x2c, 0x96,
	0x12, 0xd4, 0x43, 0x37, 0x59, 0x89, 0x21, 0x89, 0x0d, 0xec, 0x0d, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0xaf, 0x6c, 0x79, 0xd6, 0x00, 0x00, 0x00,
}
