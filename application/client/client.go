package main

import (
	"fmt"
	"github.com/arma29/mid-rpc/application/fibProxy"
	"github.com/arma29/mid-rpc/my-middleware/naming/proxy"
	"github.com/arma29/mid-rpc/shared"
	"time"
)

func pre() {
	namingService := proxy.NamingProxy{}

	fibonacciApp := namingService.Lookup("Fibonacci").(fibProxy.FibonacciProxy)

	fmt.Println("Sample,Time,Result")
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		t1 := time.Now()
		result := fibonacciApp.GetFibOf(5)
		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		fmt.Printf("%d,%f,%d\n", i, x, result)

		if i % 1000 == 0 {
			t, _:= time.ParseDuration("3s")
			time.Sleep(t)
		} 

	}
}

func main() {

	pre()
	// fmt.Scanln()

}