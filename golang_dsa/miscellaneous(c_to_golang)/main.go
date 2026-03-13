package main

import (
	"fmt"
	"stack-go/stack"
)

func main() {
	s := stack.NewStack(5)

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("Top element:", s.Peek())

	s.Display()

	s.Pop()
	s.Pop()

	fmt.Println("Current size:", s.Size())

	s.Display()
}
