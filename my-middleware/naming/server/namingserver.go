package main

import (
	"fmt"
	"github.com/arma29/mid-rpc/my-middleware/naming/invoker"
)

func main() {

	fmt.Println("Naming server running!!")

	// control loop passed to invoker
	namingInvoker := invoker.NamingInvoker{}
	namingInvoker.Invoke()
}
