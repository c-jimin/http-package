package httpClient

import (
	"github.com/c-jimin/http-package/response"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client    *http.Client
	transport *http.Transport
}

func New() *Client {
	transport := &http.Transport{}
	client := &http.Client{
		Transport: transport,
	}
	return &Client{
		client:    client,
		transport: transport,
	}
}

func (hc *Client) SetTimeout(timeout time.Duration) {
	hc.client.Timeout = timeout
}

func (hc *Client) SetProxy(proxy string) error {
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	hc.transport.Proxy = http.ProxyURL(proxyUrl)
	return nil
}

func (hc *Client) Send(req *http.Request) *response.Response {
	resp, err := hc.client.Do(req)
	return &response.Response{
		OriResp: resp,
		Error:   err,
	}
}
