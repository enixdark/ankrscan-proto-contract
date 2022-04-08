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

// TokenIndexerClient is the client API for TokenIndexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenIndexerClient interface {
	BalanceByHolder(ctx context.Context, in *BalanceByHolderRequest, opts ...grpc.CallOption) (*BalancesDetailedReply, error)
	UsdPrice(ctx context.Context, in *UsdPricesRequest, opts ...grpc.CallOption) (*UsdPricesReply, error)
	TokenHolders(ctx context.Context, in *TokenHoldersRequest, opts ...grpc.CallOption) (*TokenHoldersReply, error)
	Currencies(ctx context.Context, in *CurrenciesRequest, opts ...grpc.CallOption) (*CurrenciesReply, error)
}

type tokenIndexerClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenIndexerClient(cc grpc.ClientConnInterface) TokenIndexerClient {
	return &tokenIndexerClient{cc}
}

func (c *tokenIndexerClient) BalanceByHolder(ctx context.Context, in *BalanceByHolderRequest, opts ...grpc.CallOption) (*BalancesDetailedReply, error) {
	out := new(BalancesDetailedReply)
	err := c.cc.Invoke(ctx, "/ankrscan.tokenindexer.TokenIndexer/BalanceByHolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenIndexerClient) UsdPrice(ctx context.Context, in *UsdPricesRequest, opts ...grpc.CallOption) (*UsdPricesReply, error) {
	out := new(UsdPricesReply)
	err := c.cc.Invoke(ctx, "/ankrscan.tokenindexer.TokenIndexer/UsdPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenIndexerClient) TokenHolders(ctx context.Context, in *TokenHoldersRequest, opts ...grpc.CallOption) (*TokenHoldersReply, error) {
	out := new(TokenHoldersReply)
	err := c.cc.Invoke(ctx, "/ankrscan.tokenindexer.TokenIndexer/TokenHolders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenIndexerClient) Currencies(ctx context.Context, in *CurrenciesRequest, opts ...grpc.CallOption) (*CurrenciesReply, error) {
	out := new(CurrenciesReply)
	err := c.cc.Invoke(ctx, "/ankrscan.tokenindexer.TokenIndexer/Currencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenIndexerServer is the server API for TokenIndexer service.
// All implementations must embed UnimplementedTokenIndexerServer
// for forward compatibility
type TokenIndexerServer interface {
	BalanceByHolder(context.Context, *BalanceByHolderRequest) (*BalancesDetailedReply, error)
	UsdPrice(context.Context, *UsdPricesRequest) (*UsdPricesReply, error)
	TokenHolders(context.Context, *TokenHoldersRequest) (*TokenHoldersReply, error)
	Currencies(context.Context, *CurrenciesRequest) (*CurrenciesReply, error)
	mustEmbedUnimplementedTokenIndexerServer()
}

// UnimplementedTokenIndexerServer must be embedded to have forward compatible implementations.
type UnimplementedTokenIndexerServer struct {
}

func (UnimplementedTokenIndexerServer) BalanceByHolder(context.Context, *BalanceByHolderRequest) (*BalancesDetailedReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceByHolder not implemented")
}
func (UnimplementedTokenIndexerServer) UsdPrice(context.Context, *UsdPricesRequest) (*UsdPricesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UsdPrice not implemented")
}
func (UnimplementedTokenIndexerServer) TokenHolders(context.Context, *TokenHoldersRequest) (*TokenHoldersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenHolders not implemented")
}
func (UnimplementedTokenIndexerServer) Currencies(context.Context, *CurrenciesRequest) (*CurrenciesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Currencies not implemented")
}
func (UnimplementedTokenIndexerServer) mustEmbedUnimplementedTokenIndexerServer() {}

// UnsafeTokenIndexerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenIndexerServer will
// result in compilation errors.
type UnsafeTokenIndexerServer interface {
	mustEmbedUnimplementedTokenIndexerServer()
}

func RegisterTokenIndexerServer(s grpc.ServiceRegistrar, srv TokenIndexerServer) {
	s.RegisterService(&TokenIndexer_ServiceDesc, srv)
}

func _TokenIndexer_BalanceByHolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceByHolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenIndexerServer).BalanceByHolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.tokenindexer.TokenIndexer/BalanceByHolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenIndexerServer).BalanceByHolder(ctx, req.(*BalanceByHolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenIndexer_UsdPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UsdPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenIndexerServer).UsdPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.tokenindexer.TokenIndexer/UsdPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenIndexerServer).UsdPrice(ctx, req.(*UsdPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenIndexer_TokenHolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenHoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenIndexerServer).TokenHolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.tokenindexer.TokenIndexer/TokenHolders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenIndexerServer).TokenHolders(ctx, req.(*TokenHoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenIndexer_Currencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenIndexerServer).Currencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.tokenindexer.TokenIndexer/Currencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenIndexerServer).Currencies(ctx, req.(*CurrenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenIndexer_ServiceDesc is the grpc.ServiceDesc for TokenIndexer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenIndexer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ankrscan.tokenindexer.TokenIndexer",
	HandlerType: (*TokenIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BalanceByHolder",
			Handler:    _TokenIndexer_BalanceByHolder_Handler,
		},
		{
			MethodName: "UsdPrice",
			Handler:    _TokenIndexer_UsdPrice_Handler,
		},
		{
			MethodName: "TokenHolders",
			Handler:    _TokenIndexer_TokenHolders_Handler,
		},
		{
			MethodName: "Currencies",
			Handler:    _TokenIndexer_Currencies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "token-indexer.proto",
}
