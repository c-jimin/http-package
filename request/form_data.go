package request

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
)

type FormData struct {
	body        []byte
	len         int
	err         error
	contentType ContentType
}

func NewFormData(data map[string]interface{}) *FormData {
	d := new(FormData)
	d.contentType = ContentTypeXWwwFormUrlencoded
	if data == nil {
		d.body = nil
		d.err = io.EOF
		return d
	}

	var PostDate url.Values = make(map[string][]string)
	for k, v := range data {
		PostDate.Add(k, interface2string(v))
	}
	queryString := PostDate.Encode()
	d.body = []byte(queryString)
	d.len = len(d.body)
	return d

}

func (fd *FormData) GetBody() (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewReader(fd.body)), fd.err
}

func (fd *FormData) Len() int {
	return fd.len
}

func (fd *FormData) ContentType() ContentType {
	return fd.contentType
}

func interface2string(v interface{}) string {
	if v == nil {
		return "null"
	}
	return fmt.Sprintf("%v", v)
}
