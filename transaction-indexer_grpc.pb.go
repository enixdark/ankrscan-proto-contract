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

// TransactionIndexerClient is the client API for TransactionIndexer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionIndexerClient interface {
	GetTxsByAddress(ctx context.Context, in *GetTxsByAddressRequest, opts ...grpc.CallOption) (*SimpleTransactionsReply, error)
	GetTxsByHeightAndBlockchain(ctx context.Context, in *GetTxsByHeightAndBlockchainRequest, opts ...grpc.CallOption) (*SimpleTransactionsReply, error)
	GetBlocksByHeight(ctx context.Context, in *GetBlocksByHeightRequest, opts ...grpc.CallOption) (*SimpleBlocksReply, error)
	GetBlockDetails(ctx context.Context, in *GetBlockDetailsRequest, opts ...grpc.CallOption) (*GetBlockDetailsReply, error)
	GetTransactionDetails(ctx context.Context, in *GetTransactionDetailsRequest, opts ...grpc.CallOption) (*GetTransactionDetailsReply, error)
	GetLatestBlocks(ctx context.Context, in *GetLatestBlocksRequest, opts ...grpc.CallOption) (*SimpleBlocksReply, error)
	GetLatestTransactions(ctx context.Context, in *GetLatestTransactionsRequest, opts ...grpc.CallOption) (*SimpleTransactionsReply, error)
	GetMoneyTransfersByAddress(ctx context.Context, in *GetTxsByAddressRequest, opts ...grpc.CallOption) (*MoneyTransfersReply, error)
	GetMoneyTransfersByAddressAndTime(ctx context.Context, in *GetTxsByAddressAndTimeRequest, opts ...grpc.CallOption) (*MoneyTransfersReply, error)
	GetTxsByAddressAndTime(ctx context.Context, in *GetTxsByAddressAndTimeRequest, opts ...grpc.CallOption) (*SimpleTransactionsReply, error)
}

type transactionIndexerClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionIndexerClient(cc grpc.ClientConnInterface) TransactionIndexerClient {
	return &transactionIndexerClient{cc}
}

