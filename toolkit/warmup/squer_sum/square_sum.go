package main

import "fmt"

func Fibonachi(n int) int {
	if n < 2 {
		return n
	}

	cicle := make([]int, 2)

	cicle[0] = 0
	cicle[1] = 1

	for i := 2; i <= n; i++ {
		current := (cicle[i-1] + cicle[i-2]) % 10
		if i == n {
			return current * (current + cicle[i-1]) % 10
		}
		if cicle[1] == current && cicle[0] == cicle[i-1] {
			break
		}
		cicle = append(cicle, current)
	}

	return cicle[n%(len(cicle)-1)] * (cicle[n%(len(cicle)-1)] + cicle[(n-1)%(len(cicle)-1)]) % 10
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(Fibonachi(n))
}
