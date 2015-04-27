package protocol

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	. "quickstart/conf"
	"time"
)

func init() {
	orm.RegisterDataBase("default", "sqlite3", DATABASEPATH)
	orm.RegisterModel(new(DB_message), new(DB_transaction), new(DB_blocks))
	orm.RunSyncdb("default", false, true)
}

func Follow() {
	o := orm.NewOrm()
	o.Using("default")
	var i uint64 = 0
	for {
		message := new(DB_message)
		message.Block_index = i
		i++
		fmt.Println(o.Insert(message))
		time.Sleep(10 * time.Second)
	}
	/*
		logger.Infoln("Start... ")
		var block_index uint64
		block_index, err := dbutil.LastBlockIndex()
		if err != nil {
			logger.Errorln("Get lastblockindex failed")
			return
		}
		if block_index == 0 {
			logger.Debugln("block table in database is empty")
		}
		block_index++

				var dbtran DBTransaction
				dbtran, err = dbutil.GetLastTran()
				if err != nil {
					logger.Debugln("GetLastTran error")
				}
				logger.Debugln("dbtran=", dbtran)

			tx_index := dbtran.Tx_index + 1

		for {
			if block_index <= rpc.GetBlockCount() {
				logger.Debugln("Block: ",block_index)
				c := block_index
				requires_rollback := false
				for {
					if c == BLOCKFIRST {
						break
					}
					c_block,_ := rpc.GetBlockByIndex(c)
					bitcoind_parent = c_block.PreviousHash
					block,_ :=dbutil.GetBlock(c-1)
					if len(blocks) !=1{
						break
					}
					db_parent = block.Block_hash
					if db_parent == bitcoind_parent{
						break
					}else{
						c -= 1
						requires_rollback = true
					}
				}
				if requires_rollback{
					logger.Debugln("status:Blockchain reorganisation at block {}", c)
					dbutil.Reparse(c-1)
					block_index = c
					continue
				}
				block_hash := bitcoin.GetBlockHash(block_index)
				block:=bitcoin.GetBlock(block_hash)
				block_time := block.Time
				tx_hash_list :=block.Tx
				dbutil.InsertBlock(block_index,block_hash,block_time)
				for tx_hash in tx_hash_list{
					blocks := dbutil.GetTran(tx_hash)
					if blocks{
						tx_index += 1
						continue
					}
					tx := bitcoin.GetRawTransaction(tx_hash)
					logger.Debugln("Status: examining transaction ",tx_hash)
					source ,destination,btc_amount,fee,data := get_tx_info(tx,block_index)
					if source and (data or destination == WISHINGWALLADDRESS){
						dbutil.InsertTran(tx_index,tx_hash,block_index,block_hash,block_time,source,destination,btc_amount,fee,data)

					}
					tx_index +=1
				}
				parse_block(db,block_index,block_time)
			}
			block_count = bitcon.GetBlockCount()
			block_index +=1
		}else{
			time.sleep(2)
		}

	*/
}
