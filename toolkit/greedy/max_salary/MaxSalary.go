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

func QuickSort(numbers []int, isLarger func(i, j int) bool) {
	pv := len(numbers) - 1
	if pv == 0 {
		return
	}
	i := -1
	for j := 0; j <= pv; j++ {
		if isLarger(numbers[j], numbers[pv]) {
			i++
			if j > i {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}

	if i > 0 {
		QuickSort(numbers[0:i], isLarger)
	}
	if i < pv {
		QuickSort(numbers[i+1:], isLarger)
	}
}

func FindMaxSalary(numbers []int, sortF func([]int, func(int, int) bool)) string {
	sortF(numbers, func(i, j int) bool {
		si := strconv.Itoa(i)
		sj := strconv.Itoa(j)
		return si+sj >= sj+si
	})
	// fmt.Println(numbers)
	var res string

	for _, n := range numbers {
		// fmt.Printf("i = %d, conv i = %s\n", i, strconv.Itoa(i))
		res = res + strconv.Itoa(n)
	}
	return res
}

func main() {
	_, numbers := ReadInput()

	fmt.Println(FindMaxSalary(numbers, QuickSort))
}
