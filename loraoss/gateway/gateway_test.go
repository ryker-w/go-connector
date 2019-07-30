package gateway

import (
	"github.com/lishimeng/go-connector/loraoss"
	"github.com/lishimeng/go-connector/loraoss/connector"
	"github.com/lishimeng/go-connector/loraoss/test"
	"testing"
)

func TestListGateway(t *testing.T) {

	conn := connector.New(loraoss.ConnectorConfig{Host: test.Host, UserName: test.Username, Password: test.Password})

	var c = *conn
	_, err := c.Login()
	if err != nil {
		t.Fatal(err)
	}
	gw := *New(c)
	gw.List()
}
