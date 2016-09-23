package main

import (
	"gitlab.com/gosparom/mgtt-http-gw/mqttmodule"
	_ "gitlab.com/gosparom/mgtt-http-gw/routers"

	"github.com/astaxie/beego"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	go mqttmodule.MqttModule()
	beego.Run()
}
