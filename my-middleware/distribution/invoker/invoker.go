package invoker 

import (
	"github.com/arma29/mid-rpc/my-middleware/infrastructure/srh"
	"github.com/arma29/mid-rpc/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rpc/my-middleware/distribution/miop"
	"github.com/arma29/mid-rpc/shared"
	"github.com/arma29/mid-rpc/application"
)

type FibonacciInvoker struct {}

func NewFibonnaciInvoker() FibonacciInvoker {
	p := new(FibonacciInvoker)

	return *p
}

func (FibonacciInvoker) Invoke() {

	srhInstance := srh.SRH{ ServerHost:"localhost", ServerPort: shared.SERVER_PORT }
	marshallerInstance := marshaller.Marshaller{}

	fibonacciApp := application.FibonacciApp{}
	resultParams := make([]interface{}, 1)

	for {
		msgBytes := srhInstance.Receive()

		miopPacketRequest := marshallerInstance.Unmarshal(msgBytes)
		operation := miopPacketRequest.Body.RequestHeader.Operation

		if (operation == "GetFibo") {
			n := int32(miopPacketRequest.Body.RequestBody.Body[0].(float64))
			resultParams[0] = fibonacciApp.GetFibOf(n)
		}

		resHeader := miop.ResponseHeader{}
		resBody := miop.ResponseBody{ Body: resultParams }
		header := miop.Header{ ByteOrder: true, Size: 0 }
		body := miop.Body{ ResponseHeader: resHeader, ResponseBody: resBody }
		miopPacketResponse := miop.Packet{Header: header, Body: body}

		msgToSendBytes := marshallerInstance.Marshal(miopPacketResponse)

		srhInstance.Send(msgToSendBytes)

	}

}