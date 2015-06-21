package wallet

import (
	"fmt"
	"testing"
)

func TestCreateRawTransaction(*testing.T) {
	MsgTx, _, err := CreateRawTransaction("mofW99FBxa7g3kPaw2pAn6eXQX23EXXx2p",
		"I love you霍霍I love you 霍霍abcdefghijklmnopqrtsuvwxyzABCDEFHI love you霍霍I love you 霍霍")
	if err != nil {
		fmt.Printf("error is %v", err)
	}
	hash, err := SendRawTransaction(MsgTx)
	fmt.Printf("hash is %v\r\n", hash)
	fmt.Printf("err is %v\r\n", err)
}
