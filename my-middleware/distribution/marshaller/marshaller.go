package marshaller

import (
	"encoding/json"
	"github.com/arma29/mid-rpc/my-middleware/distribution/miop"
	"shared"
)

type Marshaller struct {}

func (Marshaller) Marshal(msg miop.Packet) []byte {

	result, err := json.Marshal(msg)
	shared.CheckError(err)

	return result
}

func (Marshaller) Unmarshal(msg []byte) miop.Packet {
	
	result := miop.Packet{}
	
	err := json.Unmarshal(msg, &result)
	shared.CheckError(err)
	
	return result
}