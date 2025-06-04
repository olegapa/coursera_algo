package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RopeNode struct {
	weight int
	left   *RopeNode
	right  *RopeNode
	str    string
}

func NewRope(s string) *RopeNode {
	if len(s) <= 5 {
		return &RopeNode{
			weight: len(s),
			str:    s,
		}
	}
	mid := len(s) / 2
	left := NewRope(s[:mid])
	right := NewRope(s[mid:])
	return &RopeNode{
		weight: len(s[:mid]),
		left:   left,
		right:  right,
	}
}

func Split(r *RopeNode, i int) (*RopeNode, *RopeNode) {
	if r == nil {
		return nil, nil
	}
	if r.left == nil && r.right == nil {
		if i >= len(r.str) {
			return r, nil
		}
		return NewRope(r.str[:i]), NewRope(r.str[i:])
	}
	if i < r.weight {
		l, r2 := Split(r.left, i)
		return l, Concat(r2, r.right)
	}
	l2, r2 := Split(r.right, i-r.weight)
	return Concat(r.left, l2), r2
}

func Concat(left, right *RopeNode) *RopeNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return &RopeNode{
		weight: left.Length(),
		left:   left,
		right:  right,
	}
}

func (r *RopeNode) Length() int {
	if r == nil {
		return 0
	}
	if r.left == nil && r.right == nil {
		return len(r.str)
	}
	return r.weight + r.right.Length()
}

func (r *RopeNode) ToString() string {
	if r == nil {
		return ""
	}
	if r.left == nil && r.right == nil {
		return r.str
	}
	return r.left.ToString() + r.right.ToString()
}

func (r *RopeNode) Index(i int) byte {
	if r.left == nil && r.right == nil {
		return r.str[i]
	}
	if i < r.weight {
		return r.left.Index(i)
	}
	return r.right.Index(i - r.weight)
}

func Process(r *RopeNode, i, j, k int) *RopeNode {
	left, midright := Split(r, i)
	mid, right := Split(midright, j-i+1)
	without := Concat(left, right)
	if k == 0 {
		return Concat(mid, without)
	}
	left2, right2 := Split(without, k)
	return Concat(Concat(left2, mid), right2)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	rope := NewRope(s)

	var q int
	fmt.Fscanln(reader, &q)
	for i := 0; i < q; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)
		if len(parts) != 3 {
			continue
		}
		i1, _ := strconv.Atoi(parts[0])
		j1, _ := strconv.Atoi(parts[1])
		k1, _ := strconv.Atoi(parts[2])
		rope = Process(rope, i1, j1, k1)
	}
	fmt.Println(rope.ToString())
}
