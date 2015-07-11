package protocol

import (
	"encoding/hex"
	"errors"
	"fmt"
	"quickstart/conf"
	"quickstart/models/bitcoinchain"
	"quickstart/models/dbutil"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcd/btcjson"
)

//this is a debug func
func get_tx_info(tx *btcjson.TxRawResult, block_index uint64) (source,
	destination string, btc_amount, fee uint64, data []string) {
	var bFound bool = false
	for _, value := range tx.Vout {
		Address := value.ScriptPubKey.Addresses

		if len(Address) != 0 {
			if Address[0] == conf.WISHINGWALLADDRESS {
				fmt.Printf("Got one valied tx %v\r\n", tx.Hex)
				fmt.Printf("the tx is %v\r\n", tx)
				bFound = true
				continue
			}
		}
		if bFound == true {
			message := strings.Split(value.ScriptPubKey.Asm, " ")
			fmt.Printf("message is %v\r\n", message)
			merge := message[1] + message[2]
			data = append(data, merge)
		}
	}

	if bFound == true {
		destination = conf.WISHINGWALLADDRESS
	} else {
		var temp []string
		return "", "", 0, 0, temp
	}

	//get source address
	fmt.Printf("tx.Vin[0].Txid %v\r\n", tx.Vin[0].Txid)
	if tx.Vin[0].Txid == "" {
		source = "coinbase"
	} else {
		SourceTx, _ := bitcoinchain.GetRawTransaction(tx.Vin[0].Txid)
		fmt.Printf("sourcetx is %v\r\n", SourceTx)
		if SourceTx == nil {
			source = "Unkonw"
		} else {
			for _, value := range SourceTx.Vout {
				if value.N == tx.Vin[0].Vout {
					source = value.ScriptPubKey.Addresses[0]
				}
			}
		}
	}

	return source, destination, 0, 0, data
}

type poll struct {
	Message_count uint64
	Message_index uint64
	Block_index   uint64
	Tx_index      uint64
	Tx_hash       string
	Account       uint64
	Source        string
	Destination   string
	Message       string
}
type polls []poll

var message_pool []polls

func UpdateMessagePool(messageS polls) ([]polls, error) {
	var newpolls []polls
	var bFound bool = false
	for _, value := range message_pool {
		if value[0].Source == messageS[0].Source {
			newpolls = append(newpolls, messageS)
			bFound = true
		} else {
			newpolls = append(newpolls, value)
		}
	}
	if false == bFound {
		newpolls = append(newpolls, messageS)
	}
	return newpolls, nil
}

func GetMessageFromPoolBySource(source string) (polls, error) {

	for _, messageS := range message_pool {
		for _, message := range messageS {
			if source == message.Source {
				return messageS, nil
			}
		}
	}
	var temp polls
	return temp, errors.New("can't find")
}

func DeleteMessageFromPoolBySource(source string) ([]polls, error) {
	var temp []polls
	for _, messageS := range message_pool {
		if messageS[0].Source != source {
			temp = append(temp, messageS)
		}
	}
	return temp, nil
}

//tran.Data  [message_count][message_index][message_body]

func GetMessageFromData(data string) (message_count,
	message_index uint64, message_body string) {
	databyte := []byte(data)
	temp, _ := strconv.Atoi(string(databyte[0:2]))
	message_count = uint64(temp)
	temp, _ = strconv.Atoi(string(databyte[2:4]))
	message_index = uint64(temp)
	converbody, _ := hex.DecodeString(string(databyte[4:]))
	//var converrune []rune = []rune(string(converbody))
	message_body = string(converbody)
	return message_count, message_index, message_body
}

func Parse_tx(tran conf.DB_transaction) error {
	if tran.Destination == conf.WISHINGWALLADDRESS {
		bAlready := dbutil.CheckWhetherRecord(tran)
		if bAlready {
			return nil
		}

		message_count, message_index, message_body := GetMessageFromData(tran.Data)

		var message poll
		message.Source = tran.Source
		message.Destination = tran.Destination
		message.Message_count = message_count
		message.Message_index = message_index
		message.Message = message_body
		message.Block_index = tran.Block_index
		message.Tx_index = tran.Tx_index
		message.Tx_hash = tran.Tx_hash
		message.Account = tran.Btc_amount

		messageS, _ := GetMessageFromPoolBySource(tran.Source)
		messageS = append(messageS, message)

		if messageS[0].Message_count == uint64(len(messageS)) {
			//insert to db
			var dbmessage conf.DB_message
			dbmessage.Message_count = messageS[0].Message_count
			dbmessage.Source = messageS[0].Source
			dbmessage.Destination = messageS[0].Destination
			for _, message := range messageS {
				dbmessage.Block_index_list = dbmessage.Block_index_list + "-" +
					strconv.FormatUint(message.Block_index, 10)
				dbmessage.Tx_index_list = dbmessage.Tx_index_list + "-" +
					strconv.FormatUint(message.Tx_index, 10)
				dbmessage.Tx_hash_list = dbmessage.Tx_hash_list + "-" + message.Tx_hash
				dbmessage.Account = dbmessage.Account + message.Account

				dbmessage.Message = dbmessage.Message + message.Message
			}

			dbutil.InsertMessage(dbmessage)
			message_pool, _ = DeleteMessageFromPoolBySource(dbmessage.Source)
		} else {
			message_pool, _ = UpdateMessagePool(messageS)
		}

	}
	return nil
}

