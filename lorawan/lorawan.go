package lorawan

import (
	"fmt"
	log "github.com/jeanphorn/log4go"
	"github.com/lishimeng/go-connector/mqtt"
)

type UpLinkListener func(data PayloadRx)

type Connector struct {
	host             string
	clientId         string
	qos              uint8
	downLinkTopicTpl string
	upLinkTopicTpl   string
	session          *mqtt.Session
	listener         UpLinkListener
}

func New(broker string, clientId string, topicUpLink string, topicDownLink string, qos uint8) (Connector, error) {

	c := Connector{
		host:             broker,
		clientId:         clientId,
		downLinkTopicTpl: topicDownLink,
		upLinkTopicTpl:   topicUpLink,
		qos:              qos,
	}

	var onConnect = func(s mqtt.Session) {
		log.Debug("lora mqtt subscribe upLink topic:%s", c.upLinkTopicTpl)
		c.session.Subscribe(c.upLinkTopicTpl, c.qos, nil)
	}
	var onConnLost = func(s mqtt.Session, reason error) {
		log.Debug("lora mqtt lost connection")
		log.Debug(reason)
		c.session.State = false
	}
	c.session = mqtt.CreateSession(c.host, c.clientId)

	c.session.OnConnected = onConnect
	c.session.OnLostConnect = onConnLost
	c.session.OnMessage = c.messageCallback

	return c, nil
}

func (c Connector) GetSession() mqtt.Session {
	return *c.session
}

func (c *Connector) Connect() {
	log.Debug("lora mqtt connect to broker %s", c.host)
	for err := c.ConnectOnce(); err != nil; {
		log.Debug(err)
	}
}

func (c *Connector) ConnectOnce() error {
	log.Debug("lora mqtt connect to broker %s", c.host)
	return c.session.ConnectAndWait()
}

func (c *Connector) SetUpLinkListener(listener UpLinkListener) {
	c.listener = listener
}

// 监听数据上传
///
func (c *Connector) messageCallback(mqSession mqtt.Session, topic string, mqttMsg []byte) {

	log.Debug("receive lora upLink data %s", topic)
	payload, err := onDataUpLink(mqttMsg)
	if err != nil {
		return
	}

	c.listener(payload)
}

func (c Connector) DownLink(appId string, deviceEUI string, payload PayloadTx) (err error) {
	data := convertJsonDownLinkData(payload) // 序列化得到json类型data字符串

	topic := fmt.Sprintf(c.downLinkTopicTpl, appId, deviceEUI)

	err = c.session.Publish(topic, c.qos, data)
	return err
}
