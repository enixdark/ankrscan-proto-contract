// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: nft-indexer.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type NftType int32

const (
	NftType_UNDEFINED NftType = 0
	NftType_ERC721    NftType = 1
	NftType_ERC1155   NftType = 2
)

// Enum value maps for NftType.
var (
	NftType_name = map[int32]string{
		0: "UNDEFINED",
		1: "ERC721",
		2: "ERC1155",
	}
	NftType_value = map[string]int32{
		"UNDEFINED": 0,
		"ERC721":    1,
		"ERC1155":   2,
	}
)

func (x NftType) Enum() *NftType {
	p := new(NftType)
	*p = x
	return p
}

func (x NftType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NftType) Descriptor() protoreflect.EnumDescriptor {
	return file_nft_indexer_proto_enumTypes[0].Descriptor()
}

func (NftType) Type() protoreflect.EnumType {
	return &file_nft_indexer_proto_enumTypes[0]
}

func (x NftType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NftType.Descriptor instead.
func (NftType) EnumDescriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{0}
}

type GetBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet         string `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`                                       // wallet address e.g. 0x134acfa283884ae586f96e24d40dd6c894a620bb
	BlockchainName string `protobuf:"bytes,2,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"` // check https://github.com/Ankr-network/ankrscan-constants
	PageToken      string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`                // field to support pagination
	PageSize       int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *GetBalanceRequest) Reset() {
	*x = GetBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRequest) ProtoMessage() {}

func (x *GetBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetBalanceRequest) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *GetBalanceRequest) GetWallet() string {
	if x != nil {
		return x.Wallet
	}
	return ""
}

func (x *GetBalanceRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *GetBalanceRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *GetBalanceRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type NftBalanceByAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet         []byte `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`                                       // wallet address e.g. 0x134acfa283884ae586f96e24d40dd6c894a620bb
	BlockchainName string `protobuf:"bytes,2,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"` // check https://github.com/Ankr-network/ankrscan-constants
	PageToken      []byte `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`                // field to support pagination
	PageSize       int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *NftBalanceByAddressRequest) Reset() {
	*x = NftBalanceByAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftBalanceByAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftBalanceByAddressRequest) ProtoMessage() {}

func (x *NftBalanceByAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftBalanceByAddressRequest.ProtoReflect.Descriptor instead.
func (*NftBalanceByAddressRequest) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *NftBalanceByAddressRequest) GetWallet() []byte {
	if x != nil {
		return x.Wallet
	}
	return nil
}

func (x *NftBalanceByAddressRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *NftBalanceByAddressRequest) GetPageToken() []byte {
	if x != nil {
		return x.PageToken
	}
	return nil
}

func (x *NftBalanceByAddressRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type NftBalanceByAddressReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nfts          []*Nft `protobuf:"bytes,1,rep,name=nfts,proto3" json:"nfts,omitempty"`
	NextPageToken []byte `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"` // field to support pagination
}

func (x *NftBalanceByAddressReply) Reset() {
	*x = NftBalanceByAddressReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftBalanceByAddressReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftBalanceByAddressReply) ProtoMessage() {}

func (x *NftBalanceByAddressReply) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftBalanceByAddressReply.ProtoReflect.Descriptor instead.
func (*NftBalanceByAddressReply) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *NftBalanceByAddressReply) GetNfts() []*Nft {
	if x != nil {
		return x.Nfts
	}
	return nil
}

func (x *NftBalanceByAddressReply) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

type NftContractDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Symbol  string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	NftType NftType `protobuf:"varint,3,opt,name=nft_type,json=nftType,proto3,enum=ankrscan.nftindexer.NftType" json:"nft_type,omitempty"`
}

func (x *NftContractDetails) Reset() {
	*x = NftContractDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftContractDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftContractDetails) ProtoMessage() {}

func (x *NftContractDetails) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftContractDetails.ProtoReflect.Descriptor instead.
func (*NftContractDetails) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{3}
}

