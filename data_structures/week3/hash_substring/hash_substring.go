package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	prime = 1000000007
	x     = 263
)


func polyHash(s string) int {
	hash := 0
	for i := len(s) - 1; i >= 0; i-- {
		hash = (hash*x + int(s[i])) % prime
	}
	return hash
}

func precomputeHashes(text string, patternLen int) []int {
	n := len(text)
	H := make([]int, n-patternLen+1)
	substr := text[n-patternLen:]
	H[n-patternLen] = polyHash(substr)
	y := 1
	for i := 0; i < patternLen; i++ {
		y = (y * x) % prime
	}
	for i := n - patternLen - 1; i >= 0; i-- {
		preHash := x*H[i+1] + int(text[i]) - y*int(text[i+patternLen])
		H[i] = ((preHash % prime) + prime) % prime
	}
	return H
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pattern, _ := reader.ReadString('\n')
	text, _ := reader.ReadString('\n')
	pattern = trimNewline(pattern)
	text = trimNewline(text)

	pLen := len(pattern)
	tLen := len(text)
	result := []int{}

	if pLen > tLen {
		fmt.Println()
		return
	}

	pHash := polyHash(pattern)
	H := precomputeHashes(text, pLen)

	for i := 0; i <= tLen-pLen; i++ {
		if H[i] == pHash && text[i:i+pLen] == pattern {
			result = append(result, i)
		}
	}

	for i, idx := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(idx)
	}
	fmt.Println()
}

func trimNewline(s string) string {
	if len(s) > 0 && (s[len(s)-1] == '\n' || s[len(s)-1] == '\r') {
		return s[:len(s)-1]
	}
	return s
}