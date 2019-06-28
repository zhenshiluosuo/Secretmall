package main

import (
	"github.com/astaxie/beego"
	//secret/routers
	//secret/utils
	_ "secret/Secretmall/routers"
	_ "secret/Secretmall/utils"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()

}
