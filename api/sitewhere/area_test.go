package api

import (
	"git.sdkeji.top/share/sqmslib/log"
	"testing"
)

func TestAreaModule_CreateAreaType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Area().CreateAreaType(CreateAreaTypeRequest{
		ContainerAreaTypeTokens: []string{},
		Description:             "test description",
		ImageURL:                "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:                    "测试区域类型",
		Token:                   "test-person-type",
		Metadata:                map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(types.ID)
}

func TestAreaModule_GetAreaTypeList(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Area().GetAreaTypeList()
	if err != nil {
		t.Error(err)
	}
	t.Log(types)
}

func TestAreaModule_UpdateAreaType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Area().UpdateAreaType("test-person-type", CreateAreaTypeRequest{
		ContainerAreaTypeTokens: []string{},
		Description:             "test description",
		ImageURL:                "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:                    "天安门",
		Token:                   "test-person-type",
		Metadata:                map[string]string{},
		Icon:                    "archway",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(types)
}

func TestAreaModule_DeleteAreaType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Area().DeleteAreaType("test-person-type")
	if err != nil {
		t.Error(err)
	}
	t.Log(types.ID)
}

func TestAreaModule_CreateArea(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Area().CreateArea(CreateAreaRequest{
		AreaTypeToken: "test-person-type",
		Description:   "test description",
		ImageURL:      "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:          "天安门线路维护",
		Token:         "test-person-type",
		Metadata:      map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}

func TestAreaModule_GetAreaList(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Area().GetAreaList()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestAreaModule_UpdateArea(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Area().UpdateArea("test-person-type", CreateAreaRequest{
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

func TestAreaModule_DeleteArea(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Area().DeleteArea("test-person-type")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}
