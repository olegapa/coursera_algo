package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key   int
	left  int
	right int
}

func isBST(tree []Node) bool {
	if len(tree) == 0 {
		return true
	}
	var inOrder func(int, int, int) bool
	prev := -1 << 31
	var prevSet bool

	inOrder = func(index int, min, max int) bool {
		if index == -1 {
			return true
		}
		node := tree[index]
		if node.key < min || node.key > max {
			return false
		}
		if !inOrder(node.left, min, node.key-1) {
			return false
		}
		if prevSet && node.key < prev {
			return false
		}
		prev = node.key
		prevSet = true
		if !inOrder(node.right, node.key, max) {
			return false
		}
		return true
	}
	return inOrder(0, -1<<31, 1<<31-1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	tree := make([]Node, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		key, _ := strconv.Atoi(parts[0])
		left, _ := strconv.Atoi(parts[1])
		right, _ := strconv.Atoi(parts[2])
		tree[i] = Node{key, left, right}
	}
	if isBST(tree) {
		fmt.Println("CORRECT")
	} else {
		fmt.Println("INCORRECT")
	}
}
