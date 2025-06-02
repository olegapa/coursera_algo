package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DSU struct {
	parent []int
	rank   []int
	size   []int
	max    int
}

func NewDSU(n int, sizes []int) *DSU {
	maxSize := 0
	for _, s := range sizes {
		if s > maxSize {
			maxSize = s
		}
	}
	return &DSU{
		parent: func() []int {
			p := make([]int, n)
			for i := range p {
				p[i] = i
			}
			return p
		}(),
		rank: make([]int, n),
		size: append([]int(nil), sizes...),
		max:  maxSize,
	}
}

func (d *DSU) Find(i int) int {
	if d.parent[i] != i {
		d.parent[i] = d.Find(d.parent[i])
	}
	return d.parent[i]
}

func (d *DSU) Union(dest, src int) int {
	destRoot := d.Find(dest)
	srcRoot := d.Find(src)
	if destRoot == srcRoot {
		return d.max
	}
	if d.rank[destRoot] < d.rank[srcRoot] {
		destRoot, srcRoot = srcRoot, destRoot
	}
	d.parent[srcRoot] = destRoot
	d.size[destRoot] += d.size[srcRoot]
	d.size[srcRoot] = 0
	if d.rank[destRoot] == d.rank[srcRoot] {
		d.rank[destRoot]++
	}
	if d.size[destRoot] > d.max {
		d.max = d.size[destRoot]
	}
	return d.max
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	parts := strings.Fields(line1)
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])

	line2, _ := reader.ReadString('\n')
	sizes := make([]int, n)
	for i, s := range strings.Fields(line2) {
		sizes[i], _ = strconv.Atoi(s)
	}

	dsu := NewDSU(n, sizes)

	for i := 0; i < m; i++ {
		opLine, _ := reader.ReadString('\n')
		ops := strings.Fields(opLine)
		dest, _ := strconv.Atoi(ops[0])
		src, _ := strconv.Atoi(ops[1])
		fmt.Println(dsu.Union(dest-1, src-1))
	}
}