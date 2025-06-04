package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() (int, int, []int) {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(-1)
	}

	spltdStr := strings.Split(strings.TrimSpace(str), " ")
	w, err := strconv.Atoi(spltdStr[0])
	if err != nil {
		os.Exit(-1)
	}

	n, err := strconv.Atoi(spltdStr[1])
	if err != nil {
		os.Exit(-1)
	}

	str, err = reader.ReadString('\n')
	if err != nil {
		os.Exit(-1)
	}

	spltdStr = strings.Split(strings.TrimSpace(str), " ")
	bars := make([]int, n)

	for i := 0; i < n; i++ {
		bars[i], err = strconv.Atoi(spltdStr[i])
		if err != nil {
			os.Exit(-1)
		}
	}
	return w, n, bars
}

func createTable(w, n int, bars []int) [][]int {
	table := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		table[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-bars[i-1] >= 0 {
				table[i][j] = max(table[i-1][j], table[i-1][j-bars[i-1]]+bars[i-1])
			} else {
				table[i][j] = table[i-1][j]
			}
		}
	}

	return table
}

// Добавить функцию max для совместимости со старыми версиями Go
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FindMaxGoldAmount(w, n int, bars []int) int {
	table := createTable(w, n, bars)
	// for i := range n + 1 {
	// 	fmt.Println(table[i])
	// }
	return table[n][w]
}

func main() {
	w, n, bars := ReadInput()
	fmt.Println(FindMaxGoldAmount(w, n, bars))
}
