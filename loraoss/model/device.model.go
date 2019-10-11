package model

import (
	"strconv"
	"time"
)

type DeviceFormWrapper struct {
	Device DeviceForm `json:"device"`
}

type DeviceForm struct {
	DevEUI        string `json:"devEUI"`
	Name          string `json:"name"`
	ApplicationID string `json:"applicationID"`
	Description   string `json:"description,omitempty"`

	DeviceProfileID   string `json:"deviceProfileID,omitempty"`
	ReferenceAltitude int    `json:"referenceAltitude,omitempty"`
	SkipFCntCheck     bool   `json:"skipFCntCheck,omitempty"`
}

type DeviceInfo struct {
	Device DeviceInfoContent `json:"device"`
}
type DeviceInfoContent struct {
	DevEUI          string `json:"devEUI"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ApplicationID   string `json:"applicationID"`
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
	DevEUI                              string `json:"devEUI"`
	Name                                string `json:"name"`
	Description                         string `json:"description,omitempty"`
	ApplicationID                       string `json:"applicationID"`
	DeviceProfileID                     string `json:"deviceProfileID,omitempty"`
	DeviceProfileName                   string `json:"deviceProfileName,omitempty"`
	DeviceStatusBattery                 int    `json:"deviceStatusBattery,omitempty"`
	DeviceStatusMargin                  int    `json:"deviceStatusMargin,omitempty"`
	DeviceStatusExternalPowerSource     bool   `json:"deviceStatusExternalPowerSource,omitempty"`
	DeviceStatusBatteryLevelUnavailable bool   `json:"deviceStatusBatteryLevelUnavailable,omitempty"`
	DeviceStatusBatteryLevel            int    `json:"deviceStatusBatteryLevel,omitempty"`
	LastSeenAt                          string `json:"lastSeenAt,omitempty"`
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
