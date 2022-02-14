// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: foundationdb.proto

package proto

import (
	_ "github.com/golang/protobuf/ptypes/any"
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
	NamespaceFdb_UnknownFoundationDdNamespace NamespaceFdb = 0
	// ankrscan-extractor
	NamespaceFdb_BlockByNumberAndHash       NamespaceFdb = 1
	NamespaceFdb_BlockByNumber              NamespaceFdb = 2
	NamespaceFdb_TransactionByBlockAndIndex NamespaceFdb = 3
	NamespaceFdb_TransactionByHash          NamespaceFdb = 4
	NamespaceFdb_ExtractorConfig            NamespaceFdb = 5
	NamespaceFdb_ConsumerOffset             NamespaceFdb = 6
)

// Enum value maps for NamespaceFdb.
var (
	NamespaceFdb_name = map[int32]string{
		0: "UnknownFoundationDdNamespace",
		1: "BlockByNumberAndHash",
		2: "BlockByNumber",
		3: "TransactionByBlockAndIndex",
		4: "TransactionByHash",
		5: "ExtractorConfig",
		6: "ConsumerOffset",
	}
	NamespaceFdb_value = map[string]int32{
		"UnknownFoundationDdNamespace": 0,
		"BlockByNumberAndHash":         1,
		"BlockByNumber":                2,
		"TransactionByBlockAndIndex":   3,
		"TransactionByHash":            4,
		"ExtractorConfig":              5,
		"ConsumerOffset":               6,
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

type ValueFdb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value      []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ExtraParts uint32 `protobuf:"varint,2,opt,name=extra_parts,json=extraParts,proto3" json:"extra_parts,omitempty"` // if entity is too large then it split and written to several independent foundation db keys
}

func (x *ValueFdb) Reset() {
	*x = ValueFdb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foundationdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueFdb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueFdb) ProtoMessage() {}

func (x *ValueFdb) ProtoReflect() protoreflect.Message {
	mi := &file_foundationdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueFdb.ProtoReflect.Descriptor instead.
func (*ValueFdb) Descriptor() ([]byte, []int) {
	return file_foundationdb_proto_rawDescGZIP(), []int{0}
}

func (x *ValueFdb) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ValueFdb) GetExtraParts() uint32 {
	if x != nil {
		return x.ExtraParts
	}
	return 0
}

var File_foundationdb_proto protoreflect.FileDescriptor

var file_foundationdb_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x64, 0x62, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x50, 0x61, 0x72, 0x74, 0x73, 0x2a, 0xbd, 0x01, 0x0a, 0x0c, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x46, 0x64, 0x62, 0x12, 0x20, 0x0a, 0x1c, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x48, 0x61, 0x73,
	0x68, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x6e, 0x64, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x48, 0x61, 0x73, 0x68, 0x10, 0x04, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x10, 0x06, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_foundationdb_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_foundationdb_proto_goTypes = []interface{}{
	(NamespaceFdb)(0), // 0: ankrscan.namespace.NamespaceFdb
	(*ValueFdb)(nil),  // 1: ankrscan.namespace.ValueFdb
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
	if !protoimpl.UnsafeEnabled {
		file_foundationdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueFdb); i {
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
			RawDescriptor: file_foundationdb_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_foundationdb_proto_goTypes,
		DependencyIndexes: file_foundationdb_proto_depIdxs,
		EnumInfos:         file_foundationdb_proto_enumTypes,
		MessageInfos:      file_foundationdb_proto_msgTypes,
	}.Build()
	File_foundationdb_proto = out.File
	file_foundationdb_proto_rawDesc = nil
	file_foundationdb_proto_goTypes = nil
	file_foundationdb_proto_depIdxs = nil
}
