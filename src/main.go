package main

import (
	"fmt"
	"queue"
	"stack"
)

func main() {
	stack := new(stack.StringStack)
	stack.Push("this")
	stack.Push("is")
	stack.Push("jon")
	stack.Print()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	queue := new(queue.StringQueue)
	queue.Add("this")
	queue.Add("is")
	queue.Add("jon")
	queue.Print()
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
	fmt.Println(queue.Poll())
}
