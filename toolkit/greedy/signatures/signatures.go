package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tenant struct {
	l, r int
}

func ReadInput() (int, []Tenant) {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Wrong input in 1 line")
		os.Exit(1)
	}

	splitedInput := strings.Split(strings.TrimSpace(str), " ")
	var n int

	n, err = strconv.Atoi(splitedInput[0])
	if err != nil {
		fmt.Println("n param is incorrect")
		os.Exit(1)
	}

	spices := make([]Tenant, n)
	for i := 0; i < n; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Wrong input for %d spice\n", i)
			os.Exit(1)
		}
		splitedInput := strings.Split(strings.TrimSpace(str), " ")
		spices[i].l, err = strconv.Atoi(splitedInput[0])
		if err != nil {
			fmt.Printf("value param is incorrect for spice %d", i)
			os.Exit(1)
		}
		spices[i].r, err = strconv.Atoi(splitedInput[1])
		if err != nil {
			fmt.Printf("weight param is incorrect for spice %d", i)
			os.Exit(1)
		}
	}

	return n, spices
}

func swapLines(spices []Tenant, i, j int) {
	spices[i], spices[j] = spices[j], spices[i]
}

func (t1 Tenant) isLefter(t2 Tenant) bool {
	if t1.l == t2.l {
		return t1.r <= t2.r
	} else {
		return t1.l <= t2.l
	}
}

func sortLines(lines []Tenant) {
	pv_idx := len(lines) - 1
	pv := lines[pv_idx]

	if pv_idx == 0 {
		return
	}

	i := -1
	for j := 0; j <= pv_idx; j++ {
		if lines[j].isLefter(pv) {
			i++
			if j > i {
				swapLines(lines, i, j)
			}
		}
	}
	if i > 0 {
		sortLines(lines[0:i])
	}
	if i < pv_idx {
		sortLines(lines[i+1 : pv_idx+1])
	}

}

func CalculateMaxLootValue(n int, lines []Tenant) []int {
	sortLines(lines)

	timeCoord := []int{lines[0].r}
	for i := 1; i < n; i++ {
		if lines[i].l > timeCoord[len(timeCoord)-1] {
			timeCoord = append(timeCoord, lines[i].r)
		}
	}
	return timeCoord
}

func main() {
	n, spicesList := ReadInput()

	coord := CalculateMaxLootValue(n, spicesList)
	fmt.Println(len(coord))
	fmt.Println(strings.Trim(fmt.Sprint(coord), "[]"))
}
