// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: health.proto

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

type MethodHealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MethodHealthRequest) Reset() {
	*x = MethodHealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodHealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodHealthRequest) ProtoMessage() {}

func (x *MethodHealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodHealthRequest.ProtoReflect.Descriptor instead.
func (*MethodHealthRequest) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{0}
}

type MethodHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Methods []*MethodHealthResponse_Method `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *MethodHealthResponse) Reset() {
	*x = MethodHealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodHealthResponse) ProtoMessage() {}

func (x *MethodHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodHealthResponse.ProtoReflect.Descriptor instead.
func (*MethodHealthResponse) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1}
}

func (x *MethodHealthResponse) GetMethods() []*MethodHealthResponse_Method {
	if x != nil {
		return x.Methods
	}
	return nil
}

type MethodHealthResponse_Method struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MethodName  string                                    `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	Blockchains []*MethodHealthResponse_Method_Blockchain `protobuf:"bytes,2,rep,name=blockchains,proto3" json:"blockchains,omitempty"`
}

func (x *MethodHealthResponse_Method) Reset() {
	*x = MethodHealthResponse_Method{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodHealthResponse_Method) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodHealthResponse_Method) ProtoMessage() {}

func (x *MethodHealthResponse_Method) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodHealthResponse_Method.ProtoReflect.Descriptor instead.
func (*MethodHealthResponse_Method) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MethodHealthResponse_Method) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *MethodHealthResponse_Method) GetBlockchains() []*MethodHealthResponse_Method_Blockchain {
	if x != nil {
		return x.Blockchains
	}
	return nil
}

type MethodHealthResponse_Method_Blockchain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Healthy        bool   `protobuf:"varint,2,opt,name=healthy,proto3" json:"healthy,omitempty"`
}

func (x *MethodHealthResponse_Method_Blockchain) Reset() {
	*x = MethodHealthResponse_Method_Blockchain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_health_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodHealthResponse_Method_Blockchain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodHealthResponse_Method_Blockchain) ProtoMessage() {}

func (x *MethodHealthResponse_Method_Blockchain) ProtoReflect() protoreflect.Message {
	mi := &file_health_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodHealthResponse_Method_Blockchain.ProtoReflect.Descriptor instead.
func (*MethodHealthResponse_Method_Blockchain) Descriptor() ([]byte, []int) {
	return file_health_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *MethodHealthResponse_Method_Blockchain) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *MethodHealthResponse_Method_Blockchain) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

var File_health_proto protoreflect.FileDescriptor

var file_health_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x22,
	0x15, 0x0a, 0x13, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb6, 0x02, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x07,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x1a, 0xd5, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73,
	0x63, 0x61, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x73, 0x1a, 0x4f,
	0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x0f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x32,
	0x64, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12,
	0x54, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x24, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73,
	0x63, 0x61, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_health_proto_rawDescOnce sync.Once
	file_health_proto_rawDescData = file_health_proto_rawDesc
)

func file_health_proto_rawDescGZIP() []byte {
	file_health_proto_rawDescOnce.Do(func() {
		file_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_health_proto_rawDescData)
	})
	return file_health_proto_rawDescData
}

var file_health_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_health_proto_goTypes = []interface{}{
	(*MethodHealthRequest)(nil),                    // 0: ankrscan.health.MethodHealthRequest
	(*MethodHealthResponse)(nil),                   // 1: ankrscan.health.MethodHealthResponse
	(*MethodHealthResponse_Method)(nil),            // 2: ankrscan.health.MethodHealthResponse.Method
	(*MethodHealthResponse_Method_Blockchain)(nil), // 3: ankrscan.health.MethodHealthResponse.Method.Blockchain
}
var file_health_proto_depIdxs = []int32{
	2, // 0: ankrscan.health.MethodHealthResponse.methods:type_name -> ankrscan.health.MethodHealthResponse.Method
	3, // 1: ankrscan.health.MethodHealthResponse.Method.blockchains:type_name -> ankrscan.health.MethodHealthResponse.Method.Blockchain
	0, // 2: ankrscan.health.MethodHealth.Check:input_type -> ankrscan.health.MethodHealthRequest
	1, // 3: ankrscan.health.MethodHealth.Check:output_type -> ankrscan.health.MethodHealthResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_health_proto_init() }
func file_health_proto_init() {
	if File_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodHealthRequest); i {
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
		file_health_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodHealthResponse); i {
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
		file_health_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodHealthResponse_Method); i {
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
		file_health_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodHealthResponse_Method_Blockchain); i {
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
			RawDescriptor: file_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_health_proto_goTypes,
		DependencyIndexes: file_health_proto_depIdxs,
		MessageInfos:      file_health_proto_msgTypes,
	}.Build()
	File_health_proto = out.File
	file_health_proto_rawDesc = nil
	file_health_proto_goTypes = nil
	file_health_proto_depIdxs = nil
}
