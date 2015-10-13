package main

import (
	_ "fmt"

	_  "quickstart/models/protocol"
	_ "quickstart/routers"

	"github.com/astaxie/beego"
)

func main() {
	//go protocol.Follow()
	//go protocol.SendLoop()
	beego.Run()
}
