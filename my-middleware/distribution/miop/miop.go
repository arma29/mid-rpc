package miop

type Packet struct {
	Header Header
	Body Body
}

type Header struct {
	ByteOrder bool // True to Little Endian
	Size int 
}

type Body struct {
	RequestHeader RequestHeader
	RequestBody RequestBody
	ResponseHeader ResponseHeader
	ResponseBody ResponseBody
}

type RequestHeader struct {
	Operation string
	ObjectID int
}

type RequestBody struct {
	Body []interface{}
}

type ResponseHeader struct {
	Context string
	RequestId int
	Status int
}

type ResponseBody struct {
	Body []interface{}
}