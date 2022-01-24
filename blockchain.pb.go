// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: blockchain.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//*
//  _____       _   _ _   _
// | ____|_ __ | |_(_) |_(_) ___  ___
// |  _| | '_ \| __| | __| |/ _ \/ __|
// | |___| | | | |_| | |_| |  __/\__ \
// |_____|_| |_|\__|_|\__|_|\___||___/
//
type ChainType int32

const (
	ChainType_CHAIN_TYPE_UNKNOWN  ChainType = 0
	ChainType_CHAIN_TYPE_ETHEREUM ChainType = 1
	ChainType_CHAIN_TYPE_POLKADOT ChainType = 2
)

// Enum value maps for ChainType.
var (
	ChainType_name = map[int32]string{
		0: "CHAIN_TYPE_UNKNOWN",
		1: "CHAIN_TYPE_ETHEREUM",
		2: "CHAIN_TYPE_POLKADOT",
	}
	ChainType_value = map[string]int32{
		"CHAIN_TYPE_UNKNOWN":  0,
		"CHAIN_TYPE_ETHEREUM": 1,
		"CHAIN_TYPE_POLKADOT": 2,
	}
)

func (x ChainType) Enum() *ChainType {
	p := new(ChainType)
	*p = x
	return p
}

func (x ChainType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChainType) Descriptor() protoreflect.EnumDescriptor {
	return file_blockchain_proto_enumTypes[0].Descriptor()
}

func (ChainType) Type() protoreflect.EnumType {
	return &file_blockchain_proto_enumTypes[0]
}

func (x ChainType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChainType.Descriptor instead.
func (ChainType) EnumDescriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{0}
}

type BlockchainProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string    `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	VerboseName    string    `protobuf:"bytes,2,opt,name=verbose_name,json=verboseName,proto3" json:"verbose_name,omitempty"`
	ChainType      ChainType `protobuf:"varint,4,opt,name=chain_type,json=chainType,proto3,enum=com.ankrscan.blockchain.ChainType" json:"chain_type,omitempty"`
	Symbol         string    `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimals       int64     `protobuf:"varint,7,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (x *BlockchainProperties) Reset() {
	*x = BlockchainProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blockchain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainProperties) ProtoMessage() {}

func (x *BlockchainProperties) ProtoReflect() protoreflect.Message {
	mi := &file_blockchain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainProperties.ProtoReflect.Descriptor instead.
func (*BlockchainProperties) Descriptor() ([]byte, []int) {
	return file_blockchain_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainProperties) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *BlockchainProperties) GetVerboseName() string {
	if x != nil {
		return x.VerboseName
	}
	return ""
}

func (x *BlockchainProperties) GetChainType() ChainType {
	if x != nil {
		return x.ChainType
	}
	return ChainType_CHAIN_TYPE_UNKNOWN
}

func (x *BlockchainProperties) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *BlockchainProperties) GetDecimals() int64 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

var File_blockchain_proto protoreflect.FileDescriptor

var file_blockchain_proto_rawDesc = []byte{
	0x0a, 0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0xd9, 0x01, 0x0a, 0x14,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x62, 0x6f, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x41, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73,
	0x63, 0x61, 0x6e, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x2a, 0x55, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13,
	0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x54, 0x48, 0x45, 0x52,
	0x45, 0x55, 0x4d, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x48, 0x41, 0x49, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x4b, 0x41, 0x44, 0x4f, 0x54, 0x10, 0x02, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_blockchain_proto_rawDescOnce sync.Once
	file_blockchain_proto_rawDescData = file_blockchain_proto_rawDesc
)

func file_blockchain_proto_rawDescGZIP() []byte {
	file_blockchain_proto_rawDescOnce.Do(func() {
		file_blockchain_proto_rawDescData = protoimpl.X.CompressGZIP(file_blockchain_proto_rawDescData)
	})
	return file_blockchain_proto_rawDescData
}

var file_blockchain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_blockchain_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_blockchain_proto_goTypes = []interface{}{
	(ChainType)(0),               // 0: com.ankrscan.blockchain.ChainType
	(*BlockchainProperties)(nil), // 1: com.ankrscan.blockchain.BlockchainProperties
}
var file_blockchain_proto_depIdxs = []int32{
	0, // 0: com.ankrscan.blockchain.BlockchainProperties.chain_type:type_name -> com.ankrscan.blockchain.ChainType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_blockchain_proto_init() }
func file_blockchain_proto_init() {
	if File_blockchain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blockchain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockchainProperties); i {
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
			RawDescriptor: file_blockchain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_blockchain_proto_goTypes,
		DependencyIndexes: file_blockchain_proto_depIdxs,
		EnumInfos:         file_blockchain_proto_enumTypes,
		MessageInfos:      file_blockchain_proto_msgTypes,
	}.Build()
	File_blockchain_proto = out.File
	file_blockchain_proto_rawDesc = nil
	file_blockchain_proto_goTypes = nil
	file_blockchain_proto_depIdxs = nil
}
