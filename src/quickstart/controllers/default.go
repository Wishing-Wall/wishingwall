package controllers

import (
	"fmt"
	. "quickstart/conf"
	"sort"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	o := orm.NewOrm()
	o.Using("default")
	// get Relaye addr
	send := new(DB_send)
	send.Message = c.GetString("clientmessage")
	fmt.Printf("client message is %v", send.Message)
	send.RelayAddr = send.Message + "testaddr"
	send.ConfirmTimes = 0
	send.CheckTimes = 0
	send.IsSent = false
	send.Succeed = false
	o.Insert(send)
	c.Ctx.Redirect(301, "addr")
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	var count int
	o.Raw("select max(id) from d_b_message").QueryRow(&count)
	fmt.Printf("count is %d\n", count)
	min := count - 10
	if min < 0 {
		min = 0
	}
	var messages DB_messages
	fmt.Printf("min %v max %v\r\n", min, count)
	o.Raw("select * from d_b_message where id >= ? and id <= ?", min, count).QueryRows(&messages)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	fmt.Printf("message before sort %v\n", messages)
	sort.Sort(sort.Reverse(messages))
	fmt.Printf("message after sort %v\n", messages)
	c.Data["messages"] = messages
	c.TplNames = "index.tpl"
}

type ShowAddrController struct {
	beego.Controller
}

func (c *ShowAddrController) Get() {
	c.TplNames = "showaddr.tpl"
}
