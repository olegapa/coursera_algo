package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const bucketCount = 5

type Query struct {
	typ  string
	s    string
	ind  int
}

type HashChains struct {
	buckets [][]string
}

func NewHashChains() *HashChains {
	return &HashChains{
		buckets: make([][]string, bucketCount),
	}
}

func (hc *HashChains) hashFunc(s string) int {
	var hash int64
	const multiplier = int64(263)
	const prime = int64(1000000007)
	for i := len(s) - 1; i >= 0; i-- {
		hash = (hash*multiplier + int64(s[i])) % prime
	}
	return int(hash) % bucketCount
}

func (hc *HashChains) Add(s string) {
	idx := hc.hashFunc(s)
	for _, v := range hc.buckets[idx] {
		if v == s {
			return
		}
	}
	hc.buckets[idx] = append([]string{s}, hc.buckets[idx]...)
}

func (hc *HashChains) Delete(s string) {
	idx := hc.hashFunc(s)
	for i, v := range hc.buckets[idx] {
		if v == s {
			hc.buckets[idx] = append(hc.buckets[idx][:i], hc.buckets[idx][i+1:]...)
			return
		}
	}
}

func (hc *HashChains) Find(s string) bool {
	idx := hc.hashFunc(s)
	for _, v := range hc.buckets[idx] {
		if v == s {
			return true
		}
	}
	return false
}

func (hc *HashChains) Check(ind int) []string {
	return hc.buckets[ind]
}

func readQuery(scanner *bufio.Scanner) Query {
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	if parts[0] == "check" {
		ind := 0
		fmt.Sscanf(parts[1], "%d", &ind)
		return Query{typ: "check", ind: ind}
	}
	return Query{typ: parts[0], s: parts[1]}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n := 0
	fmt.Sscanf(scanner.Text(), "%d", &n)
	queries := make([]Query, n)
	for i := 0; i < n; i++ {
		queries[i] = readQuery(scanner)
	}

	hc := NewHashChains()
	for _, query := range queries {
		switch query.typ {
		case "add":
			hc.Add(query.s)
		case "del":
			hc.Delete(query.s)
		case "find":
			if hc.Find(query.s) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "check":
			fmt.Println(strings.Join(hc.Check(query.ind), " "))
		}
	}
}