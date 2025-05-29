package main

import "fmt"

func CalculateInterval(first, last int, cicle []int) int {
	sum := cicle[first]
	if first <= last {
		for i := first + 1; i <= last; i++ {
			sum = (sum + cicle[i]) % 10
		}
	} else {
		for i := 0; i < first; i++ {
			sum = (sum + cicle[i]) % 10
		}
		for i := last; i < len(cicle)-1; i++ {
			sum = (sum + cicle[i]) % 10
		}
	}
	return sum
}

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
	square := 1
	for i := 2; i <= n; i++ {
		current := (cicle[i-1] + cicle[i-2]) % 10
		if cicle[1] == current && cicle[0] == cicle[i-1] {
			break
		}
		cicle = append(cicle, current)
		if i == n {
			return (CalculateInterval(m, n, cicle)) % 10
		}
		square = (square + current) % 10
	}
	element_amount := n - m + 1
	loops := (element_amount - element_amount%(len(cicle)-1)) / (len((cicle)) - 1)
	first, last := m%(len(cicle)-1), n%(len(cicle)-1)
	sum := CalculateInterval(first, last, cicle)

	return (sum + loops*square) % 10
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println(Fibonachi(n, 0))
}
