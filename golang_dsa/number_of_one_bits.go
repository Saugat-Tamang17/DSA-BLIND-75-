package main

import "fmt"

// Precomputed lookup table for 8-bit numbers
var table [256]int

// Initialize the lookup table
func initTable() {
	for i := 0; i < 256; i++ {
		table[i] = (i & 1) + table[i>>1]
	}
}

// HammingWeight calculates number of set bits in a 32-bit integer using the lookup table
func HammingWeight(n int) int {
	return table[n&0xff] + // lowest 8 bits
		table[(n>>8)&0xff] + // next 8 bits
		table[(n>>16)&0xff] + // next 8 bits
		table[(n>>24)&0xff] // highest 8 bits
}

func main() {
	// Initialize lookup table once
	initTable()

	// Test examples
	numbers := []int{11, 128, 2147483645}

	for _, n := range numbers {
		fmt.Printf("Number: %d, Set bits: %d\n", n, HammingWeight(n))
	}
}
