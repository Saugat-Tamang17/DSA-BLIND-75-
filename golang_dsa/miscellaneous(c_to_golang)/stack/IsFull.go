package stack

func (s *Stack) IsFull() bool {
	return s.top == s.size-1
}
