package main

import (
	"github.com/astaxie/beego"
	//secret/routers
	//secret/utils
	_ "secret/routers"
	_ "secret/utils"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()

}
