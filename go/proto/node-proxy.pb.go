// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: node-proxy.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BlockByNumberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	BlockHeight    uint64 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (x *BlockByNumberRequest) Reset() {
	*x = BlockByNumberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockByNumberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockByNumberRequest) ProtoMessage() {}

func (x *BlockByNumberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockByNumberRequest.ProtoReflect.Descriptor instead.
func (*BlockByNumberRequest) Descriptor() ([]byte, []int) {
	return file_node_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *BlockByNumberRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *BlockByNumberRequest) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

type BlockByNumberReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *BlockByNumberReply) Reset() {
	*x = BlockByNumberReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockByNumberReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockByNumberReply) ProtoMessage() {}

func (x *BlockByNumberReply) ProtoReflect() protoreflect.Message {
	mi := &file_node_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockByNumberReply.ProtoReflect.Descriptor instead.
func (*BlockByNumberReply) Descriptor() ([]byte, []int) {
	return file_node_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *BlockByNumberReply) GetBlock() *Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type LatestBlockHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
}

func (x *LatestBlockHeaderRequest) Reset() {
	*x = LatestBlockHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatestBlockHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatestBlockHeaderRequest) ProtoMessage() {}

func (x *LatestBlockHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatestBlockHeaderRequest.ProtoReflect.Descriptor instead.
func (*LatestBlockHeaderRequest) Descriptor() ([]byte, []int) {
	return file_node_proxy_proto_rawDescGZIP(), []int{2}
}

func (x *LatestBlockHeaderRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

type LatestBlockHeaderReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *LatestBlockHeaderReply) Reset() {
	*x = LatestBlockHeaderReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LatestBlockHeaderReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LatestBlockHeaderReply) ProtoMessage() {}

func (x *LatestBlockHeaderReply) ProtoReflect() protoreflect.Message {
	mi := &file_node_proxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LatestBlockHeaderReply.ProtoReflect.Descriptor instead.
func (*LatestBlockHeaderReply) Descriptor() ([]byte, []int) {
	return file_node_proxy_proto_rawDescGZIP(), []int{3}
}

func (x *LatestBlockHeaderReply) GetHeader() *BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

var File_node_proxy_proto protoreflect.FileDescriptor

var file_node_proxy_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x6f, 0x64,
	0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x11, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2d, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x14, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x46, 0x0a,
	0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x43, 0x0a, 0x18, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x16, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x38, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x32, 0xdd,
	0x01, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x61, 0x0a, 0x0d,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x2e,
	0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63,
	0x61, 0x6e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x6d, 0x0a, 0x11, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e,
	0x6e, 0x6f, 0x64, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_node_proxy_proto_rawDescOnce sync.Once
	file_node_proxy_proto_rawDescData = file_node_proxy_proto_rawDesc
)

func file_node_proxy_proto_rawDescGZIP() []byte {
	file_node_proxy_proto_rawDescOnce.Do(func() {
		file_node_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proxy_proto_rawDescData)
	})
	return file_node_proxy_proto_rawDescData
}

var file_node_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_node_proxy_proto_goTypes = []interface{}{
	(*BlockByNumberRequest)(nil),     // 0: ankrscan.nodeproxy.BlockByNumberRequest
	(*BlockByNumberReply)(nil),       // 1: ankrscan.nodeproxy.BlockByNumberReply
	(*LatestBlockHeaderRequest)(nil), // 2: ankrscan.nodeproxy.LatestBlockHeaderRequest
	(*LatestBlockHeaderReply)(nil),   // 3: ankrscan.nodeproxy.LatestBlockHeaderReply
	(*Block)(nil),                    // 4: ankrscan.blockstore.Block
	(*BlockHeader)(nil),              // 5: ankrscan.blockstore.BlockHeader
}
var file_node_proxy_proto_depIdxs = []int32{
	4, // 0: ankrscan.nodeproxy.BlockByNumberReply.block:type_name -> ankrscan.blockstore.Block
	5, // 1: ankrscan.nodeproxy.LatestBlockHeaderReply.header:type_name -> ankrscan.blockstore.BlockHeader
	0, // 2: ankrscan.nodeproxy.NodeProxy.BlockByNumber:input_type -> ankrscan.nodeproxy.BlockByNumberRequest
	2, // 3: ankrscan.nodeproxy.NodeProxy.LatestBlockHeader:input_type -> ankrscan.nodeproxy.LatestBlockHeaderRequest
	1, // 4: ankrscan.nodeproxy.NodeProxy.BlockByNumber:output_type -> ankrscan.nodeproxy.BlockByNumberReply
	3, // 5: ankrscan.nodeproxy.NodeProxy.LatestBlockHeader:output_type -> ankrscan.nodeproxy.LatestBlockHeaderReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_node_proxy_proto_init() }
func file_node_proxy_proto_init() {
	if File_node_proxy_proto != nil {
		return
	}
	file_block_store_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_node_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockByNumberRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockByNumberReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_proxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatestBlockHeaderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_node_proxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LatestBlockHeaderReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proxy_proto_goTypes,
		DependencyIndexes: file_node_proxy_proto_depIdxs,
		MessageInfos:      file_node_proxy_proto_msgTypes,
	}.Build()
	File_node_proxy_proto = out.File
	file_node_proxy_proto_rawDesc = nil
	file_node_proxy_proto_goTypes = nil
	file_node_proxy_proto_depIdxs = nil
}
