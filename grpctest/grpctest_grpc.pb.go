// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpctest

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GrpctestClient is the client API for Grpctest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpctestClient interface {
	GrpcConnect(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type grpctestClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpctestClient(cc grpc.ClientConnInterface) GrpctestClient {
	return &grpctestClient{cc}
}

func (c *grpctestClient) GrpcConnect(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/grpctest.grpctest/GrpcConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpctestServer is the server API for Grpctest service.
// All implementations must embed UnimplementedGrpctestServer
// for forward compatibility
type GrpctestServer interface {
	GrpcConnect(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedGrpctestServer()
}

// UnimplementedGrpctestServer must be embedded to have forward compatible implementations.
type UnimplementedGrpctestServer struct {
}

func (UnimplementedGrpctestServer) GrpcConnect(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrpcConnect not implemented")
}
func (UnimplementedGrpctestServer) mustEmbedUnimplementedGrpctestServer() {}

// UnsafeGrpctestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpctestServer will
// result in compilation errors.
type UnsafeGrpctestServer interface {
	mustEmbedUnimplementedGrpctestServer()
}

func RegisterGrpctestServer(s grpc.ServiceRegistrar, srv GrpctestServer) {
	s.RegisterService(&Grpctest_ServiceDesc, srv)
}

func _Grpctest_GrpcConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpctestServer).GrpcConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpctest.grpctest/GrpcConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpctestServer).GrpcConnect(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// Grpctest_ServiceDesc is the grpc.ServiceDesc for Grpctest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grpctest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpctest.grpctest",
	HandlerType: (*GrpctestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GrpcConnect",
			Handler:    _Grpctest_GrpcConnect_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpctest.proto",
}