package wirenboard

import (
	mqtt "github.com/alexshnup/mqtt"

	"gitlab.com/gosparom/mgtt-http-gw/mqttmodule/wirenboard/core"
)

// Wirenboard struct
type Wirenboard struct {
	client mqtt.Client
	name   string
	System *syscore.System
}

func NewWirenboard(c mqtt.Client, name string, debug bool) *Wirenboard {
	return &Wirenboard{
		client: c,
		name:   name,
		System: syscore.NewC2000(c, name, debug),
	}
}
