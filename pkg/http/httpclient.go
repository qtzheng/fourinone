package http

import (
	"net/http"
	"path"
)

type HttpClient interface {
	Do() (*http.Response, error)
}
type peerClient struct {
	client *http.Client
	urlStr string
	action string
}

func (p *peerClient) Do() (*http.Response, error) {
	req, err := http.NewRequest(p.action, p.urlStr, nil)
	if err != nil {
		return nil, err
	}
	return p.client.Do(req)
}
func NewPeerClient(action, route, url string) HttpClient {
	urlStr := url + "/" + route
	urlStr = path.Clean(urlStr)
	return &peerClient{
		client: http.DefaultClient,
		urlStr: urlStr,
		action: action,
	}
}
