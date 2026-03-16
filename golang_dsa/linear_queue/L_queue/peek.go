package L_queue

import "fmt"

func Peek(q *Lqueue) (int, error) {
	if q.Front == -1 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.Arr[q.Front], nil
}
