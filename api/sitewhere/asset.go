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

type AssetModule struct {
	api *SiteWhereAPI
}

type Asset struct {
	ID          string            `json:"id"`
	AssetTypeID string            `json:"asset_type_id"`
	CreatedDate string            `json:"createdDate"`
	Metadata    map[string]string `json:"metadata"`
	Name        string            `json:"name"` //名称
	Token       string            `json:"token"`
}

type AssetType struct {
	AssetCategory string            `json:"assetCategory"`
	CreatedDate   string            `json:"createdDate"`
	Description   string            `json:"description"`
	ID            string            `json:"id"`
	ImageURL      string            `json:"imageUrl"`
	Metadata      map[string]string `json:"metadata"`
	Name          string            `json:"name"`
	Token         string            `json:"token"`
}

type AssetTypeListResponse struct {
	NumResults int `json:"numResults"`
	Results    []struct {
		AssetType
	} `json:"results"`
}

type AssetListResponse struct {
	NumResults int `json:"numResults"`
	Results    []struct {
		AssetType
	} `json:"results"`
}

type CreateAssetRequest struct {
	BackgroundColor string            `json:"backgroundColor"` //--
	BorderColor     string            `json:"borderColor"`     //--
	ForegroundColor string            `json:"foregroundColor"` //--
	Icon            string            `json:"icon"`            //--
	AssetTypeToken  string            `json:"assetTypeToken"`
	Description     string            `json:"description"`
	ImageURL        string            `json:"imageUrl"`
	Metadata        map[string]string `json:"metadata"`
	Name            string            `json:"name"`
	Token           string            `json:"token"`
}

type CreateAssetTypeRequest struct {
	BackgroundColor string            `json:"backgroundColor"` //--
	BorderColor     string            `json:"borderColor"`     //--
	ForegroundColor string            `json:"foregroundColor"` //--
	Icon            string            `json:"icon"`            //--
	AssetCategory   string            `json:"assetCategory"`   //"Person"\"Device"\"Hardware"
	Description     string            `json:"description"`
	ImageURL        string            `json:"imageUrl"`
	Metadata        map[string]string `json:"metadata"`
	Name            string            `json:"name"`
	Token           string            `json:"token"`
}

//获取Asset类型列表
func (m AssetModule) GetAssetTypeList() (assetTypes AssetTypeListResponse, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/assettypes"))
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
	err = json.Unmarshal(b, &assetTypes)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//创建Asset类型
func (m AssetModule) CreateAssetType(request CreateAssetTypeRequest) (types AssetType, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/assettypes")
	m.api.Debug("create new assets type.", "url", URL)
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

//删除Asset类型
func (m AssetModule) DeleteAssetType(typeToken string) (types AssetType, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/assettypes/" + typeToken))
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

//更新Asset类型
func (m AssetModule) UpdateAssetType(typeToken string, request CreateAssetTypeRequest) (assetType AssetType, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/assettypes/" + typeToken)
	m.api.Debug("update asset types.", "url", URL)
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

	err = json.Unmarshal(b, &assetType)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//创建Asset
func (m AssetModule) CreateAsset(request CreateAssetRequest) (asset Asset, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/assets")
	m.api.Debug("create new asset.", "url", URL)
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

	err = json.Unmarshal(b, &asset)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//删除Asset
func (m AssetModule) DeleteAsset(assetToken string) (asset Asset, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/assets/" + assetToken))
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
	err = json.Unmarshal(b, &asset)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//获取Asset列表
func (m AssetModule) GetAssetList() (assets AssetListResponse, err error) {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/assets"))
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
	err = json.Unmarshal(b, &assets)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

//更新Asset
func (m AssetModule) UpdateAsset(typeToken string, request CreateAssetRequest) (asset Asset, err error) {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/assets/" + typeToken)
	m.api.Debug("update asset.", "url", URL)
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

	err = json.Unmarshal(b, &asset)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
