package blockchains

import (
	"bytes"
	"github.com/Ankr-network/ankrscan-proto-contract/go/proto"
	"github.com/ethereum/go-ethereum/common"
)

func AllPegs() []*proto.CurrencyPeg {
	return pegs
}

func Pegs(blockchainName string) []*proto.CurrencyPeg {
	result := make([]*proto.CurrencyPeg, 0)
	for _, peg := range pegs {
		if peg.BlockchainName == blockchainName {
			result = append(result, peg)
		}
	}
	return result
}

func UsdPegs(blockchainName string) []*proto.CurrencyPeg {
	result := make([]*proto.CurrencyPeg, 0)
	for _, peg := range Pegs(blockchainName) {
		if peg.Peg == proto.Peg_USD_PEG {
			result = append(result, peg)
		}
	}
	return result
}

func CurrencyPeg(blockchainName string, address []byte) proto.Peg {
	for _, peg := range Pegs(blockchainName) {
		if bytes.Equal(peg.Address, address) {
			return peg.Peg
		}
	}
	return proto.Peg_NO_PEG
}

var pegs = []*proto.CurrencyPeg{
	/*
		░█████╗░██████╗░██████╗░██╗████████╗██████╗░██╗░░░██╗███╗░░░███╗
		██╔══██╗██╔══██╗██╔══██╗██║╚══██╔══╝██╔══██╗██║░░░██║████╗░████║
		███████║██████╔╝██████╦╝██║░░░██║░░░██████╔╝██║░░░██║██╔████╔██║
		██╔══██║██╔══██╗██╔══██╗██║░░░██║░░░██╔══██╗██║░░░██║██║╚██╔╝██║
		██║░░██║██║░░██║██████╦╝██║░░░██║░░░██║░░██║╚██████╔╝██║░╚═╝░██║
		╚═╝░░╚═╝╚═╝░░╚═╝╚═════╝░╚═╝░░░╚═╝░░░╚═╝░░╚═╝░╚═════╝░╚═╝░░░░░╚═╝
	*/
	{
		BlockchainName: "arbitrum", // Wrapped Ether
		Address:        common.HexToAddress("0x82af49447d8a07e3bd95bd0d56f35241523fbab1").Bytes(),
		Peg:            proto.Peg_NATIVE_PEG,
	},
	{
		BlockchainName: "arbitrum", // Tether USD
		Address:        common.HexToAddress("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	{
		BlockchainName: "arbitrum", // USD Coin
		Address:        common.HexToAddress("0xff970a61a04b1ca14834a43f5de4533ebddb5cc8").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},

	/*

		░█████╗░██╗░░░██╗░█████╗░██╗░░██╗
		██╔══██╗██║░░░██║██╔══██╗╚██╗██╔╝
		███████║╚██╗░██╔╝███████║░╚███╔╝░
		██╔══██║░╚████╔╝░██╔══██║░██╔██╗░
		██║░░██║░░╚██╔╝░░██║░░██║██╔╝╚██╗
		╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░╚═╝╚═╝░░╚═╝
	*/
	{
		BlockchainName: "avalanche", // Tether USD
		Address:        common.HexToAddress("0xc7198437980c041c805a1edcba50c1ce5db95118").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	{
		BlockchainName: "avalanche", // Wrapped AVAX
		Address:        common.HexToAddress("0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7").Bytes(),
		Peg:            proto.Peg_NATIVE_PEG,
	},
	{
		BlockchainName: "avalanche", // usdc // USD Coin
		Address:        common.HexToAddress("0xB97EF9Ef8734C71904D8002F8b6Bc66Dd9c48a6E").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	/*
		██████╗░░██████╗░█████╗░
		██╔══██╗██╔════╝██╔══██╗
		██████╦╝╚█████╗░██║░░╚═╝
		██╔══██╗░╚═══██╗██║░░██╗
		██████╦╝██████╔╝╚█████╔╝
		╚═════╝░╚═════╝░░╚════╝░
	*/
	{
		BlockchainName: "bsc", // Binance-Peg BSC-USD
		Address:        common.HexToAddress("0x55d398326f99059ff775485246999027b3197955").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	{
		BlockchainName: "bsc", // Wrapped BNB
		Address:        common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c").Bytes(),
		Peg:            proto.Peg_NATIVE_PEG,
	},
	{
		BlockchainName: "bsc", // Binance-Peg USD Coin:
		Address:        common.HexToAddress("0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	/*

		███████╗████████╗██╗░░██╗
		██╔════╝╚══██╔══╝██║░░██║
		█████╗░░░░░██║░░░███████║
		██╔══╝░░░░░██║░░░██╔══██║
		███████╗░░░██║░░░██║░░██║
		╚══════╝░░░╚═╝░░░╚═╝░░╚═╝
	*/
	{
		BlockchainName: "eth", // Tether USD
		Address:        common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	{
		BlockchainName: "eth", // Wrapped Ether
		Address:        common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2").Bytes(),
		Peg:            proto.Peg_NATIVE_PEG,
	},
	{
		BlockchainName: "eth", // USD Coin
		Address:        common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	/*

		███████╗████████╗███╗░░░███╗
		██╔════╝╚══██╔══╝████╗░████║
		█████╗░░░░░██║░░░██╔████╔██║
		██╔══╝░░░░░██║░░░██║╚██╔╝██║
		██║░░░░░░░░██║░░░██║░╚═╝░██║
		╚═╝░░░░░░░░╚═╝░░░╚═╝░░░░░╚═╝
	*/
	{
		BlockchainName: "fantom", // Wrapped Fantom
		Address:        common.HexToAddress("0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83").Bytes(),
		Peg:            proto.Peg_NATIVE_PEG,
	},
	{
		BlockchainName: "fantom", // Frapped USDT
		Address:        common.HexToAddress("0x049d68029688eabf473097a2fc38ef61633a3c7a").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	{
		BlockchainName: "fantom", // USD Coin
		Address:        common.HexToAddress("0x04068da6c83afcfa0e13ba15a6696662335d5b75").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	/*

		██████╗░░█████╗░██╗░░░░░██╗░░░██╗░██████╗░░█████╗░███╗░░██╗
		██╔══██╗██╔══██╗██║░░░░░╚██╗░██╔╝██╔════╝░██╔══██╗████╗░██║
		██████╔╝██║░░██║██║░░░░░░╚████╔╝░██║░░██╗░██║░░██║██╔██╗██║
		██╔═══╝░██║░░██║██║░░░░░░░╚██╔╝░░██║░░╚██╗██║░░██║██║╚████║
		██║░░░░░╚█████╔╝███████╗░░░██║░░░╚██████╔╝╚█████╔╝██║░╚███║
		╚═╝░░░░░░╚════╝░╚══════╝░░░╚═╝░░░░╚═════╝░░╚════╝░╚═╝░░╚══╝
	*/
	{
		BlockchainName: "polygon", // (PoS) Tether USD
		Address:        common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
	{
		BlockchainName: "polygon", // Wrapped Matic
		Address:        common.HexToAddress("0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270").Bytes(),
		Peg:            proto.Peg_NATIVE_PEG,
	},
	{
		BlockchainName: "polygon", // USD Coin (PoS)
		Address:        common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174").Bytes(),
		Peg:            proto.Peg_USD_PEG,
	},
}
