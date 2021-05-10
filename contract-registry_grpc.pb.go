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

// ContractRegistryClient is the client API for ContractRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContractRegistryClient interface {
	GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractReply, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceReply, error)
}

type contractRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewContractRegistryClient(cc grpc.ClientConnInterface) ContractRegistryClient {
	return &contractRegistryClient{cc}
}

func (c *contractRegistryClient) GetContract(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractReply, error) {
	out := new(GetContractReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.ContractRegistry/GetContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractRegistryClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceReply, error) {
	out := new(GetBalanceReply)
	err := c.cc.Invoke(ctx, "/com.clover.extractor.ContractRegistry/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContractRegistryServer is the server API for ContractRegistry service.
// All implementations must embed UnimplementedContractRegistryServer
// for forward compatibility
type ContractRegistryServer interface {
	GetContract(context.Context, *GetContractRequest) (*GetContractReply, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceReply, error)
	mustEmbedUnimplementedContractRegistryServer()
}

// UnimplementedContractRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedContractRegistryServer struct {
}

func (UnimplementedContractRegistryServer) GetContract(context.Context, *GetContractRequest) (*GetContractReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContract not implemented")
}
func (UnimplementedContractRegistryServer) GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedContractRegistryServer) mustEmbedUnimplementedContractRegistryServer() {}

// UnsafeContractRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContractRegistryServer will
// result in compilation errors.
type UnsafeContractRegistryServer interface {
	mustEmbedUnimplementedContractRegistryServer()
}

func RegisterContractRegistryServer(s grpc.ServiceRegistrar, srv ContractRegistryServer) {
	s.RegisterService(&ContractRegistry_ServiceDesc, srv)
}

func _ContractRegistry_GetContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractRegistryServer).GetContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.ContractRegistry/GetContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractRegistryServer).GetContract(ctx, req.(*GetContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContractRegistry_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContractRegistryServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.clover.extractor.ContractRegistry/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContractRegistryServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContractRegistry_ServiceDesc is the grpc.ServiceDesc for ContractRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContractRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.clover.extractor.ContractRegistry",
	HandlerType: (*ContractRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContract",
			Handler:    _ContractRegistry_GetContract_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _ContractRegistry_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract-registry.proto",
}
