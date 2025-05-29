package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() (int, []int, []int) {
	var n int
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	n, err = strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	clics, prices := make([]int, n), make([]int, n)

	str, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	spltdStr := strings.Split(strings.TrimSpace(str), " ")
	for i := 0; i < n; i++ {
		prices[i], err = strconv.Atoi(spltdStr[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	str, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	spltdStr = strings.Split(strings.TrimSpace(str), " ")
	for i := 0; i < n; i++ {
		clics[i], err = strconv.Atoi(spltdStr[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return n, prices, clics
}

func SortSlice(slice []int) {
	if len(slice) <= 1 {
		return
	}

	pvt := len(slice) - 1
	i := -1
	for j := 0; j <= pvt; j++ {
		if slice[j] <= slice[pvt] {
			i++
			if j > i {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	// fmt.Printf("res = %v, i = %d, pvt = %d", slice, i, pvt)

	if i > 0 {
		SortSlice(slice[:i])
	}
	if i < pvt {
		SortSlice(slice[i+1:])
	}
}

func FindMaxRev(n int, prices, clics []int) int {
	SortSlice(prices)
	// fmt.Printf("Sorted prices: %v\n", prices)
	SortSlice(clics)
	// fmt.Printf("Sorted clics: %v\n", clics)

	product := 0
	for i := 0; i < n; i++ {
		product += clics[i] * prices[i]
	}

	return product
}

func main() {
	n, prices, clics := ReadInput()

	fmt.Println(FindMaxRev(n, prices, clics))
}
