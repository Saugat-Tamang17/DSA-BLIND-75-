package L_queue

import "fmt"

func Dequeue(q *Lqueue) int {
	if IsEmpty(q) {
		fmt.Println("Queue Underflow.Nothing to add")
		return -1
	}
	item := q.Arr[q.Front]

	q.Front++

	// means that the queueu is heading towards empty state so reset the linear queue//
	if q.Front > q.Rear {
		q.Front = -1
		q.Rear = -1
	}

	return item
}
