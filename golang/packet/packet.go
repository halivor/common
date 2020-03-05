package packet

import (
	"encoding/json"

	ce "github.com/halivor/common/golang/util/errno"
)

type Rsp struct {
	ErrCode int         `json:"error_code"`
	ErrMsg  string      `json:"error_message"`
	Data    interface{} `json:"data,omitempty"`
}

type RspRaw struct {
	ErrCode int             `json:"error_code"`
	ErrMsg  string          `json:"error_message,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func RespFail(e ce.Errno) *Rsp {
	return &Rsp{ErrCode: e.Errno(), ErrMsg: e.Error()}
}

func RespSucc(data interface{}) *Rsp {
	return &Rsp{ErrCode: 0, ErrMsg: "success", Data: data}
}

func RespString(message string) *Rsp {
	return &Rsp{ErrCode: ce.ERR_ORI, ErrMsg: message}
}

func RespError(e error) *Rsp {
	return &Rsp{ErrCode: ce.ERR_ORI, ErrMsg: e.Error()}
}

func RespRaw(data json.RawMessage) *Rsp {
	return &Rsp{ErrCode: 0, ErrMsg: "success", Data: data}
}
