package connector

import (
	"github.com/go-resty/resty/v2"
	"github.com/lishimeng/go-connector/loraoss"
)

type loraConnector struct {
	config loraoss.ConnectorConfig
	client *resty.Client
	jwt    loraoss.Token
}

func New(config loraoss.ConnectorConfig) *loraoss.Connector {

	c := loraConnector{
		config: config,
	}

	var con loraoss.Connector = &c
	return &con
}

func (c *loraConnector) Login() (token loraoss.Token, err error) {

	token, err = login(c.config.Host, "/api/internal/login", c.config.UserName, c.config.Password)
	if err == nil {
		c.jwt = token
		c.client = createRestClient(c.config.Host, token.Jwt)
	}
	return token, err
}

func (c loraConnector) Request() *resty.Request {
	return c.client.R()
}
