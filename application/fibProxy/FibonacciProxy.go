package fibProxy;

import (
	"github.com/arma29/mid-rpc/my-middleware/distribution/clientProxy"
	"github.com/arma29/mid-rpc/my-middleware/aux"
	"github.com/arma29/mid-rpc/shared"
)

type FibonacciProxy struct {
	Proxy clientProxy.ClientProxy
}

func newFibonacciProxy() FibonacciProxy {
	p := new(FibonacciProxy)

	p.Proxy.Host = "localhost"
	p.Proxy.Port = shared.SERVER_PORT

	return p
}

func (p FibonacciProxy) getFibOf(n int) int {

	param := make([]interface{}, 1)
	param[0] := n

	request := aux.Request{Op:"GetFib", Params: param}
	invoc := aux.Invocation{Host: p.Proxy.Host, Port: p.Proxy.Port, Request: request}

	// Invocando requestor
	req := requestor.Requestor{}
	res := req.Invoker(inv).([]interface{})

	return int(res[0].(int32))
}