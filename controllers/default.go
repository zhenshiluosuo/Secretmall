package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"secret/models"
	_ "secret/utils"

)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type RegController struct{
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

	if user.Password== Pwd{
		//验证成功，登录完成


















































































































































































































































		fmt.Println(Name+"登录成功"+user.Password)
		this.SetSession("loginuser", Name)
		this.Redirect("/", 301)
	}else {
		this.Redirect("/myaccount",301)
	}
}