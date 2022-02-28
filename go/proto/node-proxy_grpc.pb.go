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

// NodeProxyClient is the client API for NodeProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeProxyClient interface {
	BlockByNumber(ctx context.Context, in *BlockByNumberRequest, opts ...grpc.CallOption) (*BlockByNumberReply, error)
	LatestBlockHeader(ctx context.Context, in *LatestBlockHeaderRequest, opts ...grpc.CallOption) (*LatestBlockHeaderReply, error)
}

type nodeProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeProxyClient(cc grpc.ClientConnInterface) NodeProxyClient {
	return &nodeProxyClient{cc}
}

func (c *nodeProxyClient) BlockByNumber(ctx context.Context, in *BlockByNumberRequest, opts ...grpc.CallOption) (*BlockByNumberReply, error) {
	out := new(BlockByNumberReply)
	err := c.cc.Invoke(ctx, "/ankrscan.nodeproxy.NodeProxy/BlockByNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeProxyClient) LatestBlockHeader(ctx context.Context, in *LatestBlockHeaderRequest, opts ...grpc.CallOption) (*LatestBlockHeaderReply, error) {
	out := new(LatestBlockHeaderReply)
	err := c.cc.Invoke(ctx, "/ankrscan.nodeproxy.NodeProxy/LatestBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeProxyServer is the server API for NodeProxy service.
// All implementations must embed UnimplementedNodeProxyServer
// for forward compatibility
type NodeProxyServer interface {
	BlockByNumber(context.Context, *BlockByNumberRequest) (*BlockByNumberReply, error)
	LatestBlockHeader(context.Context, *LatestBlockHeaderRequest) (*LatestBlockHeaderReply, error)
	mustEmbedUnimplementedNodeProxyServer()
}

// UnimplementedNodeProxyServer must be embedded to have forward compatible implementations.
type UnimplementedNodeProxyServer struct {
}

func (UnimplementedNodeProxyServer) BlockByNumber(context.Context, *BlockByNumberRequest) (*BlockByNumberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockByNumber not implemented")
}
func (UnimplementedNodeProxyServer) LatestBlockHeader(context.Context, *LatestBlockHeaderRequest) (*LatestBlockHeaderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestBlockHeader not implemented")
}
func (UnimplementedNodeProxyServer) mustEmbedUnimplementedNodeProxyServer() {}

// UnsafeNodeProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeProxyServer will
// result in compilation errors.
type UnsafeNodeProxyServer interface {
	mustEmbedUnimplementedNodeProxyServer()
}

func RegisterNodeProxyServer(s grpc.ServiceRegistrar, srv NodeProxyServer) {
	s.RegisterService(&NodeProxy_ServiceDesc, srv)
}

func _NodeProxy_BlockByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockByNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeProxyServer).BlockByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.nodeproxy.NodeProxy/BlockByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeProxyServer).BlockByNumber(ctx, req.(*BlockByNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeProxy_LatestBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestBlockHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeProxyServer).LatestBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.nodeproxy.NodeProxy/LatestBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeProxyServer).LatestBlockHeader(ctx, req.(*LatestBlockHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeProxy_ServiceDesc is the grpc.ServiceDesc for NodeProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ankrscan.nodeproxy.NodeProxy",
	HandlerType: (*NodeProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BlockByNumber",
			Handler:    _NodeProxy_BlockByNumber_Handler,
		},
		{
			MethodName: "LatestBlockHeader",
			Handler:    _NodeProxy_LatestBlockHeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node-proxy.proto",
}
