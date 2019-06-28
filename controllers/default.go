package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	//"secret/models"
	"secret/Secretmall/models"
	//_ "secret/utils"
	_ "secret/Secretmall/utils"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type RegController struct {
	beego.Controller
}

func (this *RegController) Get() {
	this.TplName = "myaccount/login.tpl"
}
func (this *RegController) Post() {
	Name := this.GetString("username")
	Pwd := this.GetString("password")

	//查询用户名
	user := models.Query_single_user(Name)

	if user.Password == Pwd {
		//验证成功，登录完成
		this.SetSession("loginuser", Name)
		this.Redirect("/", 301)
	} else {
		this.Redirect("/myaccount", 301)
	}
}

type ChaoshiController struct {
	beego.Controller
}

func (this *ChaoshiController) Get() {
	item := models.Item_List(5, 0)
	for _, x := range item {
		x.Descirption = strings.Trim(x.Descirption, " ")
	}
	this.Data["item"] = item
	this.TplName = "itemlist.tpl"
}

type Reg1Controller struct {
	beego.Controller
}

func (this *Reg1Controller) Get() {
	this.TplName = "myaccount/regist.tpl"
}
