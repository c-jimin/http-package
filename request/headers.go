package request

import "net/http"

func NewHeaders(headers map[string]string) *http.Header {
	header := make(http.Header)
	for k, v := range headers {
		header[k] = []string{v}
	}
	return &header
}
