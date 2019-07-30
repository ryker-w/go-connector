package connector

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/lishimeng/go-connector/loraoss"
)

var DebugEnable = false

func login(host string, path string, username string, password string) (token loraoss.Token, err error) {

	client := resty.New().SetHostURL(host)//.SetDebug(DebugEnable)
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{"username": username, "password": password}).
		Post(path)
	if err == nil {
		raw := resp.Body()
		token = loraoss.Token{}
		err = json.Unmarshal(raw, &token)
	}
	return token, err
}

func createRestClient(host string, jwt string) *resty.Client {
	client := resty.New().
		SetHostURL(host).
		SetHeader("Grpc-Metadata-Authorization", "Bearer " + jwt).
		SetHeader("Accept", "application/json").SetDebug(DebugEnable)
	return client
}