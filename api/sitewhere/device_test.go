package api

import (
	"git.sdkeji.top/share/sqmslib/log"
	"testing"
)

func TestDeviceModule_CreateDeviceType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Device().CreateDeviceType(CreateDeviceTypeRequest{
		ContainerPolicy: "Standalone",
		Description:     "test description",
		ImageURL:        "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:            "测试工器具类别",
		Token:           "test-person-type",
		Metadata:        map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(types.ID)
}

func TestDeviceModule_GetDeviceTypeList(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Device().GetDeviceTypeList()
	if err != nil {
		t.Error(err)
	}
	t.Log(types)
}

//Q:类别列表页名称更新，但详情页类别名称不变（手动操作结果相同）
func TestDeviceModule_UpdateDeviceType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Device().UpdateDeviceType("test-person-type", CreateDeviceTypeRequest{
		ContainerPolicy: "Composite",
		Description:     "test description",
		ImageURL:        "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:            "火箭类别",
		Token:           "test-person-type",
		Metadata:        map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(types)
}

func TestDeviceModule_DeleteDeviceType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Device().DeleteDeviceType("test-person-type")
	if err != nil {
		t.Error(err)
	}
	t.Log(types.ID)
}

func TestDeviceModule_CreateDevice(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Device().CreateDevice(CreateNewDeviceRequest{
		Comments:        "测试设备",
		DeviceTypeToken: "test-person-type",
		Token:           "test-huojian-token",
		Metadata:        map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}

func TestDeviceModule_GetDeviceList(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Device().GetDeviceList()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestDeviceModule_UpdateDevice(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Device().UpdateDevice("test-huojian-token", CreateNewDeviceRequest{
		Comments:        "嫦娥三号",
		DeviceTypeToken: "test-person-type",
		Token:           "test-huojian-token",
		Metadata:        map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}

func TestDeviceModule_DeleteDevice(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Device().DeleteDevice("test-huojian-token")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}
