package main

import (
	"fmt"

	"github.com/Saugat-Tamang17/golang-dsa/Comparison/methods"
)

func main() {
	arr := []int{9, 4, 2, 7, 1}

	fmt.Println("Original:", arr)

	fmt.Println("Bubble Sort:", methods.BubbleSort(arr))
	fmt.Println("Selection Sort:", methods.Selection(arr))
	fmt.Println("Insertion Sort:", methods.Insertion(arr))
}
