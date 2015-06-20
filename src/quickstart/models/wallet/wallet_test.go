package wallet

import (
	"fmt"
	"testing"
)

func TestCreateRawTransaction(*testing.T) {
	_, _, err := CreateRawTransaction("mofW99FBxa7g3kPaw2pAn6eXQX23EXXx2p",
		"I love you霍霍I love you 霍霍abcdefghijklmnopqrtsuvwxyzABCDEFHI love you霍霍I love you 霍霍")
	if err != nil {
		fmt.Printf("error is %v", err)
	}
}
