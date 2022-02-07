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

// BlockStoreClient is the client API for BlockStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockStoreClient interface {
	// internal APIs
	GetExtractors(ctx context.Context, in *GetExtractorsRequest, opts ...grpc.CallOption) (*ExtractorConfigs, error)
	UpdateExtractors(ctx context.Context, in *UpdateExtractorsRequest, opts ...grpc.CallOption) (*ExtractorConfigs, error)
	DeleteExtractors(ctx context.Context, in *DeleteExtractorsRequest, opts ...grpc.CallOption) (*ExtractorConfigs, error)
	BlockRange(ctx context.Context, in *BlockRangeRequest, opts ...grpc.CallOption) (*BlockRangeReply, error)
	// public APIs
	Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextReply, error)
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitReply, error)
	Seek(ctx context.Context, in *SeekRequest, opts ...grpc.CallOption) (*SeekReply, error)
	LastCommit(ctx context.Context, in *LastCommitRequest, opts ...grpc.CallOption) (*LastCommitReply, error)
	BlocksByNumber(ctx context.Context, in *BlocksByNumberRequest, opts ...grpc.CallOption) (*BlocksByNumberReply, error)
	LatestBlockHeader(ctx context.Context, in *LatestBlockHeaderRequest, opts ...grpc.CallOption) (*LatestBlockHeaderReply, error)
	BlockRangeContinuous(ctx context.Context, in *BlockRangeRequest, opts ...grpc.CallOption) (*BlockRangeReply, error)
}

type blockStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockStoreClient(cc grpc.ClientConnInterface) BlockStoreClient {
	return &blockStoreClient{cc}
}

