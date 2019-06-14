package tencentim

import (
	"io/ioutil"
	"net/http"
)

var appID string

// SetAppID 设置Tencent IM appid
func SetAppID(id string) {
	appID = id
}

// Send send post to tencent im server
func Send(api API) ([]byte, error) {
	url := api.QueryString()
	err, bodyData := api.Body()
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bodyData)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
