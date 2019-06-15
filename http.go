package tencentim

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var appID string

// SetAppID 设置Tencent IM appid
func SetAppID(id string) {
	appID = id
}

type timResponse struct {
	ActionStatus string `json:"ActionStatus"`
	ErrorInfo    string `json:"ErrorInfo"`
	ErrorCode    int    `json:"ErrorCode"`
}

// Send send post to tencent im server
func Send(api API) ([]byte, error) {
	queryString := api.QueryString()
	url := fmt.Sprintf("%s/%s?%s", schema, api.URI(), queryString)
	bodyData, err := api.Body()
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bodyData)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var timRes timResponse
	if err = json.Unmarshal(content, &timRes); err != nil {
		return nil, err
	}

	if timRes.ErrorCode != 0 {
		err := fmt.Errorf("Send Tim failure, ActionStatus:%s, ErrorInfo:%s, ErrorCode:%d",
			timRes.ActionStatus, timRes.ErrorInfo, timRes.ErrorCode)
		return nil, err
	}

	return content, nil
}
