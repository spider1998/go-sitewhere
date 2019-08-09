package api

import (
	"git.sdkeji.top/share/sqmslib/log"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

type SiteWhereAPI struct {
	log.Logger
	gateway    string
	auth       AuthModule
	device     DeviceModule
	customer   CustomerModule
	asset      AssetModule
	area       AreaModule
	assignment AssignmentModule
}

func NewSiteWhereAPI(logger log.Logger, gateway string) *SiteWhereAPI {
	api := &SiteWhereAPI{
		Logger:  logger,
		gateway: gateway,
	}
	api.auth = AuthModule{api}
	api.device = DeviceModule{api}
	api.customer = CustomerModule{api}
	api.asset = AssetModule{api}
	api.area = AreaModule{api}
	api.assignment = AssignmentModule{api}
	return api
}

func (api SiteWhereAPI) createURL(path string) string {
	return api.gateway + path
}

func (api SiteWhereAPI) do(req *http.Request, token string) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	req.Header.Set("X-SiteWhere-Tenant-Auth", "ymzn")
	req.Header.Set("X-SiteWhere-Tenant-Id", "ymzn")

	return client.Do(req)
}

func (api SiteWhereAPI) Auth() AuthModule {
	return api.auth
}

func (api SiteWhereAPI) Device() DeviceModule {
	return api.device
}

func (api SiteWhereAPI) Customer() CustomerModule {
	return api.customer
}

func (api SiteWhereAPI) Asset() AssetModule {
	return api.asset
}

func (api SiteWhereAPI) Area() AreaModule {
	return api.area
}

func (api SiteWhereAPI) Assignment() AssignmentModule {
	return api.assignment
}
