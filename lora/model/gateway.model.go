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