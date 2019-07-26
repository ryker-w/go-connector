package gateway

import (
	"fmt"
	"github.com/lishimeng/go-connector/lora"
)


func New(connector lora.Connector) *lora.Gateway {

	gw := gateway{connector: connector}
	var g lora.Gateway = &gw
	return &g
}

type gateway struct {
	connector lora.Connector
}

func (gw gateway) Create() {

	//path := "/api/gateways"


}

func (gw gateway) Delete() {

}

func (gw gateway) Edit() {

}

func (gw gateway) List() {

	path := "/api/gateways"

	resp, err := gw.connector.Request().SetBody("").Get(path)
	if err != nil {
		return
	}
	if resp.StatusCode() == 401 {

		return
	}

	fmt.Println(string(resp.Body()))
}