package controllers

import (
	"github.com/astaxie/beego"
	"secret/Secretmall/models"
	"strings"
)

type CertController struct {
	beego.Controller
}

/*
   验证用户身份返回购物车
*/
func (this *CertController) Get() {
	this.TplName = "cert.tpl"
	var user models.User
	if this.GetSession("loginuser") == nil {
		this.Redirect("/myaccount", 301)
	}
	username := this.GetSession("loginuser").(string)
	user.Name = username
	items := models.Check_cert(user)
	for _, x := range items {
		x.Descirption = strings.Trim(x.Descirption, " ")
	}
	this.Data["item"] = items
}
func (this *CertController) Post() {
	this.TplName = "cert.tpl"
}
