package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type operation byte

func (op operation) Apply(left, right int) int {
	if op == '+' {
		return left + right
	} else if op == '-' {
		return left - right
	} else {
		return left * right
	}
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(-2)
	}
	return strings.TrimSpace(str)
}

func ParseString(str string) ([]int, []operation) {
	nums, ops := make([]int, (len(str)+1)/2), make([]operation, (len(str)-1)/2)

	for i, v := range str {
		if i%2 == 0 {
			var err error
			nums[i/2], err = strconv.Atoi(string(v))
			if err != nil {
				os.Exit(-1)
			}
		} else {
			ops[(i-1)/2] = operation(v)
		}
	}
	return nums, ops
}

func updateMatrixes(lIdx, rIdx int, nums []int, ops []operation, mins, maxs [][]int) {
	maximum := -9999999
	minimum := 9999999
	for i := 0; lIdx+i < rIdx; i++ {
		op := ops[lIdx+i]
		// fmt.Printf("max through %d %d %d %d %d\nwhere lIdx = %d, rIdx = %d, i = %d\n", maximum, op.Apply(maxs[lIdx][i], maxs[rIdx-i][rIdx]), op.Apply(mins[lIdx][i], maxs[rIdx-i][rIdx]), op.Apply(maxs[lIdx][i], mins[rIdx-i][rIdx]), op.Apply(mins[lIdx][i], mins[rIdx-i][rIdx]), lIdx, rIdx, i)
		// fmt.Println(string(op))
		maximum = max(maximum, op.Apply(maxs[lIdx][lIdx+i], maxs[lIdx+i+1][rIdx]), op.Apply(mins[lIdx][lIdx+i], maxs[lIdx+i+1][rIdx]), op.Apply(maxs[lIdx][lIdx+i], mins[lIdx+i+1][rIdx]), op.Apply(mins[lIdx][lIdx+i], mins[lIdx+i+1][rIdx]))
		minimum = min(minimum, op.Apply(maxs[lIdx][lIdx+i], maxs[lIdx+i+1][rIdx]), op.Apply(mins[lIdx][lIdx+i], maxs[lIdx+i+1][rIdx]), op.Apply(maxs[lIdx][lIdx+i], mins[lIdx+i+1][rIdx]), op.Apply(mins[lIdx][lIdx+i], mins[lIdx+i+1][rIdx]))
	}
	mins[lIdx][rIdx] = minimum
	maxs[lIdx][rIdx] = maximum
}

func GetMaxValue(str string) int {
	nums, ops := ParseString(str)

	mins, maxs := make([][]int, len(nums)), make([][]int, len(nums))
	for idx := range mins {
		mins[idx] = make([]int, len(nums))
		maxs[idx] = make([]int, len(nums))
		mins[idx][idx], maxs[idx][idx] = nums[idx], nums[idx]
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j+i < len(nums); j++ {
			updateMatrixes(j, i+j, nums, ops, mins, maxs)
		}
	}
	// for i := range maxs {
	// 	fmt.Println(maxs[i])
	// }
	// for i := range maxs {
	// 	fmt.Println(mins[i])
	// }
	return maxs[0][len(nums)-1]
}

func min(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}
	return m
}

func max(nums ...int) int {
	m := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}

func main() {
	str := ReadInput()
	if len(str)%2 == 0 {
		fmt.Println("Wrong input")
		os.Exit(-1)
	}
	fmt.Println(GetMaxValue(str))
}
