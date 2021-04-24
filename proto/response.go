package proto

import (
	"encoding/json"
	fmt "fmt"
)

// Guarantee RPCError satisfies the builtin error interface.
var _, _ error = &RpcError{}, (*RpcError)(nil)

// Error returns a string describing the RPC error.  This satisfies the
// builtin error interface.
func (e *RpcError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

//easyjson:json
type EasyResponse struct {
	Id     int32           `json:"id,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *RpcError       `json:"error,omitempty"`
}

//easyjson:json
type EasyRpcError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *EasyResponse) ParserResult() (result []byte, err error) {
	if r.Error != nil {
		return nil, r.Error
	}
	return r.Result, nil
}
