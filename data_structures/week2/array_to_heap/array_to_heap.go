package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Swap struct {
	Index1 int
	Index2 int
}

func getParentIndex(index int) int {
	if index == 0 {
		return -1 // No parent for the root
	}
	return (index - 1) / 2
}

func getLeftChildIndex(index int) int {
	return 2*index + 1
}

func getRightChildIndex(index int) int {
	return 2*index + 2
}

func siftDown(array []int, start, end int) []Swap {
	swaps := make([]Swap, 0)
	min_index := start
	left := getLeftChildIndex(start)
	if left <= end && array[left] < array[min_index] {
		min_index = left
	}
	right := getRightChildIndex(start)
	if right <= end && array[right] < array[min_index] {
		min_index = right
	}
	if min_index != start {
		swaps = append(swaps, Swap{start, min_index})
		array[start], array[min_index] = array[min_index], array[start]
		swaps = append(swaps, siftDown(array, min_index, end)...)
	}
	return swaps
}

func BuildHeap(array []int) []Swap {
	if len(array) == 0 {
		return []Swap{}
	}
	swaps := []Swap{}
	for i := (len(array) - 2) / 2; i >= 0; i-- {
		swaps = append(swaps, siftDown(array, i, len(array)-1)...)
	}
	return swaps
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nLine, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nLine))
	dataLine, _ := reader.ReadString('\n')
	dataStrs := strings.Fields(dataLine)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i], _ = strconv.Atoi(dataStrs[i])
	}

	swaps := BuildHeap(data)
	fmt.Println(len(swaps))
	for _, swap := range swaps {
		fmt.Println(swap.Index1, swap.Index2)
	}
	// fmt.Println(n)
	// for i := 0; i < n; i++ {
	// 	fmt.Print(data[i], " ")
	// }
}
