package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type (
	node struct {
		value any
		prev  *node
	}
	Stack struct {
		top *node
		len int
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Len() int {
	if stack == nil {
		return 0
	}
	return stack.len
}

func (stack *Stack) IsEmpty() bool {
	if stack == nil {
		return true
	}
	return stack.len == 0
}

func (stack *Stack) Top() (any, error) {
	if stack == nil {
		return nil, errors.New("recieved nil as the reciever")
	}
	if stack.len == 0 {
		return nil, nil
	}
	return stack.top.value, nil
}

func (stack *Stack) Pop() (any, error) {
	if stack == nil {
		return nil, errors.New("recieved nil as the reciever")
	}
	if stack.len == 0 {
		return nil, errors.New("cannot pop from empty stack")
	}
	topNode := stack.top
	stack.top = topNode.prev
	stack.len--
	return topNode.value, nil
}

func (stack *Stack) Push(val any) error {
	if stack == nil {
		return errors.New("recieved nil as the reciever")
	}
	topNode := stack.top
	if topNode != nil {
		if reflect.TypeOf(val) != reflect.TypeOf(topNode.value) {
			return errors.New("type missmatch in the stack")
		}
	}
	newNode := node{value: val, prev: topNode}

	stack.top = &newNode
	stack.len++
	return nil
}

type Bracket struct {
	position int
	brType   byte
}

func (b1 Bracket) Match(b2 byte) bool {
	return (b1.brType == '[' && b2 == ']') || (b1.brType == '{' && b2 == '}') || (b1.brType == '(' && b2 == ')')
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(-1)
	}
	return strings.TrimSpace(str)
}

func ValidateBrackets(str string) string {
	stack := NewStack()
	for i, c := range []byte(str) {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(Bracket{i+1, c})
		}

		if c == ')' || c == ']' || c == '}' {
			top, _ := stack.Pop()
			if top == nil {
				return strconv.Itoa(i + 1)
			}
			topBracket := top.(Bracket)
			if !Bracket(topBracket).Match(c) {
				return strconv.Itoa(i + 1)
			}
		}
	}
	if !stack.IsEmpty() {
		top, _ := stack.Top()
		topBracket := top.(Bracket)
		return strconv.Itoa(topBracket.position)
	}
	return "Success"
}

func main() {
	str := ReadInput()
	fmt.Println(ValidateBrackets(str))
}
