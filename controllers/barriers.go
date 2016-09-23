package controllers

import (
	"encoding/json"
	"strings"

	"github.com/alexshnup/mqtt-http-gw/mgtt-http-gw/models"
	"github.com/alexshnup/mqtt-http-gw/mgtt-http-gw/mqttmodule"

	_ "github.com/alexshnup/mqtt-http-gw/mgtt-http-gw/mqttmodule/conf"
	_ "github.com/alexshnup/mqtt-http-gw/mgtt-http-gw/mqttmodule/service"
	_ "github.com/alexshnup/mqtt-http-gw/mgtt-http-gw/mqttmodule/wirenboard"

	"github.com/astaxie/beego"
)

// Operations about barrier
type BarrierController struct {
	beego.Controller
}

// @Title Create
// @Description create barrier
// @Param	body		body 	models.Barrier	true		"The barrier content"
// @Success 200 {string} models.Barrier.Id
// @Failure 403 body is empty
// @router / [post]
func (o *BarrierController) Post() {
	var ob models.Barrier
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	barrierid := models.Add1(ob)
	o.Data["json"] = map[string]string{"BarrierId": barrierid}
	o.ServeJSON()
}

// @Title Cmd
// @Description create barrier
// @Param	body		body 	models.Barrier	true		"The barrier content"
// @Success 200 {string} models.Barrier.Id
// @Failure 403 body is empty
// @router /cmd [post]
func (o *BarrierController) PostCmd() {
	var ob models.Barrier
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	barrierid := models.Add1(ob)

	topicEnd := strings.Replace(ob.BarrierId, "-", "/", -1)

	mqttmodule.WB.System.Relay.PublishPayload(0, topicEnd, ob.Payload)

	o.Data["json"] = map[string]string{"BarrierId": barrierid}
	o.ServeJSON()
}

// @Title Get
// @Description find barrier by barrierid
// @Param	barrierId		path 	string	true		"the barrierid you want to get"
// @Success 200 {barrier} models.Barrier
// @Failure 403 :barrierId is empty
// @router /:barrierId [get]
func (o *BarrierController) Get() {
	barrierId := o.Ctx.Input.Param(":barrierId")
	if barrierId != "" {
		ob, err := models.Get1(barrierId)
		if err != nil {
			o.Data["json"] = err.Error()
		} else {
			o.Data["json"] = ob
		}
	}
	o.ServeJSON()
}

// @Title GetAll
// @Description get all barriers
// @Success 200 {barrier} models.Barrier
// @Failure 403 :barrierId is empty
// @router / [get]
func (o *BarrierController) GetAll() {
	obs := models.GetAllAll()
	o.Data["json"] = obs
	o.ServeJSON()
}

// @Title Update
// @Description update the barrier
// @Param	barrierId		path 	string	true		"The barrierid you want to update"
// @Param	body		body 	models.Barrier	true		"The body"
// @Success 200 {barrier} models.Barrier
// @Failure 403 :barrierId is empty
// @router /:barrierId [put]
func (o *BarrierController) Put() {
	barrierId := o.Ctx.Input.Param(":barrierId")
	var ob models.Barrier
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update1(barrierId, ob.Payload)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the barrier
// @Param	barrierId		path 	string	true		"The barrierId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 barrierId is empty
// @router /:barrierId [delete]
func (o *BarrierController) Delete() {
	barrierId := o.Ctx.Input.Param(":barrierId")
	models.Delete1(barrierId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}
