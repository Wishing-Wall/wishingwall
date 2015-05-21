package bitcoinchain

import (
	"github.com/btcsuite/btcd/btcjson"
	//"github.com/btcsuite/btcd/wire"
	"fmt"
	"github.com/btcsuite/btcrpcclient"
	//"github.com/btcsuite/btcutil"
)

var connCfg = &btcrpcclient.ConnConfig{
	Host:         "192.168.31.104:8332",
	User:         "cddiao",
	Pass:         "jaijdfakejijfkSdjfaioejfakdljfiaejfakdjf",
	HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
	DisableTLS:   true, // Bitcoin core does not provide TLS by default
}

var BlockChain, _ = btcrpcclient.New(connCfg, nil)

func GetBlockCount() (uint64, error) {
	i, err := BlockChain.GetBlockCount()

	return uint64(i), err
}

func GetBlockByIndex(block_index uint64) (*btcjson.GetBlockVerboseResult, error) {

	hash, err := BlockChain.GetBlockHash(int64(block_index))
	if err != nil {
		//return new(btcjson.GetBlockVerboseResult), err
	}
	fmt.Printf("the block[%d] hash is %s\n", block_index, hash)
	return BlockChain.GetBlockVerbose(hash, false)

}
func GetBlockHash(block_index uint64) (string, error) {
	hash, err := BlockChain.GetBlockHash(int64(block_index))
	var strHash string
	for _, value := range hash {
		strHash = strHash + string(value)
	}
	return strHash, err
}
func GetRawTransaction(tx_hash string) {

	//hash := wire.ShaHash(tx_hash)
	//tx, err := BlockChain.GetRawTransaction(&hash)
}
