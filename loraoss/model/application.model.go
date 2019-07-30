package model

type ApplicationInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	OrganizationID string `json:"organizationID"`
	ServiceProfileID string `json:"serviceProfileID"`
	ServiceProfileName string `json:"serviceProfileName"`
}

// application列表
type ApplicationPage struct {
	Total string `json:"total"`
	Items *[]ApplicationInfo `json:"result"`
}