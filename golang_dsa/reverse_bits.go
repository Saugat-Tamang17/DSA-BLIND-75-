package main

import (
	"fmt"
)

func reverseBits(n int) int {
	result := 0

	for i := 0; i < 32; i++ {
		result = (result << 1) | (n & 1)
		n = n >> 1
	}

	return result
}

func main() {
	var n int

	fmt.Print("Enter an integer: ")
	fmt.Scan(&n)

	reversed := reverseBits(n)

	fmt.Println("Original number:", n)
	fmt.Println("Reversed bits number:", reversed)
}
