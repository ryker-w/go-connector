package stomp

import (
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/stomp"
	"time"
)

type MessageMode string

const (
	defaultPort        = 61613
	defaultHost        = "localhost"
	defaultNetworkType = "tcp"
	defaultStompHost   = "/"

	Queue MessageMode = "/queue/"
	Topic MessageMode = "/topic/"
)

type MessageHandler func([]byte)

type Connector struct {
	hasAuth               bool
	username              string
	password              string
	port                  int
	host                  string
	addr                  string
	hasHeartbeat          bool
	recvTimeout           time.Duration
	sendTimeout           time.Duration
	networkType           string
	stompHost             string
	proxy                 *stomp.Conn
	defaultMessageHandler MessageHandler
	close                 chan byte
}

type Subscription struct {
	delegate *stomp.Subscription
	stop     chan byte
}

type Sender struct {
	proxy *stomp.Conn
	dest  string
}

func New() *Connector {

	return &Connector{}
}

func (c *Connector) Auth(username string, password string) *Connector {

	c.hasAuth = true
	c.username = username
	c.password = password
	return c
}

func (c *Connector) Heartbeat(sendTimeout, recvTimeout time.Duration) *Connector {

	c.hasHeartbeat = true
	c.recvTimeout = recvTimeout
	c.sendTimeout = sendTimeout
	return c
}

func (c *Connector) DefaultMessageHandler(handler MessageHandler) *Connector {

	c.defaultMessageHandler = handler
	return c
}

func (c *Connector) HostPort(host string, port int) *Connector {

	c.host = host
	c.port = port
	return c
}

func (c *Connector) Network(nw string) *Connector {
	c.networkType = nw
	return c
}

func (c *Connector) Addr(addr string) *Connector {
	c.addr = addr
	return c
}

func (c *Connector) StompHost(h string) *Connector {
	c.stompHost = h
	return c
}

func (c *Connector) Connect() (err error) {
	if c.port == 0 {
		c.port = defaultPort
	}
	if len(c.host) == 0 {
		c.host = defaultHost
	}
	if len(c.networkType) == 0 {
		c.networkType = defaultNetworkType
	}
	if len(c.addr) == 0 {
		c.addr = fmt.Sprintf("%s:%d", c.host, c.port)
	}

	if len(c.stompHost) == 0 {
		c.stompHost = defaultStompHost
	}

	c.proxy, err = stomp.Dial(c.networkType, c.addr, c.afterConnect()...)
	c.close = make(chan byte)
	return
}

func (c *Connector) afterConnect() (cb []func(*stomp.Conn) error) {

	if c.hasAuth {
		cb = append(cb, stomp.ConnOpt.Login(c.username, c.password))
	}
	cb = append(cb, stomp.ConnOpt.Host(c.stompHost))
	if c.hasHeartbeat {
		cb = append(cb, stomp.ConnOpt.HeartBeat(c.sendTimeout, c.recvTimeout))
	} else {
		cb = append(cb, stomp.ConnOpt.HeartBeat(time.Second*16, time.Second*4))
	}
	return
}

func (c *Connector) Subscriber(mode MessageMode, destination string, handler MessageHandler) (subscribe *Subscription, err error) {

	dest := string(mode) + destination
	sub, err := c.proxy.Subscribe(dest, stomp.AckAuto)
	if err != nil {
		return
	}

	stopSubscriber := make(chan byte)
	subscribe = &Subscription{delegate: sub, stop: stopSubscriber}

	if handler != nil {
		go c.subscribeLoop(subscribe, handler)
	} else {
		go c.subscribeLoop(subscribe, c.defaultMessageHandler)
	}
	return
}

func (c *Connector) subscribeLoop(sub *Subscription, handler MessageHandler) {
	for {
		select {
		case <-sub.stop:
			return
		case <-c.close:
			return
		case msg := <-sub.delegate.C:
			err := msg.Err
			if err != nil {
				log.Info(err)
				log.Info("stop subscribe loop")

				sub.stop <- 0x00
				c.LostConn(err)
				return
			}
			body := msg.Body
			handler(body)
		}
	}
}

func (c *Connector) Unsubscribe(sub *Subscription) {
	close(sub.stop)
	_ = sub.delegate.Unsubscribe(nil)
}

func (c Connector) Sender(mode MessageMode, destination string) (s *Sender) {

	s = &Sender{
		proxy: c.proxy,
		dest:  string(mode) + destination,
	}
	return
}

func (c Connector) Transaction(handler func(connector Connector) error) (err error) {
	trans := c.proxy.Begin()
	err = handler(c)
	if err != nil {
		_ = trans.Abort()
	} else {
		err = trans.Commit()
	}
	return
}

func (c Connector) Send(mode MessageMode, destination string, contentType string, data []byte) (err error) {

	dest := string(mode) + destination
	err = c.proxy.Send(dest, contentType, data, nil)
	return
}

func (c *Connector) LostConn(err error) {

}

func (c *Connector) Close() {
	close(c.close)

	if c.proxy != nil {
		_ = c.proxy.Disconnect()
	}
}
