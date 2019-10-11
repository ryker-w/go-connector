package device

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/go-connector/loraoss"
	"github.com/lishimeng/go-connector/loraoss/model"
)

func init() {

}

func New(connector loraoss.Connector, appId string) *loraoss.Device {

	dev := loraDevice{connector: connector, appId: appId}
	var h loraoss.Device = &dev
	return &h
}

type loraDevice struct {
	connector loraoss.Connector
	appId     string
}

func (d loraDevice) Create(device model.DeviceForm) (code int, err error) {

	device.ReferenceAltitude = 0
	device.SkipFCntCheck = true
	device.ApplicationID = d.appId
	req := model.DeviceFormWrapper{Device: device}

	resp, err := d.connector.Request().
		SetBody(req).
		Post("/api/devices")

	if err == nil {
		code = resp.StatusCode()
	}
	return code, err
}

func (d loraDevice) Edit(device model.DeviceForm) (code int, err error) {
	devEUI := device.DevEUI
	device.ReferenceAltitude = 0
	device.SkipFCntCheck = true
	req := model.DeviceFormWrapper{Device: device}

	resp, err := d.connector.Request().
		SetPathParams(map[string]string{"dev_eui": devEUI}).
		SetBody(req).
		Put("/api/devices/{dev_eui}")

	if err == nil {
		code = resp.StatusCode()
	}
	return code, err
}

func (d loraDevice) Delete(deviceEUI string) (int, error) {

	resp, err := d.connector.Request().SetPathParams(map[string]string{"dev_eui": deviceEUI}).Delete("/api/devices/{dev_eui}")
	if err != nil {
		return 0, err
	}
	return resp.StatusCode(), err
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
