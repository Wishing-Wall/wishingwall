package protocol

import (
	"fmt"
	"quickstart/conf"
	"quickstart/models/dbutil"
	"time"
)

func Parse_block(block_index, block_time uint64, previous_ledger_hash,
	previous_txlist_hash string) (string, string) {
	trans, err := dbutil.GetAllTransInBlock(block_index)
	if err != nil {
		var temp string
		return temp, temp
	}
	for _, tran := range trans {
		//Parse_tx(tran)

	}
	//util.BLOCK_LEDGER = []
	/*
	   cursor.execute('''SELECT * FROM transactions \
	                     WHERE block_index=? ORDER BY tx_index''',
	                 (block_index,))
	   txlist = []
	   for tx in list(cursor):
	       parse_tx(db, tx)
	       txlist.append('{}{}{}{}{}{}'.format(tx['tx_hash'], tx['source'], tx['destination'],
	                                           tx['btc_amount'], tx['fee'],
	                                           binascii.hexlify(tx['data']).decode('UTF-8')))

	   cursor.close()

	   # Consensus hashes.
	   new_txlist_hash = check.consensus_hash(db, 'txlist_hash', previous_txlist_hash, txlist)
	   new_ledger_hash = check.consensus_hash(db, 'ledger_hash', previous_ledger_hash, util.BLOCK_LEDGER)

	   return new_ledger_hash, new_txlist_hash
	*/
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
	go dbutil.DebugInsert()
	for {
		block_index, err := dbutil.LastBlockIndex()
		fmt.Printf("block_index =%d, err=%v\n", block_index, err)
		//tran, err := dbutil.GetLastTran()
		//fmt.Printf("last tran = %v, err=%v\n", tran, err)
		//block, err := dbutil.GetBlock(block_index - 2)
		//fmt.Printf("block %d is %v\n", block_index-2, block)

		//tran, err = dbutil.GetTran("test_tx_hash" + fmt.Sprintf("%d", block_index-1))
		//fmt.Printf("get tran by hash %v, %v\n", tran, err)
		//dbutil.Reinitialise(block_index - 1)
		//Reparse(block_index)
		trans, err := dbutil.GetAllTransInBlock(block_index)
		if err != nil {
			fmt.Printf("Failed to get")
		}
		for _, tran := range trans {
			fmt.Printf("tx_index is %d\n", tran.Tx_index)
		}
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
					Reparse(c-1)
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
