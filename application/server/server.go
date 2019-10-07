package main

import (
	"fmt"
	"github.com/arma29/mid-rpc/my-middleware/distribution/invoker"
)

func main() {

	fmt.Println("Server listening")
	fibonacciInvoker := invoker.NewFibonnaciInvoker()

	go fibonacciInvoker.Invoke()

	fmt.Scanln()
}