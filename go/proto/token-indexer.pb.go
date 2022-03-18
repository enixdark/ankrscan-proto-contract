// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: token-indexer.proto

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

type TokenType int32

const (
	TokenType_UNDEFINED TokenType = 0
	TokenType_NATIVE    TokenType = 1
	TokenType_ERC20     TokenType = 2
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "UNDEFINED",
		1: "NATIVE",
		2: "ERC20",
	}
	TokenType_value = map[string]int32{
		"UNDEFINED": 0,
		"NATIVE":    1,
		"ERC20":     2,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_token_indexer_proto_enumTypes[0].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_token_indexer_proto_enumTypes[0]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{0}
}

type BalanceByHolderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet         []byte `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
	BlockchainName string `protobuf:"bytes,2,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	PageToken      []byte `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize       int32  `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *BalanceByHolderRequest) Reset() {
	*x = BalanceByHolderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceByHolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceByHolderRequest) ProtoMessage() {}

func (x *BalanceByHolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceByHolderRequest.ProtoReflect.Descriptor instead.
func (*BalanceByHolderRequest) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceByHolderRequest) GetWallet() []byte {
	if x != nil {
		return x.Wallet
	}
	return nil
}

func (x *BalanceByHolderRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *BalanceByHolderRequest) GetPageToken() []byte {
	if x != nil {
		return x.PageToken
	}
	return nil
}

func (x *BalanceByHolderRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type TokenBalancesDetailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balances      []*TokenBalanceDetailed `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty"`
	NextPageToken []byte                  `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *TokenBalancesDetailed) Reset() {
	*x = TokenBalancesDetailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBalancesDetailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBalancesDetailed) ProtoMessage() {}

func (x *TokenBalancesDetailed) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBalancesDetailed.ProtoReflect.Descriptor instead.
func (*TokenBalancesDetailed) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *TokenBalancesDetailed) GetBalances() []*TokenBalanceDetailed {
	if x != nil {
		return x.Balances
	}
	return nil
}

func (x *TokenBalancesDetailed) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

type TokenBalancesReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balances      []*TokenBalance `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances,omitempty"`
	NextPageToken []byte          `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *TokenBalancesReply) Reset() {
	*x = TokenBalancesReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBalancesReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBalancesReply) ProtoMessage() {}

func (x *TokenBalancesReply) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBalancesReply.ProtoReflect.Descriptor instead.
func (*TokenBalancesReply) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *TokenBalancesReply) GetBalances() []*TokenBalance {
	if x != nil {
		return x.Balances
	}
	return nil
}

func (x *TokenBalancesReply) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

type TokenBalanceDetailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName  string    `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	ContractAddress []byte    `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Name            string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Decimals        uint64    `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Symbol          string    `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	HolderAddress   []byte    `protobuf:"bytes,6,opt,name=holder_address,json=holderAddress,proto3" json:"holder_address,omitempty"`
	Balance         []byte    `protobuf:"bytes,7,opt,name=balance,proto3" json:"balance,omitempty"`
	Type            TokenType `protobuf:"varint,8,opt,name=type,proto3,enum=ankrscan.tokenindexer.TokenType" json:"type,omitempty"`
}

func (x *TokenBalanceDetailed) Reset() {
	*x = TokenBalanceDetailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBalanceDetailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBalanceDetailed) ProtoMessage() {}

func (x *TokenBalanceDetailed) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBalanceDetailed.ProtoReflect.Descriptor instead.
func (*TokenBalanceDetailed) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{3}
}

func (x *TokenBalanceDetailed) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *TokenBalanceDetailed) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

func (x *TokenBalanceDetailed) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenBalanceDetailed) GetDecimals() uint64 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TokenBalanceDetailed) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenBalanceDetailed) GetHolderAddress() []byte {
	if x != nil {
		return x.HolderAddress
	}
	return nil
}

func (x *TokenBalanceDetailed) GetBalance() []byte {
	if x != nil {
		return x.Balance
	}
	return nil
}

func (x *TokenBalanceDetailed) GetType() TokenType {
	if x != nil {
		return x.Type
	}
	return TokenType_UNDEFINED
}

type TokenContractDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Address        []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Decimals       uint64 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Symbol         string `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (x *TokenContractDetails) Reset() {
	*x = TokenContractDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenContractDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenContractDetails) ProtoMessage() {}

func (x *TokenContractDetails) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenContractDetails.ProtoReflect.Descriptor instead.
func (*TokenContractDetails) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{4}
}

