package dbutil

import (
	//. "config"
	"database/sql"
	//"fmt"
	_ "github.com/Wishing-Wall/go-sqlite3"
	"logger"
)

func CreateTable(path string, cmd string) error {
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
