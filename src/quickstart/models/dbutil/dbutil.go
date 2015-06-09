package dbutil

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	. "quickstart/conf"
	"strconv"
	"strings"
	"time"
)

var myorm orm.Ormer

func init() {
	orm.RegisterDataBase("default", "sqlite3", DATABASEPATH)
	orm.RegisterModel(new(DB_message), new(DB_transaction), new(DB_blocks), new(DB_send))
	orm.RunSyncdb("default", false, true)
	myorm = orm.NewOrm()
	myorm.Using("default")
}

func DebugInsert() {
	var i uint64
	myorm.Raw("select count(*) from d_b_blocks").QueryRow(&i)
	i++
	for {
		var j uint64 = 1
		InsertBlock(i, "block hash", i)
		InsertTran(j, "tx hash", i, "block hash", i, "thesamesource",
			WISHINGWALLADDRESS, 1, 1, "21我爱你 ")
		InsertTran(j+1, "tx hash", i, "block hash", i, "secondsource",
			WISHINGWALLADDRESS, 1, 1, "21I Love You ")
		InsertTran(j+2, "tx hash", i, "block hash", i, "thesamesource",
			WISHINGWALLADDRESS, 1, 1, "22霍霍")
		InsertTran(j+3, "tx hash", i, "block hash", i, "secondsource",
			WISHINGWALLADDRESS, 1, 1, "22XIAO BAO BEI")
		i++
		time.Sleep(10 * time.Second)
	}
}

func LastBlockIndex() (uint64, error) {
	var i int
	myorm.Raw("Select count(*) from d_b_blocks").QueryRow(&i)

	var block DB_blocks
	block.Id = i
	err := myorm.Read(&block)
	if err == orm.ErrNoRows {
		fmt.Printf("can't find row %d\n", i)
	} else if err == orm.ErrMissPK {
		fmt.Printf("cant't find pk %d\n", i)
	}

	return block.Block_index, err
}
func GetLastTran() (DB_transaction, error) {
	var i int
	myorm.Raw("Select count(*) from d_b_transaction").QueryRow(&i)
	var tran DB_transaction
	tran.Id = i
	err := myorm.Read(&tran)
	if err == orm.ErrNoRows {
		fmt.Printf("Cant't find row %d\n", i)
	} else if err == orm.ErrMissPK {
		fmt.Printf("cant't find pk %d\n", i)
	}
	return tran, err
}
func GetTran(tx_hash string) (DB_transaction, error) {
	var trans []DB_transaction

	num, err := myorm.Raw("select * from d_b_transaction where tx_hash = ?",
		tx_hash).QueryRows(&trans)

	if err != nil || num == 0 {
		fmt.Printf("GetTran: can't find trans by tx_hash %s, %v\r\n", tx_hash, err)
		var temp DB_transaction
		return temp, errors.New("can't found")
	}

	return trans[0], nil
}

func InsertTran(tx_index uint64, tx_hash string, block_index uint64,
	block_hash string, block_time uint64, source,
	destination string, btc_amount, fee uint64, data string) error {
	tran := new(DB_transaction)
	tran.Tx_index = tx_index
	tran.Tx_hash = tx_hash
	tran.Block_index = block_index
	tran.Block_hash = block_hash
	tran.Block_time = block_time
	tran.Source = source
	tran.Destination = destination
	tran.Btc_amount = btc_amount
	tran.Fee = fee
	tran.Data = data
	_, err := myorm.Insert(tran)
	return err
}

func GetAllTransInBlock(block_index uint64) ([]DB_transaction, error) {
	var trans []DB_transaction
	num, err := myorm.Raw("select * from d_b_transaction "+
		"where block_index=? order by tx_index", block_index).QueryRows(&trans)
	if err != nil {
		fmt.Printf("Failed to get trans in block %d %v\n", block_index, err)
		return trans, err
	}
	fmt.Printf("GetAllTransInBlock: get %d trans in block %d\n", num, block_index)
	return trans, err
}

func GetAllBlocks() ([]DB_blocks, error) {
	var blocks []DB_blocks
	num, err := myorm.
		Raw("SELECT * FROM d_b_blocks ORDER BY block_index").
		QueryRows(&blocks)
	if err != nil {
		fmt.Printf("Failed to get all blocks %v\n", err)
	}
	fmt.Printf("Get all blocks in total %d\n", num)
	return blocks, err
}

func GetBlock(block_index uint64) (DB_blocks, error) {
	var blocks []DB_blocks
	num, err := myorm.Raw("select * from d_b_blocks where block_index = ?", block_index).QueryRows(&blocks)
	if err != nil || num == 0 {
		fmt.Printf("Can't find block by block_index %d, %v", block_index, err)
		var block DB_blocks
		return block, errors.New("Can't found")
	}
	return blocks[0], nil
}
func InsertBlock(block_index uint64, block_hash string, block_time uint64) error {
	block := new(DB_blocks)
	block.Block_index = block_index
	block.Block_hash = block_hash
	block.Block_time = block_time
	_, err := myorm.Insert(block)
	return err
}

func GetLastMessageIndex() (int, error) {
	var i int
	err := myorm.Raw("Select count(*) from d_b_message").QueryRow(&i)
	return i, err
}

func CheckWhetherRecord(tran DB_transaction) bool {
	var messages []DB_message

	num, err := myorm.Raw("select * from d_b_message where source = ?",
		tran.Source).QueryRows(&messages)
	if err == nil && num == 0 {
		return false
	}
	for _, message := range messages {
		if strings.Contains(message.Block_index_list, "-"+strconv.FormatUint(tran.Block_index, 10)) &&
			strings.Contains(message.Tx_index_list, "-"+strconv.FormatUint(tran.Tx_index, 10)) &&
			strings.Contains(message.Tx_hash_list, "-"+tran.Tx_hash) {
			return true
		}
	}
	return false

}

func InsertMessage(dbmessage DB_message) error {
	_, err := myorm.Insert(&dbmessage)
	return err
}

func Reinitialise(block_index uint64) error {
	//initialise(db) create missing tables
	if block_index != 0 {
		_, err := myorm.Raw("DELETE FROM d_b_transaction WHERE block_index > ?", block_index).Exec()
		if err != nil {
			fmt.Printf("Failed to delete from block_index %d in db_block\n", block_index)
		}
		_, err = myorm.Raw("DELETE FROM d_b_blocks WHERE block_index > ?", block_index).Exec()
		if err != nil {
			fmt.Printf("Failed to delete from block_index %d in transaction\n", block_index)
		}
		_, err = myorm.Raw("delete from d_b_message where block_index > ?", block_index).Exec()
		if err != nil {
			fmt.Printf("Faield to delete from block_index %d, in db_message\n", block_index)
		}
	}
	return nil
}
