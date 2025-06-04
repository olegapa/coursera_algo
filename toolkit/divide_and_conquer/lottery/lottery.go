package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// type Line struct {
// 	l, r int
// }

func ReadInput() (int, int, []int, []int, []int) {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Wrong input in 1 line")
		os.Exit(1)
	}

	splitedInput := strings.Split(strings.TrimSpace(str), " ")
	var n, m int

	n, err = strconv.Atoi(splitedInput[0])
	if err != nil {
		fmt.Println("n param is incorrect")
		os.Exit(1)
	}

	m, err = strconv.Atoi(splitedInput[1])
	if err != nil {
		fmt.Println("m param is incorrect")
		os.Exit(1)
	}

	lines_left := make([]int, n)
	lines_right := make([]int, n)
	for i := 0; i < n; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Wrong input for %d line\n", i)
			os.Exit(1)
		}
		splitedInput := strings.Split(strings.TrimSpace(str), " ")
		lines_left[i], err = strconv.Atoi(splitedInput[0])
		if err != nil {
			fmt.Printf("l incorrect for %d", i)
			os.Exit(1)
		}
		lines_right[i], err = strconv.Atoi(splitedInput[1])
		if err != nil {
			fmt.Printf("r is incorrect for %d", i)
			os.Exit(1)
		}
	}
	points := make([]int, m)
	str, err = reader.ReadString('\n')

	if err != nil {
		fmt.Println("Wrong input in points line")
		os.Exit(1)
	}

	splitedInput = strings.Split(strings.TrimSpace(str), " ")

	for i := 0; i < m; i++ {
		points[i], err = strconv.Atoi(splitedInput[i])
		if err != nil {
			fmt.Printf("Wrong point input idx = %d", i)
			os.Exit(1)
		}
	}
	return n, m, lines_left, lines_right, points
}

func sortSlice(numbers []int) {
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
		sortSlice(numbers[0:twin])
	}
	if i < length-1 {
		sortSlice(numbers[i+1:])
	}
}

func findResultPosition(value, idx int, numbers []int, isLeft bool) int {
	fmt.Printf("Find pos for %d from idx = %d, in %v . isLeft %v\n", value, idx, numbers, isLeft)

	if idx > len(numbers)-1 {
		fmt.Println(len(numbers) - 1)
		return len(numbers) - 1
	}

	if isLeft {
		v := numbers[idx]
		if value >= v {
			for idx < len(numbers)-1 {
				if numbers[idx+1] != v {
					break
				}
				idx++
			}
		} else {
			for idx > 0 {
				if numbers[idx-1] != v {
					break
				}
				idx--
			}
			idx--
		}

	} else {
		if idx < 0 {
			if value <= numbers[0] {
				fmt.Println(-1)
				return -1
			}
			fmt.Println(0)
			return 0
		}
		v := numbers[idx]

		if value > v {
			for idx < len(numbers)-1 {
				if numbers[idx+1] != v {
					break
				}
				idx++
			}
		} else {
			for idx > 0 {
				if numbers[idx-1] != v {
					break
				}
				idx--
			}
			idx--
		}
	}
	fmt.Println(idx)
	return idx
}

func binSearch(v int, numbers []int, isLeft bool) int {
	l, r, mid := 0, len(numbers)-1, 0
	for l < r {
		mid = (l + r + 1) / 2
		// fmt.Printf("%d %d %d\n", l, mid, r)
		if v == numbers[mid] {
			return findResultPosition(v, mid, numbers, isLeft)
		} else if v < numbers[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return findResultPosition(v, l, numbers, isLeft)
}

func CalculateEanings(left []int, right []int, points []int) []int {
	fmt.Printf("Input arrays \n%v,\n %v\n", left, right)
	sortSlice(left)
	sortSlice(right)
	fmt.Printf("Sorted arrays \n%v,\n %v\n", left, right)

	res := make([]int, len(points))
	for i, v := range points {
		// fmt.Printf("point = %d, left pos = %d, right pos = %d\n", v, binSearch(v, left, true), binSearch(v, right, false))
		res[i] = binSearch(v, left, true) - binSearch(v, right, false)
	}
	return res
}

func main() {
	_, _, left, right, points := ReadInput()

	coord := CalculateEanings(left, right, points)
	fmt.Println(len(coord))
	fmt.Println(strings.Trim(fmt.Sprint(coord), "[]"))
}
