// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// LoginServerClient is the client API for LoginServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServerClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type loginServerClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServerClient(cc grpc.ClientConnInterface) LoginServerClient {
	return &loginServerClient{cc}
}

func (c *loginServerClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/login_server.v1.LoginServer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServerServer is the server API for LoginServer service.
// All implementations must embed UnimplementedLoginServerServer
// for forward compatibility
type LoginServerServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedLoginServerServer()
}

// UnimplementedLoginServerServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServerServer struct {
}

func (UnimplementedLoginServerServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedLoginServerServer) mustEmbedUnimplementedLoginServerServer() {}

// UnsafeLoginServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServerServer will
// result in compilation errors.
type UnsafeLoginServerServer interface {
	mustEmbedUnimplementedLoginServerServer()
}

func RegisterLoginServerServer(s grpc.ServiceRegistrar, srv LoginServerServer) {
	s.RegisterService(&LoginServer_ServiceDesc, srv)
}

func _LoginServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login_server.v1.LoginServer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServerServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginServer_ServiceDesc is the grpc.ServiceDesc for LoginServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "login_server.v1.LoginServer",
	HandlerType: (*LoginServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _LoginServer_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/login_server/v1/login_server.proto",
}
