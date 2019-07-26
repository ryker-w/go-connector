package influx

import (
	"github.com/influxdata/influxdb1-client/v2"
	"time"
)

type Connector struct {
	proxy client.Client
}

func New(host string) (connector *Connector, err error) {

	var c client.Client
	c, err = client.NewHTTPClient(client.HTTPConfig{Addr: host})
	if err != nil {
		return connector, err
	}
	connector = &Connector{
		proxy: c,
	}
	return connector, err
}

func (connector Connector) Save(database string, measurement string, tags map[string]string, fields map[string]interface{}) (err error) {

	var bp client.BatchPoints
	var point *client.Point
	bp, err = client.NewBatchPoints(client.BatchPointsConfig{Database: database})
	if err != nil {
		return err
	}

	point, err = client.NewPoint(measurement, tags, fields, time.Now())
	if err != nil {
		return err
	}

	bp.AddPoint(point)
	err = connector.proxy.Write(bp)
	return err
}

func (connector Connector) Close() {
	if connector.proxy != nil {
		_ = connector.proxy.Close()
	}
}