func (x *NftContractDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NftContractDetails) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *NftContractDetails) GetNftType() NftType {
	if x != nil {
		return x.NftType
	}
	return NftType_UNDEFINED
}

type NftAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenUri string `protobuf:"bytes,1,opt,name=token_uri,json=tokenUri,proto3" json:"token_uri,omitempty"`
}

func (x *NftAttributes) Reset() {
	*x = NftAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftAttributes) ProtoMessage() {}

func (x *NftAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftAttributes.ProtoReflect.Descriptor instead.
func (*NftAttributes) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{4}
}

func (x *NftAttributes) GetTokenUri() string {
	if x != nil {
		return x.TokenUri
	}
	return ""
}

type NftMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName  string  `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	ContractAddress []byte  `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	TokenId         []byte  `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	NftType         NftType `protobuf:"varint,4,opt,name=nft_type,json=nftType,proto3,enum=ankrscan.nftindexer.NftType" json:"nft_type,omitempty"`
}

func (x *NftMetadata) Reset() {
	*x = NftMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftMetadata) ProtoMessage() {}

func (x *NftMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftMetadata.ProtoReflect.Descriptor instead.
func (*NftMetadata) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{5}
}

func (x *NftMetadata) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *NftMetadata) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

func (x *NftMetadata) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *NftMetadata) GetNftType() NftType {
	if x != nil {
		return x.NftType
	}
	return NftType_UNDEFINED
}

type Nft struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName  string  `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Name            string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TokenId         []byte  `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	TokenUri        string  `protobuf:"bytes,4,opt,name=token_uri,json=tokenUri,proto3" json:"token_uri,omitempty"`
	CollectionName  string  `protobuf:"bytes,5,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Symbol          string  `protobuf:"bytes,6,opt,name=symbol,proto3" json:"symbol,omitempty"`
	NftType         NftType `protobuf:"varint,7,opt,name=nft_type,json=nftType,proto3,enum=ankrscan.nftindexer.NftType" json:"nft_type,omitempty"`
	ContractAddress []byte  `protobuf:"bytes,8,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (x *Nft) Reset() {
	*x = Nft{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nft) ProtoMessage() {}

func (x *Nft) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nft.ProtoReflect.Descriptor instead.
func (*Nft) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{6}
}

func (x *Nft) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *Nft) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Nft) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *Nft) GetTokenUri() string {
	if x != nil {
		return x.TokenUri
	}
	return ""
}

func (x *Nft) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *Nft) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Nft) GetNftType() NftType {
	if x != nil {
		return x.NftType
	}
	return NftType_UNDEFINED
}

func (x *Nft) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

type NftTransfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NftMetadata *NftMetadata `protobuf:"bytes,1,opt,name=nft_metadata,json=nftMetadata,proto3" json:"nft_metadata,omitempty"`
	From        []byte       `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To          []byte       `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value       []byte       `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NftTransfer) Reset() {
	*x = NftTransfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftTransfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftTransfer) ProtoMessage() {}

func (x *NftTransfer) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftTransfer.ProtoReflect.Descriptor instead.
func (*NftTransfer) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{7}
}

func (x *NftTransfer) GetNftMetadata() *NftMetadata {
	if x != nil {
		return x.NftMetadata
	}
	return nil
}

func (x *NftTransfer) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *NftTransfer) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *NftTransfer) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type NftBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NftMetadata *NftMetadata `protobuf:"bytes,1,opt,name=nft_metadata,json=nftMetadata,proto3" json:"nft_metadata,omitempty"`
	Wallet      []byte       `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Value       []byte       `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NftBalance) Reset() {
	*x = NftBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftBalance) ProtoMessage() {}

func (x *NftBalance) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftBalance.ProtoReflect.Descriptor instead.
func (*NftBalance) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{8}
}

func (x *NftBalance) GetNftMetadata() *NftMetadata {
	if x != nil {
		return x.NftMetadata
	}
	return nil
}

