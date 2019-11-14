package deviceprofile

import (
	"fmt"
	"github.com/lishimeng/go-connector/loraoss"
	"github.com/lishimeng/go-connector/loraoss/connector"
	"github.com/lishimeng/go-connector/loraoss/test"
	"testing"
)

var _conn *loraoss.Connector

func getConnector() (c *loraoss.Connector, err error) {
	if _conn == nil {
		connector.DebugEnable = false
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

func TestDeviceProfile_List(t *testing.T) {
	conn, err := getConnector()
	if err != nil {
		t.Fatal(err)
	}

	dp := *New(*conn)
	dps, err := dp.List()
	if err != nil {
		t.Fatal(err)
		return
	}
	if len(dps.Items) > 0 {
		for index, item := range dps.Items {
			fmt.Printf("..............\n%d: %s[%s]\n", index, item.Id, item.Name)
		}
	}
}
