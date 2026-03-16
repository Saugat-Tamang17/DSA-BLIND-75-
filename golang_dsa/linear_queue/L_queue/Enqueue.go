package L_queue

import (
	"fmt"
)

func Enqueue(q *Lqueue, item int) {
	if IsFull(q) {
		fmt.Println("Queue cannot store anymore value")
		return
	}

	if IsEmpty(q) {
		q.Front = 0
	}

	q.Rear++
	q.Arr[q.Rear] = item
	fmt.Println("Added the number inside the queue")

}
