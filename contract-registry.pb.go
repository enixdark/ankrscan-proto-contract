// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: contract-registry.proto

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

type ContractStatus int32

const (
	ContractStatus_UNKNOWN  ContractStatus = 0
	ContractStatus_FOUND    ContractStatus = 1
	ContractStatus_SCANNED  ContractStatus = 2
	ContractStatus_VERIFIED ContractStatus = 3
)

// Enum value maps for ContractStatus.
var (
	ContractStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "FOUND",
		2: "SCANNED",
		3: "VERIFIED",
	}
	ContractStatus_value = map[string]int32{
		"UNKNOWN":  0,
		"FOUND":    1,
		"SCANNED":  2,
		"VERIFIED": 3,
	}
)

func (x ContractStatus) Enum() *ContractStatus {
	p := new(ContractStatus)
	*p = x
	return p
}

func (x ContractStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContractStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_contract_registry_proto_enumTypes[0].Descriptor()
}

func (ContractStatus) Type() protoreflect.EnumType {
	return &file_contract_registry_proto_enumTypes[0]
}

func (x ContractStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContractStatus.Descriptor instead.
func (ContractStatus) EnumDescriptor() ([]byte, []int) {
	return file_contract_registry_proto_rawDescGZIP(), []int{0}
}

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string         `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Address        string         `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name           string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Decimals       int32          `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Symbol         string         `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Status         ContractStatus `protobuf:"varint,6,opt,name=status,proto3,enum=com.clover.extractor.ContractStatus" json:"status,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_contract_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_contract_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Contract) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *Contract) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Contract) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Contract) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *Contract) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Contract) GetStatus() ContractStatus {
	if x != nil {
		return x.Status
	}
	return ContractStatus_UNKNOWN
}

type GetContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Address        string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetContractRequest) Reset() {
	*x = GetContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractRequest) ProtoMessage() {}

func (x *GetContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractRequest.ProtoReflect.Descriptor instead.
func (*GetContractRequest) Descriptor() ([]byte, []int) {
	return file_contract_registry_proto_rawDescGZIP(), []int{1}
}

func (x *GetContractRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *GetContractRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetContractReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contract *Contract `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *GetContractReply) Reset() {
	*x = GetContractReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContractReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContractReply) ProtoMessage() {}

func (x *GetContractReply) ProtoReflect() protoreflect.Message {
	mi := &file_contract_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContractReply.ProtoReflect.Descriptor instead.
func (*GetContractReply) Descriptor() ([]byte, []int) {
	return file_contract_registry_proto_rawDescGZIP(), []int{2}
}

func (x *GetContractReply) GetContract() *Contract {
	if x != nil {
		return x.Contract
	}
	return nil
}

var File_contract_registry_proto protoreflect.FileDescriptor

var file_contract_registry_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6c, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x22,
	0xd3, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c,
	0x6f, 0x76, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x4e,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x76, 0x65,
	0x72, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2a, 0x43,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x43, 0x41, 0x4e,
	0x4e, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x03, 0x32, 0x73, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x5f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f,
	0x76, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x2c, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0e, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_registry_proto_rawDescOnce sync.Once
	file_contract_registry_proto_rawDescData = file_contract_registry_proto_rawDesc
)

func file_contract_registry_proto_rawDescGZIP() []byte {
	file_contract_registry_proto_rawDescOnce.Do(func() {
		file_contract_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_registry_proto_rawDescData)
	})
	return file_contract_registry_proto_rawDescData
}

var file_contract_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_contract_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_contract_registry_proto_goTypes = []interface{}{
	(ContractStatus)(0),        // 0: com.clover.extractor.ContractStatus
	(*Contract)(nil),           // 1: com.clover.extractor.Contract
	(*GetContractRequest)(nil), // 2: com.clover.extractor.GetContractRequest
	(*GetContractReply)(nil),   // 3: com.clover.extractor.GetContractReply
}
var file_contract_registry_proto_depIdxs = []int32{
	0, // 0: com.clover.extractor.Contract.status:type_name -> com.clover.extractor.ContractStatus
	1, // 1: com.clover.extractor.GetContractReply.contract:type_name -> com.clover.extractor.Contract
	2, // 2: com.clover.extractor.ContractRegistry.GetContract:input_type -> com.clover.extractor.GetContractRequest
	3, // 3: com.clover.extractor.ContractRegistry.GetContract:output_type -> com.clover.extractor.GetContractReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contract_registry_proto_init() }
func file_contract_registry_proto_init() {
	if File_contract_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
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
		file_contract_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractRequest); i {
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
		file_contract_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContractReply); i {
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
			RawDescriptor: file_contract_registry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_registry_proto_goTypes,
		DependencyIndexes: file_contract_registry_proto_depIdxs,
		EnumInfos:         file_contract_registry_proto_enumTypes,
		MessageInfos:      file_contract_registry_proto_msgTypes,
	}.Build()
	File_contract_registry_proto = out.File
	file_contract_registry_proto_rawDesc = nil
	file_contract_registry_proto_goTypes = nil
	file_contract_registry_proto_depIdxs = nil
}
