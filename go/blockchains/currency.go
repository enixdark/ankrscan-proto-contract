package blockchains

import (
	"github.com/Ankr-network/ankrscan-proto-contract/go/proto"
	"github.com/ethereum/go-ethereum/common"
)

var currencies = []*proto.CurrencyDetails{
	/*
	░█████╗░██████╗░██████╗░██╗████████╗██████╗░██╗░░░██╗███╗░░░███╗
	██╔══██╗██╔══██╗██╔══██╗██║╚══██╔══╝██╔══██╗██║░░░██║████╗░████║
	███████║██████╔╝██████╦╝██║░░░██║░░░██████╔╝██║░░░██║██╔████╔██║
	██╔══██║██╔══██╗██╔══██╗██║░░░██║░░░██╔══██╗██║░░░██║██║╚██╔╝██║
	██║░░██║██║░░██║██████╦╝██║░░░██║░░░██║░░██║╚██████╔╝██║░╚═╝░██║
	╚═╝░░╚═╝╚═╝░░╚═╝╚═════╝░╚═╝░░░╚═╝░░░╚═╝░░╚═╝░╚═════╝░╚═╝░░░░░╚═╝
	*/
	{
		BlockchainName: "arbitrum",
		Name:           "Wrapped Ether",
		Symbol:         "WETH",
		Decimals:       18,
		Address:        common.HexToAddress("0x82af49447d8a07e3bd95bd0d56f35241523fbab1").Bytes(),
		Type:           proto.TokenType_ERC20,
		PegType:        proto.PegType_NATIVE_PEG,
	},
	{
		BlockchainName: "arbitrum",
		Name:           "Tether USD",
		Symbol:         "USDT",
		Decimals:       18,
		Address:        common.HexToAddress("0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9").Bytes(),
		Type:           proto.TokenType_ERC20,
		PegType:        proto.PegType_USD_PEG,
	},
	{
		BlockchainName: "arbitrum",
		Name:           "USD Coin",
		Symbol:         "USDC",
		Decimals:       18,
		Address:        common.HexToAddress("0xff970a61a04b1ca14834a43f5de4533ebddb5cc8").Bytes(),
		Type:           proto.TokenType_ERC20,
		PegType:        proto.PegType_USD_PEG,
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
		BlockchainName: "avalanche",
		Name:           "Avalanche",
		Symbol:         "AVAX",
		Decimals:       18,
		Address:        []byte{},
		Type:           proto.TokenType_NATIVE,
		PegType:        proto.PegType_NO_PEG,
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
		BlockchainName: "bsc",
		Name:           "Binance Smart Chain",
		Symbol:         "BNB",
		Decimals:       18,
		Address:        []byte{},
		Type:           proto.TokenType_NATIVE,
		PegType:        proto.PegType_NO_PEG,
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
		BlockchainName: "eth",
		Name:           "Ethereum",
		Symbol:         "ETH",
		Decimals:       18,
		Address:        []byte{},
		Type:           proto.TokenType_NATIVE,
		PegType:        proto.PegType_NO_PEG,
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
		BlockchainName: "fantom",
		Name:           "Fantom",
		Symbol:         "FTM",
		Decimals:       18,
		Address:        []byte{},
		Type:           proto.TokenType_NATIVE,
		PegType:        proto.PegType_NO_PEG,
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
		BlockchainName: "polygon",
		Name:           "Polygon",
		Symbol:         "MATIC",
		Decimals:       18,
		Address:        []byte{},
		Type:           proto.TokenType_NATIVE,
		PegType:        proto.PegType_NO_PEG,
	},
}
