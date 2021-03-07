package bitcoinrpc

import (
	"encoding/json"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"

	rpcproto "github.com/wenweih/bitcoin-rpc-golang/proto"
)

// FutureGetBlockHashResult is a future promise to deliver the result of a
// GetBlockHashAsync RPC invocation (or an applicable error).
type FutureGetBlockHashResult chan *response

// Receive waits for the response promised by the future and returns the hash of
// the block in the best block chain at the given height.
func (r FutureGetBlockHashResult) Receive() (*chainhash.Hash, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal the result as a string-encoded sha.
	var txHashStr string
	err = json.Unmarshal(res, &txHashStr)
	if err != nil {
		return nil, err
	}
	return chainhash.NewHashFromStr(txHashStr)
}

// GetBlockHashAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetBlockHash for the blocking version and more details.
func (c *Client) GetBlockHashAsync(blockHeight int64) FutureGetBlockHashResult {
	cmd := btcjson.NewGetBlockHashCmd(blockHeight)
	return c.sendCmd(cmd)
}

// GetBlockHash returns the hash of the block in the best block chain at the
// given height.
func (c *Client) GetBlockHash(blockHeight int64) (*chainhash.Hash, error) {
	return c.GetBlockHashAsync(blockHeight).Receive()
}

// waitForGetBlockRes waits for the response of a getblock request. If the
// response indicates an invalid parameter was provided, a legacy style of the
// request is resent and its response is returned instead.
func (c *Client) waitForGetBlockRes(respChan chan *response, hash string,
	verbose, verboseTx bool) ([]byte, error) {

	res, err := receiveFuture(respChan)

	// Otherwise, we can return the response as is.
	return res, err
}

// FutureGetBlockVerboseTxResult is a future promise to deliver the result of a
// GetBlockVerboseTxResult RPC invocation (or an applicable error).
type FutureGetBlockVerboseTxResult struct {
	client   *Client
	hash     string
	Response chan *response
}

// Receive waits for the response promised by the future and returns a verbose
// version of the block including detailed information about its transactions.
func (r FutureGetBlockVerboseTxResult) Receive() (*rpcproto.GetBlockVerboseTxResult, error) {
	res, err := r.client.waitForGetBlockRes(r.Response, r.hash, true, true)
	if err != nil {
		return nil, err
	}

	var blockResult rpcproto.GetBlockVerboseTxResult
	err = json.Unmarshal(res, &blockResult)
	if err != nil {
		return nil, err
	}

	return &blockResult, nil
}

// GetBlockVerboseTxAsync returns an instance of a type that can be used to get
// the result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetBlockVerboseTx or the blocking version and more details.
func (c *Client) GetBlockVerboseTxAsync(blockHash *chainhash.Hash) FutureGetBlockVerboseTxResult {
	hash := ""
	if blockHash != nil {
		hash = blockHash.String()
	}

	// From the bitcoin-cli getblock documentation:
	//
	// If verbosity is 2, returns an Object with information about block
	// and information about each transaction.
	cmd := btcjson.NewGetBlockCmd(hash, btcjson.Int(2))
	return FutureGetBlockVerboseTxResult{
		client:   c,
		hash:     hash,
		Response: c.sendCmd(cmd),
	}
}

// GetBlockVerboseTx returns a data structure from the server with information
// about a block and its transactions given its hash.
//
// See GetBlockVerbose if only transaction hashes are preferred.
// See GetBlock to retrieve a raw block instead.
func (c *Client) GetBlockVerboseTx(blockHash *chainhash.Hash) (*rpcproto.GetBlockVerboseTxResult, error) {
	return c.GetBlockVerboseTxAsync(blockHash).Receive()
}
