package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Spice struct {
	v, w int
}

func ReadInput() (int, []Spice) {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Wrong input in 1 line")
		os.Exit(1)
	}

	splitedInput := strings.Split(strings.TrimSpace(str), " ")
	var n, w int

	n, err = strconv.Atoi(splitedInput[0])
	if err != nil {
		fmt.Println("n param is incorrect")
		os.Exit(1)
	}

	w, err = strconv.Atoi(splitedInput[1])
	if err != nil {
		fmt.Println("w param is incorrect")
		os.Exit(1)
	}

	spices := make([]Spice, n)
	for i := 0; i < n; i++ {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Wrong input for %d spice\n", i)
			os.Exit(1)
		}
		splitedInput := strings.Split(strings.TrimSpace(str), " ")
		spices[i].v, err = strconv.Atoi(splitedInput[0])
		if err != nil {
			fmt.Printf("value param is incorrect for spice %d", i)
			os.Exit(1)
		}
		spices[i].w, err = strconv.Atoi(splitedInput[1])
		if err != nil {
			fmt.Printf("weight param is incorrect for spice %d", i)
			os.Exit(1)
		}
	}

	return w, spices
}

func swapSpises(spices []Spice, i, j int) {
	spices[i], spices[j] = spices[j], spices[i]
}

func isHigher(s1, s2 Spice) bool {
	return float64(s1.v)/float64(s1.w) >= float64(s2.v)/float64(s2.w)
}

func sortSpices(spices []Spice) {
	pv_idx := len(spices) - 1
	pv := spices[pv_idx]

	if pv_idx == 0 {
		return
	}

	i := -1
	for j := 0; j <= pv_idx; j++ {
		if isHigher(spices[j], pv) {
			i++
			if j > i {
				swapSpises(spices, i, j)
			}
		}
	}
	if i > 0 {
		sortSpices(spices[0:i])
	}
	if i < pv_idx {
		sortSpices(spices[i+1 : pv_idx+1])
	}

}

func CalculateMaxLootValue(w int, spices []Spice) float64 {
	sortSpices(spices)
	fmt.Println(spices)

	totalWeight, totalValue := 0, float64(0)

	for i := 0; totalWeight < w; i++ {
		weightLeft := float64(w - totalWeight)
		if weightLeft <= float64(spices[i].w) {
			totalValue += float64(spices[i].v) * float64(weightLeft/float64(spices[i].w))
			return totalValue
		}

		totalWeight += spices[i].w
		totalValue += float64(spices[i].v)
	}
	return totalValue
}

func main() {
	w, spicesList := ReadInput()

	fmt.Printf("%.4f\n", CalculateMaxLootValue(w, spicesList))

}
