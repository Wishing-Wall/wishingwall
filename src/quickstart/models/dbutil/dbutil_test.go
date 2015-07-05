package dbutil

import (
	"fmt"
	"testing"
	"time"
)

func TestGetAllUnsentMessage(*testing.T) {
	send, err := GetAllUnsentMessage(time.Hour)
	fmt.Printf("err is %v\r\n", err)
	fmt.Printf("send is %v\r\n", send)
}

func TestInsertSend(*testing.T) {
	err := InsertSend("replyaddress", "thisismessage", 10000)
	fmt.Printf("Insersned %v\r\n", err)
}
