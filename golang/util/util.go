package util

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	p "github.com/halivor/common/golang/packet"
	ce "github.com/halivor/common/golang/util/errno"
	log "github.com/halivor/goutility/logger"
)

var ul = log.NewLog("/data/logs/util.log", "[util]", log.LstdFlags, log.TRACE)

func InitLog(olog log.Logger) {
	if olog != nil {
		log.Release(ul)
		ul = olog
	}
}

func RespParse(r io.ReadCloser) (json.RawMessage, ce.Errno) {
	pb, _ := ioutil.ReadAll(r)
	rsp := &p.RspRaw{}
	json.Unmarshal(pb, rsp)
	r.Close()
	return rsp.Data, ce.Errno(rsp.ErrCode)
}

func HttpGetRaw(url string) ([]byte, ce.Errno) {
	switch resp, e := http.Get(url); {
	case e != nil:
		ul.Warn("http请求失败", e)
		return nil, ce.SRV_ERR
	case resp.StatusCode != 200:
		resp.Body.Close()
		ul.Warn("http应答码", resp.StatusCode, url)
		return nil, ce.SRV_ERR
	default:
		pb, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		return pb, ce.SUCC
	}
}

func HttpGet(url string) ([]byte, ce.Errno) {
	switch resp, e := http.Get(url); {
	case e != nil:
		ul.Warn("http请求失败", e)
		return nil, ce.SRV_ERR
	case resp.StatusCode != 200:
		resp.Body.Close()
		ul.Warn("http应答码", resp.StatusCode, e, url)
		return nil, ce.SRV_ERR
	default:
		return RespParse(resp.Body)
	}
}

func HttpPostJson(url string, data interface{}) ([]byte, ce.Errno) {
	pb, _ := json.Marshal(data)
	return HttpPost(url, "application/json", pb)
}

func HttpPost(url, ctype string, body []byte) ([]byte, ce.Errno) {
	switch resp, e := http.Post(url, ctype, bytes.NewReader(body)); {
	case e != nil:
		ul.Warn("HTTP POST FAILED", e)
		return nil, ce.SRV_ERR
	case resp.StatusCode != 200:
		ul.Warn("HTTP status code", resp.StatusCode, url)
		return nil, ce.SRV_ERR
	default:
		return RespParse(resp.Body)
	}
}
