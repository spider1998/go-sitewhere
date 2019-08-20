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

type CommandModule struct {
	api *SiteWhereAPI
}

type CreateCommandRequest struct {
	ParameterValues map[string]string `json:"parameterValues"`
	AlternateID     string            `json:"alternateId,omitempty"`
	CommandToken    string            `json:"commandToken,omitempty"`
	EventDate       string            `json:"eventDate,omitempty"`
	EventType       string            `json:"eventType,omitempty"`
	Initiator       string            `json:"initiator,omitempty"`
	InitiatorID     string            `json:"initiatorId,omitempty"`
	Metadata        map[string]string `json:"metadata,omitempty"`
	Target          string            `json:"target,omitempty"`
	TargetID        string            `json:"targetId,omitempty"`
	UpdateState     bool              `json:"updateState,omitempty"`
}

type GetCommandsResponse struct {
	NumResults int                 `json:"numResults"`
	Results    []CommandInvocation `json:"results"`
}

type CommandInvocation struct {
	AlternateID        string            `json:"alternateId"`
	AreaID             string            `json:"areaId"`
	AssetID            string            `json:"assetId"`
	CommandToken       string            `json:"commandToken"`
	CustomerID         string            `json:"customerId"`
	DeviceAssignmentID string            `json:"deviceAssignmentId"`
	DeviceID           string            `json:"deviceId"`
	EventDate          string            `json:"eventDate"`
	EventType          string            `json:"eventType"`
	ID                 string            `json:"id"`
	Initiator          string            `json:"initiator"`
	InitiatorID        string            `json:"initiatorId"`
	Metadata           map[string]string `json:"metadata"`
	ParameterValues    map[string]string `json:"parameterValues"`
	ReceiveDate        string            `json:"receiveDate"`
	Target             string            `json:"target"`
	TargetID           string            `json:"targetId"`
}

func (m CommandModule) CreateCommand(tokens string, request CreateCommandRequest) (res CommandInvocation, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	fmt.Println(request)
	URL := m.api.createURL("/sitewhere/api/assignments/" + tokens + "/invocations")
	m.api.Debug("create new areas type.", "url", URL)
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

func (m CommandModule) GetCommandsByToken(tokens string) (res GetCommandsResponse, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/assignments/" + tokens + "/invocations"))
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

type CreateCommandResponseRequest struct {
	AlternateID        string            `json:"alternateId,omitempty"`
	EventDate          string            `json:"eventDate,omitempty"`
	EventType          string            `json:"eventType,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	OriginatingEventID string            `json:"originatingEventId,omitempty"`
	Response           string            `json:"response,omitempty"`
	ResponseEventID    string            `json:"responseEventId,omitempty"`
	UpdateState        bool              `json:"updateState,omitempty"`
}

type CommandResponse struct {
	AlternateID        string            `json:"alternateId,omitempty"`
	AreaID             string            `json:"areaId,omitempty"`
	AssetID            string            `json:"assetId,omitempty"`
	CommandToken       string            `json:"commandToken,omitempty"`
	CustomerID         string            `json:"customerId,omitempty"`
	DeviceAssignmentID string            `json:"deviceAssignmentId,omitempty"`
	DeviceID           string            `json:"deviceId,omitempty"`
	EventDate          string            `json:"eventDate,omitempty"`
	EventType          string            `json:"eventType,omitempty"`
	ID                 string            `json:"id,omitempty"`
	ReceiveDate        string            `json:"receiveDate,omitempty"`
	Metadata           map[string]string `json:"metadata,omitempty"`
	Response           string            `json:"response,omitempty"`
	ResponseEventID    string            `json:"responseEventId,omitempty"`
}

func (m CommandModule) CreateCommandResponse(tokens string, request CreateCommandResponseRequest) (res CommandResponse, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	fmt.Println(request)
	URL := m.api.createURL("/sitewhere/api/assignments/" + tokens + "/responses")
	m.api.Debug("create new areas type.", "url", URL)
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
