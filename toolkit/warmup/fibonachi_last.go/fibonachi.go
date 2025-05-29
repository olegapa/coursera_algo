package main

import "fmt"

func Fibonachi(n int) int {
	if n < 2 {
		return int(n)
	}

	fib_values := make([]int, n)
	fib_values[0] = 0
	fib_values[1] = 1

	for i := 2; i < n; i++ {
		fib_values[i] = fib_values[i-1]%10 + fib_values[i-2]%10
	}

	return (fib_values[n-1] + fib_values[n-2]) % 10
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(Fibonachi(n))
}
