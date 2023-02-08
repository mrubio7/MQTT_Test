package mqttUtils

import (
	"fmt"
	"gomqtt/consts"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func CreateMqttClient() mqtt.Client {
	connectAddress := fmt.Sprintf("%s:%d", consts.Broker, consts.Port)
	client_id := fmt.Sprintf("go-client-%d", 1)

	fmt.Println("connect address: ", connectAddress)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername(consts.Username)
	opts.SetPassword(consts.Password)
	opts.SetClientID(client_id)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	// if connection failed, exit
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}

func CreateMqttClientSubscriber() mqtt.Client {
	connectAddress := fmt.Sprintf("%s:%d", consts.Broker, consts.Port)
	client_id := fmt.Sprintf("go-client-%d", 2)

	fmt.Println("connect address: ", connectAddress)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(connectAddress)
	opts.SetUsername(consts.Username)
	opts.SetPassword(consts.Password)
	opts.SetClientID(client_id)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetDefaultPublishHandler(messagePubHandler)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	// if connection failed, exit
	if token.WaitTimeout(3*time.Second) && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return client
}
