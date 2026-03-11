package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	// Create a DP slice to store the number of ways
	dp := make([]int, n+1)

	// Base cases
	dp[1] = 1
	dp[2] = 2

	// Fill the DP slice using the recurrence relation
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	var n int
	fmt.Print("Enter number of stairs: ")
	fmt.Scan(&n)

	result := climbStairs(n)
	fmt.Printf("Number of ways to climb %d stairs: %d\n", n, result)
}
