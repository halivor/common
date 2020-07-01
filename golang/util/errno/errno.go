package errno

import (
	"fmt"
)

type Errno int64

const (
	SUCC    Errno = 0
	strSucc       = "成功"
)

var errmsg = make(map[Errno]string, 1024)

func Set(no Errno, msg string) {
	errmsg[no] = msg
}

func init() {
	Set(SUCC, strSucc)
}
func (e Errno) Errno() int {
	return int(e)
}

func (e Errno) String() string {
	return fmt.Sprintf("%d->%s", e.Errno(), errmsg[e])
}

func (e Errno) Error() string {
	return errmsg[e]
}
