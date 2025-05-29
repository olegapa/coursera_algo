package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	str1, err := reader.ReadString('\n')
	str1 = strings.TrimSpace(str1)
	if err != nil {
		os.Exit(-1)
	}
	var str2 string
	str2, err = reader.ReadString('\n')
	str2 = strings.TrimSpace(str2)
	return str1, str2
}

func getMin(num ...int) int {
	min := num[0]
	for _, v := range num {
		if v < min {
			min = v
		}
	}
	return min
}

func fillDistanceMatrix(s1, s2 string) [][]int {
	matrix := make([][]int, len(s1)+1)
	for i := range len(s1) + 1 {
		matrix[i] = make([]int, len(s2)+1)
	}

	for i := 0; i < len(s1)+1; i++ {
		matrix[i][0] = i
	}

	for i := 0; i < len(s2)+1; i++ {
		matrix[0][i] = i
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			substPoint := 1
			if s1[i-1] == s2[j-1] {
				substPoint = 0
			}

			matrix[i][j] = getMin(matrix[i-1][j]+1, matrix[i][j-1]+1, matrix[i-1][j-1]+substPoint)
		}
	}
	return matrix
}

func FindEditDistance(s1, s2 string) int {
	matrix := fillDistanceMatrix(s1, s2)
	return matrix[len(s1)][len(s2)]
}

func main() {
	s1, s2 := ReadInput()

	fmt.Println(FindEditDistance(s1, s2))
}
