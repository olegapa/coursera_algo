package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInp() []int {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Wrong input")
		os.Exit(0)
	}
	var n int
	n, err = strconv.Atoi(strings.TrimSpace(line))
	slice := make([]int, n)

	line, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Wrong input")
		os.Exit(0)
	}

	str_slice := strings.Split(strings.TrimSpace(line), " ")
	for i, v := range str_slice {
		slice[i], err = strconv.Atoi(v)
		if err != nil {
			fmt.Printf("Wrong input in %d position", i)
			os.Exit(0)
		}
	}
	return slice
}

func MaxProd(slice []int) int {
	if len(slice) < 2 {
		fmt.Println("Wrong input: there less then 2 params")
		os.Exit(0)
	}
	var max1, max2 int
	if slice[0] > slice[1] {
		max1, max2 = slice[0], slice[1]
	} else {
		max2, max1 = slice[0], slice[1]
	}

	for i := 2; i < len(slice); i++ {
		elem := slice[i]
		if elem > max2 {
			if elem > max1 {
				max2 = max1
				max1 = elem
			} else {
				max2 = elem
			}
		}
	}

	return max1 * max2
}
func main() {
	slice := ReadInp()

	fmt.Println(MaxProd(slice))
	// slice = make([]int, 200000)
	// for i := 0; i < len(slice); i++ {
	// 	slice[i] = i + 1
	// }
	// fmt.Println(MaxProd(slice))
}
