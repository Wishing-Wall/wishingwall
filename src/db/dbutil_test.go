package dbutil

import (
	. "config"
	"logger"
	"testing"
)

func TestCreateTable(t *testing.T) {
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

func TestLastBlockIndex(t *testing.T) {
	block_index, err := LastBlockIndex()
	logger.Debugln("block_index ", block_index)
	if err != nil {
		t.Fail()
	} else  {
		return
	}
}
