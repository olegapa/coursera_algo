package main

import "fmt"

func FindDivisor(bigger, smaller uint64) uint64 {
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
	var n, m uint64
	fmt.Scan(&n, &m)

	fmt.Println(m*(n/FindDivisor(n, m)))
}
