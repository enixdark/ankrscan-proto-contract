// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: foundationdb.proto

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

type NamespaceFdb int32

const (
	// ankrscan-extractor
	NamespaceFdb_TransactionByBlockAndIndex NamespaceFdb = 0
	NamespaceFdb_BlockByNumberAndHash       NamespaceFdb = 1
	NamespaceFdb_BlockByNumber              NamespaceFdb = 2
	NamespaceFdb_TransactionByHash          NamespaceFdb = 4
	NamespaceFdb_ExtractorConfig            NamespaceFdb = 5
	NamespaceFdb_ConsumerOffset             NamespaceFdb = 6
	NamespaceFdb_ConsumerDetails            NamespaceFdb = 16
	NamespaceFdb_ExtractorLatest            NamespaceFdb = 7
	// ankrscan-nft-indexer
	NamespaceFdb_NftBalances        NamespaceFdb = 8
	NamespaceFdb_NftContractDetails NamespaceFdb = 9
	NamespaceFdb_NftAttributes      NamespaceFdb = 10
	// ankrscan-node-proxy
	NamespaceFdb_NodeProxyNodesConfig NamespaceFdb = 11
	// ankrscan-token-indexer
	NamespaceFdb_TokenCurrencyDetails    NamespaceFdb = 12
	NamespaceFdb_TokenBalanceByHolder    NamespaceFdb = 13
	NamespaceFdb_TokenHoldersByContract  NamespaceFdb = 14
	NamespaceFdb_TokenWalletInteractions NamespaceFdb = 15
	NamespaceFdb_TokenLiquidPoolDetails  NamespaceFdb = 17
	NamespaceFdb_TokenUsdPriceUpdate     NamespaceFdb = 18
	NamespaceFdb_TokenCurrencyMetadata   NamespaceFdb = 19 // Thumbnails and etc.
	// ankrscan-transaction-indexer
	NamespaceFdb_SimpleTransactionByAddress NamespaceFdb = 30
	NamespaceFdb_LogByTimestamp             NamespaceFdb = 31
	NamespaceFdb_LogByContractAndTopics     NamespaceFdb = 32
	NamespaceFdb_EventsAmountByTopic        NamespaceFdb = 33
	NamespaceFdb_EventsAmount               NamespaceFdb = 34
	NamespaceFdb_TransactionAmount          NamespaceFdb = 35
)

// Enum value maps for NamespaceFdb.
var (
	NamespaceFdb_name = map[int32]string{
		0:  "TransactionByBlockAndIndex",
		1:  "BlockByNumberAndHash",
		2:  "BlockByNumber",
		4:  "TransactionByHash",
		5:  "ExtractorConfig",
		6:  "ConsumerOffset",
		16: "ConsumerDetails",
		7:  "ExtractorLatest",
		8:  "NftBalances",
		9:  "NftContractDetails",
		10: "NftAttributes",
		11: "NodeProxyNodesConfig",
		12: "TokenCurrencyDetails",
		13: "TokenBalanceByHolder",
		14: "TokenHoldersByContract",
		15: "TokenWalletInteractions",
		17: "TokenLiquidPoolDetails",
		18: "TokenUsdPriceUpdate",
		19: "TokenCurrencyMetadata",
		30: "SimpleTransactionByAddress",
		31: "LogByTimestamp",
		32: "LogByContractAndTopics",
		33: "EventsAmountByTopic",
		34: "EventsAmount",
		35: "TransactionAmount",
	}
	NamespaceFdb_value = map[string]int32{
		"TransactionByBlockAndIndex": 0,
		"BlockByNumberAndHash":       1,
		"BlockByNumber":              2,
		"TransactionByHash":          4,
		"ExtractorConfig":            5,
		"ConsumerOffset":             6,
		"ConsumerDetails":            16,
		"ExtractorLatest":            7,
		"NftBalances":                8,
		"NftContractDetails":         9,
		"NftAttributes":              10,
		"NodeProxyNodesConfig":       11,
		"TokenCurrencyDetails":       12,
		"TokenBalanceByHolder":       13,
		"TokenHoldersByContract":     14,
		"TokenWalletInteractions":    15,
		"TokenLiquidPoolDetails":     17,
		"TokenUsdPriceUpdate":        18,
		"TokenCurrencyMetadata":      19,
		"SimpleTransactionByAddress": 30,
		"LogByTimestamp":             31,
		"LogByContractAndTopics":     32,
		"EventsAmountByTopic":        33,
		"EventsAmount":               34,
		"TransactionAmount":          35,
	}
)

func (x NamespaceFdb) Enum() *NamespaceFdb {
	p := new(NamespaceFdb)
	*p = x
	return p
}

func (x NamespaceFdb) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NamespaceFdb) Descriptor() protoreflect.EnumDescriptor {
	return file_foundationdb_proto_enumTypes[0].Descriptor()
}

func (NamespaceFdb) Type() protoreflect.EnumType {
	return &file_foundationdb_proto_enumTypes[0]
}

func (x NamespaceFdb) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NamespaceFdb.Descriptor instead.
func (NamespaceFdb) EnumDescriptor() ([]byte, []int) {
	return file_foundationdb_proto_rawDescGZIP(), []int{0}
}

var File_foundationdb_proto protoreflect.FileDescriptor

var file_foundationdb_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2a, 0xea, 0x04, 0x0a, 0x0c, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x46, 0x64, 0x62, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41,
	0x6e, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x48, 0x61, 0x73,
	0x68, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x10, 0x04, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x10, 0x10, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x10, 0x07,
	0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x66, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x10,
	0x08, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x66, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14,
	0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x10, 0x0b, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x10, 0x0c,
	0x12, 0x18, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x42, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x10, 0x0d, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x10, 0x0e, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x10, 0x0f, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x50, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x10, 0x11, 0x12,
	0x17, 0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x73, 0x64, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x12, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x10, 0x13, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x10, 0x1e, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x42, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x10, 0x1f, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x6f, 0x67, 0x42, 0x79,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x6e, 0x64, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x10, 0x20, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x79, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x10, 0x21, 0x12, 0x10, 0x0a, 0x0c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x10, 0x22, 0x12, 0x15,
	0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x10, 0x23, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foundationdb_proto_rawDescOnce sync.Once
	file_foundationdb_proto_rawDescData = file_foundationdb_proto_rawDesc
)

func file_foundationdb_proto_rawDescGZIP() []byte {
	file_foundationdb_proto_rawDescOnce.Do(func() {
		file_foundationdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_foundationdb_proto_rawDescData)
	})
	return file_foundationdb_proto_rawDescData
}

var file_foundationdb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_foundationdb_proto_goTypes = []interface{}{
	(NamespaceFdb)(0), // 0: ankrscan.namespace.NamespaceFdb
}
var file_foundationdb_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_foundationdb_proto_init() }
func file_foundationdb_proto_init() {
	if File_foundationdb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_foundationdb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foundationdb_proto_goTypes,
		DependencyIndexes: file_foundationdb_proto_depIdxs,
		EnumInfos:         file_foundationdb_proto_enumTypes,
	}.Build()
	File_foundationdb_proto = out.File
	file_foundationdb_proto_rawDesc = nil
	file_foundationdb_proto_goTypes = nil
	file_foundationdb_proto_depIdxs = nil
}
