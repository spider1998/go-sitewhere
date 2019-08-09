package api

import (
	"git.sdkeji.top/share/sqmslib/log"
	"testing"
)

func TestAssignmentModule_CreateAssignment(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Assignment().CreateAssignment(CreateAssignmentsRequest{
		//AreaToken:   "area-danfeng-shagnzhou",
		DeviceToken: "spider11",
		AssetToken:  "liufan",
		//CustomerToken: "costomer-111",
		Metadata: map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.AreaID)
}

func TestAssignmentModule_ReleaseAssignment(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Assignment().ReleaseAssignment("test-ass")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.DeviceID)
}

func TestAssignmentModule_GetAssignments(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Assignment().GetAssignments()
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestAssignmentModule_GetAssignmentByToken(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Assignment().GetAssignmentByToken("test-ass")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.DeviceID)
}

func TestAssignmentModule_DeleteAssignment(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Assignment().DeleteAssignment("dsdsa")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.DeviceID)
}

func TestAssignmentModule_UpdateAssignment(t *testing.T) {
	Loggers, _ = log.New(true, "test")
	Api = NewSiteWhereAPI(Loggers, "http://192.168.35.230:8080")
	res, err := Api.Assignment().UpdateAssignment("55380839-56ee-4ad6-8789-365d30c9e1a0", CreateAssignmentsRequest{
		AreaToken:     "area-danfeng-shagnzhou",
		DeviceToken:   "test-tanke-token",
		AssetToken:    "tensioner01",
		CustomerToken: "costomer-111",
		Metadata:      map[string]string{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res.DeviceID)
}
