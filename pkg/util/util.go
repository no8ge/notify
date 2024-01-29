package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"net/url"

	types "github.com/no8geo/notify/pkg"
)

type WeChat struct {
	Key string `json:"key"`
}

func (w *WeChat) SendMarkdown(m *types.MarkdownMsg) error {

	params := url.Values{}
	Url, err := url.Parse("https://qyapi.weixin.qq.com/cgi-bin/webhook/send")
	if err != nil {
		return nil
	}
	params.Set("key", w.Key)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	client := &http.Client{}
	bytesData, _ := json.Marshal(m)
	req, _ := http.NewRequest("POST", urlPath, bytes.NewReader(bytesData))

	req.Header.Add("Content-Type", "application/json")

	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
		fmt.Println(string(body))
	}

	return nil

}
