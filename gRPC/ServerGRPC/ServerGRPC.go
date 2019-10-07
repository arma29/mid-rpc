package main

import (
	"fmt"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"strconv"

	"github.com/arma29/mid-rpc/gRPC/fibonacci"
	"github.com/arma29/mid-rpc/application"
	"github.com/arma29/mid-rpc/shared"
)

type fibonacciServer struct{}

func (s *fibonacciServer) GetFibo(ctx context.Context, req *fibonacci.FibRequest) (*fibonacci.FibResponse, error) {
	return &fibonacci.FibResponse{ Number: application.CalcFibonacci(req.Number) }, nil
}

func main() {
	conn, err := net.Listen("tcp", ":"+strconv.Itoa(shared.GRPC_PORT))
	shared.CheckError(err)

	servidor := grpc.NewServer()
	fibonacci.RegisterFibonacciServer(servidor, &fibonacciServer{})

	fmt.Println("Servidor pronto ...")

	// Register reflection service on gRPC servidor.
	reflection.Register(servidor)

	err = servidor.Serve(conn);
	shared.CheckError(err)
}