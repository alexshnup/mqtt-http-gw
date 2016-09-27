package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"] = append(beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"] = append(beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"],
		beego.ControllerComments{
			Method: "PostCmd",
			Router: `/cmd`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"] = append(beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:barrierId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"] = append(beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"],
		beego.ControllerComments{
			Method: "GetStatus",
			Router: `/get-status-adc/:barrierId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"] = append(beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"] = append(beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:barrierId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"] = append(beego.GlobalControllerRouter["github.com/alexshnup/mqtt-http-gw/controllers:BarrierController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:barrierId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
