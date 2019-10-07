package application

func CalcFibonacci(n int32) int32 {
	if n < 2 {
		return n
	}
	return CalcFibonacci(n-1) + CalcFibonacci(n-2)
}

