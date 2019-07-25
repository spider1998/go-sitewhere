package api

import (
	`bytes`
	`encoding/json`
	`fmt`
	`github.com/pkg/errors`
	`io/ioutil`
	`net/http`
	`net/url`
)

type CustomerModule struct {
	api *API
}

type Customer struct {
	CreateDate string `json:"createDate"`	//创建时间
	CustomerTypeID string `json:"customerTypeId"`
	Description string `json:"description"`			//描述
	ID string `json:"id"`							//ID
	ImageURL string `json:"imageUrl"`				//图像链接
	Metadata map[string]string `json:"metadata"`
	Name string `json:"name"`						//类型名称
	Token string `json:"token"`						//类型token
	UpdatedDate string `json:"updatedDate"`			//更新时间
}

type CustomerType struct {
	ID string `json:"id"`							//ID
	ContainedCustomerTypeIDs []string `json:"containedCustomerTypeIds"`
	CreateDate string `json:"createDate"`	//创建时间
	Description string `json:"description"`			//描述
	ImageURL string `json:"imageUrl"`				//图像链接
	Metadata map[string]string `json:"metadata"`
	Name string `json:"name"`						//类型名称
	Token string `json:"token"`						//类型token
	UpdatedDate string `json:"updatedDate"`			//更新时间
	Icon string `json:"icon"`
}

type CustomerList struct {
	NumResults int `json:"numResults"`
	Results []Customer `json:"results"`
}

type CreateCustomerRequest struct {
	BackgroundColor string `json:"backgroundColor"`
	BorderColor string `json:"borderColor"`
	CreateDate string `json:"createDate"`			//创建时间
	CustomerTypeToken string `json:"customerTypeToken"`
	Description string `json:"description"`			//描述
	ForegroundColor string `json:"foregroundColor"`
	Icon string `json:"icon"`
	ImageURL string `json:"imageUrl"`				//图像链接
	Metadata map[string]string `json:"metadata"`
	Name string `json:"name"`						//类型名称
	ParentCustomerToken string `json:"parentCustomerToken"`
	Token string `json:"token"`						//类型token
}

func (m CustomerModule)GetCustomer(customerToken string)(customer Customer,err error)  {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/customers/"+customerToken))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	reqs, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	token,err := m.api.auth.Authorization("admin","password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(reqs,key)
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
	err = json.Unmarshal(b, &customer)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (m CustomerModule)GetCustomerList()(customers CustomerList,err error)  {
	URL, err := url.Parse(m.api.createURL("/sitewhere/api/customers"))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	reqs, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	token,err := m.api.auth.Authorization("admin","password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(reqs,key)
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
	err = json.Unmarshal(b, &customers)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}

func (m CustomerModule)CreateNewCustomer(request CreateCustomerRequest)(customer Customer,err error)  {

	b, err := json.Marshal(request)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	URL := m.api.createURL("/sitewhere/api/customers")
	//m.api.Debug("create new device.", "url", URL)
	req, err := http.NewRequest(http.MethodPost, URL, bytes.NewReader(b))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	token,err := m.api.auth.Authorization("admin","password")
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	key := Bearer + " " + token

	resp, err := m.api.do(req,key)
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

	err = json.Unmarshal(b, &customer)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}




