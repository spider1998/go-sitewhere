package api

import (
	"git.sdkeji.top/share/sdlib/log"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

type API struct {
	log.Logger
	gateway string
	auth  AuthModule
	device	DeviceModule
	customer	 CustomerModule
}

func NewAPI(logger log.Logger, gateway string) *API {
	api := &API{
		Logger:  logger,
		gateway: gateway,
	}
	api.auth = AuthModule{api}
	api.device = DeviceModule{api}
	api.customer =  CustomerModule{api}
	return api
}

func (api API) createURL(path string) string {
	return api.gateway + path
}

func (api API) do(req *http.Request,token string) (*http.Response, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	req.Header.Set("X-SiteWhere-Tenant-Auth", "ymzn")
	req.Header.Set("X-SiteWhere-Tenant-Id", "ymzn")

	return client.Do(req)
}

func (api API) Auth() AuthModule {
	return api.auth
}

func (api API) Device() DeviceModule {
	return api.device
}

func (api API)  Customer()  CustomerModule {
	return api.customer
}