package packet

import (
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	ce "github.com/halivor/common/golang/util/errno"
)

var (
	ver    int32   = 0
	appId  int32   = 0
	header *Header = nil
	ext            = false
)

func InitProto(ver int32, appid int32, trace bool) {
	header = &Header{
		Ver:   ver,
		AppId: appid,
		Type:  MessageType_PROTO,
	}
	ext = trace
}

func NewRequest(data proto.Message) *Request {
	pb, _ := proto.Marshal(data)
	var extension *Ext
	if ext {
		extension = &Ext{
			TraceId: uuid.New().String(),
			SpanId:  uuid.New().String(),
		}
	}
	return &Request{
		Header: header,
		Body:   pb,
		Ext:    extension,
	}
}

func NewResponse(req *Request, msg proto.Message, en ce.Errno) *Response {
	pb, _ := proto.Marshal(msg)
	var extension Ext
	if ext {
		extension = *req.Ext
	}
	return &Response{
		Header: req.GetHeader(),
		Body:   pb,
		Ext:    &extension,
		Errno:  int32(en),
	}
}
