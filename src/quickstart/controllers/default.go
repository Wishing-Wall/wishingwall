package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "quickstart/conf"
	"sort"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	var count int
	o.Raw("select count(*) from d_b_message").QueryRow(&count)
	fmt.Printf("count is %d\n", count)
	min := count - 10
	if min < 0 {
		min = 0
	}
	var messages DB_messages
	o.Raw("select id,"+
		"Message_index,"+
		"Block_index,"+
		"Block_hash,"+
		"Block_time,"+
		"Tx_index,"+
		"Tx_hash,"+
		"Account,"+
		"Source,"+
		"Destination,"+
		"Message "+
		"from d_b_message where id >= ? and id <= ?", min, count).QueryRows(&messages)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//fmt.Printf("message before sort %v\n", messages)
	sort.Sort(sort.Reverse(messages))
	//fmt.Printf("message after sort %v\n", messages)
	c.Data["messages"] = messages
	c.TplNames = "index.tpl"
}
