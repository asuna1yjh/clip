package request

import (
	"bytes"
	"log"
	"net/http"
)

var declareUrl = "https://poke.migu.cn/poke/media/upload/v1.0"

const (
	baskCookie = ""
)

// Declare 申报文件上传的方法
func Declare(cookie string, parameter []byte) {
	reader := bytes.NewReader(parameter)
	client := &http.Client{}
	req, err := http.NewRequest("POST", declareUrl, reader)
	if err != nil {
		log.Printf("http NewRequest to failed, err:%v\n", err)
		return
	}
	req.Header.Add("Cookie", cookie)
	res, err := client.Do(req)
	if err != nil {
		log.Printf("client.Do to failed, err:%v\n", err)
		return
	}
	defer res.Body.Close()
	return
}
