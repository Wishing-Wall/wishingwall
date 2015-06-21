package wallet

import (
	"errors"
	"fmt"
	"quickstart/conf"
	"strings"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcrpcclient"
	_ "github.com/btcsuite/btcutil"
)

const TranVersion = "01000000"
const TranSequence = "ffffffff"

type Input struct {
	PrevOutputHash string
	PrevOutIndex   string
	ScriptLen      string
	ScriptSig      string
	Sequence       string
}

type Output struct {
	Type         string
	Value        string
	ScriptLen    string
	ScriptPubKey string
}

type RawTransaction struct {
	Version       string
	InputCount    string
	InputList     []Input
	OutputCount   string
	OutputList    []Output
	BlockLockTime string
}

func (this *RawTransaction) ToSerialize() (string, error) {
	var serialize string
	serialize = this.Version
	serialize += this.InputCount
	for _, temp := range this.InputList {
		serialize += HashRevert(temp.PrevOutputHash)
		serialize += temp.PrevOutIndex
		serialize += temp.ScriptLen
		serialize += temp.ScriptSig
		serialize += temp.Sequence
	}
	serialize += this.OutputCount
	for _, t := range this.OutputList {
		serialize += t.Value
		serialize += t.ScriptLen
		serialize += t.ScriptPubKey
	}
	serialize += "00000000"
	serialize = strings.Replace(serialize, " ", "", -1)
	return serialize, nil

}

func HashRevert(hash string) string {
	var temp []byte = []byte(hash)
	var revert string

	for i := len(hash) - 1; i > 0; i -= 2 {
		revert += string(temp[i-1])
		revert += string(temp[i])
	}
	return revert
}

func GetEmptyRawTransacton() RawTransaction {
	var temp RawTransaction
	temp.Version = TranVersion
	return temp
}

func InsertInput(RawTran RawTransaction, tx []btcjson.ListUnspentResult) (RawTransaction, error) {
	var inputcount int = 0
	for _, temp := range tx {
		var input Input
		input.PrevOutputHash = temp.TxID
		input.PrevOutIndex = HashRevert(fmt.Sprintf("%.8d", temp.Vout))
		input.ScriptLen = "00" // hard code ,I don't know why
		input.Sequence = TranSequence
		RawTran.InputList = append(RawTran.InputList, input)
		inputcount++
	}
	RawTran.InputCount = fmt.Sprintf("%.2x", inputcount)
	return RawTran, nil
}

func InsertOutputMinMoney(RawTran RawTransaction, Message string) (RawTransaction, uint64, error) {
	var m []byte = []byte(Message)
	var output Output
	fee := conf.MESSAGEFEE
	output.Value = HashRevert(fmt.Sprintf("%.16x", conf.COIN))
	output.ScriptPubKey = conf.PAYTOWISHINGWALL
	output.ScriptLen = fmt.Sprintf("%x", len(output.ScriptPubKey)/2)
	RawTran.OutputList = append(RawTran.OutputList, output)

	// total[1] + index[1] + 31 + 33
	var total = len(m) / 64

	var tail = len(m) % 64
	if tail != 0 {
		total++
		for i := 0; i < (64 - tail); i++ {
			m = append(m, 0xff)
		}
	}

	for i := 0; i < total; i++ {
		base := i * 64
		var message Output
		message.Value = HashRevert(fmt.Sprintf("%.16x", conf.MESSAGEFEE))
		fee += conf.MESSAGEFEE
		message.ScriptPubKey = conf.OP_1 + "21" +
			fmt.Sprintf("%.2x", total) + fmt.Sprintf("%.2x", i) +
			fmt.Sprintf("%x", string(m[base:base+31])) + "21" +
			fmt.Sprintf("%x", string(m[base+31:base+64])) +
			conf.OP_2 + conf.OP_CHECKMULTISIG

		message.ScriptLen = fmt.Sprintf("%x", len(message.ScriptPubKey)/2)
		RawTran.OutputList = append(RawTran.OutputList, message)
	}
	RawTran.OutputCount = fmt.Sprintf("%.2x", total+1)

	return RawTran, fee, nil
}

