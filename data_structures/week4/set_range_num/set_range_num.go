package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SegmentTree struct {
	n    int
	tree []int
	lazy []int
	set  []bool
}

func NewSegmentTree(size int) *SegmentTree {
	n := 1
	for n < size {
		n <<= 1
	}
	return &SegmentTree{
		n:    n,
		tree: make([]int, 2*n),
		lazy: make([]int, 2*n),
		set:  make([]bool, 2*n),
	}
}

func (st *SegmentTree) push(v, l, r int) {
	if st.set[v] {
		st.tree[v] = (r - l) * st.lazy[v]
		if v < st.n {
			st.lazy[2*v] = st.lazy[v]
			st.lazy[2*v+1] = st.lazy[v]
			st.set[2*v] = true
			st.set[2*v+1] = true
		}
		st.set[v] = false
	}
}

func (st *SegmentTree) rangeSet(v, l, r, ql, qr, value int) {
	st.push(v, l, r)
	if ql >= r || qr <= l {
		return
	}
	if ql <= l && r <= qr {
		st.lazy[v] = value
		st.set[v] = true
		st.push(v, l, r)
		return
	}
	m := (l + r) / 2
	st.rangeSet(2*v, l, m, ql, qr, value)
	st.rangeSet(2*v+1, m, r, ql, qr, value)
	st.tree[v] = st.tree[2*v] + st.tree[2*v+1]
}

func (st *SegmentTree) rangeSum(v, l, r, ql, qr int) int {
	st.push(v, l, r)
	if ql >= r || qr <= l {
		return 0
	}
	if ql <= l && r <= qr {
		return st.tree[v]
	}
	m := (l + r) / 2
	return st.rangeSum(2*v, l, m, ql, qr) + st.rangeSum(2*v+1, m, r, ql, qr)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	st := NewSegmentTree(n)

	for i := 0; i < m; i++ {
		line, _ := reader.ReadString('\n')
		tokens := strings.Fields(line)
		if tokens[0] == "set" {
			l, _ := strconv.Atoi(tokens[1])
			r, _ := strconv.Atoi(tokens[2])
			val, _ := strconv.Atoi(tokens[3])
			st.rangeSet(1, 0, st.n, l, r, val)
		} else if tokens[0] == "sum" {
			l, _ := strconv.Atoi(tokens[1])
			r, _ := strconv.Atoi(tokens[2])
			res := st.rangeSum(1, 0, st.n, l, r)
			fmt.Fprintln(writer, res)
		}
	}
}
