package controllers

import (
	"fmt"
	"quickstart/conf"
	"quickstart/models/dbutil"
	"quickstart/models/protocol"
	"quickstart/models/wallet"
	"sort"

	"github.com/astaxie/beego"
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
	pagestart, _ := c.GetInt("s")
	pageend, _ := c.GetInt("e")

	count := dbutil.LastMessageIndex()
	if pagestart == 0 && pageend == 0 {
		pageend = count
	}

	if pageend > count {
		pageend = count
	}
	if pagestart < pageend-10 {
		pagestart = pageend - 10
	}
	if pagestart < 0 {
		pagestart = 0
	}

	messages := dbutil.GetMessages(pagestart, pageend)

	nextend := pageend + 10
	if nextend > count {
		nextend = count
	}
	nextstart := nextend - 10
	if nextstart < pagestart {
		nextstart = pagestart
	}
	prestart := pagestart - 10
	if prestart < 0 {
		prestart = 0
	}
	preend := prestart + 10
	if preend > pageend {
		preend = pageend
	}

	sort.Sort(sort.Reverse(messages))
	c.Data["messages"] = messages
	c.Data["prestart"] = prestart
	c.Data["preend"] = preend
	c.Data["nextstart"] = nextstart
	c.Data["nextend"] = nextend
	c.Data["maxblockindex"] = protocol.GlobalLastBlockIndex
	c.Data["parsedblockindex"] = protocol.GlobalParsedBlockIndex
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
