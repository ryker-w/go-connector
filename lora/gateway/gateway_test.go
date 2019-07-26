package gateway

import (
	"github.com/lishimeng/go-connector/lora"
	"github.com/lishimeng/go-connector/lora/connector"
	"github.com/lishimeng/go-connector/lora/test"
	"testing"
)

func TestListGateway(t *testing.T) {

	conn := connector.New(lora.ConnectorConfig{Host: test.Host, UserName: test.Username, Password: test.Password})

	var c = *conn
	_, err := c.Login()
	if err != nil {
		t.Fatal(err)
	}
	gw := *New(c)
	gw.List()
}
