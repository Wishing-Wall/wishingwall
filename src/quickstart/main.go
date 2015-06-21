package main

import (
	_ "fmt"

	"quickstart/models/dbutil"
	"quickstart/models/protocol"
	"quickstart/models/wallet"
	_ "quickstart/routers"

	"github.com/astaxie/beego"
)

func main() {
	go protocol.Follow()
	go protocol.SendLoop()
	_, Min, _ := wallet.CreateRawTransaction("mi5mCfD3r84UKEbkexHr429j1rpJvWGt6U",
		"I miss you HUO")
	dbutil.InsertSend("mi5mCfD3r84UKEbkexHr429j1rpJvWGt6U", "I miss you HUO", Min)
	beego.Run()
}
