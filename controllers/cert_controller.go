package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"secret/models"
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
	username := this.GetSession("loginuser").(string)
	user.Name = username
	var items []models.Item
	items = models.Check_cert(user)
	fmt.Println(items)
}
func (this *CertController) Post() {
	this.TplName = "myaccount/cert.tpl"
}
