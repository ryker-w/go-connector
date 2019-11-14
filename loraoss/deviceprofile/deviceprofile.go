package deviceprofile

import (
	"encoding/json"
	"github.com/lishimeng/go-connector/loraoss"
	"github.com/lishimeng/go-connector/loraoss/model"
)

func New(connector loraoss.Connector) *loraoss.DeviceProfile {

	gw := deviceProfile{connector}
	var g loraoss.DeviceProfile = &gw
	return &g
}

type deviceProfile struct {
	loraoss.Connector
}

func (dp deviceProfile) List() (dps model.DeviceProfilePage, err error) {

	path := "/api/device-profiles"

	resp, err := dp.Connector.Request().SetQueryParams(map[string]string{"limit": "100"}).Get(path)
	if err != nil {
		return dps, err
	}

	dps = model.DeviceProfilePage{}
	body := resp.Body()
	err = json.Unmarshal(body, &dps)
	if err != nil {
		return dps, err
	}
	return dps, err
}
