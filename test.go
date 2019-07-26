package main

import (
	`fmt`/*
	`git.sdkeji.top/share/sdlib/log`
	`sdkeji/go_mqtt/api`*/
	`time`
)
func main() {
	/*var API   *api.API
	var Logger log.Logger
	Logger, err := log.New(true, "test")
	if err != nil {
		panic(err)
	}
	API = api.NewAPI(Logger, "http://192.168.35.230:8080")*/
	s := time.Now()
	fmt.Println(s.Local().Unix())
	fmt.Println(s.Unix())




	//----------------------------------------------------------------------------------------------------------------------

	/*	//身份验证测试	---------------------TEST OK!-----------------
	token, _ := API.Auth().Authorization("admin", "password")
	fmt.Println(token)
	*/

	//----------------------------------------------------------------------------------------------------------------------

	/*//创建设备测试	---------------------TEST OK!-----------------
	req := api.CreateNewDeviceRequest{
		Comments:"rest test4",
		DeviceTypeToken:"bracelet",
		Token:"rest-token-4",
	}
	res, err := API.Device().CreateNewDevice(req)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)*/

//----------------------------------------------------------------------------------------------------------------------

	/*//添加关联工作详情（assignments）		---------------------TEST OK!-----------------
	req := api.AddAssignmentsRequest{
		AreaToken:"area-danfeng-shagnzhou",
		AssetToken:"100003",
		CustomerToken:"750-shenmu",
		DeviceToken:"rest-token-4",
	}
	res,err := API.Device().DeviceAddAssignments(req)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)*/

//----------------------------------------------------------------------------------------------------------------------


	/*//创建Customer  ）		---------------------TEST OK!-----------------
	req := api.CreateCustomerRequest{
		CustomerTypeToken:"mva-330",
		Description:"test444",
		Icon:"adjust",
		ImageURL:"http://epaper.shaoyangnews.net/epaper/sywb/html/2011/11/15/01/images/3.jpg",
		Name:"圣点2",
		Token:"customers-222",
	}
	res,err := API.Customer().CreateNewCustomer(req)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)*/

//----------------------------------------------------------------------------------------------------------------------
/*
	//获取设备列表		---------------------TEST OK!-----------------
	res,err := API.Device().GetDeviceList()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res.Results[0].DeviceType.ImageURL)*/

//----------------------------------------------------------------------------------------------------------------------

	/*//获取某个项目信息（Customer）		---------------------TEST OK!-----------------
	res,err := API.Customer().GetCustomer("line-danfeng-shangzhou")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res)*/

//----------------------------------------------------------------------------------------------------------------------

	/*//获取项目列表（Customers））		---------------------TEST OK!-----------------
	res,err := API.Customer().GetCustomerList()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(res.Results[0].CustomerType.ImageURL)*/


}