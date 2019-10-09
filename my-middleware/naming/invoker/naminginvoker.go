package invoker

import (
	//"fmt"
	"github.com/arma29/mid-rpc/my-middleware/infrastructure/srh"
	"github.com/arma29/mid-rpc/my-middleware/distribution/marshaller"
	"github.com/arma29/mid-rpc/shared"
	"github.com/arma29/mid-rpc/my-middleware/distribution/miop"
	clientProxy "github.com/arma29/mid-rpc/my-middleware/distribution/clientProxy"
	"github.com/arma29/mid-rpc/my-middleware/naming"

)

type NamingInvoker struct {}

func (NamingInvoker) Invoke(){
	srhImpl := srh.SRH{ServerHost:"",ServerPort:shared.NS_PORT}
	marshallerImpl := marshaller.Marshaller{}
	namingImpl := naming.NamingService{}
	miopPacketReply := miop.Packet{}
	replyParams := make([]interface{},1)

	// control loop
	for {
		// receive request packet
		rcvMsgBytes := srhImpl.Receive()

		// unmarshall request packet
		miopPacketRequest := marshallerImpl.Unmarshal(rcvMsgBytes)

		// extract operation name
		operation := miopPacketRequest.Body.RequestHeader.Operation

		// demux request
		switch operation {
		case "Register" :
			_p1 := miopPacketRequest.Body.RequestBody.Body[0].(string)
			_map := miopPacketRequest.Body.RequestBody.Body[1].(map[string]interface{})
			_proxyTemp := _map["Proxy"].(map[string]interface{})
			_p2 := clientProxy.ClientProxy{Host:_proxyTemp["Host"].(string),Port:int(_proxyTemp["Port"].(float64))}

			// dispatch request
			replyParams[0] = namingImpl.Register(_p1,_p2)
		case "Lookup":
			_p1 := miopPacketRequest.Body.RequestBody.Body[0].(string)

			// dispatch request
			replyParams[0] = namingImpl.Lookup(_p1)
		}

		// assembly reply packet
		repHeader := miop.ResponseHeader{}
		repBody := miop.ResponseBody{Body:replyParams}
		header := miop.Header{ByteOrder:true,Size: 4}
		body := miop.Body{ResponseHeader:repHeader,ResponseBody:repBody}
		miopPacketReply = miop.Packet{Header:header,Body:body}

		// marshall reply packet
		msgToClientBytes := marshallerImpl.Marshal(miopPacketReply)

		// send reply packet
		srhImpl.Send(msgToClientBytes)
	}
}