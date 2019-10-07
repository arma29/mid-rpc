package main

import (
	"fmt"
	"github.com/arma29/mid-rpc/my-middleware/distribution/invoker"
	"github.com/arma29/mid-rpc/application/fibProxy"
	"github.com/arma29/mid-rpc/my-middleware/naming/proxy"
)

func main() {

	namingProxy := proxy.NamingProxy{}

	fibonacciApp := fibProxy.NewFibonacciProxy()

	namingProxy.Register("Fibonacci", fibonacciApp)

	fmt.Println("Server listening")
	fibonacciInvoker := invoker.NewFibonnaciInvoker()
	fibonacciInvoker.Invoke()

	fmt.Scanln()
}