func (c *transactionIndexerClient) GetTxsByAddress(ctx context.Context, in *GetTxsByAddressRequest, opts ...grpc.CallOption) (*SimpleTransactionsReply, error) {
	out := new(SimpleTransactionsReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetTxsByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetTxsByHeightAndBlockchain(ctx context.Context, in *GetTxsByHeightAndBlockchainRequest, opts ...grpc.CallOption) (*SimpleTransactionsReply, error) {
	out := new(SimpleTransactionsReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetTxsByHeightAndBlockchain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetBlocksByHeight(ctx context.Context, in *GetBlocksByHeightRequest, opts ...grpc.CallOption) (*SimpleBlocksReply, error) {
	out := new(SimpleBlocksReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetBlocksByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetBlockDetails(ctx context.Context, in *GetBlockDetailsRequest, opts ...grpc.CallOption) (*GetBlockDetailsReply, error) {
	out := new(GetBlockDetailsReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetBlockDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetTransactionDetails(ctx context.Context, in *GetTransactionDetailsRequest, opts ...grpc.CallOption) (*GetTransactionDetailsReply, error) {
	out := new(GetTransactionDetailsReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetTransactionDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetLatestBlocks(ctx context.Context, in *GetLatestBlocksRequest, opts ...grpc.CallOption) (*SimpleBlocksReply, error) {
	out := new(SimpleBlocksReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetLatestBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetLatestTransactions(ctx context.Context, in *GetLatestTransactionsRequest, opts ...grpc.CallOption) (*SimpleTransactionsReply, error) {
	out := new(SimpleTransactionsReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetLatestTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetMoneyTransfersByAddress(ctx context.Context, in *GetTxsByAddressRequest, opts ...grpc.CallOption) (*MoneyTransfersReply, error) {
	out := new(MoneyTransfersReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetMoneyTransfersByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetMoneyTransfersByAddressAndTime(ctx context.Context, in *GetTxsByAddressAndTimeRequest, opts ...grpc.CallOption) (*MoneyTransfersReply, error) {
	out := new(MoneyTransfersReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetMoneyTransfersByAddressAndTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionIndexerClient) GetTxsByAddressAndTime(ctx context.Context, in *GetTxsByAddressAndTimeRequest, opts ...grpc.CallOption) (*SimpleTransactionsReply, error) {
	out := new(SimpleTransactionsReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.TransactionIndexer/GetTxsByAddressAndTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionIndexerServer is the server API for TransactionIndexer service.
// All implementations must embed UnimplementedTransactionIndexerServer
// for forward compatibility
type TransactionIndexerServer interface {
	GetTxsByAddress(context.Context, *GetTxsByAddressRequest) (*SimpleTransactionsReply, error)
	GetTxsByHeightAndBlockchain(context.Context, *GetTxsByHeightAndBlockchainRequest) (*SimpleTransactionsReply, error)
	GetBlocksByHeight(context.Context, *GetBlocksByHeightRequest) (*SimpleBlocksReply, error)
	GetBlockDetails(context.Context, *GetBlockDetailsRequest) (*GetBlockDetailsReply, error)
	GetTransactionDetails(context.Context, *GetTransactionDetailsRequest) (*GetTransactionDetailsReply, error)
	GetLatestBlocks(context.Context, *GetLatestBlocksRequest) (*SimpleBlocksReply, error)
	GetLatestTransactions(context.Context, *GetLatestTransactionsRequest) (*SimpleTransactionsReply, error)
	GetMoneyTransfersByAddress(context.Context, *GetTxsByAddressRequest) (*MoneyTransfersReply, error)
	GetMoneyTransfersByAddressAndTime(context.Context, *GetTxsByAddressAndTimeRequest) (*MoneyTransfersReply, error)
	GetTxsByAddressAndTime(context.Context, *GetTxsByAddressAndTimeRequest) (*SimpleTransactionsReply, error)
	mustEmbedUnimplementedTransactionIndexerServer()
}

// UnimplementedTransactionIndexerServer must be embedded to have forward compatible implementations.
type UnimplementedTransactionIndexerServer struct {
}

func (UnimplementedTransactionIndexerServer) GetTxsByAddress(context.Context, *GetTxsByAddressRequest) (*SimpleTransactionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxsByAddress not implemented")
}
func (UnimplementedTransactionIndexerServer) GetTxsByHeightAndBlockchain(context.Context, *GetTxsByHeightAndBlockchainRequest) (*SimpleTransactionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxsByHeightAndBlockchain not implemented")
}
func (UnimplementedTransactionIndexerServer) GetBlocksByHeight(context.Context, *GetBlocksByHeightRequest) (*SimpleBlocksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocksByHeight not implemented")
}
func (UnimplementedTransactionIndexerServer) GetBlockDetails(context.Context, *GetBlockDetailsRequest) (*GetBlockDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockDetails not implemented")
}
func (UnimplementedTransactionIndexerServer) GetTransactionDetails(context.Context, *GetTransactionDetailsRequest) (*GetTransactionDetailsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionDetails not implemented")
}
func (UnimplementedTransactionIndexerServer) GetLatestBlocks(context.Context, *GetLatestBlocksRequest) (*SimpleBlocksReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlocks not implemented")
}
func (UnimplementedTransactionIndexerServer) GetLatestTransactions(context.Context, *GetLatestTransactionsRequest) (*SimpleTransactionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestTransactions not implemented")
}
func (UnimplementedTransactionIndexerServer) GetMoneyTransfersByAddress(context.Context, *GetTxsByAddressRequest) (*MoneyTransfersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoneyTransfersByAddress not implemented")
}
func (UnimplementedTransactionIndexerServer) GetMoneyTransfersByAddressAndTime(context.Context, *GetTxsByAddressAndTimeRequest) (*MoneyTransfersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoneyTransfersByAddressAndTime not implemented")
}
func (UnimplementedTransactionIndexerServer) GetTxsByAddressAndTime(context.Context, *GetTxsByAddressAndTimeRequest) (*SimpleTransactionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxsByAddressAndTime not implemented")
}
func (UnimplementedTransactionIndexerServer) mustEmbedUnimplementedTransactionIndexerServer() {}

// UnsafeTransactionIndexerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionIndexerServer will
// result in compilation errors.
type UnsafeTransactionIndexerServer interface {
	mustEmbedUnimplementedTransactionIndexerServer()
}

func RegisterTransactionIndexerServer(s grpc.ServiceRegistrar, srv TransactionIndexerServer) {
	s.RegisterService(&TransactionIndexer_ServiceDesc, srv)
}

func _TransactionIndexer_GetTxsByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetTxsByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetTxsByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetTxsByAddress(ctx, req.(*GetTxsByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetTxsByHeightAndBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsByHeightAndBlockchainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetTxsByHeightAndBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetTxsByHeightAndBlockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetTxsByHeightAndBlockchain(ctx, req.(*GetTxsByHeightAndBlockchainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetBlocksByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksByHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetBlocksByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetBlocksByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetBlocksByHeight(ctx, req.(*GetBlocksByHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetBlockDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetBlockDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetBlockDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetBlockDetails(ctx, req.(*GetBlockDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetTransactionDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetTransactionDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetTransactionDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetTransactionDetails(ctx, req.(*GetTransactionDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetLatestBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetLatestBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetLatestBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetLatestBlocks(ctx, req.(*GetLatestBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetLatestTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetLatestTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetLatestTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetLatestTransactions(ctx, req.(*GetLatestTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetMoneyTransfersByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetMoneyTransfersByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetMoneyTransfersByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetMoneyTransfersByAddress(ctx, req.(*GetTxsByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetMoneyTransfersByAddressAndTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsByAddressAndTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetMoneyTransfersByAddressAndTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetMoneyTransfersByAddressAndTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetMoneyTransfersByAddressAndTime(ctx, req.(*GetTxsByAddressAndTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransactionIndexer_GetTxsByAddressAndTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxsByAddressAndTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionIndexerServer).GetTxsByAddressAndTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.TransactionIndexer/GetTxsByAddressAndTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionIndexerServer).GetTxsByAddressAndTime(ctx, req.(*GetTxsByAddressAndTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionIndexer_ServiceDesc is the grpc.ServiceDesc for TransactionIndexer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionIndexer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.clover.extractor.TransactionIndexer",
	HandlerType: (*TransactionIndexerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTxsByAddress",
			Handler:    _TransactionIndexer_GetTxsByAddress_Handler,
		},
		{
			MethodName: "GetTxsByHeightAndBlockchain",
			Handler:    _TransactionIndexer_GetTxsByHeightAndBlockchain_Handler,
		},
		{
			MethodName: "GetBlocksByHeight",
			Handler:    _TransactionIndexer_GetBlocksByHeight_Handler,
		},
		{
			MethodName: "GetBlockDetails",
			Handler:    _TransactionIndexer_GetBlockDetails_Handler,
		},
		{
			MethodName: "GetTransactionDetails",
			Handler:    _TransactionIndexer_GetTransactionDetails_Handler,
		},
		{
			MethodName: "GetLatestBlocks",
			Handler:    _TransactionIndexer_GetLatestBlocks_Handler,
		},
		{
			MethodName: "GetLatestTransactions",
			Handler:    _TransactionIndexer_GetLatestTransactions_Handler,
		},
		{
			MethodName: "GetMoneyTransfersByAddress",
			Handler:    _TransactionIndexer_GetMoneyTransfersByAddress_Handler,
		},
		{
			MethodName: "GetMoneyTransfersByAddressAndTime",
			Handler:    _TransactionIndexer_GetMoneyTransfersByAddressAndTime_Handler,
		},
		{
			MethodName: "GetTxsByAddressAndTime",
			Handler:    _TransactionIndexer_GetTxsByAddressAndTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction-indexer.proto",
}
