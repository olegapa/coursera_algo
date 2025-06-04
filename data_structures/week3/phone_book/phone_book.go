package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	Name string
	Num  int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	phoneBook := make(map[int]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		parts := strings.Fields(scanner.Text())
		switch parts[0] {
		case "add":
			num, _ := strconv.Atoi(parts[1])
			name := parts[2]
			phoneBook[num] = name
		case "del":
			num, _ := strconv.Atoi(parts[1])
			delete(phoneBook, num)
		case "find":
			num, _ := strconv.Atoi(parts[1])
			if name, ok := phoneBook[num]; ok {
				fmt.Println(name)
			} else {
				fmt.Println("not found")
			}
		}
	}
}