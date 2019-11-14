package device

import (
	"encoding/json"
	"github.com/lishimeng/go-connector/loraoss/model"
)

func (d loraDevice) SetOTAAKeys(keys model.DeviceKeys) (code int, err error) {

	keysParam := model.DeviceOTAAKeys{DeviceKeys: keys}
	devEUI := keys.DevEUI

	resp, err := d.Connector.Request().
		SetBody(keysParam).
		SetPathParams(map[string]string{"dev_eui": devEUI}).
		Post("/api/devices/{dev_eui}/keys")

	if err == nil {
		code = resp.StatusCode()
	}

	return code, err
}

func (d loraDevice) UpdateOTAAKeys(keys model.DeviceKeys) (code int, err error) {

	keysParam := model.DeviceOTAAKeys{DeviceKeys: keys}
	devEUI := keys.DevEUI

	resp, err := d.Connector.Request().
		SetBody(keysParam).
		SetPathParams(map[string]string{"dev_eui": devEUI}).
		Put("/api/devices/{dev_eui}/keys")

	if err == nil {
		code = resp.StatusCode()
	}

	return code, err
}

func (d loraDevice) GetOTAAKeys(devEUI string) (keys model.DeviceKeys, code int, err error) {
	resp, err := d.Connector.Request().
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
