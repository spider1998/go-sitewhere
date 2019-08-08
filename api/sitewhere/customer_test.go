package api

import (
	"git.sdkeji.top/share/sqmslib/log"
	"testing"
)

func TestCustomerModule_CreateCustomerType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Customer().CreateCustomerType(CreateCustomerTypeRequest{
		ContainerCustomerTypeTokens: []string{},
		Description:                 "test description",
		ImageURL:                    "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:                        "测试项目",
		Token:                       "test-person-type",
		Metadata:                    map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(types.ID)
}

func TestCustomerModule_GetCustomerTypeList(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Customer().GetCustomerList()
	if err != nil {
		t.Error(err)
	}
	t.Log(types)
}

func TestCustomerModule_UpdateCustomerType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Customer().UpdateCustomerType("test-person-type", CreateCustomerTypeRequest{
		ContainerCustomerTypeTokens: []string{},
		Description:                 "test description",
		ImageURL:                    "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:                        "天安门扩建项目",
		Token:                       "test-person-type",
		Metadata:                    map[string]string{},
		Icon:                        "archway",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(types)
}

func TestCustomerModule_DeleteCustomerType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Customer().DeleteCustomerType("test-person-type")
	if err != nil {
		t.Error(err)
	}
	t.Log(types.ID)
}

func TestCustomerModule_CreateCustomer(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Customer().CreateCustomer(CreateCustomerRequest{
		CustomerTypeToken: "test-person-type",
		Description:       "test description",
		ImageURL:          "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:              "天安门线路维护",
		Token:             "test-person-type",
		Metadata:          map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}

func TestCustomerModule_GetCustomerList(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Customer().GetCustomerList()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestCustomerModule_UpdateCustomer(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Customer().UpdateCustomer("test-person-type", CreateCustomerRequest{
		Description: "test description",
		ImageURL:    "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:        "刘大凡",
		Token:       "test-person-type",
		Metadata:    map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}

func TestCustomerModule_DeleteCustomer(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Customer().DeleteCustomer("test-person-type")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}
