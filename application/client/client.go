package main

import (
	"fmt"
	"github.com/arma29/mid-rpc/application/fibProxy"
)


func main() {

	// namingService := proxy.NamingProxy{}

	// namingService.Lookup("Fibonacci").(proxy.FibonacciProxy)

	fibonacciApp := fibProxy.NewFibonacciProxy()

	result := fibonacciApp.GetFibOf(5)

	fmt.Print(result)

	fmt.Scanln()

}