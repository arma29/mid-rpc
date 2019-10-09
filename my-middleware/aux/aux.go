package aux

type Request struct {
	Op string
	Params []interface{}
}

type Invocation struct {
	Host string
	Port int
	ObjectID int
	Request Request
}