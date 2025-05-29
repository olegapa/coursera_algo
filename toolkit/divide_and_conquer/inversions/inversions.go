package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() (int, []int) {
	reader := bufio.NewReader(os.Stdin)
	var n int

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

	str, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numbers := make([]int, n)

	splittedStr := strings.Split(strings.TrimSpace(str), " ")

	for i := 0; i < n; i++ {
		numbers[i], err = strconv.Atoi(splittedStr[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return n, numbers
}

func FindInversions(numbers []int) int {
	if len(numbers) <= 1 {
		return 0
	}
	mid, inv := len(numbers)/2, 0
	inv += FindInversions(numbers[:mid])
	inv += FindInversions(numbers[mid:])
	tempSlice := make([]int, len(numbers))
	copy(tempSlice, numbers)
	// fmt.Printf("Input slice: %v, inv = %d\n", numbers, inv)
	for k, i, j := 0, 0, mid; k < len(numbers); k++ {
		if i == mid {
			for j < len(numbers) {
				numbers[k] = tempSlice[j]
				j++
			}
			break
		} else if j == len(numbers) {
			for i < mid {
				numbers[k] = tempSlice[i]
				i++
			}
			break
		} else if tempSlice[i] <= tempSlice[j] {
			numbers[k] = tempSlice[i]
			i++
		} else {
			numbers[k] = tempSlice[j]
			j++
			inv += mid - i
		}
	}
	// fmt.Printf("Res slice: %v, inv = %d\n", numbers, inv)

	return inv
}

func main() {
	_, numbers := ReadInput()

	fmt.Println(FindInversions(numbers))
}
