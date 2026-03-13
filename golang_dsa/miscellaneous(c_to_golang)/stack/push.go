package stack

import "fmt"

func (s *Stack) Push(value int) {
	if s.IsFull() {
		fmt.Println("Stacked Overflowed brochaho! cannot fit more values:", value)
		return
	}

	s.top++
	s.data[s.top] = value
	fmt.Println("Pushed:", value)
}
