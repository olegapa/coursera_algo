package main

import "fmt"

func GetChange(value int) int {
	tens := value / 10
	rest := value - tens*10

	fifes := rest / 5
	rest -= fifes * 5

	return tens + fifes + rest
}

func main() {
	var inp int

	fmt.Scan(&inp)

	fmt.Println(GetChange(inp))
}
