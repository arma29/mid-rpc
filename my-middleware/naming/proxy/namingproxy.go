package proxy

import (
	clientProxy "github.com/arma29/mid-rpc/my-middleware/distribution/clientProxy"
	requestor "github.com/arma29/mid-rpc/my-middleware/distribution/requestor"
	"github.com/arma29/mid-rpc/my-middleware/aux"
	"github.com/arma29/mid-rpc/shared"
	"github.com/arma29/mid-rpc/application/fibProxy"
	//"fmt"
)

type NamingProxy struct{}

func CheckRepository(proxy clientProxy.ClientProxy) interface{}{
	var clientProxy interface{}

	fiboProxy := fibProxy.NewFibonacciProxy()
	fiboProxy.Proxy.Host = proxy.Host
	fiboProxy.Proxy.Port = proxy.Port
	clientProxy = fiboProxy

	return clientProxy
}

func (NamingProxy) Register(p1 string, proxy interface{}) bool {

	//prepare
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = proxy
	namingproxy := clientProxy.ClientProxy{Host:"", Port:shared.NS_PORT}
	request := aux.Request{Op: "Register", Params: params}
	inv := aux.Invocation{Host: namingproxy.Host, Port: namingproxy.Port,Request: request}

	// requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})
	
	return ter[0].(bool)
}

func (NamingProxy) Lookup(p1 string) interface{} {
	// prepare invocation
	params := make([]interface{}, 1)
	params[0] = p1
	namingproxy := clientProxy.ClientProxy{Host:"",Port:shared.NS_PORT}
	request := aux.Request{Op: "Lookup", Params: params}
	inv := aux.Invocation{Host:namingproxy.Host,Port:namingproxy.Port,Request:request}

	// invoke requestor
	req := requestor.Requestor{}
	ter := req.Invoke(inv).([]interface{})

	// process reply
	proxyTemp := ter[0].(map[string]interface{})
	clientProxyTemp := clientProxy.ClientProxy{Host:proxyTemp["Host"].(string),Port:int(proxyTemp["Port"].(float64))}
	clientProxy := CheckRepository(clientProxyTemp)

	return clientProxy
}
