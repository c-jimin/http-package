package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type JsonData struct {
	body        []byte
	len         int
	err         error
	contentType ContentType
}

func NewJsonData(data map[string]interface{}) *JsonData {
	jd := new(JsonData)
	jd.contentType = ContentTypeJSON
	if data == nil {
		jd.body = nil
		jd.err = io.EOF
		return jd
	}

	postData, err := json.Marshal(data)
	if err != nil {
		jd.err = err
		return jd
	}
	fmt.Println(string(postData))
	jd.body = postData
	jd.len = len(jd.body)
	return jd

}

func (jd *JsonData) GetBody() (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewReader(jd.body)), jd.err
}

func (jd *JsonData) Len() int {
	return jd.len
}

func (jd *JsonData) ContentType() ContentType {
	return jd.contentType
}
