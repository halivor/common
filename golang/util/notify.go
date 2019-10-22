package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var (
	TopUpWeChatUrl = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?"
)

func InitNotifyWeChatUrl(url string) {
	TopUpWeChatUrl = url
}

func NotifyWeChatMarkDown(markdown string) {
	p := struct {
		Type     string            `json:"msgtype"`
		Markdown map[string]string `json:"markdown"`
	}{"markdown", map[string]string{
		"content": markdown,
	}}
	pb, _ := json.Marshal(p)
	ul.Trace("notify wechat", string(pb))
	http.Post(TopUpWeChatUrl, "application/json", bytes.NewReader(pb))
}

func NotifyWeChatText(text string) {
	p := struct {
		Type string                 `json:"msgtype"`
		Text map[string]interface{} `json:"text"`
	}{"text", map[string]interface{}{
		"content": text,
	}}
	pb, _ := json.Marshal(p)
	http.Post(TopUpWeChatUrl, "application/json", bytes.NewReader(pb))
}
