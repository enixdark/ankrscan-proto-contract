// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.13.0
// source: nft-indexer.proto

package proto

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

// NftIndexerClient is the client API for NftIndexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NftIndexerClient interface {
	// internal APIs
	NftBalanceByAddress(ctx context.Context, in *NftBalanceByAddressRequest, opts ...grpc.CallOption) (*NftBalanceByAddressReply, error)
	NftGetMetadata(ctx context.Context, in *NftGetMetadataRequest, opts ...grpc.CallOption) (*NftGetMetadataReply, error)
	// public APIs
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*NftBalanceByAddressReply, error)
	GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*NftGetMetadataReply, error)
}

type nftIndexerClient struct {
	cc grpc.ClientConnInterface
}

func NewNftIndexerClient(cc grpc.ClientConnInterface) NftIndexerClient {
	return &nftIndexerClient{cc}
}

func (c *nftIndexerClient) NftBalanceByAddress(ctx context.Context, in *NftBalanceByAddressRequest, opts ...grpc.CallOption) (*NftBalanceByAddressReply, error) {
	out := new(NftBalanceByAddressReply)
	err := c.cc.Invoke(ctx, "/ankrscan.nftindexer.NftIndexer/NftBalanceByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftIndexerClient) NftGetMetadata(ctx context.Context, in *NftGetMetadataRequest, opts ...grpc.CallOption) (*NftGetMetadataReply, error) {
	out := new(NftGetMetadataReply)
	err := c.cc.Invoke(ctx, "/ankrscan.nftindexer.NftIndexer/NftGetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftIndexerClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*NftBalanceByAddressReply, error) {
	out := new(NftBalanceByAddressReply)
	err := c.cc.Invoke(ctx, "/ankrscan.nftindexer.NftIndexer/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftIndexerClient) GetMetadata(ctx context.Context, in *GetMetadataRequest, opts ...grpc.CallOption) (*NftGetMetadataReply, error) {
	out := new(NftGetMetadataReply)
	err := c.cc.Invoke(ctx, "/ankrscan.nftindexer.NftIndexer/GetMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NftIndexerServer is the server API for NftIndexer service.
// All implementations must embed UnimplementedNftIndexerServer
// for forward compatibility
type NftIndexerServer interface {
	// internal APIs
	NftBalanceByAddress(context.Context, *NftBalanceByAddressRequest) (*NftBalanceByAddressReply, error)
	NftGetMetadata(context.Context, *NftGetMetadataRequest) (*NftGetMetadataReply, error)
	// public APIs
	GetBalance(context.Context, *GetBalanceRequest) (*NftBalanceByAddressReply, error)
	GetMetadata(context.Context, *GetMetadataRequest) (*NftGetMetadataReply, error)
	mustEmbedUnimplementedNftIndexerServer()
}

// UnimplementedNftIndexerServer must be embedded to have forward compatible implementations.
type UnimplementedNftIndexerServer struct {
}

func (UnimplementedNftIndexerServer) NftBalanceByAddress(context.Context, *NftBalanceByAddressRequest) (*NftBalanceByAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NftBalanceByAddress not implemented")
}
func (UnimplementedNftIndexerServer) NftGetMetadata(context.Context, *NftGetMetadataRequest) (*NftGetMetadataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NftGetMetadata not implemented")
}
func (UnimplementedNftIndexerServer) GetBalance(context.Context, *GetBalanceRequest) (*NftBalanceByAddressReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedNftIndexerServer) GetMetadata(context.Context, *GetMetadataRequest) (*NftGetMetadataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedNftIndexerServer) mustEmbedUnimplementedNftIndexerServer() {}

// UnsafeNftIndexerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NftIndexerServer will
// result in compilation errors.
type UnsafeNftIndexerServer interface {
	mustEmbedUnimplementedNftIndexerServer()
}

func RegisterNftIndexerServer(s grpc.ServiceRegistrar, srv NftIndexerServer) {
	s.RegisterService(&NftIndexer_ServiceDesc, srv)
}

func _NftIndexer_NftBalanceByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NftBalanceByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftIndexerServer).NftBalanceByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.nftindexer.NftIndexer/NftBalanceByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftIndexerServer).NftBalanceByAddress(ctx, req.(*NftBalanceByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NftIndexer_NftGetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NftGetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftIndexerServer).NftGetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.nftindexer.NftIndexer/NftGetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftIndexerServer).NftGetMetadata(ctx, req.(*NftGetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NftIndexer_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftIndexerServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.nftindexer.NftIndexer/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftIndexerServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NftIndexer_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftIndexerServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.nftindexer.NftIndexer/GetMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftIndexerServer).GetMetadata(ctx, req.(*GetMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NftIndexer_ServiceDesc is the grpc.ServiceDesc for NftIndexer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NftIndexer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ankrscan.nftindexer.NftIndexer",
	HandlerType: (*NftIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NftBalanceByAddress",
			Handler:    _NftIndexer_NftBalanceByAddress_Handler,
		},
		{
			MethodName: "NftGetMetadata",
			Handler:    _NftIndexer_NftGetMetadata_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _NftIndexer_GetBalance_Handler,
		},
		{
			MethodName: "GetMetadata",
			Handler:    _NftIndexer_GetMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nft-indexer.proto",
}
