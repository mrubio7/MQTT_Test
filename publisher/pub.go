package main

import (
	"gomqtt/mqttUtils"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

func main() {
	client = mqttUtils.CreateMqttClient()

	for {
		client.Publish("test/a", byte(2), false, "Enviado!")
		time.Sleep(time.Second * 2)
	}
}
