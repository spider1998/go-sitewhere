package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.sdkeji.top/share/sqmslib/api"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AssignmentModule struct {
	api *SiteWhereAPI
}

//绑定单个对象时其余字段必须为null(omitempty)
type CreateAssignmentsRequest struct {
	AreaToken     string            `json:"areaToken,omitempty"`     //区域token
	AssetToken    string            `json:"assetToken,omitempty"`    //领用人token
	CustomerToken string            `json:"customerToken,omitempty"` //项目token
	DeviceToken   string            `json:"deviceToken"`             //设备token
	Metadata      map[string]string `json:"metadata"`
}

type Assignment struct {
	ID           string            `json:"id"`
	ActiveDate   string            `json:"activeDate"`
	AreaID       string            `json:"areaId"`
	AssetID      string            `json:"assetId"`
	CustomerID   string            `json:"customerId"`
	DeviceID     string            `json:"deviceId"`
	DeviceTypeID string            `json:"deviceTypeId"`
	Status       string            `json:"status"`
	Metadata     map[string]string `json:"metadata"`
	ReleasedDate string            `json:"releasedDate"`
	CreatedBy    string            `json:"createdBy"`
	CreatedDate  string            `json:"createdDate"`
}

type AssignmentListResponse struct {
	NumResults int          `json:"numResults"`
	Results    []Assignment `json:"results"`
}

//创建Assignment
func (m AssignmentModule) CreateAssignment(request CreateAssignmentsRequest) (res Assignment, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/assignments/")
	m.api.Debug("create new device.", "url", URL)
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(b))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	token, err := m.api.auth.Authorization("admin", "password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(req, key)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		var result api.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//查询Assignment
func (m AssignmentModule) GetAssignments() (res AssignmentListResponse, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/assignments"))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	reqs, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	token, err := m.api.auth.Authorization("admin", "password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(reqs, key)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		//m.api.Warn("received error response.", "response", string(b))
		var result api.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//获取单个Assignment（按Token）
func (m AssignmentModule) GetAssignmentByToken(req string) (res Assignment, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/assignments/" + req))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	reqs, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	token, err := m.api.auth.Authorization("admin", "password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(reqs, key)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		//m.api.Warn("received error response.", "response", string(b))
		var result api.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//修改Assignment
func (m AssignmentModule) UpdateAssignment(tokens string, request CreateAssignmentsRequest) (res Assignment, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/assignments/" + tokens)
	m.api.Debug("update assignment.", "url", URL)
	req, err := http.NewRequest(http.MethodPut, URL, bytes.NewReader(b))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	token, err := m.api.auth.Authorization("admin", "password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(req, key)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		var result api.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//删除Assignment
func (m AssignmentModule) DeleteAssignment(req string) (res Assignment, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/assignments/" + req))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	reqs, err := http.NewRequest(http.MethodDelete, URL.String(), nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	token, err := m.api.auth.Authorization("admin", "password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(reqs, key)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		//m.api.Warn("received error response.", "response", string(b))
		var result api.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//释放活动设备分配
func (m AssignmentModule) ReleaseAssignment(tokens string) (res Assignment, err error) {

	URL := m.api.createURL("/sitewhere/api/assignments/" + tokens + "/end")
	m.api.Debug("create new device.", "url", URL)
	req, err := http.NewRequest(http.MethodPost, URL, nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	token, err := m.api.auth.Authorization("admin", "password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(req, key)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		var result api.APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
