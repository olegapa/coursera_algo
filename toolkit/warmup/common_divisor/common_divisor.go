package main

import "fmt"

func FindDivisor(bigger, smaller int) int {
	if bigger%smaller == 0 {
		return smaller
	}
	if bigger-smaller > smaller {
		return FindDivisor(bigger-smaller, smaller)
	} else {
		return FindDivisor(smaller, bigger-smaller)
	}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	fmt.Println(FindDivisor(n, m))
}