func (x *TokenContractDetails) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *TokenContractDetails) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TokenContractDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenContractDetails) GetDecimals() uint64 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TokenContractDetails) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

type TokenBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName  string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	ContractAddress []byte `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	HolderAddress   []byte `protobuf:"bytes,3,opt,name=holder_address,json=holderAddress,proto3" json:"holder_address,omitempty"`
	Balance         []byte `protobuf:"bytes,4,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *TokenBalance) Reset() {
	*x = TokenBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBalance) ProtoMessage() {}

func (x *TokenBalance) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBalance.ProtoReflect.Descriptor instead.
func (*TokenBalance) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{5}
}

func (x *TokenBalance) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *TokenBalance) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

func (x *TokenBalance) GetHolderAddress() []byte {
	if x != nil {
		return x.HolderAddress
	}
	return nil
}

func (x *TokenBalance) GetBalance() []byte {
	if x != nil {
		return x.Balance
	}
	return nil
}

type TokenContractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Address        []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *TokenContractRequest) Reset() {
	*x = TokenContractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenContractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenContractRequest) ProtoMessage() {}

func (x *TokenContractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenContractRequest.ProtoReflect.Descriptor instead.
func (*TokenContractRequest) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{6}
}

func (x *TokenContractRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *TokenContractRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

type TokenBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Address        []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	BlockHeight    uint64 `protobuf:"varint,3,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (x *TokenBalanceRequest) Reset() {
	*x = TokenBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenBalanceRequest) ProtoMessage() {}

func (x *TokenBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenBalanceRequest.ProtoReflect.Descriptor instead.
func (*TokenBalanceRequest) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{7}
}

func (x *TokenBalanceRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *TokenBalanceRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TokenBalanceRequest) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

type TokenLiquidPoolDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Address        []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Decimals       uint64 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Symbol         string `protobuf:"bytes,5,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Token0Address  []byte `protobuf:"bytes,6,opt,name=token0_address,json=token0Address,proto3" json:"token0_address,omitempty"`
	Token1Address  []byte `protobuf:"bytes,7,opt,name=token1_address,json=token1Address,proto3" json:"token1_address,omitempty"`
}

func (x *TokenLiquidPoolDetails) Reset() {
	*x = TokenLiquidPoolDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_indexer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenLiquidPoolDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenLiquidPoolDetails) ProtoMessage() {}

func (x *TokenLiquidPoolDetails) ProtoReflect() protoreflect.Message {
	mi := &file_token_indexer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenLiquidPoolDetails.ProtoReflect.Descriptor instead.
func (*TokenLiquidPoolDetails) Descriptor() ([]byte, []int) {
	return file_token_indexer_proto_rawDescGZIP(), []int{8}
}

func (x *TokenLiquidPoolDetails) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *TokenLiquidPoolDetails) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TokenLiquidPoolDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TokenLiquidPoolDetails) GetDecimals() uint64 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *TokenLiquidPoolDetails) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *TokenLiquidPoolDetails) GetToken0Address() []byte {
	if x != nil {
		return x.Token0Address
	}
	return nil
}

func (x *TokenLiquidPoolDetails) GetToken1Address() []byte {
	if x != nil {
		return x.Token1Address
	}
	return nil
}

var File_token_indexer_proto protoreflect.FileDescriptor

