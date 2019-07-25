package api

import (
	`encoding/base64`
	`encoding/json`
	`github.com/pkg/errors`
	`io/ioutil`
	`net/http`
	`net/url`
)

const (
	BasicKey ="Basic"
	Bearer = "Bearer"
	AuthorizationKey = "X-Sitewhere-Jwt"
)

type AuthModule struct {
	api *API
}

func (m AuthModule)Authorization(name,pwd string)(token string,err error)  {
	URL, err := url.Parse(m.api.createURL("/sitewhere/authapi/jwt"))
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	reqs, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}

	input := []byte(name+":"+pwd)
	// 演示base64编码
	encodeString := base64.StdEncoding.EncodeToString(input)
	key := BasicKey + " " + encodeString

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
	token = resp.Header.Get(AuthorizationKey)
	return
}


