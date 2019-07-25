package main

import (
	`encoding/json`
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
	DeviceToken string `json:"device_Token"`
	Originator string `json:"originator"`
	Type string `json:"type"`
	Request struct{
		Latitude string `json:"latitude"`
		Longitude string `json:"longitude"`
		Elevation string `json:"elevation"`
		UpdateState bool `json:"updateState"`
		EventDate string `json:"eventDate"`
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
	opts.SetClientID("sample1")
/*	opts.SetDefaultPublishHandler(f)*/
	opts.SetConnectTimeout(time.Duration(60) * time.Second)

	//创建连接
	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.WaitTimeout(time.Duration(60)*time.Second) && token.Wait() && token.Error() != nil {
		fail_nums++
		fmt.Printf("taskId:%d,fail_nums:%d,error:%s \n", taskId, fail_nums, token.Error())
		return
	}

	//每隔5秒向topic发送一条消息
		var msg Testmsg
		msg.DeviceToken= "8a504a07-aad7-4e23-a6de-8e84a5a82ce2"
		msg.Originator = "100002"
		msg.Type = "DeviceLocation"
		msg.Request.Elevation = "0"
		msg.Request.EventDate = "2019-07-25T19:40:03.390Z"
		msg.Request.Latitude = "40"
		msg.Request.Longitude = "-80"
		msg.Request.UpdateState = true
		fmt.Println(msg)
		s,_ := json.Marshal(msg)
		token := c.Publish("SiteWhere/default/input/json", 0, false, s)
		if !token.Wait(){
			panic(token.Error())
		}

	c.Disconnect(250)
	fmt.Println("task ok!!")
}