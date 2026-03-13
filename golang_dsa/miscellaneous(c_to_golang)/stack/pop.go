package stack

import "fmt"

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack Underflow sarkaar ! Nothing to pop.")
		return -1
	}
	value := s.data[s.top]
	s.top--
	fmt.Println("Popped:", value)
	return value

}
