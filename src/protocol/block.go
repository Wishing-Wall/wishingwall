package protocol

import (
	. "config"

	"dbutil"
	"fmt"
	"logger"
)

func init() {
	tables := [...]string{
		DATABASE_CREATE_BLOCK,
		DATABASE_CREATE_BLOCK_INDEX,
		DATABASE_CREATE_BLOCK_HASH_INDEX,
		DATABASE_TRANSACTION,
		DATABASE_CREATE_TRANSACTION_INDEX,
		DATABASE_CREATE_TRANSACTION_INDEX2,
		DATABASE_CREATE_TRANSACTION_INDEX3,
		DATABASE_CREATE_TRANSACTION_INDEX4,
		DATABASE_CREATE_MESSAGE,
		DATABASE_CREATE_MESSAGE_INDEX,
		DATABASE_CREATE_MESSAGE_INDEX1,
		DATABASE_CREATE_MESSAGE_INDEX2,
		DATABASE_CREATE_MESSAGE_INDEX3,
	}
	for _, v := range tables {
		err := CreateTable(DATABASEPATH, v)
		if err != nil {
			t.Fail()
		} else {
			return
		}
	}
}

func Follow() {
	logger.Infoln("Start... ")
	//	block_index,err := dbutil.LastBlockIndex()
	//		if err != nil{
	//			logger.Errorln("Get lastblockindex failed")
	//			return
	//		}
	//		if block_index == 0{
	//			logger.Debugln("block table in database is empty")
	//		}
	//		block_index ++
	//
	//		var dbtran DBTransaction
	//		dbtran,err = ListLastTran()
	//		if err!=nil{
	//			logger.Debugln("ListLastTran error")
	//		}
	//		logger.Debugln("dbtran.Tx_index=",dbtran.Tx_index)
	//		tx_index := dbtran.Tx_index + 1

}
