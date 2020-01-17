package main

import (
	"fmt"
	"queue"
	"stack"
	"crainsort"
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

	coolUnsortedArray := []int{5, 3, 2, 2, 1, 4}
	fmt.Println(crainsort.IntMergesort(coolUnsortedArray))

	coolerUnsortedArray := []int{5, 3, 2, 2, 1, 4, 0}
	crainsort.IntQuicksort(coolerUnsortedArray)
	fmt.Println(coolerUnsortedArray)
}
