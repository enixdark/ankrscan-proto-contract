// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MethodResolverClient is the client API for MethodResolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MethodResolverClient interface {
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveReply, error)
}

type methodResolverClient struct {
	cc grpc.ClientConnInterface
}

func NewMethodResolverClient(cc grpc.ClientConnInterface) MethodResolverClient {
	return &methodResolverClient{cc}
}

func (c *methodResolverClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveReply, error) {
	out := new(ResolveReply)
	err := c.cc.Invoke(ctx, "/com.clover.methodresolver.MethodResolver/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MethodResolverServer is the server API for MethodResolver service.
// All implementations must embed UnimplementedMethodResolverServer
// for forward compatibility
type MethodResolverServer interface {
	Resolve(context.Context, *ResolveRequest) (*ResolveReply, error)
	mustEmbedUnimplementedMethodResolverServer()
}

// UnimplementedMethodResolverServer must be embedded to have forward compatible implementations.
type UnimplementedMethodResolverServer struct {
}

func (UnimplementedMethodResolverServer) Resolve(context.Context, *ResolveRequest) (*ResolveReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolve not implemented")
}
func (UnimplementedMethodResolverServer) mustEmbedUnimplementedMethodResolverServer() {}

// UnsafeMethodResolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MethodResolverServer will
// result in compilation errors.
type UnsafeMethodResolverServer interface {
	mustEmbedUnimplementedMethodResolverServer()
}

func RegisterMethodResolverServer(s grpc.ServiceRegistrar, srv MethodResolverServer) {
	s.RegisterService(&MethodResolver_ServiceDesc, srv)
}

func _MethodResolver_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MethodResolverServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.methodresolver.MethodResolver/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MethodResolverServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MethodResolver_ServiceDesc is the grpc.ServiceDesc for MethodResolver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MethodResolver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.clover.methodresolver.MethodResolver",
	HandlerType: (*MethodResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resolve",
			Handler:    _MethodResolver_Resolve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "method-resolver.proto",
}
