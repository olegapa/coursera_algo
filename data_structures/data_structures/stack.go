package datastructures

import (
	"errors"
	"reflect"
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

func (stack *Stack) Len() (int) {
	if stack == nil {
		return 0
	}
	return stack.len
}

func (stack *Stack) IsEmpty() (bool) {
	if stack == nil{
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

func (stack *Stack) Pop() (any, error){
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

func (stack *Stack) Push(val any) error{
	if stack == nil {
		return errors.New("recieved nil as the reciever")
	}
	topNode := stack.top
	if topNode != nil{
		if reflect.TypeOf(val) != reflect.TypeOf(topNode.value) {
			return errors.New("type missmatch in the stack")
		}
	}
	newNode := node{value: val, prev: topNode}

	stack.top = &newNode
	stack.len++
	return nil
}