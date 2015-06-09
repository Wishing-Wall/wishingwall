package protocol

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"quickstart/conf"
	"quickstart/models/bitcoinchain"
	"quickstart/models/dbutil"
	"strconv"
	_ "time"
)

//this is a debug func
func get_tx_info(tx *btcjson.TxRawResult, block_index uint64) (source,
	destination string, btc_amount, fee uint64, data string) {
	return
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
	fmt.Printf("the message_pool is %T, %v\n", message_pool, message_pool)
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

	message_count = uint64(data[0] - '0')
	message_index = uint64(data[1] - '0')
	message_body = data[2:]
	return message_count, message_index, message_body
}

func Parse_tx(tran conf.DB_transaction) error {
	if tran.Destination == conf.WISHINGWALLADDRESS {
		//fmt.Printf("before call GetMEssageBySource\n")
		bAlready := dbutil.CheckWhetherRecord(tran)
		if bAlready {
			return nil
		}
		//_, err := dbutil.GetMessageBySource(tran.Source)
		//fmt.Printf("GetMessage return err = %v\n", err)
		//if nil == err {
		////	fmt.Printf("return from parse_tx\n")
		//	return nil
		//}
		message_count, message_index, message_body := GetMessageFromData(tran.Data)
		fmt.Printf("get from data count=%d, index =%d body=%s\n", message_count, message_index, message_body)
		//like a temporary storage
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
		fmt.Printf("message len=%d, should be %d\n", len(messageS), messageS[0].Message_count)
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
	block_index, _ := bitcoinchain.GetBlockCount()
	fmt.Printf("block count is %d\n", block_index)
	block, err := bitcoinchain.GetBlockByIndex(116219)
	fmt.Printf("block is %+v %v\n", block, err)
	tran, _ := bitcoinchain.GetRawTransaction(block.Tx[0])
	fmt.Printf("tran is %+v\n", tran)
	/*
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

		var dbtran conf.DB_transaction
		dbtran, err = dbutil.GetLastTran()
		if err != nil {
			fmt.Println("GetLastTran error")
		}
		fmt.Printf("dbtran=%v\n", dbtran)

		tx_index := dbtran.Tx_index + 1

		for {
			tempblockcount, err := bitcoinchain.GetBlockCount()
			fmt.Printf("block count is %d\n", tempblockcount)
			if err != nil {
				continue
			}
			if block_index <= tempblockcount {
				fmt.Printf("Block: %d\n", block_index)
				c := block_index
				requires_rollback := false
				for {
					if c == 3 {
						break
					}
					c_block, _ := bitcoinchain.GetBlockByIndex(c)
					bitcoind_parent := c_block.PreviousHash
					block, err := dbutil.GetBlock(c - 1)
					if err != nil {
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
				dbutil.InsertBlock(block_index, block_hash, uint64(block_time))
				for _, tx_hash := range tx_hash_list {
					_, err := dbutil.GetTran(tx_hash)
					if err == nil {
						tx_index += 1
						continue
					}
					tx, _ := bitcoinchain.GetRawTransaction(tx_hash)
					source, destination, btc_amount, fee, data := get_tx_info(tx, block_index)
					if source != "" && (data != "" || destination == conf.WISHINGWALLADDRESS) {
						dbutil.InsertTran(tx_index, tx_hash, block_index, block_hash, uint64(block_time), source, destination, btc_amount, fee, data)

					}
					tx_index += 1
				}
				Parse_block(block_index, uint64(block_time), "", "")
			}
			block_index += 1
		}
		time.Sleep(2)
	*/
}
