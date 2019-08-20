package api

import (
	"git.sdkeji.top/share/sqmslib/log"
	"testing"
)

func TestCommandModule_GetCommandsByToken(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Command().GetCommandsByToken("testre")
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestCommandModule_CreateCommand(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Command().CreateCommand("blb7jpm21bii4tl4srhg", CreateCommandRequest{
		CommandToken: "181ea0e1-49d0-404d-808d-ac3aa425fac9",
		Initiator:    "REST",
		ParameterValues: map[string]string{
			"time": "lll",
		},
		EventType:   "Location",
		InitiatorID: "admin",
		Target:      "Assignment",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestCommandModule_CreateCommandResponse(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Command().CreateCommandResponse("spider544", CreateCommandResponseRequest{
		AlternateID: "spider",
		//EventDate:          "2019-08-20T02:07:21.292Z",
		OriginatingEventID: "181ea0e1-49d0-404d-808d-ac3aa425fac3",
		Response:           "test response",
		ResponseEventID:    "181ea0e1-49d0-404d-808d-ac3aa425fac9",
		EventType:          "Location",
		UpdateState:        false,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
