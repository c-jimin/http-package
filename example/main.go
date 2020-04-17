package main

import (
	"bytes"
	"fmt"
	"github.com/c-jimin/http-package/client"
	"github.com/c-jimin/http-package/request"
	"io"
)

func main() {
	client := httpClient.New()
	//err := client.SetProxy("socks5://127.0.0.1:1080")
	//if err != nil {
	//	panic(err)
	//}
	resp := request.Request{
		Method: "GET",
		URL:    "http://ori.codetech.top",
		Body: request.NewFormData(map[string]interface{}{
			"string": "string",
			"bool":   true,
			"int":    123,
			"float":  1.234,
			"null":   nil,
		}),
	}.SendWithClient(client)
	if resp.HasError() {
		panic(resp.Error)
	}
	defer resp.Close()
	b := bytes.NewBuffer(nil)
	_, _ = io.Copy(b, resp)
	fmt.Println(b.String())
}
