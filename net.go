package bitcoinrpc

import (
	"encoding/json"

	"github.com/btcsuite/btcd/btcjson"
)

// FutureGetNetworkInfoResult is a future promise to deliver the result of a
// GetNetworkInfoAsync RPC invocation (or an applicable error).
type FutureGetNetworkInfoResult chan *response

// Receive waits for the response promised by the future and returns data about
// the current network.
func (r FutureGetNetworkInfoResult) Receive() (*btcjson.GetNetworkInfoResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as an array of getpeerinfo result objects.
	var networkInfo btcjson.GetNetworkInfoResult
	err = json.Unmarshal(res, &networkInfo)
	if err != nil {
		return nil, err
	}

	return &networkInfo, nil
}

// GetNetworkInfoAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetNetworkInfo for the blocking version and more details.
func (c *Client) GetNetworkInfoAsync() FutureGetNetworkInfoResult {
	cmd := btcjson.NewGetNetworkInfoCmd()
	return c.sendCmd(cmd)
}

// GetNetworkInfo returns data about the current network.
func (c *Client) GetNetworkInfo() (*btcjson.GetNetworkInfoResult, error) {
	return c.GetNetworkInfoAsync().Receive()
}
