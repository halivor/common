package util

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	qnSpc string
	qnKey string
	qnSec string
)

func QnInit(key, sec, space string) {
	qnSpc = space
	qnKey = key
	qnSec = sec
}

func QiniuToken(name string) string {
	putPolicy := storage.PutPolicy{
		Scope: qnSpc + name,
	}
	putPolicy.Expires = 60
	mac := qbox.NewMac(qnKey, qnSec)
	return putPolicy.UploadToken(mac)
}
