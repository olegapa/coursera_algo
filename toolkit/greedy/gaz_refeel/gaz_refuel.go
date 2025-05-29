package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateMinRefills(m, d, stopsLen int, stops []int) int {
	var stopNum, driven int
	for n := 0; driven+m < d && n < stopsLen-1; n++ {
		if stops[n+1] > driven+m {
			driven = stops[n]
			if stops[n+1]-driven > m {
				return -1
			}
			stopNum++
			// fmt.Printf("stop at %d\n", driven)
		}
	}

	if d-driven > m {
		if d-stops[stopsLen-1] < m {
			stopNum++
		} else {
			return -1
		}
	}
	return stopNum
}

func ReadInput() (int, int, int, []int) {
	reader := bufio.NewReader(os.Stdin)
	var m, d, n int

	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d, err = strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	str, err = reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	m, err = strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	str, err = reader.ReadString('\n')
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

	stops := make([]int, n)

	splittedStr := strings.Split(strings.TrimSpace(str), " ")

	for i := 0; i < n; i++ {
		stops[i], err = strconv.Atoi(splittedStr[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return m, d, n, stops
}

func main() {
	m, d, n, stops := ReadInput()

	fmt.Println(CalculateMinRefills(m, d, n, stops))
}
