package application

import (
	"github.com/lishimeng/go-connector/loraoss"
)

func New(connector loraoss.Connector) *loraoss.Application {

	app := loraApplication{connector: connector}
	var h loraoss.Application = &app
	return &h
}

type loraApplication struct {
	connector loraoss.Connector
}

func (app loraApplication) Create() {

}

func (app loraApplication) Delete() {

}

func (app loraApplication) Edit() {

}

func (app loraApplication) List() {

}