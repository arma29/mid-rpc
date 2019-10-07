package requestor

import (
	"github.com/arma29/mid-rpc/shared"
	"github.com/arma29/mid-rpc/my-middleware/aux"
	"github.com/arma29/mid-rpc/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rpc/my-middleware/distribution/crh"
	"github.com/arma29/mid-rpc/my-middleware/distribution/miop"
	"unsafe"
)

type Requestor struct{}

func (requestor Requestor) Invoke(inv aux.Invocation) interface{} {
	
	marshallerInstance := marshaller.Marshaller{}
	crhInstance := crh.CRH{ServerHost: inv.Host, ServerPort: inv.Port}

	reqHeader := miop.RequestHeader{RequestID: 0, Operation: inv.Request.Op}
	reqBody := miop.RequestBody{Body:inv.Request.Params}
	header := miop.Header{ByteOrder: true, Size: unsafe.Sizeof(inv.Request.Params) }
	body := miop.Body{RequestHeader: reqHeader, RequestBody: reqBody}
	packetRequest := miop.Packet{Header: header, Body: body}

	msgRequestBytes := marshallerInstance.Marshal(packetRequest)

	msgResponseBytes := crhInstance.SendReceive(msgRequestBytes)
	msgResponsePacket := marshallerInstance.Unmarshal(msgResponseBytes)

	result = msgResponsePacket.ResponseBody.Body

	return result
}