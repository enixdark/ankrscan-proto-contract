package proto

import (
	"crypto/md5"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	"time"
)

func (x *EthBlock) MixHashAsHash() common.Hash {
	return common.BytesToHash(x.MixHash)
}

func (x *EthBlock) ReceiptsRootAsHash() common.Hash {
	return common.BytesToHash(x.ReceiptsRoot)
}

func (x *EthBlock) UnclesAsHash() []common.Hash {
	uncles := make([]common.Hash, 0)
	for _, uncle := range x.Uncles {
		uncles = append(uncles, common.BytesToHash(uncle))
	}
	return uncles
}

func (x *EthBlock) Sha3UnclesAsHash() common.Hash {
	return common.BytesToHash(x.Sha3Uncles)

}
func (x *EthBlock) TransactionsRootAsHash() common.Hash {
	return common.BytesToHash(x.TransactionsRoot)

}
func (x *EthBlock) StateRootAsHash() common.Hash {
	return common.BytesToHash(x.StateRoot)

}
func (x *EthBlock) MinerAsAddress() common.Address {
	return common.BytesToAddress(x.Miner)
}

func (x *EthBlock) DifficultyAsInt() *big.Int {
	z := new(big.Int)
	z.SetBytes(x.Difficulty)
	return z
}

func (x *EthBlock) TotalDifficultyAsInt() *big.Int {
	z := new(big.Int)
	z.SetBytes(x.TotalDifficulty)
	return z
}

func (x *BlockHeader) BlockHashAsHash() common.Hash {
	return common.BytesToHash(x.BlockHash)
}

func (x *BlockHeader) ParentHashAsHash() common.Hash {
	return common.BytesToHash(x.ParentHash)
}

func (x *BlockHeader) TimestampAsDuration() time.Duration {
	return time.Second * time.Duration(x.Timestamp)
}

func (x *Transaction) BlockHashAsHash() common.Hash {
	return common.BytesToHash(x.BlockHash)
}

func (x *Transaction) TimestampAsDuration() time.Duration {
	return time.Second * time.Duration(x.Timestamp)
}

func (x *Transaction) TransactionHashAsHash() common.Hash {
	return common.BytesToHash(x.TransactionHash)
}

func (x *EthTransaction) FromAsAddress() common.Address {
	return common.BytesToAddress(x.From)
}

func (x *EthTransaction) ToAsAddress() (address common.Address, exists bool) {
	return common.BytesToAddress(x.To), len(x.To) > 0
}

func (x *EthTransaction) ToAsString() *string {
	toAddress, exists := x.ToAsAddress()
	if !exists {
		return nil
	} else {
		asString := strings.ToLower(toAddress.String())
		return &asString
	}
}

func (x *EthTransaction) ValueAsInt() *big.Int {
	z := new(big.Int)
	z.SetBytes(x.Value)
	return z
}

func (x *EthTransaction) VAsInt() *big.Int {
	z := new(big.Int)
	z.SetBytes(x.V)
	return z
}

func (x *EthTransaction) RAsInt() *big.Int {
	z := new(big.Int)
	z.SetBytes(x.R)
	return z
}

func (x *EthTransaction) SAsInt() *big.Int {
	z := new(big.Int)
	z.SetBytes(x.S)
	return z
}

func (x *EthTransaction) GasPriceAsInt() *big.Int {
	z := new(big.Int)
	z.SetBytes(x.GasPrice)
	return z
}

func (x *EthTransaction) ContractAddressAsAddress() (address common.Address, exists bool) {
	return common.BytesToAddress(x.ContractAddress), len(x.ContractAddress) > 0
}

func (x *EthTransaction) ContractAddressAsString() *string {
	toAddress, exists := x.ContractAddressAsAddress()
	if !exists {
		return nil
	} else {
		asString := strings.ToLower(toAddress.String())
		return &asString
	}
}

func (x *Transaction) EthLogs() []*types.Log {
	logs := make([]*types.Log, 0, 300)
	if x.GetEthTx() != nil {
		for i, protoLog := range x.GetEthTx().Logs {
			topics := make([]common.Hash, len(protoLog.Topics))

			for i, topic := range protoLog.Topics {
				topics[i] = common.BytesToHash(topic)
			}
			logs = append(logs, &types.Log{
				Address:     common.BytesToAddress(protoLog.Address),
				Topics:      topics,
				Data:        protoLog.Data,
				Removed:     protoLog.Removed,
				Index:       uint(i),
				TxIndex:     uint(x.TransactionIndex),
				TxHash:      common.BytesToHash(x.TransactionHash),
				BlockHash:   common.BytesToHash(x.BlockHash),
				BlockNumber: x.BlockHeight,
			})
		}
	}
	return logs
}

func (x *EthLog) TopicsAsHash() []common.Hash {
	topics := make([]common.Hash, 0)
	for _, topic := range x.Topics {
		topics = append(topics, common.BytesToHash(topic))
	}
	return topics
}

func (x *EthLog) AddressAsAddress() common.Address {
	return common.BytesToAddress(x.Address)
}

func (x *BlockConsumer) ConsumerId() string {
	hasher := md5.New()
	hasher.Write([]byte(x.UserId))
	return fmt.Sprintf("%s-%s-%x", x.BlockchainName, x.ConsumerName, hasher.Sum(nil))
}

func (x *BlockConsumer) FullName() string {
	return fmt.Sprintf("%s@%s", x.ConsumerName, x.BlockchainName)
}
