package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key, prior  int
	left, right *Node
	sum         int64
}

func update(n *Node) {
	if n == nil {
		return
	}
	n.sum = int64(n.key)
	if n.left != nil {
		n.sum += n.left.sum
	}
	if n.right != nil {
		n.sum += n.right.sum
	}
}

func split(n *Node, key int) (*Node, *Node) {
	if n == nil {
		return nil, nil
	}
	if key < n.key {
		l, r := split(n.left, key)
		n.left = r
		update(n)
		return l, n
	} else {
		l, r := split(n.right, key)
		n.right = l
		update(n)
		return n, r
	}
}

func merge(l, r *Node) *Node {
	if l == nil || r == nil {
		if l != nil {
			return l
		}
		return r
	}
	if l.prior > r.prior {
		l.right = merge(l.right, r)
		update(l)
		return l
	} else {
		r.left = merge(l, r.left)
		update(r)
		return r
	}
}

func insert(n *Node, key int) *Node {
	if exists(n, key) {
		return n
	}
	node := &Node{key: key, prior: rand.Int(), sum: int64(key)}
	return insertNode(n, node)
}

func insertNode(n, node *Node) *Node {
	if n == nil {
		return node
	}
	if node.prior > n.prior {
		node.left, node.right = split(n, node.key)
		update(node)
		return node
	}
	if node.key < n.key {
		n.left = insertNode(n.left, node)
	} else {
		n.right = insertNode(n.right, node)
	}
	update(n)
	return n
}

func erase(n *Node, key int) *Node {
	if n == nil {
		return nil
	}
	if n.key == key {
		return merge(n.left, n.right)
	}
	if key < n.key {
		n.left = erase(n.left, key)
	} else {
		n.right = erase(n.right, key)
	}
	update(n)
	return n
}

func exists(n *Node, key int) bool {
	for n != nil {
		if n.key == key {
			return true
		}
		if key < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return false
}

func rangeSum(n *Node, l, r int) int64 {
	var t1, t2, t3 *Node
	t1, t2 = split(n, l-1)
	t2, t3 = split(t2, r)
	var res int64
	if t2 != nil {
		res = t2.sum
	}
	n = merge(t1, merge(t2, t3))
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var root *Node
	var last_sum int64 = 0

	var n int
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			i--
			continue
		}
		tokens := strings.Fields(line)
		switch tokens[0] {
		case "+":
			x, _ := strconv.Atoi(tokens[1])
			x = int((int64(x) + last_sum) % 1000000001)
			root = insert(root, x)
		case "-":
			x, _ := strconv.Atoi(tokens[1])
			x = int((int64(x) + last_sum) % 1000000001)
			root = erase(root, x)
		case "?":
			x, _ := strconv.Atoi(tokens[1])
			x = int((int64(x) + last_sum) % 1000000001)
			if exists(root, x) {
				fmt.Fprintln(writer, "Found")
			} else {
				fmt.Fprintln(writer, "Not found")
			}
		case "s":
			l, _ := strconv.Atoi(tokens[1])
			r, _ := strconv.Atoi(tokens[2])
			l = int((int64(l) + last_sum) % 1000000001)
			r = int((int64(r) + last_sum) % 1000000001)
			if l > r {
				last_sum = 0
				fmt.Fprintln(writer, 0)
			} else {
				last_sum = rangeSum(root, l, r)
				fmt.Fprintln(writer, last_sum)
			}
		}
	}
}
