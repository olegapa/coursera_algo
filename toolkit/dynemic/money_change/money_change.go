package main

import (
	"fmt"
)

func ReadInput() int {
	var n int
	fmt.Scan(&n)
	return n
}

func getMin(num ...int) int {
	min := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] < min {
			min = num[i]
		}
	}
	return min
}

func fillBuff(n int, buf []int) {
	if n <= 4 {
		if n == 2 {
			buf[1] = 2
		} else {
			buf[n-1] = 1
		}
		return
	}
	buf[0], buf[1], buf[2], buf[3] = 1, 2, 1, 1

	for i := 4; i < n; i++ {
		buf[i] = getMin(buf[i-4], buf[i-3], buf[i-1]) + 1
	}
}

func CalculateMinChange(n int) int {
	buf := make([]int, n)
	fillBuff(n, buf)
	return buf[n-1]
}

func main() {
	n := ReadInput()
	fmt.Println(CalculateMinChange(n))
}
