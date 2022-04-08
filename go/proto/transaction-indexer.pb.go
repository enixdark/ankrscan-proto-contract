// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: transaction-indexer.proto

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

// Transaction of each type is stored separately and accessed separately
type SimpleTransactionType int32

const (
	SimpleTransactionType_REGULAR_TRANSACTION SimpleTransactionType = 0
	SimpleTransactionType_TRANSFER_TOKEN      SimpleTransactionType = 1
)

// Enum value maps for SimpleTransactionType.
var (
	SimpleTransactionType_name = map[int32]string{
		0: "REGULAR_TRANSACTION",
		1: "TRANSFER_TOKEN",
	}
	SimpleTransactionType_value = map[string]int32{
		"REGULAR_TRANSACTION": 0,
		"TRANSFER_TOKEN":      1,
	}
)

func (x SimpleTransactionType) Enum() *SimpleTransactionType {
	p := new(SimpleTransactionType)
	*p = x
	return p
}

func (x SimpleTransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SimpleTransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_transaction_indexer_proto_enumTypes[0].Descriptor()
}

func (SimpleTransactionType) Type() protoreflect.EnumType {
	return &file_transaction_indexer_proto_enumTypes[0]
}

func (x SimpleTransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SimpleTransactionType.Descriptor instead.
func (SimpleTransactionType) EnumDescriptor() ([]byte, []int) {
	return file_transaction_indexer_proto_rawDescGZIP(), []int{0}
}

type TransactionStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
}

func (x *TransactionStatsRequest) Reset() {
	*x = TransactionStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStatsRequest) ProtoMessage() {}

func (x *TransactionStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStatsRequest.ProtoReflect.Descriptor instead.
func (*TransactionStatsRequest) Descriptor() ([]byte, []int) {
	return file_transaction_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionStatsRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

type TransactionsStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionsAmount uint64 `protobuf:"varint,1,opt,name=transactions_amount,json=transactionsAmount,proto3" json:"transactions_amount,omitempty"`
	EventsAmount       uint64 `protobuf:"varint,2,opt,name=events_amount,json=eventsAmount,proto3" json:"events_amount,omitempty"`
}

func (x *TransactionsStats) Reset() {
	*x = TransactionsStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionsStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionsStats) ProtoMessage() {}

func (x *TransactionsStats) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionsStats.ProtoReflect.Descriptor instead.
func (*TransactionsStats) Descriptor() ([]byte, []int) {
	return file_transaction_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionsStats) GetTransactionsAmount() uint64 {
	if x != nil {
		return x.TransactionsAmount
	}
	return 0
}

func (x *TransactionsStats) GetEventsAmount() uint64 {
	if x != nil {
		return x.EventsAmount
	}
	return 0
}

type SimpleTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName    string                `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	TransactionHash   []byte                `protobuf:"bytes,2,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	TransactionId     []byte                `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	BlockHeight       uint64                `protobuf:"varint,4,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	BlockHash         []byte                `protobuf:"bytes,5,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	FromAddress       []byte                `protobuf:"bytes,6,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ContractAddress   []byte                `protobuf:"bytes,7,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	ToAddress         []byte                `protobuf:"bytes,8,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Value             []byte                `protobuf:"bytes,9,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp         uint64                `protobuf:"varint,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Success           bool                  `protobuf:"varint,11,opt,name=success,proto3" json:"success,omitempty"`
	ActionDescription string                `protobuf:"bytes,12,opt,name=action_description,json=actionDescription,proto3" json:"action_description,omitempty"`
	Type              SimpleTransactionType `protobuf:"varint,13,opt,name=type,proto3,enum=ankrscan.transactionindexer.SimpleTransactionType" json:"type,omitempty"`
	TransactionIndex  uint64                `protobuf:"varint,14,opt,name=transaction_index,json=transactionIndex,proto3" json:"transaction_index,omitempty"`
	LogIndex          uint64                `protobuf:"varint,15,opt,name=log_index,json=logIndex,proto3" json:"log_index,omitempty"`
	IsLog             bool                  `protobuf:"varint,16,opt,name=is_log,json=isLog,proto3" json:"is_log,omitempty"`
}

func (x *SimpleTransaction) Reset() {
	*x = SimpleTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_indexer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleTransaction) ProtoMessage() {}

func (x *SimpleTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_indexer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleTransaction.ProtoReflect.Descriptor instead.
func (*SimpleTransaction) Descriptor() ([]byte, []int) {
	return file_transaction_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *SimpleTransaction) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *SimpleTransaction) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

func (x *SimpleTransaction) GetTransactionId() []byte {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

func (x *SimpleTransaction) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *SimpleTransaction) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *SimpleTransaction) GetFromAddress() []byte {
	if x != nil {
		return x.FromAddress
	}
	return nil
}

