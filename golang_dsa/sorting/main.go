package main

import(
	"fmt"
	"github.com/Saugat-Tamang17/golang-dsa/Comparison/methods"

	"github.com/Saugat-Tamang17/golang-dsa/Comparison/Methods"
)

type Complexity struct{
	BestTime string
	AverageTime string
	WorstTime string
	Space string
}

var complexityMap = map[string]Complexity{
"bubble": {
		BestTime:    "O(n)      — already sorted",
		AverageTime: "O(n²)",
		WorstTime:   "O(n²)     — reverse sorted",
		Space:       "O(1)      — in-place",
	},
	"selection": {
		BestTime:    "O(n²)",
		AverageTime: "O(n²)",
		WorstTime:   "O(n²)",
		Space:       "O(1)      — in-place",
	},
	"insertion": {
		BestTime:    "O(n)      — already sorted",
		AverageTime: "O(n²)",
		WorstTime:   "O(n²)     — reverse sorted",
		Space:       "O(1)      — in-place",
	},
	"quicksort": {
		BestTime:    "O(n log n)",
		AverageTime: "O(n log n)",
		WorstTime:   "O(n²)     — bad pivot (already sorted)",
		Space:       "O(log n)  — recursive call stack",
	},
	"mergesort": {
		BestTime:    "O(n log n)",
		AverageTime: "O(n log n)",
		WorstTime:   "O(n log n)",
		Space:       "O(n)      — auxiliary array",
	},
}

func printComplexity(algo string){
	c:=complexityMap[algo]
	fmt.Println("complexity for :%s\n",algo)
	fmt.Printf("  Best Time    :%s\n",c.BestTime)
	fmt.Printf("  Average Time : %s\n", c.AverageTime)
	fmt.Printf("  Worst Time   : %s\n", c.WorstTime)
	fmt.Printf("  Space        : %s\n", c.Space)
}

//input handlers ( for array and algo input ) //
func getArray() []int{
	var n int
	fmt.Println("Enter the Number of elements:")
	fmt.Scan(&n)

	arr := make([]int,n)\
	fmt.Printf("Enter %d integers (space-separated): ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr

}
