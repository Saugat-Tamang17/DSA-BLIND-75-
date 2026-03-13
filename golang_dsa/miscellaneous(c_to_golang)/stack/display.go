package stack

import "fmt"

func (s *Stack) Display() {
	if s.IsEmpty() {
		fmt.Println("nothing to show brochacho")
		return
	}
	fmt.Println("Stack Elements:")
	for i := s.top; i >= 0; i-- {
		fmt.Println(s.data[i])
	}
}
