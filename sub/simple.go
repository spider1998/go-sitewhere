package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	`github.com/gogo/protobuf/proto`
	"os"
	"os/signal"
	pro "sdkeji/go_mqtt/proto"
	"syscall"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func readMessage(client MQTT.Client, msg MQTT.Message) {
		stReceive := &pro.TestInfo{}
		//protobuf解码
		err := proto.Unmarshal(msg.Payload(), stReceive)
		if err != nil {
			panic(err)
		}
		fmt.Printf("receive: %s\n",stReceive)
}


func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.35.190:1883")
	opts.SetClientID("sample5")
	opts.SetDefaultPublishHandler(f)
	topic := "SiteWhere/default/input/json"
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