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
	linkedListNode struct {
		value *TreeNode
		next  *linkedListNode
		prev  *linkedListNode
	}
	linkedList struct {
		front  *linkedListNode
		back   *linkedListNode
		length int
	}
	TreeNode struct {
		Value    any
		children linkedList
		parent   *TreeNode
	}
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

func (node *TreeNode) GetChildren() []*TreeNode {
	childrenList := node.children
	res := make([]*TreeNode, childrenList.length)
	// fmt.Println(childrenList)
	currentItem := childrenList.front
	for i := 0; currentItem != nil; i++ {
		// fmt.Println(currentItem.next)
		res[i] = currentItem.value
		currentItem = currentItem.next
	}
	return res
}

func (node *TreeNode) AddChild(childNode *TreeNode) {
	var newLinkedListNode linkedListNode
	if node.children.length == 0 {
		newLinkedListNode = linkedListNode{value: childNode, prev: nil, next: nil}
		node.children.front = &newLinkedListNode
	} else {
		newLinkedListNode = linkedListNode{value: childNode, prev: node.children.back, next: nil}
		node.children.back.next = &newLinkedListNode
	}
	node.children.back = &newLinkedListNode
	node.children.length++
	childNode.parent = node
}

func NewNode(value any) *TreeNode {
	return &TreeNode{Value: value, children: linkedList{front: nil, back: nil, length: 0}, parent: nil}
}

func (root TreeNode) Height() int {
	type nodeDepth struct {
		node  TreeNode
		depth int
	}
	height := 1
	proccessStack := NewStack()
	proccessStack.Push(nodeDepth{root, 1})
	for !proccessStack.IsEmpty() {
		item, _ := proccessStack.Pop()
		curNode := item.(nodeDepth)
		if height < curNode.depth {
			height = curNode.depth
		}
		children := curNode.node.GetChildren()
		// fmt.Printf("children = %v, len = %d, node_cal = %v\n", children, len(children), curNode.node.Value)
		for i := 0; i < len(children); i++ {
			proccessStack.Push(nodeDepth{*children[i], curNode.depth + 1})
		}
	}
	return height
}

func ReadInput() (int, []int) {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(str))

	nums := make([]int, n)

	str, _ = reader.ReadString('\n')
	splStr := strings.Split(strings.TrimSpace(str), " ")
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(splStr[i])
	}

	return n, nums
}

func CreateTree(n int, nums []int) *TreeNode {
	nodes := make([]*TreeNode, n)
	var root *TreeNode
	for i := 0; i < n; i++ {
		if nodes[i] == nil {
			nodes[i] = NewNode(i)
		}
		if nums[i] == -1 {
			root = nodes[i]
		} else {
			parNode := nums[i]
			if nodes[parNode] == nil {
				nodes[parNode] = NewNode(parNode)
			}
			nodes[parNode].AddChild(nodes[i])
		}
	}
	return root
}

func main() {
	n, nums := ReadInput()
	root := CreateTree(n, nums)
	fmt.Println(root.Height())
}
