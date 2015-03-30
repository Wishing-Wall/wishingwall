package dbutil

import (
	. "config"
	"database/sql"
	"errors"
	//"fmt"
	_ "github.com/Wishing-Wall/go-sqlite3"
	"logger"
)

func CreateTable(path string, cmd string) error {
	logger.Debugln("path = ", path)
	logger.Debugln("cmd = ", cmd)
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		logger.Errorln(err)
		return err
	}
	defer db.Close()

	_, err = db.Exec(cmd)
	if err != nil {
		logger.Errorln("%q: %s\n", err, cmd)
		return err
	}
	db.Close()
	return nil
}

func LastBlockIndex() (uint64, error) {
	db, err := sql.Open("sqlite3", DATABASEPATH)
	if err != nil {
		logger.Errorln("open DATABASE error")
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM blocks WHERE block_index = (SELECT MAX(block_index) from blocks)")
	defer rows.Close()

	if err != nil {
		logger.Errorln("LastBlockIndex error")
		return 0, errors.New("Query error")
	}
	var block_index uint64
	rows.Scan(&block_index)
	db.Close()
	return block_index, nil

}

func ListLastTran() (DBTransaction, error) {
	db, err := sql.Open("sqlite3", DATABASEPATH)
	if err != nil {
		logger.Errorln("open DATABASE error")
	}
	defer db.Close()
	var data DBTransaction
	rows, err := db.Query("SELECT * FROM transactions WHERE tx_index = (SELECT MAX(tx_index) from transactions)")
	defer rows.Close()

	if err != nil {
		logger.Errorln("LastTranIndex error")
		return data, errors.New("Query error")
	}

	rows.Scan(&data.Tx_index, &data.Tx_hash, &data.Block_index, &data.Block_hash, &data.Block_time,
		&data.Source, &data.Destination, &data.Btc_amount, &data.Fee, &data.Data, &data.Supported)
	db.Close()
	return data, nil
}
