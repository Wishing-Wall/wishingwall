package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "quickstart/conf"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	message := new(DB_message)
	message.Id = 10
	o.Read(message)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["Message"] = message.Message
	c.TplNames = "index.tpl"
}
