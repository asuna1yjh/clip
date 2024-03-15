package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type UploadResp struct {
	Code string `json:"code"`
	Info string `json:"info"`
	Data []struct {
		Url          interface{} `json:"url"`
		ErrMsg       interface{} `json:"errMsg"`
		Date         interface{} `json:"date"`
		UuIdFileName interface{} `json:"uuIdFileName"`
		LocalPath    string      `json:"localPath"`
		ServiceType  string      `json:"serviceType"`
		Duration     int         `json:"duration"`
		VideoLength  string      `json:"videoLength"`
		Format       string      `json:"format"`
		Width        int         `json:"width"`
		Height       int         `json:"height"`
		FrameRate    float64     `json:"frameRate"`
	} `json:"data"`
}

func UploadFile(fileName string) (*UploadResp, error) {
	url := "https://poke.migu.cn/poke/media/upload/v1.0"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("serviceType", "01")
	file, err := os.Open(fileName)
	defer file.Close()
	part, err := writer.CreateFormFile("files", filepath.Base(fileName))
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var u = new(UploadResp)
	err = json.Unmarshal(body, u)
	if err != nil {
		fmt.Println("Unmarshal to failed")
		return nil, err
	}
	return u, err
}
