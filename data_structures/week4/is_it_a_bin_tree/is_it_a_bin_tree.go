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

func isBST(tree []Node, idx int, min, max int) bool {
	if idx == -1 {
		return true
	}
	node := tree[idx]
	if node.key < min || node.key > max {
		return false
	}
	return isBST(tree, node.left, min, node.key-1) &&
		isBST(tree, node.right, node.key, max)
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
	if n == 0 || isBST(tree, 0, -1<<31, 1<<31-1) {
		fmt.Println("CORRECT")
	} else {
		fmt.Println("INCORRECT")
	}
}
