package mqtt

import (
	Mqtt"github.com/eclipse/paho.mqtt.golang"
)

type MessageCallback func(session Session, topic string, msg []byte)
type ConnectedCallback func(session Session)
type ConnectionLostCallback func(session Session, reason error)

type Session struct {
	client *Mqtt.Client
	opts *Mqtt.ClientOptions
	OnMessage MessageCallback
	OnConnected ConnectedCallback
	OnLostConnect ConnectionLostCallback
	State bool
	ErrorMessage string
}

type Payload struct {
	Topic string
	Qos uint8
	Payload []byte
}

func CreateSession(broker string, clientId string) *Session {
	session := Session{}
	opts := Mqtt.NewClientOptions()

	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetCleanSession(true)
	opts.SetDefaultPublishHandler(session.DefaultMessageHandler)
	opts.SetConnectionLostHandler(session.onConnLost)
	opts.SetOnConnectHandler(session.onConned)

	c := Mqtt.NewClient(opts)
	session.client = &c
	session.opts = opts
	return &session
}

func (session *Session) AddBroker(broker string) *Session {
	session.opts.AddBroker(broker)
	return session
}

func (session *Session) SetAuth(userName string, password string) *Session {
	session.opts.Username = userName
	session.opts.Password = password
	return session
}

func (session *Session) CleanSession(clean bool) *Session {
	session.opts.CleanSession = clean
	return session
}

func (session *Session) onConnLost(c Mqtt.Client, reason error) {
	if session.OnLostConnect != nil {
		session.OnLostConnect(*session, reason)
	}
}

func (session *Session) onConned(c Mqtt.Client) {
	if session.OnConnected != nil {
		session.OnConnected(*session)
	}
}

func (session *Session) Publish(topic string, qos uint8, payload string) (err error) {

	client := *session.client
	if token := client.Publish(topic, qos, false, payload); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	return err
}

func (session *Session) Subscribe(topic string, qos uint8, callback MessageCallback) bool {
	client := *session.client

	var cb Mqtt.MessageHandler = nil
	if callback != nil {
		cb = func(client Mqtt.Client, message Mqtt.Message) {
			callback(*session, message.Topic(), message.Payload())
		}
	}

	if token := client.Subscribe(topic, qos, cb); token.Wait() && token.Error() != nil {
		session.ErrorMessage = token.Error().Error()
		session.State = false
		return false
	} else {
		session.State = true
		return true
	}
}

func (session *Session) Unsubscribe(topic string) bool {
	client := *session.client
	if token := client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return false
	} else {
		return true
	}
}

func (session *Session) Connect() () {

	client := Mqtt.NewClient(session.opts)
	session.client = &client

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		session.ErrorMessage = "connect failed"
	} else {
		session.State = true
		session.ErrorMessage = ""
	}
}

func (session *Session) Close() () {
	client := *session.client
	client.Disconnect(250)
}

func (session *Session) DefaultMessageHandler(client Mqtt.Client, msg Mqtt.Message) {
	if session.OnMessage != nil {
		session.OnMessage(*session, msg.Topic(), msg.Payload())
	}
}
