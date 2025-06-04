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

func binSearch(searchSlice []int, numbers []int, indexies []int, res []int, offset int) {
	if len(searchSlice) == 0 || len(indexies) == 0 {
		return
	}
	mid := (len(searchSlice) - 1) / 2
	left_idx, right_idx := make([]int, 0), make([]int, 0)

	for _, idx := range indexies {
		if searchSlice[mid] == numbers[idx] {
			res[idx] = mid + offset
		} else if len(searchSlice) == 1 {
			res[idx] = -1
		} else if numbers[idx] < searchSlice[mid] {
			left_idx = append(left_idx, idx)
		} else {
			right_idx = append(right_idx, idx)
		}
	}
	if len(left_idx) > 0 && mid > 0 {
		binSearch(searchSlice[:mid], numbers, left_idx, res, offset)
	}
	if len(right_idx) > 0 && mid+1 < len(searchSlice) {
		binSearch(searchSlice[mid+1:], numbers, right_idx, res, offset+mid+1)
	}
	// Исправление: если элемент не найден, res[idx] должен быть -1 (а не 0)
	// Это уже делается выше, но есть случай, когда число встречается несколько раз в numbers.
	// Поэтому после binSearch можно пройтись по res и если res[i] == 0 и numbers[i] != searchSlice[0], то res[i] = -1
}

func PerformBinSearch(searchSlice []int, numbers []int) []int {
	indexes := make([]int, len(numbers))
	res := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		indexes[i] = i
	}
	binSearch(searchSlice, numbers, indexes, res, 0)
	// Исправление: если res[i] == 0 и numbers[i] != searchSlice[0], то res[i] = -1
	for i := 0; i < len(res); i++ {
		if res[i] == 0 && (len(searchSlice) == 0 || numbers[i] != searchSlice[0]) {
			res[i] = -1
		}
	}
	return res
}

func main() {
	_, array := ReadInput()

	// Читаем строку с количеством чисел для поиска
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Читаем строку с числами для поиска
	str, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	splittedStr := strings.Split(strings.TrimSpace(str), " ")
	numbers := make([]int, len(splittedStr))
	for i := 0; i < len(splittedStr); i++ {
		numbers[i], err = strconv.Atoi(splittedStr[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(PerformBinSearch(array, numbers)), "[]"))
}
