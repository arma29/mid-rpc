package lcm 

import (
	app "github.com/arma29/mid-rpc/application"
	"github.com/arma29/mid-rpc/my-middleware/naming/proxy"
	"github.com/arma29/mid-rpc/application/fibProxy"
	// clientProxy "github.com/arma29/mid-rpc/my-middleware/distribution/clientProxy"
	// "github.com/arma29/mid-rpc/app/fibProxy"
)

var poolGlobal []interface{}

var lcmMAP = make(map[app.FibonacciApp]string)

type LCM struct{}

type ProxyMaker func() interface{}

func PutObjectState(object app.FibonacciApp, state string){
	lcmMAP[object] = state
}

func IsObjectStateCreated(object app.FibonacciApp) (bool){
	boolean := false
	_, ok := lcmMAP[object]
	if ok {
		boolean = false
	} else {
		boolean = true
	}
	return boolean
}

func (lcm LCM) GetRemoteObjectByID(id int) interface{} {
	for i := 0; i < len(poolGlobal); i++ {
		fibApp := poolGlobal[i].(*app.FibonacciApp)
		if fibApp.ObjectID == id{
			return poolGlobal[i]
		}
	}

	return nil
}

func (lcm LCM) GetPool() []interface{} {
	var poolLength = 5
	pool := make([]interface{}, poolLength)
	
	for i := 0; i < poolLength; i++ {
		pool[i] = app.GetFibonacciApp()
	}

	poolGlobal = pool
	return pool
}


func (lcm LCM) RegisterFibonacci() {
	namingProxy := proxy.NamingProxy{}

	pool := lcm.GetPool()

	
	for i := 1; i < len(pool); i++ {
		objectID := pool[i].(*app.FibonacciApp).ObjectID
		fibonacciProxy := fibProxy.NewFibonacciProxy(objectID)
		namingProxy.Register("Fibonacci", fibonacciProxy)
		fiboStruct := app.FibonacciApp{}
		fiboStruct.ObjectID = objectID
		PutObjectState(fiboStruct, "Created")
	}
		
}