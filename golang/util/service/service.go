package service

import (
	"context"
	"sync"

	"github.com/golang/protobuf/proto"
	cp "github.com/halivor/common/golang/packet"
	ce "github.com/halivor/common/golang/util/errno"
)

type method func(context.Context, *cp.Request) (cp.Response, error)

type Service interface {
	SetUp(name string, m method)
	Call(name string, req proto.Message, rsp proto.Message) ce.Errno
}

type service struct {
	sync.RWMutex
	mms map[string]method
}

func New() Service {
	return &service{}
}

func (svc *service) SetUp(name string, m method) {
	svc.Lock()
	defer svc.Unlock()
	svc.mms[name] = m
	// TODO: 同步服务注册
}

func (svc *service) Call(name string, req proto.Message, rsp proto.Message) ce.Errno {
	svc.RLock()
	m, ok := svc.mms[name]
	svc.RUnlock()
	if ok {
		return ce.DATA_INVALID
	}
	rst, e := m(context.Background(), cp.NewRequest(req))
	if e != nil {
		return ce.SRV_ERR
	}
	proto.Unmarshal(rst.GetBody(), rsp)
	return ce.Errno(rst.Errno)
}
