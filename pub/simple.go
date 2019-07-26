package main

import (
	"flag"
	"fmt"
	p `github.com/golang/protobuf/proto`
	`sdkeji/go_mqtt/proto`
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
	/*var msg Testmsg
	msg.DeviceToken = "bracelet-002"
	msg.Originator = "admin"
	msg.Type = "DeviceLocation"
	msg.Request.Elevation = "0"
	msg.Request.EventDate = time.Now().Format(time.RFC3339)
	msg.Request.Latitude = "40"
	msg.Request.Longitude = "-80"
	msg.Request.UpdateState = true*/
	var header proto.DeviceEvent_Header
	header.DeviceToken = &proto.GOptionalString{Value:"rest-token-4"}
	header.Originator = &proto.GOptionalString{Value:"admin"}
	//header.Command = proto.DeviceEvent_SendLocation
	header.Command = proto.DeviceEvent_SendLocation

	var body1 proto.DeviceEvent_DeviceMeasurement
	body1.MeasurementName = &proto.GOptionalString{Value:"tmp"}
	body1.MeasurementValue= &proto.GOptionalDouble{Value:42}
	//body.Latitude = &proto.GOptionalDouble{Value:float64(39.31)}
	//body.Longitude = &proto.GOptionalDouble{Value:float64(-85.52)}
	body1.UpdateState = &proto.GOptionalBoolean{Value:true}
	//body.Elevation = &proto.GOptionalDouble{Value:float64(2)}
	t := uint64(time.Now().Unix())
	body1.EventDate = &proto.GOptionalFixed64{Value:t}
	/*body.Metadata = map[string]string{
		"latitude" :"30",
		"longitude":"-98",
	}*/



	/*var header proto.DeviceEvent_Header
	header.Command = "SendLocation"
	header.Originator = "admin"
	header.DeviceToken = "bracelet-002"
	var msg proto.DeviceEvent_DeviceLocation
	msg.Elevation = float64(0)
	msg.EventDate = uint64(time.Now().Unix())
	msg.Latitude = float64(30)
	msg.Longitude = float64(-60)
	msg.UpdateState = true*/
	send1,err := p.Marshal(&header)
	if err != nil{
		panic(err)
	}
	send2,err := p.Marshal(&body1)
	if err != nil{
		panic(err)
	}
	var send []byte
	send = append(send,byte(len(send1)))
	send = append(send,send1...)
	send = append(send,byte(len(send2)))
	send = append(send,send2...)
	fmt.Println(send)
	token := c.Publish("SiteWhere/ymzn/input/protobuf", 1, true,send)
	if !token.Wait() {
		panic(token.Error())
	}
	c.Disconnect(250)
	fmt.Println("task ok!!")
}
