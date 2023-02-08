package main

import (
	"fmt"
	"gomqtt/mqttUtils"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var client mqtt.Client

func main() {
	fmt.Printf("STARTING CONSUMER\n")

	client = mqttUtils.CreateMqttClientSubscriber()

	topic := "test/a"

	go client.Subscribe(topic, byte(1), nil)
	for {
		time.Sleep(time.Second * 1)
	}
}