func (x *SimpleTransaction) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

func (x *SimpleTransaction) GetToAddress() []byte {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *SimpleTransaction) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SimpleTransaction) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *SimpleTransaction) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SimpleTransaction) GetActionDescription() string {
	if x != nil {
		return x.ActionDescription
	}
	return ""
}

func (x *SimpleTransaction) GetType() SimpleTransactionType {
	if x != nil {
		return x.Type
	}
	return SimpleTransactionType_REGULAR_TRANSACTION
}

func (x *SimpleTransaction) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *SimpleTransaction) GetLogIndex() uint64 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *SimpleTransaction) GetIsLog() bool {
	if x != nil {
		return x.IsLog
	}
	return false
}

type TransactionByTypeAndAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName string                `protobuf:"bytes,1,opt,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Address        []byte                `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Type           SimpleTransactionType `protobuf:"varint,3,opt,name=type,proto3,enum=ankrscan.transactionindexer.SimpleTransactionType" json:"type,omitempty"`
	Timestamp      uint64                `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	OrderAsc       bool                  `protobuf:"varint,5,opt,name=order_asc,json=orderAsc,proto3" json:"order_asc,omitempty"`
	PageToken      []byte                `protobuf:"bytes,6,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize       int32                 `protobuf:"varint,7,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *TransactionByTypeAndAddressRequest) Reset() {
	*x = TransactionByTypeAndAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_indexer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionByTypeAndAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionByTypeAndAddressRequest) ProtoMessage() {}

func (x *TransactionByTypeAndAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_indexer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionByTypeAndAddressRequest.ProtoReflect.Descriptor instead.
func (*TransactionByTypeAndAddressRequest) Descriptor() ([]byte, []int) {
	return file_transaction_indexer_proto_rawDescGZIP(), []int{3}
}

func (x *TransactionByTypeAndAddressRequest) GetBlockchainName() string {
	if x != nil {
		return x.BlockchainName
	}
	return ""
}

func (x *TransactionByTypeAndAddressRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TransactionByTypeAndAddressRequest) GetType() SimpleTransactionType {
	if x != nil {
		return x.Type
	}
	return SimpleTransactionType_REGULAR_TRANSACTION
}

func (x *TransactionByTypeAndAddressRequest) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *TransactionByTypeAndAddressRequest) GetOrderAsc() bool {
	if x != nil {
		return x.OrderAsc
	}
	return false
}

func (x *TransactionByTypeAndAddressRequest) GetPageToken() []byte {
	if x != nil {
		return x.PageToken
	}
	return nil
}

func (x *TransactionByTypeAndAddressRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type SimpleTransactionsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions  []*SimpleTransaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	NextPageToken []byte               `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *SimpleTransactionsReply) Reset() {
	*x = SimpleTransactionsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_indexer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleTransactionsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleTransactionsReply) ProtoMessage() {}

func (x *SimpleTransactionsReply) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_indexer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleTransactionsReply.ProtoReflect.Descriptor instead.
func (*SimpleTransactionsReply) Descriptor() ([]byte, []int) {
	return file_transaction_indexer_proto_rawDescGZIP(), []int{4}
}

func (x *SimpleTransactionsReply) GetTransactions() []*SimpleTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *SimpleTransactionsReply) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

type TransactionByTypeAndAddressAndBlockchainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockchainName []string              `protobuf:"bytes,1,rep,name=blockchain_name,json=blockchainName,proto3" json:"blockchain_name,omitempty"`
	Address        []byte                `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Type           SimpleTransactionType `protobuf:"varint,3,opt,name=type,proto3,enum=ankrscan.transactionindexer.SimpleTransactionType" json:"type,omitempty"`
	PageToken      []byte                `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	PageSize       int32                 `protobuf:"varint,5,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *TransactionByTypeAndAddressAndBlockchainRequest) Reset() {
	*x = TransactionByTypeAndAddressAndBlockchainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_indexer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionByTypeAndAddressAndBlockchainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionByTypeAndAddressAndBlockchainRequest) ProtoMessage() {}

