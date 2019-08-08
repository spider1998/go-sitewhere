package api

import (
	"git.sdkeji.top/share/sqmslib/log"
	"testing"
)

var (
	Api     *SiteWhereAPI
	Loggers log.Logger
)

func TestAssetModule_CreateAssetType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Asset().CreateAssetType(CreateAssetTypeRequest{
		AssetCategory: "Person",
		Description:   "test description",
		ImageURL:      "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:          "测试人",
		Token:         "test-person-type",
		Metadata:      map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(types.ID)
}

func TestAssetModule_GetAssetTypeList(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Asset().GetAssetList()
	if err != nil {
		t.Error(err)
	}
	t.Log(types)
}

func TestAssetModule_UpdateAssetType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Asset().UpdateAssetType("test-person-type", CreateAssetTypeRequest{
		AssetCategory: "Person",
		Description:   "test description",
		ImageURL:      "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:          "包工头",
		Token:         "test-person-type",
		Metadata:      map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(types)
}

func TestAssetModule_DeleteAssetType(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	types, err := Api.Asset().DeleteAssetType("test-person-type")
	if err != nil {
		t.Error(err)
	}
	t.Log(types.ID)
}

func TestAssetModule_CreateAsset(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Asset().CreateAsset(CreateAssetRequest{
		AssetTypeToken: "person",
		Description:    "test description",
		ImageURL:       "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:           "刘凡",
		Token:          "test-person-type",
		Metadata:       map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}

func TestAssetModule_GetAssetList(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Asset().GetAssetList()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestAssetModule_UpdateAsset(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Asset().UpdateAsset("test-person-type", CreateAssetRequest{
		AssetTypeToken: "person",
		Description:    "test description",
		ImageURL:       "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=1856187420,2118550874&fm=26&gp=0.jpg",
		Name:           "刘大凡",
		Token:          "test-person-type",
		Metadata:       map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}

func TestAssetModule_DeleteAsset(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Asset().DeleteAsset("test-person-type")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.ID)
}
