package model

import (
	"strconv"
	"time"
)

type DeviceForm struct {
	Device DeviceFormContent `json:"device"`
}

type DeviceFormContent struct {
	ApplicationID string `json:"applicationID"`
	Description string `json:"description"`
	DevEUI string `json:"devEUI"`
	DeviceProfileID string `json:"deviceProfileID"`
	Name string `json:"name"`
	ReferenceAltitude int `json:"referenceAltitude"`
	SkipFCntCheck bool `json:"skipFCntCheck"`
}

type DeviceInfo struct {
	Device DeviceInfoContent `json:"device"`
}
type DeviceInfoContent struct {
	DevEUI string `json:"devEUI"`
	Name string `json:"name"`
	Description string `json:"description"`
	ApplicationID string `json:"applicationID"`
	DeviceProfileID string `json:"deviceProfileID"`
}

type DeviceOTAAKeys struct {
	DeviceKeys DeviceKeys `json:"deviceKeys"`
}

type DeviceKeys struct {
	DevEUI string `json:"devEUI"`
	AppKey string `json:"nwkKey"`
}

type DeviceItemInfo struct {
	DevEUI string `json:"devEUI"`
	Name string `json:"name"`
	Description string `json:"description"`
	ApplicationID string `json:"applicationID"`
	DeviceProfileID string `json:"deviceProfileID"`
	DeviceProfileName string `json:"deviceProfileName"`
	DeviceStatusBattery string `json:"deviceStatusBattery"`
	DeviceStatusMargin string `json:"deviceStatusMargin"`
	DeviceStatusExternalPowerSource bool `json:"deviceStatusExternalPowerSource"`
	DeviceStatusBatteryLevelUnavailable bool `json:"deviceStatusBatteryLevelUnavailable"`
	DeviceStatusBatteryLevel string `json:"deviceStatusBatteryLevel"`
	LastSeenAt string `json:"lastSeenAt"`
}

// application列表
type DevicePage struct {
	Total string            `json:"total"`
	Items *[]DeviceItemInfo `json:"result"`
}

type DeviceRequestBuilder struct {

	params map[string]string
}

func NewDeviceRequestBuilder() *DeviceRequestBuilder {
	params := map[string]string{}
	params["random"] = strconv.FormatInt(time.Now().Unix(), 10)
	return &DeviceRequestBuilder{params: params}
}

func (b *DeviceRequestBuilder) ApplicationID(appId string) *DeviceRequestBuilder {
	b.params["applicationID"] = appId
	return b
}

func (b *DeviceRequestBuilder) DeviceID(devId string) *DeviceRequestBuilder {
	b.params["search"] = devId
	return b
}

func (b *DeviceRequestBuilder) Limit(limit string) *DeviceRequestBuilder {
	b.params["limit"] = limit
	return b
}

func (b *DeviceRequestBuilder) Offset(offset string) *DeviceRequestBuilder {
	b.params["offset"] = offset
	return b
}

func (b *DeviceRequestBuilder) Build() map[string]string {
	return b.params
}