package util

import (
	"bytes"
	"encoding/json"
	"net/http"

	p "github.com/halivor/common/golang/packet"
)

var bcDm string

func BcInit(domain string) {
	bcDm = domain
}

func UniCast(uid uint64, cmd p.CmdID, message interface{}) error {
	pb, _ := json.Marshal([]interface{}{uid, cmd, message})
	switch resp, e := http.Post(bcDm+"/sender/v1s/user",
		"application/json", bytes.NewReader(pb)); {
	case e != nil:
		return e
	default:
		resp.Body.Close()
	}
	return nil
}

func RoomCast(rid uint64, cmd p.CmdID, message interface{}) error {
	pb, _ := json.Marshal([]interface{}{rid, cmd, message})
	switch resp, e := http.Post(bcDm+"/sender/v1s/room",
		"application/json", bytes.NewReader(pb)); {
	case e != nil:
		return e
	default:
		resp.Body.Close()
	}
	return nil
}

func BroadCast(cmd p.CmdID, message interface{}) error {
	pb, _ := json.Marshal([]interface{}{cmd, message})
	switch resp, e := http.Post(bcDm+"/sender/v1s/broadcast", "application/json", bytes.NewReader(pb)); {
	case e != nil:
		return e
	default:
		resp.Body.Close()
	}
	return nil
}
