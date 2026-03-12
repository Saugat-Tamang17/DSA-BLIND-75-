package main

import "fmt"

// Definition of a singly linked list node //
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to print a linked list
func printList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("NULL")
}

// Function to reverse the linked list
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

func main() {
	// Create linked list: 1 -> 2 -> 3 -> 4 -> 5 -> NULL
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("Original Linked List:")
	printList(head)

	// Reverse the linked list
	head = reverseList(head)

	fmt.Println("Reversed Linked List:")
	printList(head)
}
