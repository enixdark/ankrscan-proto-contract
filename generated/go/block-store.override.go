package proto

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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

func (x *EthBlock) DifficultyAsInt() big.Int {
	z := new(big.Int)
	z.SetBytes(x.Difficulty)
	return *z
}

func (x *EthBlock) TotalDifficultyAsInt() big.Int {
	z := new(big.Int)
	z.SetBytes(x.TotalDifficulty)
	return *z
}
func (x *BlockHeader) BlockHashAsHash() common.Hash {
	return common.BytesToHash(x.BlockHash)
}
func (x *BlockHeader) ParentHashAsHash() common.Hash {
	return common.BytesToHash(x.ParentHash)
}
