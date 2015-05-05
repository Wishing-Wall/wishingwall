package dbutil

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	. "quickstart/conf"
	"strconv"
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
	myorm.Raw("select count(*) from d_b_message").QueryRow(&i)
	i++
	for {
		message := new(DB_message)
		message.Block_index = i
		message.Message = "hello " + strconv.FormatUint(i, 10)
		i++
		fmt.Println(myorm.Insert(message))

		InsertBLock(i, "test_hello_"+fmt.Sprintf("%d", i), i)

		InsertTran(i, "test_tx_hash"+fmt.Sprintf("%d", i),
			i, "test_block_hash"+fmt.Sprintf("%d", i),
			i, "test_source"+fmt.Sprintf("%d", i),
			"test_destination"+fmt.Sprintf("%d", i),
			i, i, "test_data"+fmt.Sprintf("%d", i))
		InsertTran(i+1, "test_tx_hash"+fmt.Sprintf("%d", i),
			i, "test_block_hash"+fmt.Sprintf("%d", i),
			i, "test_source"+fmt.Sprintf("%d", i),
			"test_destination"+fmt.Sprintf("%d", i),
			i, i, "test_data"+fmt.Sprintf("%d", i))

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
		fmt.Printf("can't find trans by tx_hash %s, %v", tx_hash, err)
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
	fmt.Printf("get %d trans in block %d\n", num, block_index)
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
func InsertBLock(block_index uint64, block_hash string, block_time uint64) error {
	block := new(DB_blocks)
	block.Block_index = block_index
	block.Block_hash = block_hash
	block.Block_time = block_time
	_, err := myorm.Insert(block)
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
