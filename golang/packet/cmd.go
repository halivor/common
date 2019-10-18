package packet

import (
	"strconv"
)

type CmdID uint32

const (
	C_AUTH      CmdID = 1000
	C_AUTH_SUCC CmdID = 1001
	C_PING      CmdID = 1010
	C_PONG      CmdID = 1011
)

func (id CmdID) ToString() string {
	return strconv.Itoa(int(id))
}
