package mqttmodule

import
// "flag"
(
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexshnup/mqtt-http-gw/mqttmodule/conf"
	"github.com/alexshnup/mqtt-http-gw/mqttmodule/service"
	"github.com/alexshnup/mqtt-http-gw/mqttmodule/wirenboard"
)

// checkError check error
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var WB *wirenboard.Wirenboard

func MqttModule() {
	fmt.Printf("%v", conf.Config.Mqtt.Address)

	// interrupt
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	// open mqtt connection
	client, err := service.NewMqttClient(
		conf.Config.Mqtt.Protocol,
		conf.Config.Mqtt.Address,
		conf.Config.Mqtt.Port,
		0,
	)
	checkError(err)

	// new instance of mqtt client
	WB = wirenboard.NewWirenboard(client, conf.Config.Name, conf.Config.Debug)

	// Run publisher

	// WB.System.Memory.Publish(Config.Timeout, 0)

	// Run subscribing
	WB.System.Relay.Subscribe(2)
	// WB.System.Relay1.Subscribe(2)
	// WB.System.Relay2.Subscribe(2)
	// WB.System.Relay3.Subscribe(2)
	// WB.System.Relay4.Subscribe(2)

	// wait for terminating
	for {
		select {
		case <-interrupt:
			log.Println("Clean and terminating...")

			// Unsubscribe when terminating
			WB.System.Relay.UnSubscribe()
			// WB.System.Relay1.UnSubscribe()
			// WB.System.Relay2.UnSubscribe()
			// WB.System.Relay3.UnSubscribe()
			// WB.System.Relay4.UnSubscribe()

			// disconnecting
			client.Disconnect(250)

			os.Exit(0)
		}
	}

	// LoadJSONConfig()
	//
	// fmt.Println(ConfigJSON.Address)

	// for _, m := range ConfigJSON.Relays {
	// 	fmt.Println(m)
	// 	Relay_ON(ConfigJSON.Address, m)
	// 	StatusRelay(ConfigJSON.Address, m)
	// 	time.Sleep(100 * time.Millisecond)
	// 	Relay_OFF(ConfigJSON.Address, m)
	// 	StatusRelay(ConfigJSON.Address, m)
	// }

}
