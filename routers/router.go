package routers

import (
	"github.com/astaxie/beego"
	"secret/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/myaccount", &controllers.RegController{})
}