func InsertOutputPay(RawTran RawTransaction, PaytoMe uint64, Message string) (RawTransaction, error) {
	var m []byte = []byte(Message)
	var output Output
	output.Value = HashRevert(fmt.Sprintf("%.16x", PaytoMe))
	output.ScriptPubKey = conf.PAYTOWISHINGWALL
	output.ScriptLen = fmt.Sprintf("%x", len(output.ScriptPubKey)/2)
	RawTran.OutputList = append(RawTran.OutputList, output)

	// total[1] + index[1] + 31 + 33
	var total = len(m) / 64

	var tail = len(m) % 64
	if tail != 0 {
		total++
		for i := 0; i < (64 - tail); i++ {
			m = append(m, 0xff)
		}
	}

	for i := 0; i < total; i++ {
		base := i * 64
		var message Output
		message.Value = HashRevert(fmt.Sprintf("%.16x", conf.MESSAGEFEE))

		message.ScriptPubKey = conf.OP_1 + "21" +
			fmt.Sprintf("%.2x", total) + fmt.Sprintf("%.2x", i) +
			fmt.Sprintf("%x", string(m[base:base+31])) + "21" +
			fmt.Sprintf("%x", string(m[base+31:base+64])) +
			conf.OP_2 + conf.OP_CHECKMULTISIG

		message.ScriptLen = fmt.Sprintf("%x", len(message.ScriptPubKey)/2)
		RawTran.OutputList = append(RawTran.OutputList, message)
	}
	RawTran.OutputCount = fmt.Sprintf("%.2x", total+1)

	return RawTran, nil
}

var connCfg = &btcrpcclient.ConnConfig{
	Host:         "192.168.31.104:19011",
	User:         "admin2",
	Pass:         "123",
	HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
	DisableTLS:   true, // Bitcoin core does not provide TLS by default
}

var BlockChain, _ = btcrpcclient.New(connCfg, nil)

func SendRawTransaction(MsgTx *wire.MsgTx) (*wire.ShaHash, error) {
	hash, err := BlockChain.SendRawTransaction(MsgTx, false)
	fmt.Printf("hash is %v err is %v\r\n", hash, err)
	return hash, err
}

func CreateRawTransaction(PayAddress string, Message string) (*wire.MsgTx, uint64, error) {
	UnspentList, err := BlockChain.ListUnspent()
	if err != nil {
		return nil, 0, errors.New("ListUnspent error")
	}
	var tx []btcjson.ListUnspentResult
	var totalmoney uint64
	for _, temp := range UnspentList {
		if temp.Address == PayAddress {
			tx = append(tx, temp)
			totalmoney += uint64(temp.Amount * float64(conf.COIN))
		}
	}

	RawTranTry := GetEmptyRawTransacton()
	RawTranTry, _ = InsertInput(RawTranTry, tx)

	RawTranTry, MinMoney, _ := InsertOutputMinMoney(RawTranTry, Message)
	serializeTry, _ := RawTranTry.ToSerialize()
	MsgTx, _, _ := BlockChain.SignRawTransactionCMD(serializeTry)

	MinMoney += uint64(MsgTx.SerializeSize())/1000*conf.FEE + conf.FEE

	//hash, err := BlockChain.SendRawTransaction(MsgTx, true)
	//fmt.Printf("hash is %v err is %v\r\n", hash, err)

	if totalmoney < MinMoney {
		return nil, MinMoney, errors.New(fmt.Sprintf("Do not Have enough money in address %s has %v need %v", PayAddress,
			totalmoney, MinMoney))
	}

	RawTranReal := GetEmptyRawTransacton()
	RawTranReal, _ = InsertInput(RawTranReal, tx)

	RawTranReal, _ = InsertOutputPay(RawTranReal, (totalmoney - MinMoney + conf.MESSAGEFEE), Message)
	serializeReal, _ := RawTranReal.ToSerialize()
	MsgTxReal, _, _ := BlockChain.SignRawTransactionCMD(serializeReal)
	MinMoney += uint64(MsgTxReal.SerializeSize()) / 1000 * conf.FEE
	return MsgTxReal, MinMoney, nil

}
