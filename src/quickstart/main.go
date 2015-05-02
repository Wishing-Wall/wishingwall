package main

import (
	. "LastBlockIndex"
	_ "fmt"
	"github.com/astaxie/beego"
	_ "quickstart/routers"
)

func main() {
	go Follow()
	beego.Run()
}
