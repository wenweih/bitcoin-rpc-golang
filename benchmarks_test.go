package bitcoinrpc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	rpcproto "github.com/wenweih/bitcoin-rpc-golang/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

func readTestData() []byte {
	blockFile, err := os.Open("./testdata/600000.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer blockFile.Close()
	byteValue, _ := ioutil.ReadAll(blockFile)
	return byteValue
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	byteValue := readTestData()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var blockResult rpcproto.GetBlockVerboseTxResult
		err := json.Unmarshal(byteValue, &blockResult)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkProtojsonUnmarshal(b *testing.B) {
	byteValue := readTestData()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var blockResult rpcproto.GetBlockVerboseTxResult
		err := protojson.Unmarshal(byteValue, &blockResult)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkEasyjsonUnmarshal(b *testing.B) {
	byteValue := readTestData()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var blockResult rpcproto.GetBlockVerboseTxResult
		err := blockResult.UnmarshalJSON(byteValue)
		if err != nil {
			b.Fatal(err)
		}
	}
}
