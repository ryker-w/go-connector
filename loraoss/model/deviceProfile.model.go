package model

type DeviceProfilePage struct {
	Total string              `json:"totalCount"`
	Items []DeviceProfileInfo `json:"result"`
}

type DeviceProfileInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (dpp DeviceProfilePage) ToMap() (m map[string]string) {

	m = make(map[string]string)
	if len(dpp.Items) > 0 {
		for _, item := range dpp.Items {
			m[item.Name] = item.Id
		}
	}
	return m
}
