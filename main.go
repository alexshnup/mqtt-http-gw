package main

import (
	"github.com/alexshnup/mqtt-http-gw/mqttmodule"
	_ "github.com/alexshnup/mqtt-http-gw/routers"

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
