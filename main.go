package main

import (
	"github.com/astaxie/beego"
	_ "secret/routers"
	_ "secret/utils"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()

}
