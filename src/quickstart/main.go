package main

import (
	_ "fmt"
	"github.com/astaxie/beego"
	_ "quickstart/models/protocol"
	_ "quickstart/routers"
)

func main() {

	beego.Run()
}
