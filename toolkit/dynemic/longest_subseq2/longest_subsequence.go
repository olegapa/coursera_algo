package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() ([]int, []int) {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	n, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		os.Exit(1)
	}

	seq1 := make([]int, n)

	str, err = reader.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	splStr := strings.Split(strings.TrimSpace(str), " ")

	for i := range n {
		seq1[i], err = strconv.Atoi(splStr[i])
		if err != nil {
			os.Exit(1)
		}
	}

	str, err = reader.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	m, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		os.Exit(1)
	}

	seq2 := make([]int, m)

	str, err = reader.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	splStr = strings.Split(strings.TrimSpace(str), " ")

	for i := range m {
		seq2[i], err = strconv.Atoi(splStr[i])
		if err != nil {
			os.Exit(1)
		}
	}

	return seq1, seq2
}

func getMin(num ...int) (int, int) {
	min, k := num[0], 0
	for i, v := range num {
		if v <= min {
			min, k = v, i
		}
	}
	return min, k
}

func fillDistanceMatrix(s1, s2 []int) [][]int {
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
			if s1[i-1] == s2[j-1] {
				matrix[i][j], _ = getMin(matrix[i-1][j]+1, matrix[i][j-1]+1, matrix[i-1][j-1])
			} else {
				matrix[i][j], _ = getMin(matrix[i-1][j]+1, matrix[i][j-1]+1)
			}

		}
	}
	return matrix
}

func countMatches(matrix [][]int) int {
	matches := 0
	for i, j := len(matrix)-1, len(matrix[0])-1; i > 0 && j > 0; {
		var k int
		if matrix[i][j] == matrix[i-1][j-1] {
			_, k = getMin(matrix[i-1][j], matrix[i][j-1], matrix[i-1][j-1])
		} else {
			_, k = getMin(matrix[i-1][j], matrix[i][j-1])
		}
		switch k {
		case 0:
			i--
		case 1:
			j--
		case 2:
			// fmt.Printf("Case 2: prev = %d, curr = %d\n", matrix[i-1][j-1], matrix[i][j])
			matches++
			j--
			i--
		}

	}
	return matches
}

func FindLongestSubseq(s1, s2 []int) int {
	matrix := fillDistanceMatrix(s1, s2)
	// for _, v := range matrix {
	// 	fmt.Println(v)
	// }
	return countMatches(matrix)
}

func main() {
	seq1, seq2 := ReadInput()
	// fmt.Printf("Input seq1 = %v seq2 = %v\n", seq1, seq2)
	fmt.Println(FindLongestSubseq(seq1, seq2))
}
