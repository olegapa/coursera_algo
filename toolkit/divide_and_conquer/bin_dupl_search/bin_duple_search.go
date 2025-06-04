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

func PerformBinSearch(searchSlice []int, numbers []int) []int {
	indexes := make([]int, len(numbers))
	res := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		indexes[i] = i
	}
	binSearch(searchSlice, numbers, indexes, res, 0)
	return res
}

func findEdges(searchSlice []int, idx int) (int, int) {
	left, right, last_idx := idx, idx, len(searchSlice)-1
	if idx > 0 {
		for searchSlice[idx] == searchSlice[left-1] {
			left--
			if left-1 < 0 {
				break
			}
		}
	}
	if idx < last_idx {
		for searchSlice[idx] == searchSlice[right+1] {
			right++
			if right+1 >= last_idx {
				break
			}
		}
	}
	return left, right
}

func binSearch(searchSlice []int, numbers []int, indexies []int, res []int, offset int) {
	mid := (len(searchSlice) - 1) / 2
	left_idx, right_idx := make([]int, 0), make([]int, 0)

	leftEdge, rightEdge := findEdges(searchSlice, mid)

	// fmt.Printf("search slice = %v, indexes = %v, mid = %d. left_idx = %d, right_idx = %d\n", searchSlice, indexies, mid, leftEdge, rightEdge)
	for _, idx := range indexies {
		// fmt.Printf("%d %d\n", mid, idx)
		if searchSlice[mid] == numbers[idx] {
			res[idx] = leftEdge + offset
		} else if len(searchSlice) == 1{
			res[idx] = -1
		} else if numbers[idx] < searchSlice[mid] {
			left_idx = append(left_idx, idx)
		} else {
			right_idx = append(right_idx, idx)
		}
	}

	if len(left_idx) > 0 {
		if leftEdge == 0 {
			for _, idx := range left_idx {
				res[idx] = -1
			}
		} else {
			binSearch(searchSlice[:leftEdge], numbers, left_idx, res, offset)
		}
	}

	if len(right_idx) > 0 {
		if rightEdge == len(searchSlice)-1 {
			for _, idx := range right_idx {
				res[idx] = -1
			}
		} else {
			binSearch(searchSlice[rightEdge+1:], numbers, right_idx, res, offset+rightEdge+1)
		}
	}
}

func main() {
	_, array := ReadInput()

	_, numbers := ReadInput()

	fmt.Println(strings.Trim(fmt.Sprint(PerformBinSearch(array, numbers)), "[]"))
}
