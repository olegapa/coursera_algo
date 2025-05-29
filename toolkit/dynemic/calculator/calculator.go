package main

import (
	"fmt"
	"strings"
)

type Number struct {
	prev       int
	operations int
}

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

func fillBuff(n int, buf []Number) {
	buf[0] = Number{1, 0}

	for i := 0; i < n-1; i++ {
		if 3*(i+1) <= n {
			if buf[3*(i+1)-1].operations == 0 || buf[3*(i+1)-1].operations > buf[i].operations+1 {
				buf[3*(i+1)-1].operations, buf[3*(i+1)-1].prev = buf[i].operations+1, i+1
			}
		}
		if 2*(i+1) <= n {
			if buf[2*(i+1)-1].operations == 0 || buf[2*(i+1)-1].operations > buf[i].operations+1 {
				buf[2*(i+1)-1].operations, buf[2*(i+1)-1].prev = buf[i].operations+1, i+1
			}
		}
		if buf[i+1].operations == 0 || buf[i+1].operations > buf[i].operations+1 {
			buf[i+1].operations, buf[i+1].prev = buf[i].operations+1, i+1
		}
	}
}

func getSequence(buf []Number) []int {
	n := len(buf)
	res := []int{n}
	for i := buf[n-1].prev; i != 1; i = buf[i-1].prev {
		res = append(res, i)
	}
	if n > 1 {
		res = append(res, 1)
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res

}

func CalculateMinChange(n int) (int, []int) {
	buf := make([]Number, n)
	fillBuff(n, buf)

	return buf[n-1].operations, getSequence(buf)
}

func main() {
	n := ReadInput()
	operations, resSequence := CalculateMinChange(n)
	fmt.Printf("%d\n%s", operations, strings.Trim(fmt.Sprint(resSequence), "[]"))
}
