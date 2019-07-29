package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DeviceModule struct {
	api *API
}

type CreateNewDeviceRequest struct {
	Comments        string            `json:"comments"`        //备注名
	DeviceTypeToken string            `json:"deviceTypeToken"` //类型token
	Metadata        map[string]string `json:"metadata"`
	Token           string            `json:"token"` //Token
}

type Area struct {
	ID         string `json:"id"`
	AreaTypeID string `json:"areaTypeId"` //区域类型id
	Bounds     []struct {
		Elevation int    `json:"elevation"` //海拔
		Latitude  string `json:"latitude"`  //纬度
		Longitude string `json:"longitude"` ///经度
	} `json:"bounds"` //位置信息
	CreatedDate string            `json:"createdDate"` //创建时间
	Description string            `json:"description"` //描述
	ImageURL    string            `json:"imageUrl"`    //图像链接
	Metadata    map[string]string `json:"metadata"`
	Name        string            `json:"name"` //名称
	Token       string            `json:"token"`
	UpdatedDate string            `json:"updatedDate"`
}

type Asset struct {
	ID          string            `json:"id"`
	AssetTypeID string            `json:"asset_type_id"`
	CreateDate  string            `json:"createDate"`
	Metadata    map[string]string `json:"metadata"`
	Name        string            `json:"name"` //名称
	Token       string            `json:"token"`
}

type Device struct {
	ID                    string              `json:"id"`
	Comments              string              `json:"comments"`   //备注名
	CreateDate            string              `json:"createDate"` //创建时间
	DeviceElementMappings []map[string]string `json:"deviceElementMappings"`
	DeviceTypeID          string              `json:"deviceTypeId"`
	DeviceAssignmentID    string              `json:"deviceAssignmentId"`
	Metadata              map[string]string   `json:"metadata"`
	Token                 string              `json:"token"`
}

type DeviceType struct {
	ContainerPolice string            `json:"containerPolice"`
	CreateDate      string            `json:"createDate"`  //创建时间
	Description     string            `json:"description"` //描述
	ID              string            `json:"id"`          //ID
	ImageURL        string            `json:"imageUrl"`    //图像链接
	Metadata        map[string]string `json:"metadata"`
	Name            string            `json:"name"`        //类型名称
	Token           string            `json:"token"`       //类型token
	UpdatedDate     string            `json:"updatedDate"` //更新时间

}

type Assignments struct {
	ID         string            `json:"id"` //ID
	ActiveDate string            `json:"active_date"`
	AreaID     string            `json:"areaId"`
	AssetID    string            `json:"assetId"`
	DeviceID   string            `json:"deviceId"`
	Metadata   map[string]string `json:"metadata"`
	Status     string            `json:"status"`
	Token      string            `json:"token"` //设备token
}

type DeviceListResponse struct {
	NumResults int `json:"numResults"`
	Results    []struct {
		Device
	} `json:"results"`
}

func (m DeviceModule) GetDeviceList() (deviceList DeviceListResponse, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/devices"))
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
		var result APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}
	err = json.Unmarshal(b, &deviceList)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (m DeviceModule) CreateNewDevice(request CreateNewDeviceRequest) (device Device, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/devices")
	//m.api.Debug("create new device.", "url", URL)
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
		var result APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}

	err = json.Unmarshal(b, &device)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

type AddAssignmentsRequest struct {
	AreaToken     string            `json:"areaToken"`     //位置token
	AssetToken    string            `json:"assetToken"`    //领用人token
	CustomerToken string            `json:"customerToken"` //项目token
	DeviceToken   string            `json:"deviceToken"`   //设备token
	Metadata      map[string]string `json:"metadata"`
}

func (m DeviceModule) DeviceAddAssignments(request AddAssignmentsRequest) (assignments Assignments, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/assignments")
	//m.api.Debug("create new device.", "url", URL)
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
		var result APIError
		err = json.Unmarshal(b, &result)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
		err = result.WithStatus(resp.StatusCode)
		return
	}

	err = json.Unmarshal(b, &assignments)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
