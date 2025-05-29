package main

import (
	"bufio"
	"fmt"
	"math/rand"
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

func QuickSort(numbers []int) {
	length := len(numbers)
	if length <= 1 {
		return
	}
	// fmt.Println(length)
	pv := rand.Intn(length)
	numbers[pv], numbers[length-1] = numbers[length-1], numbers[pv]
	// fmt.Printf("Sorting %v\n", numbers)
	i := -1
	twin := -1
	for j := 0; j <= length-1; j++ {
		if numbers[j] <= numbers[length-1] {
			i++
			if j > i {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
			if numbers[i] != numbers[length-1] {
				twin++
				if i > twin {
					numbers[i], numbers[twin] = numbers[twin], numbers[i]
				}
			}
		}
	}
	twin++
	// fmt.Printf("Result %v\n", numbers)

	if twin > 0 {
		QuickSort(numbers[0:twin])
	}
	if i < length-1 {
		QuickSort(numbers[i+1:])
	}
}

func main() {
	_, numbers := ReadInput()
	QuickSort(numbers)
	fmt.Println(strings.Trim(fmt.Sprint(numbers), "[]"))
}