func Parse_block(block_index, block_time uint64, previous_ledger_hash,
	previous_txlist_hash string) (string, string) {
	fmt.Printf("Start to parse block[%d]\r\n", block_index)
	trans, err := dbutil.GetAllTransInBlock(block_index)
	fmt.Printf("all tran in block %d, %v\n", block_index, trans)
	if err != nil {
		var temp string
		return temp, temp
	}
	for _, tran := range trans {
		Parse_tx(tran)
	}

	var temp string
	return temp, temp
}

func Reparse(block_index uint64) error {
	dbutil.Reinitialise(block_index)

	blocks, err := dbutil.GetAllBlocks()
	if err != nil {
		return err
	}
	var previous_ledger_hash, previous_txlist_hash string
	for _, block := range blocks {
		conf.CURRENT_BLOCK_INDEX = block.Block_index
		previous_ledger_hash, previous_txlist_hash =
			Parse_block(block.Block_index, block.Block_time,
				previous_ledger_hash, previous_txlist_hash)
	}

	return nil
}

func Follow() {
	var block_index uint64
	block_index, err := dbutil.LastBlockIndex()
	if err != nil {
		fmt.Println("Get lastblockindex failed")
		//return
		block_index = 0
	}
	if block_index == 0 {
		fmt.Println("block table in database is empty")
	}
	block_index++
        fmt.Print("The last block_index++ in db is ", block_index)
	if block_index < 363919 {
		block_index = 363919
        }
	fmt.Printf("block_index is %v\r\n", block_index)

	var dbtran conf.DB_transaction
	dbtran, err = dbutil.GetLastTran()
	if err != nil {
		fmt.Println("GetLastTran error")
	}

	tx_index := dbtran.Tx_index + 1
	for {
		tempblockcount, err := bitcoinchain.GetBlockCount()

		if err != nil {
			fmt.Printf("get tempblockcount failed %v\r\n", err)
			continue
		}
		if block_index <= tempblockcount {
			fmt.Printf("block_index[%d] < tempblockcount[%d]\r\n", block_index, tempblockcount)
			c := block_index
			requires_rollback := false
			for {
				if c == 363919 {
					fmt.Printf("Reach the start...")
					break
				}
				c_block, _ := bitcoinchain.GetBlockByIndex(c)
				bitcoind_parent := c_block.PreviousHash
				block, err := dbutil.GetBlock(c - 1)
				if err != nil {
					fmt.Printf("dbutil.Getblock failed [%d]\r\n", c-1)
					break
				}
				db_parent := block.Block_hash
				if db_parent == bitcoind_parent {
					break
				} else {
					c -= 1
					requires_rollback = true
				}
			}
			if requires_rollback {
				fmt.Printf("status:Blockchain reorganisation at block %v\n", c)
				Reparse(c - 1)
				block_index = c
				continue
			}
			block_hash, _ := bitcoinchain.GetBlockHashString(block_index)
			block, _ := bitcoinchain.GetBlockByIndex(block_index)
			block_time := block.Time
			tx_hash_list := block.Tx
			fmt.Printf("Insert block[%d] into db\r\n", block_index)

			for _, tx_hash := range tx_hash_list {
				fmt.Printf("Try to parse tx %v\r\n", tx_hash)
				_, err := dbutil.GetTran(tx_hash)
				if err == nil {
					tx_index += 1
					fmt.Printf("the tx_index[%d] already in db, next\r\n", tx_index-1)
					continue
				}
				tx, err := bitcoinchain.GetRawTransaction(tx_hash)
				if err != nil {
					fmt.Printf("GetRawTransaction %v failed %v\r\n", tx_hash, err)
					continue
				}
				source, destination, btc_amount, fee, data := get_tx_info(tx, block_index)
				for _, value := range data {
					if source != "" && (value != "" || destination == conf.WISHINGWALLADDRESS) {

						dbutil.InsertTran(tx_index, tx_hash, block_index, block_hash, uint64(block_time), source, destination, btc_amount, fee, value)
					}
				}
				tx_index += 1
			}
			dbutil.InsertBlock(block_index, block_hash, uint64(block_time))
			Parse_block(block_index, uint64(block_time), "", "")
			block_index++
		}
		if block_index >= 363919 {
			fmt.Printf("sleep 1 second")
			time.Sleep(1 * time.Second)
		}
	}

}
