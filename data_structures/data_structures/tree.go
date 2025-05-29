package datastructures

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

func (node *TreeNode) GetChildren() []*TreeNode {
	childrenList := node.children
	res := make([]*TreeNode, childrenList.length)
	currentItem := childrenList.front
	for i := 0; i < childrenList.length; i++ {
		res[i] = currentItem.value
		currentItem = currentItem.next
	}
	return res
}

func (node *TreeNode) AddChild(childNode *TreeNode) {
	var prev *linkedListNode = nil
	if node.children.length != 0 {
		prev = node.children.back
	}
	newLinkedListNode := linkedListNode{value: childNode, prev: prev, next: nil}
	node.children.back.next = &newLinkedListNode
	if node.children.length == 0 {
		node.children.front = &newLinkedListNode
	}
	node.children.length++
	childNode.parent = node
}

func NewNode(value any) TreeNode {
	return TreeNode{Value: value, children: linkedList{front: nil, back: nil, length: 0}, parent: nil}
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
		for i := 0; i < len(children); i++ {
			proccessStack.Push(nodeDepth{*children[i], curNode.depth + 1})
		}
	}
	return height
}