func (c *blockStoreClient) GetExtractors(ctx context.Context, in *GetExtractorsRequest, opts ...grpc.CallOption) (*ExtractorConfigs, error) {
	out := new(ExtractorConfigs)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/GetExtractors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) UpdateExtractors(ctx context.Context, in *UpdateExtractorsRequest, opts ...grpc.CallOption) (*ExtractorConfigs, error) {
	out := new(ExtractorConfigs)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/UpdateExtractors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) DeleteExtractors(ctx context.Context, in *DeleteExtractorsRequest, opts ...grpc.CallOption) (*ExtractorConfigs, error) {
	out := new(ExtractorConfigs)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/DeleteExtractors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) BlockRange(ctx context.Context, in *BlockRangeRequest, opts ...grpc.CallOption) (*BlockRangeReply, error) {
	out := new(BlockRangeReply)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/BlockRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) Next(ctx context.Context, in *NextRequest, opts ...grpc.CallOption) (*NextReply, error) {
	out := new(NextReply)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/Next", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitReply, error) {
	out := new(CommitReply)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) Seek(ctx context.Context, in *SeekRequest, opts ...grpc.CallOption) (*SeekReply, error) {
	out := new(SeekReply)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/Seek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) LastCommit(ctx context.Context, in *LastCommitRequest, opts ...grpc.CallOption) (*LastCommitReply, error) {
	out := new(LastCommitReply)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/LastCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) BlocksByNumber(ctx context.Context, in *BlocksByNumberRequest, opts ...grpc.CallOption) (*BlocksByNumberReply, error) {
	out := new(BlocksByNumberReply)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/BlocksByNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) LatestBlockHeader(ctx context.Context, in *LatestBlockHeaderRequest, opts ...grpc.CallOption) (*LatestBlockHeaderReply, error) {
	out := new(LatestBlockHeaderReply)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/LatestBlockHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockStoreClient) BlockRangeContinuous(ctx context.Context, in *BlockRangeRequest, opts ...grpc.CallOption) (*BlockRangeReply, error) {
	out := new(BlockRangeReply)
	err := c.cc.Invoke(ctx, "/ankrscan.blockstore.BlockStore/BlockRangeContinuous", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockStoreServer is the server API for BlockStore service.
// All implementations must embed UnimplementedBlockStoreServer
// for forward compatibility
type BlockStoreServer interface {
	// internal APIs
	GetExtractors(context.Context, *GetExtractorsRequest) (*ExtractorConfigs, error)
	UpdateExtractors(context.Context, *UpdateExtractorsRequest) (*ExtractorConfigs, error)
	DeleteExtractors(context.Context, *DeleteExtractorsRequest) (*ExtractorConfigs, error)
	BlockRange(context.Context, *BlockRangeRequest) (*BlockRangeReply, error)
	// public APIs
	Next(context.Context, *NextRequest) (*NextReply, error)
	Commit(context.Context, *CommitRequest) (*CommitReply, error)
	Seek(context.Context, *SeekRequest) (*SeekReply, error)
	LastCommit(context.Context, *LastCommitRequest) (*LastCommitReply, error)
	BlocksByNumber(context.Context, *BlocksByNumberRequest) (*BlocksByNumberReply, error)
	LatestBlockHeader(context.Context, *LatestBlockHeaderRequest) (*LatestBlockHeaderReply, error)
	BlockRangeContinuous(context.Context, *BlockRangeRequest) (*BlockRangeReply, error)
	mustEmbedUnimplementedBlockStoreServer()
}

// UnimplementedBlockStoreServer must be embedded to have forward compatible implementations.
type UnimplementedBlockStoreServer struct {
}

func (UnimplementedBlockStoreServer) GetExtractors(context.Context, *GetExtractorsRequest) (*ExtractorConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExtractors not implemented")
}
func (UnimplementedBlockStoreServer) UpdateExtractors(context.Context, *UpdateExtractorsRequest) (*ExtractorConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExtractors not implemented")
}
func (UnimplementedBlockStoreServer) DeleteExtractors(context.Context, *DeleteExtractorsRequest) (*ExtractorConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExtractors not implemented")
}
func (UnimplementedBlockStoreServer) BlockRange(context.Context, *BlockRangeRequest) (*BlockRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockRange not implemented")
}
func (UnimplementedBlockStoreServer) Next(context.Context, *NextRequest) (*NextReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedBlockStoreServer) Commit(context.Context, *CommitRequest) (*CommitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedBlockStoreServer) Seek(context.Context, *SeekRequest) (*SeekReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seek not implemented")
}
func (UnimplementedBlockStoreServer) LastCommit(context.Context, *LastCommitRequest) (*LastCommitReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastCommit not implemented")
}
func (UnimplementedBlockStoreServer) BlocksByNumber(context.Context, *BlocksByNumberRequest) (*BlocksByNumberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlocksByNumber not implemented")
}
func (UnimplementedBlockStoreServer) LatestBlockHeader(context.Context, *LatestBlockHeaderRequest) (*LatestBlockHeaderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestBlockHeader not implemented")
}
func (UnimplementedBlockStoreServer) BlockRangeContinuous(context.Context, *BlockRangeRequest) (*BlockRangeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockRangeContinuous not implemented")
}
func (UnimplementedBlockStoreServer) mustEmbedUnimplementedBlockStoreServer() {}

// UnsafeBlockStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockStoreServer will
// result in compilation errors.
type UnsafeBlockStoreServer interface {
	mustEmbedUnimplementedBlockStoreServer()
}

func RegisterBlockStoreServer(s grpc.ServiceRegistrar, srv BlockStoreServer) {
	s.RegisterService(&BlockStore_ServiceDesc, srv)
}

func _BlockStore_GetExtractors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExtractorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).GetExtractors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/GetExtractors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).GetExtractors(ctx, req.(*GetExtractorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_UpdateExtractors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExtractorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).UpdateExtractors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/UpdateExtractors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).UpdateExtractors(ctx, req.(*UpdateExtractorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_DeleteExtractors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExtractorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).DeleteExtractors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/DeleteExtractors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).DeleteExtractors(ctx, req.(*DeleteExtractorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_BlockRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).BlockRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/BlockRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).BlockRange(ctx, req.(*BlockRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/Next",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).Next(ctx, req.(*NextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).Commit(ctx, req.(*CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_Seek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeekRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).Seek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/Seek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).Seek(ctx, req.(*SeekRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_LastCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LastCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).LastCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/LastCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).LastCommit(ctx, req.(*LastCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_BlocksByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlocksByNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).BlocksByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/BlocksByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).BlocksByNumber(ctx, req.(*BlocksByNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_LatestBlockHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestBlockHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).LatestBlockHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/LatestBlockHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).LatestBlockHeader(ctx, req.(*LatestBlockHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockStore_BlockRangeContinuous_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockStoreServer).BlockRangeContinuous(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ankrscan.blockstore.BlockStore/BlockRangeContinuous",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockStoreServer).BlockRangeContinuous(ctx, req.(*BlockRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockStore_ServiceDesc is the grpc.ServiceDesc for BlockStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ankrscan.blockstore.BlockStore",
	HandlerType: (*BlockStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetExtractors",
			Handler:    _BlockStore_GetExtractors_Handler,
		},
		{
			MethodName: "UpdateExtractors",
			Handler:    _BlockStore_UpdateExtractors_Handler,
		},
		{
			MethodName: "DeleteExtractors",
			Handler:    _BlockStore_DeleteExtractors_Handler,
		},
		{
			MethodName: "BlockRange",
			Handler:    _BlockStore_BlockRange_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _BlockStore_Next_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _BlockStore_Commit_Handler,
		},
		{
			MethodName: "Seek",
			Handler:    _BlockStore_Seek_Handler,
		},
		{
			MethodName: "LastCommit",
			Handler:    _BlockStore_LastCommit_Handler,
		},
		{
			MethodName: "BlocksByNumber",
			Handler:    _BlockStore_BlocksByNumber_Handler,
		},
		{
			MethodName: "LatestBlockHeader",
			Handler:    _BlockStore_LatestBlockHeader_Handler,
		},
		{
			MethodName: "BlockRangeContinuous",
			Handler:    _BlockStore_BlockRangeContinuous_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "block-store.proto",
}
