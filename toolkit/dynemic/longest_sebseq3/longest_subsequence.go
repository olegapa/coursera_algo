package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() ([]int, []int, []int) {
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

	for i := 0; i < n; i++ {
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

	for i := 0; i < m; i++ {
		seq2[i], err = strconv.Atoi(splStr[i])
		if err != nil {
			os.Exit(1)
		}
	}

	str, err = reader.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	p, err := strconv.Atoi(strings.TrimSpace(str))
	if err != nil {
		os.Exit(1)
	}

	seq3 := make([]int, p)

	str, err = reader.ReadString('\n')

	if err != nil {
		os.Exit(1)
	}

	splStr = strings.Split(strings.TrimSpace(str), " ")

	for i := 0; i < p; i++ {
		seq3[i], err = strconv.Atoi(splStr[i])
		if err != nil {
			os.Exit(1)
		}
	}

	return seq1, seq2, seq3
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

func fillEdges(matrix [][][]int, xlen, ylen, zlen int, s1, s2, s3 []int) {
	for i := 0; i < xlen; i++ {
		matrix[i][0][0] = i
	}
	for i := 0; i < zlen; i++ {
		matrix[0][i][0] = i
	}
	for i := 0; i < ylen; i++ {
		matrix[0][0][i] = i
	}
	for i := 1; i < zlen; i++ {
		for j := 1; j < ylen; j++ {
			matrix[0][i][j], _ = getMin(matrix[0][i-1][j]+1, matrix[0][i][j-1]+1)
		}
	}
	for i := 1; i < xlen; i++ {
		for j := 1; j < ylen; j++ {
			matrix[i][0][j], _ = getMin(matrix[i-1][0][j]+1, matrix[i][0][j-1]+1)
		}
	}
	for i := 1; i < xlen; i++ {
		for j := 1; j < zlen; j++ {
			matrix[i][j][0], _ = getMin(matrix[i-1][j][0]+1, matrix[i][j-1][0]+1)
		}
	}
}

func fillDistanceMatrix(s1, s2, s3 []int) ([][][]int, int, int, int) {
	xlen, ylen, zlen := len(s1)+1, len(s3)+1, len(s2)+1
	matrix := make([][][]int, xlen)
	for i := 0; i < xlen; i++ {
		matrix[i] = make([][]int, zlen)
		for j := 0; j < zlen; j++ {
			matrix[i][j] = make([]int, ylen)
		}
	}
	fillEdges(matrix, xlen, ylen, zlen, s1, s2, s3)

	for i := 1; i < xlen; i++ {
		for j := 1; j < zlen; j++ {
			for k := 1; k < ylen; k++ {
				if s1[i-1] == s2[j-1] && s1[i-1] == s3[k-1] {
					matrix[i][j][k], _ = getMin(matrix[i-1][j][k]+1, matrix[i][j-1][k]+1, matrix[i][j][k-1]+1, matrix[i-1][j-1][k-1])
				} else {
					matrix[i][j][k], _ = getMin(matrix[i-1][j][k]+1, matrix[i][j-1][k]+1, matrix[i][j][k-1]+1)
				}
			}

		}
	}

	return matrix, xlen, ylen, zlen
}

func countMatches(matrix [][][]int, xlen, ylen, zlen int) int {
	matches := 0
	for i, j, k := xlen-1, zlen-1, ylen-1; i > 0 && j > 0 && k > 0; {
		var idx int
		if matrix[i][j][k] == matrix[i-1][j-1][k-1] {
			_, idx = getMin(matrix[i-1][j][k], matrix[i][j-1][k], matrix[i][j][k-1], matrix[i-1][j-1][k-1])
		} else {
			_, idx = getMin(matrix[i-1][j][k], matrix[i][j-1][k], matrix[i][j][k-1])
		}
		switch idx {
		case 0:
			i--
		case 1:
			j--
		case 2:
			k--
		case 3:
			// fmt.Printf("Case 2: prev = %d, curr = %d\n", matrix[i-1][j-1], matrix[i][j])
			matches++
			j--
			i--
			k--
		}

	}
	return matches
}

func FindLongestSubseq(s1, s2, s3 []int) int {
	matrix, xlen, ylen, zlen := fillDistanceMatrix(s1, s2, s3)
	// for _, v := range matrix {
	// 	fmt.Println(v)
	// }
	return countMatches(matrix, xlen, ylen, zlen)
}

func main() {
	seq1, seq2, seq3 := ReadInput()
	// fmt.Printf("Input seq1 = %v seq2 = %v\n", seq1, seq2)
	fmt.Println(FindLongestSubseq(seq1, seq2, seq3))
}
