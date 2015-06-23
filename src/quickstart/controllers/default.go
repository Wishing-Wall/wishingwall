package controllers

import (
	"fmt"
	. "quickstart/conf"
	"quickstart/models/dbutil"
	"quickstart/models/wallet"
	"sort"

	"quickstart/conf"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Post() {
	/*
		o := orm.NewOrm()
		o.Using("default")
		// get Relaye addr
		send := new(DB_send)
		send.Message = c.GetString("clientmessage")
		fmt.Printf("client message is %v", send.Message)
		send.RelayAddr = send.Message + "testaddr"

		send.IsSent = false

		o.Insert(send)
	*/
	message := c.GetString("clientmessage")
	replyaddress, err := wallet.GetNewAddress()
	if err != nil {
		return
	}

	_, minmoney, err := wallet.CreateRawTransaction(replyaddress, message)

	dbutil.InsertSend(replyaddress, message, minmoney)

	min := float64(minmoney) / float64(conf.COIN)

	c.Ctx.Redirect(301, replyaddress+string('/')+fmt.Sprintf("%v", min))
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	var count int
	o.Raw("select max(id) from d_b_message").QueryRow(&count)
	//fmt.Printf("count is %d\n", count)
	min := count - 10
	if min < 0 {
		min = 0
	}
	var messages DB_messages
	//fmt.Printf("min %v max %v\r\n", min, count)
	o.Raw("select * from d_b_message where id >= ? and id <= ?", min, count).QueryRows(&messages)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//fmt.Printf("message before sort %v\n", messages)
	sort.Sort(sort.Reverse(messages))
	//fmt.Printf("message after sort %v\n", messages)
	c.Data["messages"] = messages
	c.TplNames = "index.tpl"
}

type ShowAddrController struct {
	beego.Controller
}

func (c *ShowAddrController) Get() {
	replyaddress := c.GetString(":address")
	minmoney, _ := c.GetFloat(":minmoney")
	c.Data["address"] = replyaddress
	c.Data["minmoney"] = minmoney

	c.TplNames = "showaddr.tpl"
}
