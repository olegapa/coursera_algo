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

func mergeSortAndCount(arr, temp []int, left, right int) int {
	if right-left <= 1 {
		return 0
	}
	mid := (left + right) / 2
	inv := mergeSortAndCount(arr, temp, left, mid)
	inv += mergeSortAndCount(arr, temp, mid, right)

	i, j, k := left, mid, left
	for i < mid && j < right {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			inv += mid - i
			j++
		}
		k++
	}
	for i < mid {
		temp[k] = arr[i]
		i++
		k++
	}
	for j < right {
		temp[k] = arr[j]
		j++
		k++
	}
	for l := left; l < right; l++ {
		arr[l] = temp[l]
	}
	return inv
}

func FindInversions(numbers []int) int {
	n := len(numbers)
	temp := make([]int, n)
	return mergeSortAndCount(numbers, temp, 0, n)
}

func main() {
	_, numbers := ReadInput()

	fmt.Println(FindInversions(numbers))
}
