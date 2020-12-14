package main

import (
	"fmt"

	bitcoinrpc "github.com/wenweih/bitcoin-rpc-golang"
)

func main() {
	btcClient, err := bitcoinrpc.New(&bitcoinrpc.ConnConfig{
		Host:       "host:port",
		User:       "user",
		Pass:       "pass",
		DisableTLS: true,
	})
	if err != nil {
		fmt.Println("Connection error: " + err.Error())
		return
	}
	// result, err := btcClient.GetNetworkInfo()
	// result, err := btcClient.GetBlockChainInfo()
	// result, err := btcClient.GetBlockHash(500000)
	blockHash, err := btcClient.GetBlockHash(100000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, err := btcClient.GetBlockVerboseTx(blockHash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println("result" + result.Bestblockhash)
	fmt.Println(result)
}
