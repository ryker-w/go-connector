package model

type GatewayInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	OrganizationID string `json:"organizationID"`
	NetworkServerID string `json:"networkServerID"`
}

// application列表
type GatewayPage struct {
	Total string `json:"total"`
	Items *[]ApplicationInfo `json:"result"`
}

type GatewayFormWrapper struct {
	Gateway GatewayForm `json:"gateway"`
}

type GatewayForm struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	OrganizationID string `json:"organizationID"`
	NetworkServerID string `json:"networkServerID"`

	GatewayProfileID string `json:"gatewayProfileID,omitempty"`
	DiscoveryEnabled bool `json:"discoveryEnabled,omitempty"`

	Location GatewayLocation `json:"location,omitempty"`

	Boards GatewayBoard `json:"boards,omitempty"`
}

type GatewayLocation struct {
	Accuracy float64 `json:"accuracy"`
	Altitude float64 `json:"altitude"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Source string `json:"source"`
}
type GatewayBoard struct {
	FpgaID string `json:"fpgaID"`
	FineTimestampKey string `json:"fineTimestampKey"`
}