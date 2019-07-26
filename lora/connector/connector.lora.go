package connector

import (
	"github.com/go-resty/resty/v2"
	"github.com/lishimeng/go-connector/lora"
)

type loraConnector struct {
	config lora.ConnectorConfig
	client *resty.Client
	jwt lora.Token
}

func New(config lora.ConnectorConfig) *lora.Connector{

	c := loraConnector{
		config: config,
	}

	var con lora.Connector = &c
	return &con
}

func (c *loraConnector) Login() (token lora.Token, err error) {

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