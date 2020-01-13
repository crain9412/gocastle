package queue

import (
	"strings"
)

const EmptyQueue = "∅"
const LeftParen = "["
const RightParen = "]"
const Connector = " <-> "
const EmptyString = ""

type StringQueue struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	prev *Node
	data string
}

func (queue *StringQueue) Add(data string) {
	createdNode := Node{nil, nil, data}

	if queue.head == nil {
		queue.head = &createdNode
		queue.tail = &createdNode
	} else {
		oldHead := queue.head
		queue.head = &createdNode
		queue.head.next = oldHead
		oldHead.prev = queue.head
	}

	queue.head = &createdNode
}

func (queue *StringQueue) Print() string {
	var stringBuilder strings.Builder
	current := queue.head
	if current == nil {
		stringBuilder.WriteString(EmptyQueue)
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

func (queue *StringQueue) Poll() string {
	oldTail := queue.tail
	if oldTail == nil {
		return EmptyQueue
	}
	if oldTail.prev == nil {
		queue.head = nil
		queue.tail = nil
	} else {
		queue.tail = oldTail.prev
	}
	return oldTail.data
}
