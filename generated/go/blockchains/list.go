package blockchains

import (
	"github.com/Ankr-network/ankrscan-proto-contract/generated/go/proto"
	"github.com/pkg/errors"
)

func AllBlockchainNames() []string {
	names := make([]string, 0)
	for _, chain := range blockchains {
		names = append(names, chain.BlockchainName)
	}
	return names
}

func AllBlockchains() []*proto.BlockchainProperties {
	return blockchains
}

func Blockchain(name string) (*proto.BlockchainProperties, error) {
	var result *proto.BlockchainProperties
	for _, blockchain := range blockchains {
		if blockchain.BlockchainName == name {
			result = blockchain
		}
	}
	if result == nil {
		return nil, errors.Errorf("no blockchain with name %s", name)
	}
	return result, nil
}

var blockchains = []*proto.BlockchainProperties{
	{
		BlockchainName: "arbitrum",
		VerboseName:    "Arbitrum",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "ETH",
		Decimals:       18,
	},
	{
		BlockchainName: "avalanche",
		VerboseName:    "Avalanche C-chain",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "AVAX",
		Decimals:       18,
	},
	{
		BlockchainName: "bsc",
		VerboseName:    "Binance Smart Chain",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "BNB",
		Decimals:       18,
	},
	{
		BlockchainName: "eth",
		VerboseName:    "Ethereum",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "ETH",
		Decimals:       18,
	},
	{
		BlockchainName: "fantom",
		VerboseName:    "Fantom",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "FTM",
		Decimals:       18,
	},
	{
		BlockchainName: "polygon",
		VerboseName:    "Polygon",
		ChainType:      proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:         "MATIC",
		Decimals:       18,
	},
}
