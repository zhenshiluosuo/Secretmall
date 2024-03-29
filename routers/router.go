package routers

import (
	"github.com/astaxie/beego"
	"secret/Secretmall/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/myaccount", &controllers.RegController{})
	beego.Router("/myaccount/reg", &controllers.Reg1Controller{})
	beego.Router("/chaoshi", &controllers.ChaoshiController{})
	beego.Router("/cert", &controllers.CertController{})
}
