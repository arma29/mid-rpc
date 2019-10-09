package invoker 

import (
	"reflect"
	app "github.com/arma29/mid-rpc/application"

)

type LCM struct{}

func (lcm LCM) getRemoteObjectByID(id int) interface{} {
	return app.GetFibonacciApp()
}

func (lcm LCM) getPool(proxyType reflect.Type) []interface{} {

	var poolLength = 5
	pool := make([]interface{}, poolLength)
	
	for i := 0; i < poolLength; i++ {
	}

	return pool
}