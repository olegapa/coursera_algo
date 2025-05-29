package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() {

}

func CalculateValue() (int, int, []int, []int) {
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Wrong first string input")
	}

	slpittedInput := strings.Split(strings.TrimSpace(str), " ")
	n, w := strconv.Atoi(slpittedInput[0]), strconv.Atoi(slpittedInput[1])

	weights := make(int[], n)
	values := make(int[], n)
	for i = range n{

	}
	str, err = reader.ReadString('\n')


	return n, w, 
}

func main() {

}
