package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() ([]int, int) {
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

	return seq1, n
}

func getPartitionMatrix(set []int, n, sum int) [][][]bool {
	matrix := make([][][]bool, n+1)

	for i := range matrix {
		matrix[i] = make([][]bool, sum+1)
		for j := range matrix[i] {
			matrix[i][j] = make([]bool, sum+1)
		}
	}
	matrix[0][0][0] = true

	for k := 1; k <= n; k++ {
		for i := range matrix[k] {
			for j := range matrix[k][i] {
				iOffset, jOffset := false, false
				if i-set[k-1] >= 0 {
					iOffset = matrix[k-1][i-set[k-1]][j]
				} else {
					iOffset = false
				}

				if j-set[k-1] >= 0 {
					jOffset = matrix[k-1][i][j-set[k-1]]
				} else {
					jOffset = false
				}
				matrix[k][i][j] = (matrix[k-1][i][j] || iOffset || jOffset)
			}
		}
	}

	return matrix
}

func IsSplittable(set []int, n int) int {
	sum := 0
	for _, v := range set {
		sum += v
	}

	if sum%3 != 0 {
		return 0
	}

	expectedSum := sum / 3

	matrix := getPartitionMatrix(set, n, expectedSum)
	if matrix[n][expectedSum][expectedSum] {
		return 1
	}
	return 0
}

func main() {
	set, n := ReadInput()

	fmt.Println(IsSplittable(set, n))
}
