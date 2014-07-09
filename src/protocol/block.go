package protocol

import (
	. "config"
	"database/sql"
	"fmt"
	_ "github.com/Wishing-Wall/go-sqlite3"
	//"log"
)

func init() {
	_, err := sql.Open("sqlite3", DATABASEPATH+"block.db")
	if err != nil {
		fmt.Println("open block.db failed")
	} else {
		fmt.Println("open block.db succeed")
	}

}
