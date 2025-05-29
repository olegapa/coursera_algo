package main

import "fmt"

func Fibonachi(n, m int) int {
	if n < 2 {
		return n 
	}
	if m == 1 {
		return 0
	}

	cicle := make([]int, 2)

	cicle[0] = 0
	cicle[1] = 1

	for i := 2; i <= n; i++ {
		current := (cicle[i-1] + cicle[i-2]) % m
		if i == n {
			return current
		}
		if cicle[1] == current && cicle[0] == cicle[i-1] {
			break
		}
		cicle = append(cicle, current)
	}

	return cicle[n%(len(cicle)-1)]
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Println(Fibonachi(n, m))
}
