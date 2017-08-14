// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/client.proto

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	proto/client.proto

It has these top-level messages:
	VoidRequest
	HelloRequest
	HelloReply
*/
package helloworld

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

// The request message containing the user's name.
type VoidRequest struct {
}

func (m *VoidRequest) Reset()                    { *m = VoidRequest{} }
func (m *VoidRequest) String() string            { return proto.CompactTextString(m) }
func (*VoidRequest) ProtoMessage()               {}
func (*VoidRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*VoidRequest)(nil), "helloworld.VoidRequest")
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SampleProject service

type SampleProjectClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHelloEmpty(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type sampleProjectClient struct {
	cc *grpc.ClientConn
}

func NewSampleProjectClient(cc *grpc.ClientConn) SampleProjectClient {
	return &sampleProjectClient{cc}
}

func (c *sampleProjectClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.SampleProject/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sampleProjectClient) SayHelloEmpty(ctx context.Context, in *VoidRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.SampleProject/SayHelloEmpty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SampleProject service

type SampleProjectServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloEmpty(context.Context, *VoidRequest) (*HelloReply, error)
}

func RegisterSampleProjectServer(s *grpc.Server, srv SampleProjectServer) {
	s.RegisterService(&_SampleProject_serviceDesc, srv)
}

func _SampleProject_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleProjectServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.SampleProject/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleProjectServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SampleProject_SayHelloEmpty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleProjectServer).SayHelloEmpty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.SampleProject/SayHelloEmpty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleProjectServer).SayHelloEmpty(ctx, req.(*VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SampleProject_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.SampleProject",
	HandlerType: (*SampleProjectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SampleProject_SayHello_Handler,
		},
		{
			MethodName: "SayHelloEmpty",
			Handler:    _SampleProject_SayHelloEmpty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/client.proto",
}

func init() { proto.RegisterFile("proto/client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x03, 0x73, 0x84, 0xb8, 0x32, 0x52, 0x73,
	0x72, 0xf2, 0xcb, 0xf3, 0x8b, 0x72, 0x52, 0xa4, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5,
	0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21,
	0x2a, 0x95, 0x78, 0xb9, 0xb8, 0xc3, 0xf2, 0x33, 0x53, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x94, 0x94, 0xb8, 0x78, 0x3c, 0x40, 0x5a, 0xa1, 0x7c, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x8d, 0x8b, 0x0b, 0xaa, 0xa6,
	0x20, 0xa7, 0x52, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6, 0x08, 0xc6,
	0x35, 0xda, 0xc4, 0xc8, 0xc5, 0x1b, 0x9c, 0x98, 0x5b, 0x90, 0x93, 0x1a, 0x50, 0x94, 0x9f, 0x95,
	0x9a, 0x5c, 0x22, 0x14, 0xcc, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xd6, 0x2c, 0x24, 0xa1, 0x87, 0x70,
	0xa3, 0x1e, 0xb2, 0x9d, 0x52, 0x62, 0x58, 0x64, 0x0a, 0x72, 0x2a, 0x95, 0x44, 0x9b, 0x2e, 0x3f,
	0x99, 0xcc, 0xc4, 0x2f, 0xc4, 0xab, 0x0f, 0x96, 0xd7, 0xaf, 0x06, 0xb9, 0xa6, 0x56, 0xc8, 0x07,
	0x64, 0x0b, 0xc4, 0x50, 0xd7, 0xdc, 0x82, 0x92, 0x4a, 0x21, 0x71, 0x64, 0xfd, 0x48, 0x9e, 0xc3,
	0x69, 0x30, 0x17, 0xd8, 0x60, 0x16, 0x21, 0x26, 0x7d, 0x2d, 0x27, 0x3d, 0x2e, 0xe9, 0xcc, 0x7c,
	0xbd, 0xf4, 0xa2, 0x82, 0x64, 0xbd, 0xd4, 0x0a, 0xb0, 0xe3, 0x8b, 0x91, 0x74, 0x39, 0xa1, 0x7a,
	0x28, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x8c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x20,
	0x6b, 0xfd, 0x86, 0x01, 0x00, 0x00,
}