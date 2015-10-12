package routers

import (
	"quickstart/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/:address/:minmoney", &controllers.ShowAddrController{})
	beego.Router("/toolongmessage", &controllers.TooLongMessage{})
}
