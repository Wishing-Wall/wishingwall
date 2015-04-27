package main

import (
	_ "fmt"
	"github.com/astaxie/beego"
	. "quickstart/models/protocol"
	_ "quickstart/routers"
)

func main() {
	go Follow()
	beego.Run()
}
