package syscore

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	mqtt "github.com/alexshnup/mqtt"
	"github.com/alexshnup/mqtt-http-gw/mgtt-http-gw/models"
)

/*
Struct relay provides system LED[0,1] control

Topics:
	Subscribe:
		name + "/SYSTEM/LED[0,1]/ACTION		{0, 1, STATUS}
	Publish:
		name + "/SYSTEM/LED[0,1]/STATUS		{0, 1}

Methods:
	Subscribe
	Unsubscribe
	PublishPayload

Functions:
	Set trigger to [none] when subscribe
		echo none | sudo tee /sys/class/relays/relay0/trigger
	Set trigger to [mmc0] when unsubscribe
		echo mmc0 | sudo tee /sys/class/relays/relay0/trigger
	Set brightness to 1 when ON
		echo 1 | sudo tee /sys/class/relays/relay0/brightness
	Set brightness to 0 when OFF
		echo 0 | sudo tee /sys/class/relays/relay0/brightness
	Get brightness status
		sudo cat /sys/class/relays/relay0/brightness

TODO:
[ ] catch errors in relayMessageHandler

*/
type RelaY struct {
	client   mqtt.Client
	debug    bool
	topic    string
	status   string
	relayID  string
	deviceID string
}

// a[len(a)-1:] last char

// newRelay return new RelaY object.
func newRelay(c mqtt.Client, topic string, debug bool) *RelaY {
	return &RelaY{
		client:   c,
		debug:    debug,
		topic:    topic,
		status:   "0",
		relayID:  topic[len(topic)-1:],
		deviceID: topic[len(topic)-3:],
	}
}

// Subscribe to topic
func (L *RelaY) Subscribe(qos byte) {

	topic := L.topic + "/#"

	log.Println("[RUN] Subscribing:", qos, topic)

	if token := L.client.Subscribe(topic, qos, L.relayMessageHandler); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

// UnSubscribe from topic
func (L *RelaY) UnSubscribe() {

	topic := L.topic + "/#"

	log.Println("[RUN] UnSubscribing:", topic)

	if token := L.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}
}

// PublishPayload Relay status
func (L *RelaY) PublishPayload(qos byte, topicEnd string, payload string) {
	log.Println("test PUB")

	topic := L.topic + "/" + topicEnd

	// publish result
	if token := L.client.Publish(topic, qos, false, payload); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// debug
	if L.debug {
		log.Println("[PUB]", qos, topic, L.status)
	}
}

// PublishPayload Relay status
func (L *RelaY) PublishPayloadSensor(qos byte, deviceID, relayID string) {

	topic := L.topic + "/" + deviceID + "/" + relayID + "/status/sensor"

	// publish result
	if token := L.client.Publish(topic, qos, false, L.status); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// debug
	if L.debug {
		log.Println("[PUB]", qos, topic, L.status)
	}
}

// PublishPayload Relay status
func (L *RelaY) PublishADC(qos byte, deviceID, adcID string) {

	topic := L.topic + "/" + deviceID + "/" + adcID + "/status/adc"

	// publish result
	if token := L.client.Publish(topic, qos, false, L.status); token.Wait() && token.Error() != nil {
		log.Println(token.Error())
	}

	// debug
	if L.debug {
		log.Println("[PUB]", qos, topic, L.status)
	}
}

// relayMessageHandler set Relay to ON or OFF and get STATUS
func (L *RelaY) relayMessageHandler(client mqtt.Client, msg mqtt.Message) {

	// debug
	if L.debug {
		log.Println("[SUB]", msg.Qos(), msg.Topic(), string(msg.Payload()))
	}

	s1 := strings.Replace(msg.Topic(), L.topic, "", 1)
	s1 = strings.Replace(s1, "/", " ", -1)

	s_fields := strings.Fields(s1)
	// device_id, _ := strconv.ParseUint(s_fields[0], 10, 64)
	relay_id, _ := strconv.ParseUint(s_fields[1], 10, 64)

	switch len(s_fields) {
	case 3:
		er := models.Update1(s_fields[0]+"-"+s_fields[1]+"-"+s_fields[2], string(msg.Payload()))
		if er != nil {
			br := models.Barrier{}
			br.BarrierId = s_fields[0] + "-" + s_fields[1] + "-" + s_fields[2]
			br.Payload = string(msg.Payload())
			models.Add1(br)
		}
	case 4:
		er := models.Update1(s_fields[0]+"-"+s_fields[1]+"-"+s_fields[2]+"-"+s_fields[3], string(msg.Payload()))
		if er != nil {
			br := models.Barrier{}
			br.BarrierId = s_fields[0] + "-" + s_fields[1] + "-" + s_fields[2] + "-" + s_fields[3]
			br.Payload = string(msg.Payload())
			models.Add1(br)
		}

	}

	fmt.Println(s_fields[0], relay_id)

	// switch s_fields[2] {
	// case "on":
	// 	// receive message and DO
	// 	switch string(msg.Payload()) {
	// 	case "0":
	// 		// logic when OFF
	//
	// 		L.status = RelayOnOff(uint8(device_id), uint8(relay_id), 0)
	// 		log.Println("L.status", L.status)
	//
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	case "1":
	// 		// logic when ON
	//
	// 		L.status = RelayOnOff(uint8(device_id), uint8(relay_id), 1)
	// 		log.Println("L.status", L.status)
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	case "3":
	// 		// logic when ON
	//
	// 		L.status = RelayWhile(uint8(device_id), uint8(relay_id), 3)
	// 		log.Println("L.status", L.status)
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	}
	//
	// case "sensor":
	// 	// publish status
	// 	L.status = Payload(uint8(device_id), uint8(relay_id))
	// 	log.Println("L.status", L.status)
	// 	L.PublishPayloadSensor(0, s_fields[0], s_fields[1])
	//
	// case "setrelaydefaultmode":
	// 	// receive message and DO
	// 	switch string(msg.Payload()) {
	// 	case "off":
	// 		// publish status
	// 		L.status = SetConfig(uint8(device_id), uint8(relay_id), 2)
	// 		log.Println("L.status", L.status)
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	case "on":
	// 		// publish status
	// 		L.status = SetConfig(uint8(device_id), uint8(relay_id), 1)
	// 		log.Println("L.status", L.status)
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	case "blink":
	// 		// publish status
	// 		L.status = SetConfig(uint8(device_id), uint8(relay_id), 9)
	// 		log.Println("L.status", L.status)
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	case "pcn":
	// 		// publish status
	// 		L.status = SetConfig(uint8(device_id), uint8(relay_id), 10)
	// 		log.Println("L.status", L.status)
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	}
	//
	// case "setrelaytime":
	// 	// receive message and DO
	// 	v, _ := strconv.ParseUint(string(msg.Payload()), 10, 64)
	// 	if v >= 1 && v <= 60 {
	// 		L.status = SetConfig(uint8(device_id), uint8(relay_id)+4, uint8(v))
	// 		log.Println("L.status", L.status)
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	}
	//
	// case "changeaddress":
	// 	// receive message and DO
	// 	var newaddr uint8 = uint8(relay_id)
	// 	if newaddr >= 1 && newaddr <= 127 {
	// 		L.status = ChangeAddress(uint8(device_id), newaddr)
	// 		log.Println("L.status", L.status)
	// 		L.PublishPayload(0, s_fields[0], s_fields[1])
	// 	}
	//
	// case "adc":
	// 	// publish status
	// 	L.status = ADC(uint8(device_id), uint8(relay_id))
	// 	log.Println("L.status", L.status)
	// 	L.PublishADC(0, s_fields[0], s_fields[1])
	// }

}
