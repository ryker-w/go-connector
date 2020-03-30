package stomp

import "encoding/json"

func (s Sender) Send(payload []byte, contentType string) error {
	return s.proxy.Send(s.dest, contentType, payload, nil)
}

func (s Sender) SendText(payload string) error {
	return s.Send([]byte(payload), "text/plain")
}

// payload: json(bytes/string/struct)
func (s Sender) SendJson(payload interface{}) error {

	contentType := "application/json"
	if payload == nil {
		return nil
	}
	bytes, ok := payload.([]byte)
	if ok {
		return s.Send(bytes, contentType)
	}
	if str, ok := payload.(string); ok {
		return s.Send([]byte(str), contentType)
	}

	bs, err := json.Marshal(&payload)
	if err != nil {
		return err
	}
	return s.Send(bs, contentType)
}
