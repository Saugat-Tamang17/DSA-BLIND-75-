package main

import (
	"fmt"

	"github.com/Saugat-Tamang17/DSA-BLIND-75-/L_queue"
)

func main() {
	var q L_queue.Lqueue

	// Initialize the queue
	L_queue.Initialize(&q)
	fmt.Println("Queue initialized.")

	// Check if queue is empty
	fmt.Println("Is queue empty?", L_queue.IsEmpty(&q))

	// Enqueue elements
	L_queue.Enqueue(&q, 10)
	L_queue.Enqueue(&q, 20)
	L_queue.Enqueue(&q, 30)
	fmt.Println("Added 10, 20, 30 to the queue.")

	// Peek at the front element
	front, err := L_queue.Peek(&q)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front element:", front)
	}

	// Check if queue is full
	fmt.Println("Is queue full?", L_queue.IsFull(&q))

	// Dequeue an element
	deq1 := L_queue.Dequeue(&q)
	fmt.Println("Dequeued element:", deq1)

	// Dequeue another element
	deq2 := L_queue.Dequeue(&q)
	fmt.Println("Dequeued element:", deq2)

	// Peek again after dequeue
	front, err = L_queue.Peek(&q)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Front element now:", front)
	}

	// Enqueue more elements
	L_queue.Enqueue(&q, 40)
	L_queue.Enqueue(&q, 50)
	fmt.Println("Added 40 and 50 to the queue.")

	// Dequeue all elements to empty the queue
	for !L_queue.IsEmpty(&q) {
		fmt.Println("Dequeued:", L_queue.Dequeue(&q))
	}

	// Try peeking on empty queue
	front, err = L_queue.Peek(&q)
	if err != nil {
		fmt.Println("Peek after emptying queue:", err)
	}

	// Final check
	fmt.Println("Is queue empty at the end?", L_queue.IsEmpty(&q))
	fmt.Println("Is queue full at the end?", L_queue.IsFull(&q))
}
