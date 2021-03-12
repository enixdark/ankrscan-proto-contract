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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TxoCheckerClient is the client API for TxoChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TxoCheckerClient interface {
	GetTxoInfoByTxHash(ctx context.Context, in *GetTxoInfoByTxHashRequest, opts ...grpc.CallOption) (*GetTxoInfoReply, error)
	GetTxoInfoByTxHashAndVout(ctx context.Context, in *GetTxoInfoByTxHashAndVoutRequest, opts ...grpc.CallOption) (*GetTxoInfoReply, error)
}

type txoCheckerClient struct {
	cc grpc.ClientConnInterface
}

func NewTxoCheckerClient(cc grpc.ClientConnInterface) TxoCheckerClient {
	return &txoCheckerClient{cc}
}

func (c *txoCheckerClient) GetTxoInfoByTxHash(ctx context.Context, in *GetTxoInfoByTxHashRequest, opts ...grpc.CallOption) (*GetTxoInfoReply, error) {
	out := new(GetTxoInfoReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TxoChecker/GetTxoInfoByTxHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *txoCheckerClient) GetTxoInfoByTxHashAndVout(ctx context.Context, in *GetTxoInfoByTxHashAndVoutRequest, opts ...grpc.CallOption) (*GetTxoInfoReply, error) {
	out := new(GetTxoInfoReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TxoChecker/GetTxoInfoByTxHashAndVout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TxoCheckerServer is the server API for TxoChecker service.
// All implementations must embed UnimplementedTxoCheckerServer
// for forward compatibility
type TxoCheckerServer interface {
	GetTxoInfoByTxHash(context.Context, *GetTxoInfoByTxHashRequest) (*GetTxoInfoReply, error)
	GetTxoInfoByTxHashAndVout(context.Context, *GetTxoInfoByTxHashAndVoutRequest) (*GetTxoInfoReply, error)
	mustEmbedUnimplementedTxoCheckerServer()
}

// UnimplementedTxoCheckerServer must be embedded to have forward compatible implementations.
type UnimplementedTxoCheckerServer struct {
}

func (UnimplementedTxoCheckerServer) GetTxoInfoByTxHash(context.Context, *GetTxoInfoByTxHashRequest) (*GetTxoInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxoInfoByTxHash not implemented")
}
func (UnimplementedTxoCheckerServer) GetTxoInfoByTxHashAndVout(context.Context, *GetTxoInfoByTxHashAndVoutRequest) (*GetTxoInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxoInfoByTxHashAndVout not implemented")
}
func (UnimplementedTxoCheckerServer) mustEmbedUnimplementedTxoCheckerServer() {}

// UnsafeTxoCheckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TxoCheckerServer will
// result in compilation errors.
type UnsafeTxoCheckerServer interface {
	mustEmbedUnimplementedTxoCheckerServer()
}

func RegisterTxoCheckerServer(s grpc.ServiceRegistrar, srv TxoCheckerServer) {
	s.RegisterService(&TxoChecker_ServiceDesc, srv)
}

func _TxoChecker_GetTxoInfoByTxHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxoInfoByTxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxoCheckerServer).GetTxoInfoByTxHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TxoChecker/GetTxoInfoByTxHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxoCheckerServer).GetTxoInfoByTxHash(ctx, req.(*GetTxoInfoByTxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TxoChecker_GetTxoInfoByTxHashAndVout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxoInfoByTxHashAndVoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TxoCheckerServer).GetTxoInfoByTxHashAndVout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TxoChecker/GetTxoInfoByTxHashAndVout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TxoCheckerServer).GetTxoInfoByTxHashAndVout(ctx, req.(*GetTxoInfoByTxHashAndVoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TxoChecker_ServiceDesc is the grpc.ServiceDesc for TxoChecker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TxoChecker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.clover.extractor.TxoChecker",
	HandlerType: (*TxoCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTxoInfoByTxHash",
			Handler:    _TxoChecker_GetTxoInfoByTxHash_Handler,
		},
		{
			MethodName: "GetTxoInfoByTxHashAndVout",
			Handler:    _TxoChecker_GetTxoInfoByTxHashAndVout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "txo-checker.proto",
}
