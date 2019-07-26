package lora

import (
	"github.com/go-resty/resty/v2"
	"github.com/lishimeng/go-connector/lora/model"
)

type Token struct {
	Jwt string `json:"jwt"`
}

type ConnectorConfig struct {
	Host string
	UserName string
	Password string
}

type Connector interface {

	Login() (token Token, err error)

	Request() *resty.Request
}

type Gateway interface {
	Create()
	Delete()
	Edit()
	List()
}

type Application interface {
	Create()
	Delete()
	Edit()
	List()
}

type Device interface {
	Create()
	Delete(devEUI string) (int, error)
	Edit()
	List(*model.DeviceRequestBuilder) (model.DevicePage, error)

	GetOTAAKeys(devEUI string) (keys model.DeviceKeys, code int, err error)
	SetOTAAKeys(keys model.DeviceKeys) (code int, err error)
}
