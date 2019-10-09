package application

var remoteFibID = 0

// For my-middleware
type FibonacciApp struct {
	ObjectID int
}

func GetFibonacciApp() *FibonacciApp {
	app := new(FibonacciApp)
	app.ObjectID = remoteFibID

	remoteFibID += 1
	
	return app
}

func (f FibonacciApp) GetFibOf(n int32) int32 {
	if n < 2 {
		return n
	}
	return f.GetFibOf(n-1) + f.GetFibOf(n-2)
}

// For gRPC
func CalcFibonacci(n int32) int32 {
	if n < 2 {
		return n
	}
	return CalcFibonacci(n-1) + CalcFibonacci(n-2)
}

