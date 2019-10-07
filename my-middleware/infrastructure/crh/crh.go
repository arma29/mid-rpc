package crh

import (
	"net"
	"strconv"
	"encoding/binary"
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

	defer conn.Close()

	// Send message to Server
	msgLengthBytes := make([]byte, 4)
	length := uint32(len(msg))

	binary.LittleEndian.PutUint32(msgLengthBytes, length)
	conn.Write(msgLengthBytes)
	
	_, err = conn.Write(msg)
	shared.CheckError(err)
	

	// Receiver Message
	msgReceivedLengthBytes := make([]byte, 4)
	_, err = conn.Read(msgReceivedLengthBytes)
	shared.CheckError(err)

	msgReceivedLengthInt := binary.LittleEndian.Uint32(msgReceivedLengthBytes)

	msgFromServer := make([]byte, msgReceivedLengthInt)
	_, err = conn.Read(msgFromServer)
	shared.CheckError(err)

	return msgFromServer
}