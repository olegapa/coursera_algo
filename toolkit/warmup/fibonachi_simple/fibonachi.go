package main

import "fmt"

func Fibonachi(n int) uint64 {
	if n < 2 {
		return uint64(n)
	}

	fib_values := make([]uint64, n)
	fib_values[0] = 0
	fib_values[1] = 1

	for i := 2; i < n; i++ {
		fib_values[i] = fib_values[i-1] + fib_values[i-2]
	}

	return fib_values[n-1] + fib_values[n-2]
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(Fibonachi(n))
}
