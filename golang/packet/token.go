package packet

import (
	"fmt"
)

type Token struct {
	Uid  int64  `json:"uid"`
	Key  string `json:"key"`
	Sign string `json:"sign"`
	Exp  int64  `json:"exp"`
}

func (t *Token) String() string {
	return fmt.Sprintf("%d|%s|%d", t.Uid, t.Sign, t.Exp)
}
