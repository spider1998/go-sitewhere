package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"sync"
	"time"
	//导入mqtt包
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
var fail_nums int = 0

type Testmsg struct {
	DeviceToken string `json:"deviceToken"`
	Originator  string `json:"originator"`
	Type        string `json:"type"`
	Request     struct {
		Latitude    string `json:"latitude"`
		Longitude   string `json:"longitude"`
		Elevation   string `json:"elevation"`
		UpdateState bool   `json:"updateState"`
		EventDate   string `json:"eventDate"`
	} `json:"request"`
}

func main() {
	//生成连接的客户端数
	c := flag.Uint64("n", 1, "client nums")
	flag.Parse()
	nums := int(*c)
	wg := sync.WaitGroup{}
	for i := 0; i < nums; i++ {
		wg.Add(1)
		time.Sleep(5 * time.Millisecond)
		go createTask(i, &wg)
	}
	wg.Wait()
}

func createTask(taskId int, wg *sync.WaitGroup) {
	fmt.Println("ddd")
	defer wg.Done()
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.35.230:1883")
	opts.SetCleanSession(true)
	opts.SetClientID("your_clientID")
	/*	opts.SetDefaultPublishHandler(f)*/
	opts.SetConnectTimeout(time.Duration(5) * time.Second)

	//创建连接
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		fail_nums++
		fmt.Printf("taskId:%d,fail_nums:%d,error:%s \n", taskId, fail_nums, token.Error())
		return
	}

	//json消息示例
	var msg Testmsg
	msg.DeviceToken = "bracelet-002"
	msg.Originator = "admin"
	msg.Type = "DeviceLocation"
	msg.Request.Elevation = "0"
	msg.Request.EventDate = time.Now().Format(time.RFC3339)
	msg.Request.Latitude = "40"
	msg.Request.Longitude = "-80"
	msg.Request.UpdateState = true
	send, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	token := c.Publish("SiteWhere/ymzn/input/json", 1, true, send)
	if !token.Wait() {
		panic(token.Error())
	}
	fmt.Println("task ok!!")
}
