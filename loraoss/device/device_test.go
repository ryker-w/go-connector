package device

import (
	"fmt"
	"github.com/lishimeng/go-connector/loraoss"
	"github.com/lishimeng/go-connector/loraoss/connector"
	"github.com/lishimeng/go-connector/loraoss/model"
	"github.com/lishimeng/go-connector/loraoss/test"
	"testing"
)

var _conn *loraoss.Connector

func getConnector() (c *loraoss.Connector, err error) {
	if _conn == nil {
		connector.DebugEnable = true
		conn := connector.New(loraoss.ConnectorConfig{Host: test.Host, UserName: test.Username, Password: test.Password})
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
		//DeviceID("0fb7789000000a60").
		Limit("2")
	ds, err := dev.List(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ds.Total)
	items := *ds.Items
	for _, item := range items {
		fmt.Printf("%s:%s\n", item.DevEUI, item.Name)
	}
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

func TestLoraDevice_Create(t *testing.T) {
	conn, err := getConnector()
	if err != nil {
		t.Fatal(err)
	}

	dev := *New(*conn, "1")
	params := model.DeviceForm{
		DevEUI:          "0fb7789000000a60",
		Name:            "B001",
		Description:     "B001 Beacon",
		DeviceProfileID: "16cb3289-3b00-47de-8e51-e68b3c16ba72",
	}
	code, err := dev.Create(params)
	if err != nil {
		t.Fatal(err)
	}
	if code != 200 {
		t.Fatal(code)
	}
}

func TestLoraDevice_SetOTAAKeys(t *testing.T) {
	conn, err := getConnector()
	if err != nil {
		t.Fatal(err)
	}

	dev := *New(*conn, "1")

	fmt.Println("test OTAA keys")
	keys := model.DeviceKeys{DevEUI: "0fb7789000000a60", AppKey: "2deb83b7137b8c42f328a3ee879b4af4"}

	code, err := dev.SetOTAAKeys(keys)
	if err != nil {
		t.Fatal(err)
	}
	if code != 200 {
		t.Fatal(code)
	}
}

func TestLoraDevice_Delete(t *testing.T) {
	conn, err := getConnector()
	if err != nil {
		t.Fatal(err)
	}

	dev := *New(*conn, "1")
	code, err := dev.Delete("0fb7789000000a60")
	if err != nil {
		t.Fatal(err)
	}
	if code != 200 {
		t.Fatal(code)
	}
}
