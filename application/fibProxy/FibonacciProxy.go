package fibProxy;

import (
	clientProxy "github.com/arma29/mid-rpc/my-middleware/distribution/clientProxy"
	requestor "github.com/arma29/mid-rpc/my-middleware/distribution/requestor"
	"github.com/arma29/mid-rpc/my-middleware/aux"
	"github.com/arma29/mid-rpc/shared"
)

type FibonacciProxy struct {
	Proxy clientProxy.ClientProxy
}

func NewFibonacciProxy() FibonacciProxy {
	p := new(FibonacciProxy)

	p.Proxy.Host = "localhost"
	p.Proxy.Port = shared.SERVER_PORT

	return *p
}

func (p FibonacciProxy) GetFibOf(n int) int {

	param := make([]interface{}, 1)
	param[0] = n

	request := aux.Request{Op:"GetFibo", Params: param}
	invoc := aux.Invocation{Host: p.Proxy.Host, Port: p.Proxy.Port, Request: request}

	// Invocando requestor
	req := requestor.Requestor{}
	res := req.Invoke(invoc).([]interface{})
	
	return int(res[0].(float64))
}