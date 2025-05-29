package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateSubsequences(s []int, substings map[int][][]int) {
	// fmt.Println(s)
	k, ok := substings[len(s)]
	if ok == false {
		substings[len(s)] = make([][]int, 0)
		k = substings[len(s)]
	}
	k = append(k, s)
	substings[len(s)] = k
	if len(s) != 1 {
		for i := 0; i < len(s); i++ {
			subs := make([]int, len(s)-1)
			for j, k := 0, 0; j < len(s)-1; k++ {
				if k == i {
					continue
				}
				subs[j] = s[k]
				j++
			}
			generateSubsequences(subs, substings)
		}
	}
}

func NaiveLongestSubseq(str1, str2 []int) int {
	if len(str1) == 0 || len(str2) == 0 {
		return 0
	}
	map1, map2 := make(map[int][][]int), make(map[int][][]int)
	generateSubsequences(str1, map1)
	generateSubsequences(str2, map2)
	max := 0
	for k, list1 := range map1 {
		list2, ok := map2[k]
		if ok {
			for _, s1 := range list1 {
				for _, s2 := range list2 {
					equals := true
					for i := 0; i < len(s1); i++ {
						if s1[i] != s2[i] {
							equals = false
							break
						}

					}
					if len(s1) > max && equals == true {
						max = len(s1)
					}
				}
			}
		}
	}
	return max
}

func TestFindLongestSubseq(t *testing.T) {

	t.Run("Naive algorithm comparison", func(t *testing.T) {
		for i := range 50 {
			n1, n2 := rand.Intn(9), rand.Intn(9)
			s1, s2 := make([]int, n1), make([]int, n2)
			for i := 0; i < n1; i++ {
				s1[i] = rand.Intn(2)
			}
			for i := 0; i < n2; i++ {
				s2[i] = rand.Intn(3)
			}
			fmt.Println(s1)
			fmt.Println(s2)
			actual, expected := FindLongestSubseq(s1, s2), NaiveLongestSubseq(s1, s2)
			fmt.Println(strconv.Itoa(i) + ". Actual = " + strconv.Itoa(actual) + " Expected = " + strconv.Itoa(expected))
			assert.Equal(t, expected, actual)
		}
	})
}
