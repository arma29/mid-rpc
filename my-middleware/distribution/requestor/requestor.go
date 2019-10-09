package requestor

import (
	"github.com/arma29/mid-rpc/my-middleware/aux"
	"github.com/arma29/mid-rpc/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rpc/my-middleware/infrastructure/crh"
	"github.com/arma29/mid-rpc/my-middleware/distribution/miop"

)

type Requestor struct{}

func (requestor Requestor) Invoke(inv aux.Invocation) interface{} {
	
	marshallerInstance := marshaller.Marshaller{}
	crhInstance := crh.CRH{ServerHost: inv.Host, ServerPort: inv.Port}

	reqHeader := miop.RequestHeader{Operation: inv.Request.Op, ObjectID: inv.ObjectID}
	reqBody := miop.RequestBody{Body:inv.Request.Params}
	header := miop.Header{ByteOrder: true, Size: 4 }
	body := miop.Body{RequestHeader: reqHeader, RequestBody: reqBody}
	packetRequest := miop.Packet{Header: header, Body: body}

	msgRequestBytes := marshallerInstance.Marshal(packetRequest)

	msgResponseBytes := crhInstance.SendReceive(msgRequestBytes)

	msgResponsePacket := marshallerInstance.Unmarshal(msgResponseBytes)

	result := msgResponsePacket.Body.ResponseBody.Body


	return result
}