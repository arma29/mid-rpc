package main;

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"strconv"
	"fmt"
	"time"
	"os"
	
	"github.com/arma29/mid-rpc/gRPC/fibonacci"
	"github.com/arma29/mid-rpc/shared"
)

func main() {
	// Get Argument from command Line
	if len(os.Args) != 3 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	ipContainer := os.Args[1]
	var i int32

	conn, err := grpc.Dial(ipContainer + ":" + 
		strconv.Itoa(shared.GRPC_PORT), grpc.WithInsecure())
	shared.CheckError(err)

	defer conn.Close()

	fib := fibonacci.NewFibonacciClient(conn)

	// Contacta o servidor
	ctx, cancel := context.WithTimeout(context.Background(), 
		time.Minute) // havia um problema com o time.Second . 1s -> 1m
	defer cancel()

	number, _ := strconv.Atoi(os.Args[2])

	fmt.Println("Fibonacci,Answer,Time")
	for i = 0; i < shared.SAMPLE_SIZE; i++ {
		t1 := time.Now()

		// Invoca operação remota
		msgReply, err := fib.GetFibo(ctx, &fibonacci.FibRequest{ Number: int32(number)})
		shared.CheckError(err)

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000

		s := fmt.Sprintf("%d,%d,%f", number, msgReply.Number, x) 
		fmt.Println(s)
	}
}
