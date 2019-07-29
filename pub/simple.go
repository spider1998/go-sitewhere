package main

import (
	"flag"
	"fmt"
	proto2 `github.com/gogo/protobuf/proto`
	//p `github.com/golang/protobuf/proto`
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
	opts := MQTT.NewClientOptions().AddBroker("tcp://192.168.35.230:1883").SetUsername("admin").SetPassword("password")
	opts.SetCleanSession(true)
	opts.SetClientID("sample1")
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
	header.DeviceToken = &proto.GOptionalString{Value:"rest-token-4"}	//设备token
	header.Originator = &proto.GOptionalString{Value:"admin"}			//操作人
	//header.Command = proto.DeviceEvent_SendLocation
	header.Command = proto.DeviceEvent_SendLocation						//命令

	var body1 proto.DeviceEvent_DeviceLocation
	//body1.MeasurementName = &proto.GOptionalString{Value:"Latitude"}
	//body1.MeasurementValue= &proto.GOptionalDouble{Value:42}
	body1.Latitude = &proto.GOptionalDouble{Value:float64(39.31)}
	body1.Longitude = &proto.GOptionalDouble{Value:float64(-85.52)}
	body1.UpdateState = &proto.GOptionalBoolean{Value:true}
	body1.Elevation = &proto.GOptionalDouble{Value:float64(2)}
	t := uint64(time.Now().UnixNano()/1e6)
	body1.EventDate = &proto.GOptionalFixed64{Value:t}
	body1.Metadata = map[string]string{
		"Latitude" :"30",
		"Longitude":"-98",
	}

	send1,err :=proto2.Marshal(&header)
	if err != nil{
		panic(err)
	}
	send2,err := proto2.Marshal(&body1)
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
	//c.Disconnect(250)
	fmt.Println("task ok!!")
}



//[27 8 2 18 14 10 12 114 101 115 116 45 116 111 107 101 110 45 52 26 7 10 5 97 100 109 105 110
// 38 10 10 10 8 76 97 116 105 116 117 100 101 18 9 9 0 0 0 0 0 0 69 64 26 9 9 2 209
//  58 93 0 0 0 0 34 2 8 1]



//[27 8 2 18 14 10 12 114 101 115 116 45 116 111 107 101 110 45 52 26 7 10 5 97 100 109 105 110
// 48 10 9 9 72 225 122 20 174 167 67 64 18 9 9 225 122 20 174 71 97 85 192 26 9
// 9 0 0 0 0 0 0 0 64 34 9 9
// 239 216
// 58 93 0 0 0 0 42 2 8 1]



//[27 8 2 18 14 10 12 114 101 115 116 45 116 111 107 101 110 45 52 26 7 10 5 97 100 109 105 110
// 48 10 9 9 72 225 122 20 174 167 67 64 18 9 9 225 122 20 174 71 97 85 192 26 9
// 9 0 0 0 0 0 0 0 64 34 9 9
// 92 217
// 58 93 0 0 0 0 42 2 8 1]

