package request

import (
	"bytes"
	"io"
	"io/ioutil"
)

type Body interface {
	GetBody() (io.ReadCloser, error)
	Len() int
	ContentType() ContentType
}

type noBody struct{}

func (n noBody) GetBody() (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewBuffer(nil)), io.EOF
}

func (n noBody) Len() int {
	return 0
}

func (n noBody) ContentType() ContentType {
	return ""
}
