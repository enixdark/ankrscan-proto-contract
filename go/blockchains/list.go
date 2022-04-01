package blockchains

import "ankr/ankrscan-proto-contract/go/proto"

var blockchains = []*proto.BlockchainProperties{
	{
		BlockchainName:      "arbitrum",
		VerboseName:         "Arbitrum",
		ChainType:           proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:              "ETH",
		Decimals:            18,
		BlockchainNumber:    proto.BlockchainNumber_BLOCKCHAIN_ARBITRUM,
		BlockConfirmedAfter: 12,
		ThumbnailUrl:        "",
	},
	{
		BlockchainName:      "avalanche",
		VerboseName:         "Avalanche C-chain",
		ChainType:           proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:              "AVAX",
		Decimals:            18,
		BlockchainNumber:    proto.BlockchainNumber_BLOCKCHAIN_AVALANCHE,
		BlockConfirmedAfter: 12,
		ThumbnailUrl:        "https://assets.coingecko.com/coins/images/12559/large/coin-round-red.png",
	},
	{
		BlockchainName:      "bsc",
		VerboseName:         "Binance Smart Chain",
		ChainType:           proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:              "BNB",
		Decimals:            18,
		BlockchainNumber:    proto.BlockchainNumber_BLOCKCHAIN_BSC,
		BlockConfirmedAfter: 12,
		ThumbnailUrl:        "https://assets.coingecko.com/coins/images/825/large/bnb-icon2_2x.png",
	},
	{
		BlockchainName:      "eth",
		VerboseName:         "Ethereum",
		ChainType:           proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:              "ETH",
		Decimals:            18,
		BlockchainNumber:    proto.BlockchainNumber_BLOCKCHAIN_ETH,
		BlockConfirmedAfter: 12,
		ThumbnailUrl:        "https://assets.coingecko.com/coins/images/279/large/ethereum.png",
	},
	{
		BlockchainName:      "fantom",
		VerboseName:         "Fantom",
		ChainType:           proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:              "FTM",
		Decimals:            18,
		BlockchainNumber:    proto.BlockchainNumber_BLOCKCHAIN_FANTOM,
		BlockConfirmedAfter: 12,
		ThumbnailUrl:        "https://assets.coingecko.com/coins/images/4001/large/Fantom.png",
	},
	{
		BlockchainName:      "polygon",
		VerboseName:         "Polygon",
		ChainType:           proto.ChainType_CHAIN_TYPE_ETHEREUM,
		Symbol:              "MATIC",
		Decimals:            18,
		BlockchainNumber:    proto.BlockchainNumber_BLOCKCHAIN_POLYGON,
		BlockConfirmedAfter: 12,
		ThumbnailUrl:        "https://assets.coingecko.com/coins/images/4713/large/matic-token-icon.png",
	},
}

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

func Blockchain(name string) (*proto.BlockchainProperties, bool) {
	var result *proto.BlockchainProperties
	for _, blockchain := range blockchains {
		if blockchain.BlockchainName == name {
			result = blockchain
		}
	}
	if result == nil {
		return nil, false
	}
	return result, true
}

func BlockConfirmedAfter(blockchainName string) uint64 {
	blockchainProperties, ok := Blockchain(blockchainName)
	if ok {
		return blockchainProperties.BlockConfirmedAfter
	} else {
		return 12
	}
}

func Decimals(blockchainName string) uint64 {
	blockchainProperties, ok := Blockchain(blockchainName)
	if ok {
		return blockchainProperties.Decimals
	} else {
		return 18
	}
}
