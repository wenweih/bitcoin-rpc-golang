package bitcoinrpc

import (
	"github.com/btcsuite/btcd/btcjson"
	rpcproto "github.com/wenweih/bitcoin-rpc-golang/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

// GetBlockChainInfo returns information related to the processing state of
// various chain-specific details such as the current difficulty from the tip
// of the main chain.
func (c *Client) GetBlockChainInfo() (*rpcproto.GetBlockChainInfoResult, error) {
	return c.GetBlockChainInfoAsync().Receive()
}

// FutureGetBlockChainInfoResult is a promise to deliver the result of a
// GetBlockChainInfoAsync RPC invocation (or an applicable error).
type FutureGetBlockChainInfoResult struct {
	client   *Client
	Response chan *response
}

// Receive waits for the response promised by the future and returns chain info
// result provided by the server.
func (r FutureGetBlockChainInfoResult) Receive() (*rpcproto.GetBlockChainInfoResult, error) {
	res, err := receiveFuture(r.Response)
	if err != nil {
		return nil, err
	}
	var chainInfo rpcproto.GetBlockChainInfoResult
	if err := protojson.Unmarshal(res, &chainInfo); err != nil {
		return nil, err
	}

	return &chainInfo, nil
}

// GetBlockChainInfoAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function
// on the returned instance.
//
// See GetBlockChainInfo for the blocking version and more details.
func (c *Client) GetBlockChainInfoAsync() FutureGetBlockChainInfoResult {
	cmd := btcjson.NewGetBlockChainInfoCmd()
	return FutureGetBlockChainInfoResult{
		client:   c,
		Response: c.sendCmd(cmd),
	}
}
