package application

import "github.com/lishimeng/go-connector/lora"

func New(connector lora.Connector) *lora.Application {

	app := loraApplication{connector: connector}
	var h lora.Application = &app
	return &h
}

type loraApplication struct {
	connector lora.Connector
}

func (app loraApplication) Create() {

}

func (app loraApplication) Delete() {

}

func (app loraApplication) Edit() {

}

func (app loraApplication) List() {

}