package device

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/go-connector/lora"
	"github.com/lishimeng/go-connector/lora/model"
)

func New(connector lora.Connector, appId string) *lora.Device {

	dev := loraDevice{connector: connector, appId: appId}
	var h lora.Device = &dev
	return &h
}

type loraDevice struct {
	connector lora.Connector
	appId string
}

func (d loraDevice) Create() {

}

func (d loraDevice) Delete(deviceEUI string) (int, error) {

	resp, err := d.connector.Request().SetPathParams(map[string]string{"dev_eui": deviceEUI}).Delete("/api/devices/{dev_eui}")
	if err != nil {
		return 0, err
	}
	return resp.StatusCode(), err
}

func (d loraDevice) Edit() {

}

func (d loraDevice) List(param *model.DeviceRequestBuilder) (devices model.DevicePage, err error) {

	param.ApplicationID(d.appId)

	resp, err := d.connector.Request().SetQueryParams(param.Build()).Get("/api/devices")
	if err != nil {
		fmt.Println(err)
		return
	}

	devices = model.DevicePage{}
	body := resp.Body()
	err = json.Unmarshal(body, &devices)
	if err != nil {
		return
	}
	return devices, err
}

func (d loraDevice) SetOTAAKeys(keys model.DeviceKeys) (code int, err error) {

	keysParam := model.DeviceOTAAKeys{DeviceKeys: keys}
	devEUI := keys.DevEUI

	resp, err := d.connector.Request().
		SetBody(&keysParam).
		SetPathParams(map[string]string{"dev_eui": devEUI}).
		Put("/api/devices/{dev_eui}/keys")

	if err == nil {
		code = resp.StatusCode()
	}

	return code, err
}

func (d loraDevice) GetOTAAKeys(devEUI string) (keys model.DeviceKeys, code int, err error) {
	resp, err := d.connector.Request().
		SetPathParams(map[string]string{"dev_eui": devEUI}).
		Get("/api/devices/{dev_eui}/keys")

	if err == nil {
		code = resp.StatusCode()
		body := resp.Body()
		keysParam := model.DeviceOTAAKeys{}
		err = json.Unmarshal(body, &keysParam)
		if err == nil {
			keys = keysParam.DeviceKeys
		}
	}
	return keys, code, err
}