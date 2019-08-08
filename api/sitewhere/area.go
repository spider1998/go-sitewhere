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

type AreaModule struct {
	api *SiteWhereAPI
}

type AreaListResponse struct {
	NumResults int    `json:"numResults"`
	Results    []Area `json:"results"`
}

type Area struct {
	ID         string `json:"id"`
	AreaTypeID string `json:"areaTypeId"` //区域类型id
	Bounds     []struct {
		Elevation float64 `json:"elevation"` //海拔
		Latitude  float64 `json:"latitude"`  //纬度
		Longitude float64 `json:"longitude"` ///经度
	} `json:"bounds"` //位置信息
	CreatedDate string            `json:"createdDate"` //创建时间
	Description string            `json:"description"` //描述
	ImageURL    string            `json:"imageUrl"`    //图像链接
	Metadata    map[string]string `json:"metadata"`
	Name        string            `json:"name"` //名称
	Token       string            `json:"token"`
	UpdatedDate string            `json:"updatedDate"`
}

type AreaType struct {
	ContainerAreaTypeIDS []string          `json:"containerAreaTypeIds"`
	CreatedDate          string            `json:"createdDate"` //创建时间
	Description          string            `json:"description"` //描述
	Icon                 string            `json:"icon"`
	ID                   string            `json:"id"`
	ImageURL             string            `json:"imageUrl"` //图像链接
	Metadata             map[string]string `json:"metadata"`
	Name                 string            `json:"name"` //名称
	Token                string            `json:"token"`
	UpdatedDate          string            `json:"updatedDate"`
}

type AreaTypeListResponse struct {
	NumResults int        `json:"numResults"`
	Results    []AreaType `json:"results"`
}

type CreateAreaTypeRequest struct {
	BackgroundColor         string            `json:"backgroundColor"`
	BorderColor             string            `json:"borderColor"`
	ForegroundColor         string            `json:"foregroundColor"`
	Icon                    string            `json:"icon"`
	ContainerAreaTypeTokens []string          `json:"containerAreaTypeTokens"`
	Description             string            `json:"description"`
	ImageURL                string            `json:"imageUrl"`
	Metadata                map[string]string `json:"metadata"`
	Name                    string            `json:"name"`
	Token                   string            `json:"token"`
}

type CreateAreaRequest struct {
	BackgroundColor string            `json:"backgroundColor"`
	BorderColor     string            `json:"borderColor"`
	CreatedDate     string            `json:"createdDate"` //创建时间
	AreaTypeToken   string            `json:"areaTypeToken"`
	Description     string            `json:"description"` //描述
	ForegroundColor string            `json:"foregroundColor"`
	Icon            string            `json:"icon"`
	ImageURL        string            `json:"imageUrl"` //图像链接
	Metadata        map[string]string `json:"metadata"`
	Name            string            `json:"name"` //类型名称
	ParentAreaToken string            `json:"parentAreaToken"`
	Token           string            `json:"token"` //类型token
}

//获取Area类型列表
func (m AreaModule) GetAreaTypeList() (areaTypes AreaTypeListResponse, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/areatypes"))
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
	err = json.Unmarshal(b, &areaTypes)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//创建Area类型
func (m AreaModule) CreateAreaType(request CreateAreaTypeRequest) (types AreaType, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/areatypes")
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

	err = json.Unmarshal(b, &types)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//删除Area类型
func (m AreaModule) DeleteAreaType(typeToken string) (types AreaType, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/areatypes/" + typeToken))
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
	err = json.Unmarshal(b, &types)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//更新Area类型
func (m AreaModule) UpdateAreaType(typeToken string, request CreateAreaTypeRequest) (types AreaType, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/areatypes/" + typeToken)
	m.api.Debug("update area types.", "url", URL)
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

	err = json.Unmarshal(b, &types)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//创建Area
func (m AreaModule) CreateArea(request CreateAreaRequest) (res Area, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/areas")
	m.api.Debug("create new area.", "url", URL)
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

//删除Area
func (m AreaModule) DeleteArea(typeToken string) (res Area, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/areas/" + typeToken))
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

//获取Area列表
func (m AreaModule) GetAreaList() (res AreaListResponse, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/areas"))
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

//更新Area
func (m AreaModule) UpdateArea(typeToken string, request CreateAreaRequest) (res Area, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/areas/" + typeToken)
	m.api.Debug("update area.", "url", URL)
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
