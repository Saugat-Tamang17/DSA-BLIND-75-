package l_queue

const MAX_SIZE = 100

type lqueue struct {
	Arr   [MAX_SIZE]int
	Front int
	Rear  int
}
