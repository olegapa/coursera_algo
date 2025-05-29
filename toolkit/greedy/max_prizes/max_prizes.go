package main

import (
	"fmt"
	// "strings"
)

func CuontMaxPizes(n int) (int, []int) {
	i, sum := 0, 0
	res := make([]int, 0)
	for n > sum {
		if n < sum+2*i+3 {
			res = append(res, n-sum)
			return len(res), res
		}
		i += 1
		sum += i
		res = append(res, i)
	}
	return len(res), res
}

func main() {
	var n int
	fmt.Scan(&n)

	k, _ := CuontMaxPizes(n)
	fmt.Println(k)
	// fmt.Println(strings.Trim(fmt.Sprint(prizes), "[]"))
}
