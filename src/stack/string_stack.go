package stack

import (
	"strings"
)

const EmptyStack = "_"
const LeftParen = "["
const RightParen = "]"
const Connector = " -> "
const EmptyString = ""

type StringStack struct {
	head *Node
}

type Node struct {
	next *Node
	data string
}

func (stack *StringStack) Push(data string) {
	createdNode := Node{nil, data}
	if stack.head == nil {
		stack.head = &createdNode
	} else {
		oldHead := stack.head
		stack.head = &createdNode
		stack.head.next = oldHead
	}
}

func (stack *StringStack) Print() string {
	var stringBuilder strings.Builder
	current := stack.head
	if current == nil {
		stringBuilder.WriteString(EmptyStack)
	}
	for current != nil {
		end := EmptyString
		if current.next != nil {
			end = Connector
		}
		stringBuilder.WriteString(LeftParen)
		stringBuilder.WriteString(current.data)
		stringBuilder.WriteString(RightParen)
		stringBuilder.WriteString(end)
		current = current.next
	}
	return stringBuilder.String()
}

func (stack *StringStack) Pop() string {
	oldHead := stack.head
	if oldHead == nil {
		return EmptyStack
	}
	if oldHead.next == nil {
		stack.head = nil
	} else {
		stack.head = oldHead.next
	}
	return oldHead.data
}
