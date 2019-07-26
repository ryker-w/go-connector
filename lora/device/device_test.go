package device

import (
	"fmt"
	"github.com/lishimeng/go-connector/lora"
	"github.com/lishimeng/go-connector/lora/connector"
	"github.com/lishimeng/go-connector/lora/model"
	"github.com/lishimeng/go-connector/lora/test"
	"testing"
)

var _conn *lora.Connector
func getConnector() (c *lora.Connector, err error) {
	if _conn == nil {
		connector.DebugEnable = true
		conn := connector.New(lora.ConnectorConfig{Host: test.Host, UserName: test.Username, Password: test.Password})
		var c = *conn
		_, err = c.Login()
		if err == nil {
			_conn = conn
		}
	}
	c = _conn
	return c, err
}

func TestLoraDevice_List(t *testing.T) {

	conn, err := getConnector()
	if err != nil {
		t.Fatal(err)
	}

	dev := *New(*conn, "4")
	params := model.NewDeviceRequestBuilder().
		DeviceID("0fb7789000000a60").
		Limit("1")
	ds, err := dev.List(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ds.Total)
	fmt.Println(len(*ds.Items))
}

func TestLoraDevice_GetOTAAKeys(t *testing.T) {
	conn, err := getConnector()
	if err != nil {
		t.Fatal(err)
	}

	dev := *New(*conn, "4")

	res, code, err := dev.GetOTAAKeys("0fb7789000000a60")
	if err != nil {
		t.Fatal(err)
	}
	if code != 200 {
		t.Fatal(code)
	}
	fmt.Println(res.DevEUI)
	fmt.Println(res.AppKey)
}
