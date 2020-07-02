package packet

import (
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	ce "github.com/halivor/common/golang/util/errno"
)

var (
	ver    int32   = 0
	appid  int32   = 0
	header *Header = nil
	ext            = false
)

func InitHeader(ver int64, id int64, trace bool) {
	header = &Header{
		Ver:   ver,
		AppId: id,
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

func NewInRequest(req *Request, data proto.Message) *Request {
	pb := []byte("")
	if data != nil {
		pb, _ = proto.Marshal(data)
	}
	return &Request{
		Header: req.GetHeader(),
		Body:   pb,
		Ext:    req.GetExt(),
	}
}

func NewResponse(req *Request, msg proto.Message, en ce.Errno) *Response {
	pb := []byte("")
	if msg != nil {
		pb, _ = proto.Marshal(msg)
	}
	var extension Ext
	if ext {
		extension = *req.Ext
	}
	return &Response{
		Header: req.GetHeader(),
		Body:   pb,
		Ext:    &extension,
		Errno:  int64(en),
	}
}