var file_token_indexer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x22, 0x95, 0x01, 0x0a,
	0x16, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x15, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x47,
	0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x52, 0x08, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x7d, 0x0a, 0x12, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3f, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63,
	0x61, 0x6e, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa9,
	0x02, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x14, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0xa3,
	0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x68, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x59, 0x0a, 0x14, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x7b, 0x0a, 0x13, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xf1, 0x01, 0x0a,
	0x16, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x50, 0x6f, 0x6f, 0x6c,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x30, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x30, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x31, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x31, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2a, 0x31, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a,
	0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x43, 0x32,
	0x30, 0x10, 0x02, 0x32, 0x7e, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x72, 0x12, 0x6e, 0x0a, 0x0f, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79,
	0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61,
	0x6e, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x79, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_indexer_proto_rawDescOnce sync.Once
	file_token_indexer_proto_rawDescData = file_token_indexer_proto_rawDesc
)

func file_token_indexer_proto_rawDescGZIP() []byte {
	file_token_indexer_proto_rawDescOnce.Do(func() {
		file_token_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_indexer_proto_rawDescData)
	})
	return file_token_indexer_proto_rawDescData
}

var file_token_indexer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_token_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_token_indexer_proto_goTypes = []interface{}{
	(TokenType)(0),                 // 0: ankrscan.tokenindexer.TokenType
	(*BalanceByHolderRequest)(nil), // 1: ankrscan.tokenindexer.BalanceByHolderRequest
	(*TokenBalancesDetailed)(nil),  // 2: ankrscan.tokenindexer.TokenBalancesDetailed
	(*TokenBalancesReply)(nil),     // 3: ankrscan.tokenindexer.TokenBalancesReply
	(*TokenBalanceDetailed)(nil),   // 4: ankrscan.tokenindexer.TokenBalanceDetailed
	(*TokenContractDetails)(nil),   // 5: ankrscan.tokenindexer.TokenContractDetails
	(*TokenBalance)(nil),           // 6: ankrscan.tokenindexer.TokenBalance
	(*TokenContractRequest)(nil),   // 7: ankrscan.tokenindexer.TokenContractRequest
	(*TokenBalanceRequest)(nil),    // 8: ankrscan.tokenindexer.TokenBalanceRequest
	(*TokenLiquidPoolDetails)(nil), // 9: ankrscan.tokenindexer.TokenLiquidPoolDetails
}
var file_token_indexer_proto_depIdxs = []int32{
	4, // 0: ankrscan.tokenindexer.TokenBalancesDetailed.balances:type_name -> ankrscan.tokenindexer.TokenBalanceDetailed
	6, // 1: ankrscan.tokenindexer.TokenBalancesReply.balances:type_name -> ankrscan.tokenindexer.TokenBalance
	0, // 2: ankrscan.tokenindexer.TokenBalanceDetailed.type:type_name -> ankrscan.tokenindexer.TokenType
	1, // 3: ankrscan.tokenindexer.TokenIndexer.BalanceByHolder:input_type -> ankrscan.tokenindexer.BalanceByHolderRequest
	2, // 4: ankrscan.tokenindexer.TokenIndexer.BalanceByHolder:output_type -> ankrscan.tokenindexer.TokenBalancesDetailed
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_token_indexer_proto_init() }
func file_token_indexer_proto_init() {
	if File_token_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceByHolderRequest); i {
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
		file_token_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBalancesDetailed); i {
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
		file_token_indexer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBalancesReply); i {
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
		file_token_indexer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBalanceDetailed); i {
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
		file_token_indexer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenContractDetails); i {
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
		file_token_indexer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBalance); i {
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
		file_token_indexer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenContractRequest); i {
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
		file_token_indexer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenBalanceRequest); i {
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
		file_token_indexer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenLiquidPoolDetails); i {
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
			RawDescriptor: file_token_indexer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_indexer_proto_goTypes,
		DependencyIndexes: file_token_indexer_proto_depIdxs,
		EnumInfos:         file_token_indexer_proto_enumTypes,
		MessageInfos:      file_token_indexer_proto_msgTypes,
	}.Build()
	File_token_indexer_proto = out.File
	file_token_indexer_proto_rawDesc = nil
	file_token_indexer_proto_goTypes = nil
	file_token_indexer_proto_depIdxs = nil
}
