package bitcoinchain

import (
	_ "fmt"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcrpcclient"
	"github.com/btcsuite/btcutil"
)

/*
var connCfg = &btcrpcclient.ConnConfig{
	Host:         "192.168.31.104:8332",
	User:         "cddiao",
	Pass:         "jaijdfakejijfkSdjfaioejfakdljfiaejfakdjf",
	HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
	DisableTLS:   true, // Bitcoin core does not provide TLS by default
}*/
var connCfg = &btcrpcclient.ConnConfig{
	Host:         "127.0.0.1:19011",
	User:         "johnsonbernoulli",
	Pass:         "BitcoinIsTheFuture",
	HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
	DisableTLS:   true, // Bitcoin core does not provide TLS by default
}

var BlockChain, _ = btcrpcclient.New(connCfg, nil)

func GetBlockCount() (uint64, error) {
	i, err := BlockChain.GetBlockCount()

	return uint64(i), err
}

/*
{
Hash:0000000000007c639f2cbb23e4606a1d022fa4206353b9d92e99f5144bd74611
Confirmations:160713
Size:1536
Height:116219
Version:1
MerkleRoot:587fefd748f899f84d0fa1d8a3876fdb406a4bb8f54a31445cb72564701daea6
Tx:[
	be8f08d7f519eb863a68cf292ca51dbab7c9b49f50a96d13f2db32e432db363e
	a387039eca66297ba51ef2da3dcc8a0fc745bcb511e20ed9505cc6762be037bb
	2bd83162e264abf59f9124ca517050065f8c8eed2a21fbf85d454ee4e0e4c267
	028cfae228f8a4b0caee9c566bd41aed36bcd237cdc0eb18f0331d1e87111743
	3a06b6615756dc3363a8567fbfa8fe978ee0ba06eb33fd844886a0f01149ad62
	]
RawTx:[]
Time:1301705313
Nonce:1826107553
Bits:1b00f339
Difficulty:68977.78463021
PreviousHash:00000000000010d549135eb39bd3bbb1047df8e1512357216e8a85c57a1efbfb
NextHash:000000000000e9fcc59a6850f64a94476a30f5fe35d6d8c4b4ce0b1b04103a77
}
*/

func GetBlockByIndex(block_index uint64) (*btcjson.GetBlockVerboseResult, error) {

	hash, err := BlockChain.GetBlockHash(int64(block_index))
	if err != nil {
		//return new(btcjson.GetBlockVerboseResult), err
	}

	return BlockChain.GetBlockVerbose(hash, true)
}

func GetRawBlock(block_index uint64) (*btcutil.Block, error) {
	hash, err := BlockChain.GetBlockHash(int64(block_index))
	if err != nil {
		//return new(btcjson.GetBlockVerboseResult), err
	}

	return BlockChain.GetBlock(hash)
}

func GetBlockHashString(block_index uint64) (string, error) {
	hash, err := BlockChain.GetBlockHash(int64(block_index))
	/*
		var strHash string
		for _, value := range hash {
			strHash = strHash + string(value)
		}
	*/
	return hash.String(), err
}

/*
{
Hex:01000000010000000000000000000000000000000000000000000000000000000000
	000000ffffffff070439f3001b0134ffffffff014034152a010000004341045b3aaa284d169c5ae2
	d20d0b0673468ed3506aa8fea5976eacaf1ff304456f6522fbce1a646a24005b8b8e771a671f564c
	a6c03e484a1c394bf96e2a4ad01dceac00000000
Txid:be8f08d7f519eb863a68cf292ca51dbab7c9b49f50a96d13f2db32e432db363e
Version:1
LockTime:0
Vin:
[
	{
		Coinbase:0439f3001b0134
		Txid:
		Vout:0
		ScriptSig:<nil>
		Sequence:4294967295
	}
]
Vout:
[
	{
		Value:50.01
		N:0
		ScriptPubKey:{
			Asm:045b3aaa284d169c5ae2d20d0b0673468ed3506aa8fea5976eacaf1ff304456f65
22fbce1a646a24005b8b8e771a671f564ca6c03e484a1c394bf96e2a4ad01dce OP_CHECKSIG
			Hex:41045b3aaa284d169c5ae2d20d0b0673468ed3506aa8fea5976eacaf1ff304456f6522fbce1a646
a24005b8b8e771a671f564ca6c03e484a1c394bf96e2a4ad01dceac
			ReqSigs:1
			Type:pubkey
			Addresses:
			[1LgZTvoTJ6quJNCURmBUaJJkWWQZXkQnDn]
		}
	}
]
BlockHash:0000000000007c639f2cbb23e4606a1d022fa4206353b9d92e99f5144bd74611
Confirmations:160713
Time:1301705313
Blocktime:1301705313
}

{
Hex:0100000001764a1102d3ba6c4c5e8132161ebaf9101de3366f55b3421bd9759a2100
	8c828d010000008b48304502210099ea7cf845cbd60df939fa9ddb29357e1768e9ed88a9bfa051f4
	d33812c00a0a0220793811fe370be6ec9ba3893be44771ed022487ef7297c3df6cc72b4800d2ccc0
	014104e36253417cc23aa0f1e3781b941a128b70941a8ac698a7ceecd2da8cf206f0541175c15d49
	65c4808aec687e4140388ca3b87d324e41b35752d4d1473badffecffffffff01c068780400000000
	1976a9141add1d2e4145f8da4e8b21d395ebb0419a4c640788ac00000000
Txid:a387039eca66297ba51ef2da3dcc8a0fc745bcb511e20ed9505cc6762be037bb
Version:1
LockTime:0
Vin:
[
	{
		Coinbase:
		Txid:8d828c00219a75d91b42b3556f36e31d10f9ba1e1632815e4c6cbad302114a76
		Vout:1
		ScriptSig:0xc0820a94a0
		Sequence:4294967295
	}
]
Vout:
[
	{
		Value:0.75
		N:0
		ScriptPubKey:{
			Asm:OP_DUP OP_HASH160 1add1d2e4145f8da4e8b21d395ebb0419a4c6407 OP_EQUALVERIFY OP_CHECKSIG
			Hex:76a9141add1d2e4145f8da4e8b21d395ebb0419a4c640788ac
			ReqSigs:1
			Type:pubkeyhash
			Addresses:
				[13T3TFthRHfPXbBFu2ky7UEvNm9iNu6Tfo]
		}
	}
]
BlockHash:0000000000007c639f2cbb23e4606a1d022fa4206353b9d92e99f5144bd74611
Confirmations:164875
Time:1301705313
Blocktime:1301705313
}




*/
func GetRawTransaction(tx_hash string) (*btcjson.TxRawResult, error) {
	hash, err := wire.NewShaHashFromStr(tx_hash)
	tx, err := BlockChain.GetRawTransactionVerbose(hash)
	return tx, err
}
