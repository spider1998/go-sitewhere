package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	p "github.com/gogo/protobuf/proto"
	"os"
	"os/signal"
	"sdkeji/go_mqtt/proto"
	"syscall"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
	var ss proto.DeviceEvent
	_ = p.Unmarshal(msg.Payload(), &ss)
	fmt.Println(ss)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.35.230:1883")
	opts.SetClientID("sample5")
	opts.SetDefaultPublishHandler(f)
	topic := "SiteWhere/ymzn/input/protobuf"
	opts.OnConnect = func(c MQTT.Client) {
		if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to server\n")
	}
	<-c
}