func (x *NftBalance) GetWallet() []byte {
	if x != nil {
		return x.Wallet
	}
	return nil
}

func (x *NftBalance) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type NftIndexerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BlockchainName       string `protobuf:"bytes,2,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	LatestProcessedBlock int64  `protobuf:"varint,3,opt,name=latest_processed_block,json=latestProcessedBlock,proto3" json:"latest_processed_block,omitempty"` // set to -1 to start from origin
	NodeUrl              string `protobuf:"bytes,8,opt,name=node_url,json=nodeUrl,proto3" json:"node_url,omitempty"`
}

func (x *NftIndexerConfig) Reset() {
	*x = NftIndexerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nft_indexer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftIndexerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftIndexerConfig) ProtoMessage() {}

func (x *NftIndexerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_nft_indexer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftIndexerConfig.ProtoReflect.Descriptor instead.
func (*NftIndexerConfig) Descriptor() ([]byte, []int) {
	return file_nft_indexer_proto_rawDescGZIP(), []int{9}
}

func (x *NftIndexerConfig) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NftIndexerConfig) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *NftIndexerConfig) GetLatestProcessedBlock() int64 {
	if x != nil {
		return x.LatestProcessedBlock
	}
	return 0
}

func (x *NftIndexerConfig) GetNodeUrl() string {
	if x != nil {
		return x.NodeUrl
	}
	return ""
}

var File_nft_indexer_proto protoreflect.FileDescriptor

var file_nft_indexer_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6e, 0x66, 0x74, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66,
	0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x1a, 0x4e, 0x66,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x70, 0x0a, 0x18, 0x4e, 0x66, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x66, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66, 0x74, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x4e, 0x66, 0x74, 0x52, 0x04, 0x6e, 0x66, 0x74, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x79, 0x0a, 0x12, 0x4e, 0x66, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x37, 0x0a, 0x08, 0x6e, 0x66, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x6e,
	0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x72, 0x2e, 0x4e, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x66, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x2c, 0x0a, 0x0d, 0x4e, 0x66, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x69,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x55, 0x72, 0x69,
	0x22, 0xb5, 0x01, 0x0a, 0x0b, 0x4e, 0x66, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12,
	0x37, 0x0a, 0x08, 0x6e, 0x66, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66, 0x74,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x4e, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x6e, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9f, 0x02, 0x0a, 0x03, 0x4e, 0x66, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x55, 0x72, 0x69, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x37, 0x0a, 0x08, 0x6e, 0x66, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73,
	0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x4e,
	0x66, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6e, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x4e,
	0x66, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0c, 0x6e, 0x66,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66, 0x74, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x4e, 0x66, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0b, 0x6e, 0x66, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x0a, 0x4e, 0x66, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x6e, 0x66, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66, 0x74, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x72, 0x2e, 0x4e, 0x66, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x0b, 0x6e, 0x66, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x10, 0x4e,
	0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x16, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x19,
	0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x55, 0x72, 0x6c, 0x2a, 0x31, 0x0a, 0x07, 0x4e, 0x66, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x52, 0x43, 0x37, 0x32, 0x31, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x45, 0x52, 0x43, 0x31, 0x31, 0x35, 0x35, 0x10, 0x02, 0x32, 0x88, 0x02, 0x0a,
	0x0a, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x12, 0x75, 0x0a, 0x13, 0x4e,
	0x66, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2f, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66,
	0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x4e, 0x66, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e,
	0x66, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x4e, 0x66, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x82, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x26, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66, 0x74,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x6e, 0x6b, 0x72,
	0x73, 0x63, 0x61, 0x6e, 0x2e, 0x6e, 0x66, 0x74, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e,
	0x4e, 0x66, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nft_indexer_proto_rawDescOnce sync.Once
	file_nft_indexer_proto_rawDescData = file_nft_indexer_proto_rawDesc
)

func file_nft_indexer_proto_rawDescGZIP() []byte {
	file_nft_indexer_proto_rawDescOnce.Do(func() {
		file_nft_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_nft_indexer_proto_rawDescData)
	})
	return file_nft_indexer_proto_rawDescData
}

var file_nft_indexer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_nft_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_nft_indexer_proto_goTypes = []interface{}{
	(NftType)(0),                       // 0: ankrscan.nftindexer.NftType
	(*GetBalanceRequest)(nil),          // 1: ankrscan.nftindexer.GetBalanceRequest
	(*NftBalanceByAddressRequest)(nil), // 2: ankrscan.nftindexer.NftBalanceByAddressRequest
	(*NftBalanceByAddressReply)(nil),   // 3: ankrscan.nftindexer.NftBalanceByAddressReply
	(*NftContractDetails)(nil),         // 4: ankrscan.nftindexer.NftContractDetails
	(*NftAttributes)(nil),              // 5: ankrscan.nftindexer.NftAttributes
	(*NftMetadata)(nil),                // 6: ankrscan.nftindexer.NftMetadata
	(*Nft)(nil),                        // 7: ankrscan.nftindexer.Nft
	(*NftTransfer)(nil),                // 8: ankrscan.nftindexer.NftTransfer
	(*NftBalance)(nil),                 // 9: ankrscan.nftindexer.NftBalance
	(*NftIndexerConfig)(nil),           // 10: ankrscan.nftindexer.NftIndexerConfig
}
var file_nft_indexer_proto_depIdxs = []int32{
	7, // 0: ankrscan.nftindexer.NftBalanceByAddressReply.nfts:type_name -> ankrscan.nftindexer.Nft
	0, // 1: ankrscan.nftindexer.NftContractDetails.nft_type:type_name -> ankrscan.nftindexer.NftType
	0, // 2: ankrscan.nftindexer.NftMetadata.nft_type:type_name -> ankrscan.nftindexer.NftType
	0, // 3: ankrscan.nftindexer.Nft.nft_type:type_name -> ankrscan.nftindexer.NftType
	6, // 4: ankrscan.nftindexer.NftTransfer.nft_metadata:type_name -> ankrscan.nftindexer.NftMetadata
	6, // 5: ankrscan.nftindexer.NftBalance.nft_metadata:type_name -> ankrscan.nftindexer.NftMetadata
	2, // 6: ankrscan.nftindexer.NftIndexer.NftBalanceByAddress:input_type -> ankrscan.nftindexer.NftBalanceByAddressRequest
	1, // 7: ankrscan.nftindexer.NftIndexer.GetBalance:input_type -> ankrscan.nftindexer.GetBalanceRequest
	3, // 8: ankrscan.nftindexer.NftIndexer.NftBalanceByAddress:output_type -> ankrscan.nftindexer.NftBalanceByAddressReply
	3, // 9: ankrscan.nftindexer.NftIndexer.GetBalance:output_type -> ankrscan.nftindexer.NftBalanceByAddressReply
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_nft_indexer_proto_init() }
func file_nft_indexer_proto_init() {
	if File_nft_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nft_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRequest); i {
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
		file_nft_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftBalanceByAddressRequest); i {
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
		file_nft_indexer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftBalanceByAddressReply); i {
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
		file_nft_indexer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftContractDetails); i {
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
		file_nft_indexer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftAttributes); i {
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
		file_nft_indexer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftMetadata); i {
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
		file_nft_indexer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nft); i {
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
		file_nft_indexer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftTransfer); i {
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
		file_nft_indexer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftBalance); i {
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
		file_nft_indexer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftIndexerConfig); i {
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
			RawDescriptor: file_nft_indexer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nft_indexer_proto_goTypes,
		DependencyIndexes: file_nft_indexer_proto_depIdxs,
		EnumInfos:         file_nft_indexer_proto_enumTypes,
		MessageInfos:      file_nft_indexer_proto_msgTypes,
	}.Build()
	File_nft_indexer_proto = out.File
	file_nft_indexer_proto_rawDesc = nil
	file_nft_indexer_proto_goTypes = nil
	file_nft_indexer_proto_depIdxs = nil
}
