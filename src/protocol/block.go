package protocol

import (
	. "config"
	"database/sql"
	"fmt"
	_ "github.com/Wishing-Wall/go-sqlite3"
	"logger"
)

func init() {
	db, err := sql.Open("sqlite3", DATABASEPATH+"block.db")
	if err != nil {
		logger.Errorln("open block.db failed")
	}
	defer db.Close()

	_, err = db.Exec(DATABASE_CREATE_BLOCK)
	if err != nil {
		logger.Errorln("Create table blocks failed")
	}
}
