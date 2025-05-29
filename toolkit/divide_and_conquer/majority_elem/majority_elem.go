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

func IsMajority(numbers []int, n int) bool {
	res, _, _ := getMajority(numbers, n)
	return res
}

func countNumber(numbers []int, n int) int {
	res := 0
	for _, v := range numbers {
		if v == n {
			res++
		}
	}
	return res
}

func getMajority(numbers []int, n int) (bool, int, int) {
	if n == 1 {
		return true, numbers[0], 1
	}

	mid := n / 2

	hasLeft, nLeft, countLeft := getMajority(numbers[0:mid], mid)
	hasRight, nRight, countRight := getMajority(numbers[mid:n], n-mid)

	if hasLeft {
		countLeft += countNumber(numbers[mid:n], nLeft)
		if countLeft > n/2 {
			return true, nLeft, countLeft
		}
	}
	if hasRight {
		countRight += countNumber(numbers[0:mid], nRight)
		if countRight > n/2 {
			return true, nRight, countRight
		}
	}

	return false, 0, 0
}

func main() {
	n, numbers := ReadInput()

	fmt.Println(IsMajority(numbers, n))
}
