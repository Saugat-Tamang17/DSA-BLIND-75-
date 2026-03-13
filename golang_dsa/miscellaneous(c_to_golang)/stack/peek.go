package stack

import "fmt"

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty.Nothing to peek blud")
		return -1

	}
	return s.data[s.top]
}
