package controllers

import (
	"fmt"
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
	fmt.Println("session", this.GetSession("loginuser"))
	if this.GetSession("loginuser") == nil {
		this.Redirect("/myaccount", 302)
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
	name := this.GetSession("loginuser").(string)
	id, _ := this.GetInt("item_id")
	count, _ := this.GetInt("count")

	u := models.Query_single_user(name)
	models.Add_to_cert(0, u.Id, id, count)

}
