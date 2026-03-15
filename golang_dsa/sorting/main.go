package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 4, 2, 7, 1}

	fmt.Println("Original:", arr)

	fmt.Println("Bubble Sort:", sorting_methods.BubbleSort(arr))
	fmt.Println("Selection Sort:", sorting_methods.SelectionSort(arr))
	fmt.Println("Insertion Sort:", sorting_methods.InsertionSort(arr))
}
