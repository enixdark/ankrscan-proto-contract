// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: token-indexer.proto

package proto

import (
	"fmt"
)

func (x *TokenContractId) AsString() string {
	return fmt.Sprintf("%s/0x%x", x.BlockchainName, x.Address)
}