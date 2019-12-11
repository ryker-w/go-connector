package mqtt

import (
	delegate "github.com/eclipse/paho.mqtt.golang"
)

type MessageCallback func(session Session, topic string, msg []byte)
type ConnectedCallback func(session Session)
type ConnectionLostCallback func(session Session, reason error)

type Session struct {
	client         *delegate.Client
	opts           *delegate.ClientOptions
	OnMessage      MessageCallback
	OnConnected    ConnectedCallback
	OnLostConnect  ConnectionLostCallback
	ErrorMessage   string
	offlineMessage []byte
	onlineMessage  []byte
}

type Payload struct {
	Topic   string
	Qos     uint8
	Payload []byte
}

func CreateSession(clean bool, clientId string, brokers ...string) *Session {
	session := Session{}
	opts := delegate.NewClientOptions()

	if len(brokers) > 0 {
		for _, broker := range brokers {
			opts.AddBroker(broker)
		}
	}
	opts.SetClientID(clientId)
	opts.SetCleanSession(clean)
	opts.SetDefaultPublishHandler(session.DefaultMessageHandler)
	opts.SetConnectionLostHandler(session.onConnLost)
	opts.SetOnConnectHandler(session.onConned)

	c := delegate.NewClient(opts)
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

func (session *Session) SetWill(qos byte, retained bool, topic string, onLineMessage []byte, offlineMessage []byte) {
	session.opts.WillEnabled = true
	session.opts.WillQos = qos
	session.opts.WillRetained = retained
	session.opts.WillTopic = topic
	session.opts.WillPayload = offlineMessage
	session.onlineMessage = onLineMessage
	session.offlineMessage = offlineMessage
}

func (session *Session) CleanSession(clean bool) *Session {
	session.opts.CleanSession = clean
	return session
}

func (session *Session) onConnLost(_ delegate.Client, reason error) {
	if session.OnLostConnect != nil {
		session.OnLostConnect(*session, reason)
	}
}

func (session *Session) onConned(_ delegate.Client) {

	session.online()
	if session.OnConnected != nil {
		session.OnConnected(*session)
	}
}

func (session *Session) Publish(topic string, qos uint8, retained bool, payload interface{}) (err error) {

	client := *session.client
	if token := client.Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
		err = token.Error()
	}
	return err
}

func (session *Session) SimplePublish(topic string, payload interface{}) (err error) {

	return session.Publish(topic, 0, false, payload)
}

func (session *Session) Subscribe(topic string, qos uint8, callback MessageCallback) bool {
	client := *session.client

	var cb delegate.MessageHandler = nil
	if callback != nil {
		cb = func(client delegate.Client, message delegate.Message) {
			callback(*session, message.Topic(), message.Payload())
		}
	}

	if token := client.Subscribe(topic, qos, cb); token.Wait() && token.Error() != nil {
		session.ErrorMessage = token.Error().Error()
		return false
	} else {
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

func (session *Session) Connect() {

	client := delegate.NewClient(session.opts)

	session.client = &client

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		session.ErrorMessage = "connect failed"
	} else {
		session.ErrorMessage = ""
	}
}

func (session *Session) ConnectAndWait() (err error) {

	client := delegate.NewClient(session.opts)
	session.client = &client

	token := client.Connect()
	token.Wait()
	err = token.Error()

	return err
}

func (session *Session) Close() {
	client := *session.client
	session.offline()
	client.Disconnect(250)
}

func (session *Session) online() {
	session.status(true)
}

func (session *Session) offline() {
	session.status(false)
}

func (session *Session) status(online bool) {
	var payload []byte
	if online {
		payload = session.onlineMessage
	} else {
		payload = session.opts.WillPayload
	}

	if len(session.opts.WillTopic) > 0 && len(payload) > 0 {
		topic := session.opts.WillTopic
		qos := session.opts.WillQos
		retained := session.opts.WillRetained
		_ = session.Publish(topic, qos, retained, payload)
	}
}

func (session *Session) DefaultMessageHandler(_ delegate.Client, msg delegate.Message) {
	if session.OnMessage != nil {
		session.OnMessage(*session, msg.Topic(), msg.Payload())
	}
}
