package main

import (
	"fmt"
	"github.com/arma29/mid-rpc/my-middleware/distribution/invoker"
	"github.com/arma29/mid-rpc/my-middleware/distribution/lcm"
)

func main() {

	lcmInstance := lcm.LCM{}
	lcmInstance.RegisterFibonacci()

	fmt.Println("Server listening")
	fibonacciInvoker := invoker.NewFibonnaciInvoker()
	fibonacciInvoker.Invoke()

	fmt.Scanln()
}