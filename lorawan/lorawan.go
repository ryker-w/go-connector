package lorawan

import (
	"fmt"
	"github.com/lishimeng/go-connector/mqtt"
	"github.com/lishimeng/go-libs/log"
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

func New(broker string, clientId string, topicUpLink string, topicDownLink string, qos uint8) (*Connector, error) {

	c := Connector{
		host:             broker,
		clientId:         clientId,
		downLinkTopicTpl: topicDownLink,
		upLinkTopicTpl:   topicUpLink,
		qos:              qos,
	}

	var onConnect = func(s mqtt.Session) {
		log.Fine("lora subscribe upLink topic:%s", c.upLinkTopicTpl)
		c.session.Subscribe(c.upLinkTopicTpl, c.qos, nil)
	}
	var onConnLost = func(s mqtt.Session, reason error) {
		log.Fine("lora lost connection")
		log.Fine(reason)
	}
	c.session = mqtt.CreateSession(false, c.clientId, c.host)

	c.session.OnConnected = onConnect
	c.session.OnLostConnect = onConnLost
	c.session.OnMessage = c.messageCallback

	return &c, nil
}

func (c Connector) GetSession() *mqtt.Session {
	return c.session
}

func (c *Connector) Connect() {
	log.Fine("lora connect %s", c.host)
	for err := c.ConnectOnce(); err != nil; {
		log.Fine(err)
	}
}

func (c *Connector) ConnectOnce() error {
	log.Fine("lora connect %s", c.host)
	return c.session.ConnectAndWait()
}

func (c *Connector) SetUpLinkListener(listener UpLinkListener) {
	c.listener = listener
}

// 监听数据上传
///
func (c *Connector) messageCallback(_ mqtt.Session, topic string, msg []byte) {

	log.Fine("lora upLink:%s", topic)
	payload, err := onDataUpLink(msg)
	if err != nil {
		return
	}

	c.listener(payload)
}

func (c Connector) DownLink(appId string, deviceEUI string, payload PayloadTx) (err error) {
	data := convertJsonDownLinkData(payload) // 序列化得到json类型data字符串

	topic := fmt.Sprintf(c.downLinkTopicTpl, appId, deviceEUI)

	err = c.session.Publish(topic, c.qos, false, data)
	return err
}
