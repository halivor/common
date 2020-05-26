package packet

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	Uid  int64  `json:"uid"`
	Key  string `json:"key"`
	Sign string `json:"sign"`
	Rand int64  `json:"rand"`
	Exp  int64  `json:"exp"`
}

func (t *Token) String() string {
	return fmt.Sprintf("%d|%s|%s|%d|%d", t.Uid, t.Key, t.Sign, t.Exp, t.Rand)
}

func (t *Token) Parse(token string) {
	fields := strings.Split(token, "|")
	if len(fields) != 5 {
		return
	}
	t.Uid, _ = strconv.ParseInt(fields[0], 10, 0)
	t.Key = fields[1]
	t.Sign = fields[2]
	t.Exp, _ = strconv.ParseInt(fields[3], 10, 0)
	t.Rand, _ = strconv.ParseInt(fields[4], 10, 0)
}
