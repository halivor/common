package packet

import (
	"fmt"
)

type Token struct {
	Uid    int64  `json:"uid"`
	Key    string `json:"-"`
	Sign   string `json:"sign"`
	Expire int64  `json:"expire"`
}

func (t *Token) String() string {
	return fmt.Sprintf("%d|%s|%d", t.Uid, t.Sign, t.Expire)
}
