package protocol

import (
	"database/sql"
	"fmt"
	_ "github.com/Wishing-Wall/go-sqlite3"
)

func block() {
	db, err := sql.Open("sqlite3", "./block.db")
	if err != nil {
		fmt.Println("open block.db failed")

	} else {

		fmt.Println("open block.db succeed")
	}

}
