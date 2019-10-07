package crh

import (
	"net"
	"strconv"
	"enconding/binary"
	"github.com/arma29/mid-rpc/shared"
)

type CRH struct {
	ServerHost string
	ServerPort int
}

func (crh CRH) SendReceive(msg []byte) []byte {

	var conn net.Conn
	var err error

	for {
		conn, _ = net.Dial("tcp", crh.ServerHost + ":" + strconv.Itoa(crh.ServerPort))

		if err == nil {
			break
		}
	}

	defer conn.Close

	// Send message to Server
	msgSize := make([]byte, 4)
	length := uint32(len(msg))

	binary.LittleEndian.PutUint32(sizeMsgToServer, length)
	_, err := conn.Write(length)
	shared.CheckError(err)
	
	_, err = conn.Write(msg)
	shared.CheckError(err)
	

	// Receiver Message
	msgReceivedLengthBytes := make([]byte, 4)
	_, err = conn.Read(msgReceivedLength)
	shared.CheckError(err)

	msgReceivedLengthInt := binary.LittleEndian.Uint32(msgReceivedLengthBytes)

	msgFromServer := make([]byte, msgReceivedLengthInt)
	_, err = conn.Read(msgFromServer)
	shared.CheckError(err)

	return msgFromServer
}