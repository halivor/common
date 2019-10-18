package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var (
	TopUpWeChatUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?" +
		"key=504171d0-f32f-440b-8554-c3e0dd4c255c"
)

func InitWeChatTopUpUrl(url string) {
	TopUpWeChatUrl = url
}

func NotifyTopUpWeChat(data string) {
	p := struct {
		Type     string            `json:"msgtype"`
		Markdown map[string]string `json:"markdown"`
	}{"markdown", map[string]string{
		"content": data,
	}}
	pb, _ := json.Marshal(p)
	ul.Trace("notify wechat", string(pb))
	http.Post(TopUpWeChatUrl, "application/json", bytes.NewReader(pb))
}
