package request

import (
	"context"
	"github.com/c-jimin/http-package/client"
	"github.com/c-jimin/http-package/response"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	Method    string
	URL       string
	Params    map[string]string
	Headers   map[string]string
	UserAgent string
	Body      Body
	Ctx       context.Context
}

func (r *Request) makeHttpRequest() (*http.Request, error) {
	_url, err := url.Parse(r.URL)
	if err != nil {
		return nil, err
	}
	if r.Params != nil {
		query := _url.Query()
		for k, v := range r.Params {
			query.Add(k, v)
		}
		_url.RawQuery = query.Encode()
	}
	if r.Body == nil {
		r.Body = &noBody{}
	}
	body, err := r.Body.GetBody()
	if err != nil && err != io.EOF {
		return nil, err
	}

	headers := NewHeaders(r.Headers)
	if r.Body.ContentType() != "" {
		headers.Set("Content-Type", string(r.Body.ContentType()))
	}
	if r.UserAgent != "" {
		headers.Set("User-Agent", r.UserAgent)
	}

	req := &http.Request{
		Method:        r.Method,
		URL:           _url,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        *headers,
		Body:          body,
		Host:          _url.Host,
		ContentLength: int64(r.Body.Len()),
		GetBody:       r.Body.GetBody,
	}
	if r.Ctx != nil {
		req = req.WithContext(r.Ctx)
	}
	return req, nil
}

func (r Request) Send() (*response.Response, error) {
	req, err := r.makeHttpRequest()
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		OriResp: resp,
	}, nil
}

func (r Request) SendWithClient(client *httpClient.Client) *response.Response {
	req, err := r.makeHttpRequest()
	if err != nil {
		return &response.Response{
			Error: err,
		}
	}
	return client.Send(req)
}
