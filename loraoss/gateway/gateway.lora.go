package gateway

import (
	"fmt"
	"github.com/lishimeng/go-connector/loraoss"
	"github.com/lishimeng/go-connector/loraoss/model"
)

func New(connector loraoss.Connector) *loraoss.Gateway {

	gw := gateway{connector}
	var g loraoss.Gateway = &gw
	return &g
}

type gateway struct {
	loraoss.Connector
}

func (gw gateway) Create(param model.GatewayForm) (code int, err error) {

	req := model.GatewayFormWrapper{Gateway: param}

	resp, err := gw.Connector.Request().
		SetBody(req).
		Post("/api/gateways")

	if err == nil {
		code = resp.StatusCode()
	}
	return code, err
}

func (gw gateway) Delete(id string) (int, error) {
	resp, err := gw.Connector.Request().SetPathParams(map[string]string{"id": id}).Delete("/api/gateways/{id}")
	if err != nil {
		return 0, err
	}
	return resp.StatusCode(), err
}

func (gw gateway) Edit() {

}

func (gw gateway) List() {

	path := "/api/gateways"

	resp, err := gw.Connector.Request().SetBody("").Get(path)
	if err != nil {
		return
	}
	if resp.StatusCode() == 401 {

		return
	}

	fmt.Println(string(resp.Body()))
}
