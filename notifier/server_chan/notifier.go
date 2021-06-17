package server_chan

import (
	"fmt"
	"net/http"
	"net/url"
)

type Notifier struct {
	sendKey string
	client *http.Client
}

func NewNotifier(key string) *Notifier {
	return &Notifier{
		sendKey: key,
		client: &http.Client{},
	}
}

func (n *Notifier) Notify(title, message string) error {
	serverChanUrl := fmt.Sprintf("https://sctapi.ftqq.com/%s.send?title=%s&desp=%s",
		n.sendKey, url.QueryEscape(title), url.QueryEscape(message))
	_, err := n.client.Get(serverChanUrl)
	return err
}

