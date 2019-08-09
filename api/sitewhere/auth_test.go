package api

import (
	"git.sdkeji.top/share/sqmslib/log"
	"testing"
)

func TestAuthModule_Authorization(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Auth().Authorization("admin", "password")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