func (x *TransactionByTypeAndAddressAndBlockchainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_indexer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionByTypeAndAddressAndBlockchainRequest.ProtoReflect.Descriptor instead.
func (*TransactionByTypeAndAddressAndBlockchainRequest) Descriptor() ([]byte, []int) {
	return file_transaction_indexer_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionByTypeAndAddressAndBlockchainRequest) GetBlockchainName() []string {
	if x != nil {
		return x.BlockchainName
	}
	return nil
}

func (x *TransactionByTypeAndAddressAndBlockchainRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *TransactionByTypeAndAddressAndBlockchainRequest) GetType() SimpleTransactionType {
	if x != nil {
		return x.Type
	}
	return SimpleTransactionType_REGULAR_TRANSACTION
}

func (x *TransactionByTypeAndAddressAndBlockchainRequest) GetPageToken() []byte {
	if x != nil {
		return x.PageToken
	}
	return nil
}

func (x *TransactionByTypeAndAddressAndBlockchainRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

var File_transaction_indexer_proto protoreflect.FileDescriptor

var file_transaction_indexer_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x6e, 0x6b,
	0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x22, 0x42, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a, 0x11,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x2f, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x12,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xe3, 0x04, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x74, 0x6f,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x32, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6c, 0x6f, 0x67,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4c, 0x6f, 0x67, 0x22, 0xa6, 0x02,
	0x0a, 0x22, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x17, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x52, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73,
	0x63, 0x61, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xf8,
	0x01, 0x0a, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6e,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x46, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x72, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x2a, 0x44, 0x0a, 0x15, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x47, 0x55, 0x4c, 0x41, 0x52, 0x5f, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x32,
	0x8e, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x12, 0x78, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x34, 0x2e, 0x61, 0x6e, 0x6b,
	0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2e, 0x2e, 0x61, 0x6e, 0x6b, 0x72, 0x73, 0x63, 0x61, 0x6e, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_indexer_proto_rawDescOnce sync.Once
	file_transaction_indexer_proto_rawDescData = file_transaction_indexer_proto_rawDesc
)

func file_transaction_indexer_proto_rawDescGZIP() []byte {
	file_transaction_indexer_proto_rawDescOnce.Do(func() {
		file_transaction_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_indexer_proto_rawDescData)
	})
	return file_transaction_indexer_proto_rawDescData
}

var file_transaction_indexer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transaction_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_transaction_indexer_proto_goTypes = []interface{}{
	(SimpleTransactionType)(0),                              // 0: ankrscan.transactionindexer.SimpleTransactionType
	(*TransactionStatsRequest)(nil),                         // 1: ankrscan.transactionindexer.TransactionStatsRequest
	(*TransactionsStats)(nil),                               // 2: ankrscan.transactionindexer.TransactionsStats
	(*SimpleTransaction)(nil),                               // 3: ankrscan.transactionindexer.SimpleTransaction
	(*TransactionByTypeAndAddressRequest)(nil),              // 4: ankrscan.transactionindexer.TransactionByTypeAndAddressRequest
	(*SimpleTransactionsReply)(nil),                         // 5: ankrscan.transactionindexer.SimpleTransactionsReply
	(*TransactionByTypeAndAddressAndBlockchainRequest)(nil), // 6: ankrscan.transactionindexer.TransactionByTypeAndAddressAndBlockchainRequest
}
var file_transaction_indexer_proto_depIdxs = []int32{
	0, // 0: ankrscan.transactionindexer.SimpleTransaction.type:type_name -> ankrscan.transactionindexer.SimpleTransactionType
	0, // 1: ankrscan.transactionindexer.TransactionByTypeAndAddressRequest.type:type_name -> ankrscan.transactionindexer.SimpleTransactionType
	3, // 2: ankrscan.transactionindexer.SimpleTransactionsReply.transactions:type_name -> ankrscan.transactionindexer.SimpleTransaction
	0, // 3: ankrscan.transactionindexer.TransactionByTypeAndAddressAndBlockchainRequest.type:type_name -> ankrscan.transactionindexer.SimpleTransactionType
	1, // 4: ankrscan.transactionindexer.TransactionIndexer.TransactionStats:input_type -> ankrscan.transactionindexer.TransactionStatsRequest
	2, // 5: ankrscan.transactionindexer.TransactionIndexer.TransactionStats:output_type -> ankrscan.transactionindexer.TransactionsStats
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_transaction_indexer_proto_init() }
func file_transaction_indexer_proto_init() {
	if File_transaction_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionStatsRequest); i {
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
		file_transaction_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionsStats); i {
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
		file_transaction_indexer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleTransaction); i {
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
		file_transaction_indexer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionByTypeAndAddressRequest); i {
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
		file_transaction_indexer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleTransactionsReply); i {
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
		file_transaction_indexer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionByTypeAndAddressAndBlockchainRequest); i {
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
			RawDescriptor: file_transaction_indexer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transaction_indexer_proto_goTypes,
		DependencyIndexes: file_transaction_indexer_proto_depIdxs,
		EnumInfos:         file_transaction_indexer_proto_enumTypes,
		MessageInfos:      file_transaction_indexer_proto_msgTypes,
	}.Build()
	File_transaction_indexer_proto = out.File
	file_transaction_indexer_proto_rawDesc = nil
	file_transaction_indexer_proto_goTypes = nil
	file_transaction_indexer_proto_depIdxs = nil
}