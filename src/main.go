package main

import (
	"fmt"
	"stack"
)

func main() {
	stack := new(stack.StringStack)
	stack.Push("world")
	stack.Push("hello")
	stack.Push("jon")
	stack.Push("is")
	stack.Push("this")
	stack.Print()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
