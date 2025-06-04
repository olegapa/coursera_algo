package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeOrders struct {
	n     int
	key   []int
	left  []int
	right []int
}

func (t *TreeOrders) read() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t.n, _ = strconv.Atoi(scanner.Text())
	t.key = make([]int, t.n)
	t.left = make([]int, t.n)
	t.right = make([]int, t.n)
	for i := 0; i < t.n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		t.key[i], _ = strconv.Atoi(parts[0])
		t.left[i], _ = strconv.Atoi(parts[1])
		t.right[i], _ = strconv.Atoi(parts[2])
	}
}

func (t *TreeOrders) inOrder() []int {
	result := []int{}
	var inOrderRec func(int)
	inOrderRec = func(index int) {
		if index == -1 {
			return
		}
		inOrderRec(t.left[index])
		result = append(result, t.key[index])
		inOrderRec(t.right[index])
	}
	inOrderRec(0)
	return result
}

func (t *TreeOrders) preOrder() []int {
	result := []int{}
	var preOrderRec func(int)
	preOrderRec = func(index int) {
		if index == -1 {
			return
		}
		result = append(result, t.key[index])
		preOrderRec(t.left[index])
		preOrderRec(t.right[index])
	}
	preOrderRec(0)
	return result
}

func (t *TreeOrders) postOrder() []int {
	result := []int{}
	var postOrderRec func(int)
	postOrderRec = func(index int) {
		if index == -1 {
			return
		}
		postOrderRec(t.left[index])
		postOrderRec(t.right[index])
		result = append(result, t.key[index])
	}
	postOrderRec(0)
	return result
}

func printSlice(a []int) {
	for i, v := range a {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	tree := TreeOrders{}
	tree.read()
	printSlice(tree.inOrder())
	printSlice(tree.preOrder())
	printSlice(tree.postOrder())
}
