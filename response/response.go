package response

import (
	"bytes"
	"io"
	"net/http"
)

type Response struct {
	OriResp *http.Response
	Error   error
}

func (r *Response) StatusCode() int {
	return r.OriResp.StatusCode
}

func (r *Response) Status() string {
	return r.OriResp.Status
}

func (r *Response) Proto() string {
	return r.OriResp.Proto
}

func (r *Response) Read(p []byte) (n int, err error) {
	return r.OriResp.Body.Read(p)
}

func (r *Response) Close() error {
	return r.OriResp.Body.Close()
}

func (r *Response) Headers() map[string][]string {
	return r.OriResp.Header
}

func (r *Response) Buffer() *bytes.Buffer {
	buffer := bytes.NewBuffer(nil)
	_, _ = io.Copy(buffer, r)
	return buffer
}

func (r *Response) HasError() bool {
	return r.Error != nil
}
