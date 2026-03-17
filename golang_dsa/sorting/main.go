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

func getMethodChoice() int {
	fmt.Println("Choose the sorting method:\n")
	fmt.Println("  1. Comparison Based  (Bubble, Selection, Insertion)")
	fmt.Println("  2. Divide & Conquer  (Quick Sort, Merge Sort)")
	fmt.Println("enter the choice ( 1 or 2 ): ")

	var choice int
	fmt.Scan(&n)
	return choice
}

func getAlgorithmChoice(method int) string{
case 1:
	fmt.Println("You choosed the Sorting method (Sorting By Comparison)")
	fmt.Println("Choose the desired algorithm")
	fmt.Println(" 1. Bubble Sort")
	fmt.Println(" 2. Selection Sort")
	fmt.Println(" 3. Insertion Sort")
	fmt.Println("Enter the choice (1-3):")
	var choice int
	fmt.Scan(&choice)

	switch choice{
	case 1:
		return "bubble"
	case 2:
			return "selection"
	case 3:
			return "insertion"
	default:
			fmt.Println("Invalid choice, defaulting to Bubble Sort.")
	return "bubble"
	}

case 2:
	fmt.Println("You chose the sorting method ( By Divide&Conquer)")
	fmt.Println("Choose the sorting algorithm:")
	fmt.Println("  1. Quick Sort")
	fmt.Println("  2. Merge Sort")
	fmt.Print("Enter choice (1 or 2): ")
		var choice int
		fmt.Scan(&choice)
 
		switch choice {
		case 1:
			return "quicksort"
		case 2:
			return "mergesort"
		default:
			fmt.Println("Invalid choice, defaulting to Quick Sort.")
			return "quicksort"
		}
 
	default:
		fmt.Println("Invalid category, defaulting to Bubble Sort.")
		return "bubble"


	}

	func main(){
		fmt.Println("--------------------------------------")
		fmt.Println("Sorting Algorithms")
		fmt.Println("=======================================")
		arr :=getArray()

		fmt.Println("\n original Array : %v\t",arr)

		method:=getMethodChoice()

		algo :=getAlgorithmChoice(method)
				
		printComplexity(algo)

		copyArr := make([]int, len(arr))
	copy(copyArr, arr)
 
	fmt.Printf("Original Array : %v\n", arr)
 
	switch algo {
	case "bubble":
		fmt.Printf("Bubble Sort    : %v\n", methods.BubbleSort(copyArr))
	case "selection":
		fmt.Printf("Selection Sort : %v\n", methods.Selection(copyArr))
	case "insertion":
		fmt.Printf("Insertion Sort : %v\n", methods.Insertion(copyArr))
	case "quicksort":
		methods.QuickSort(copyArr, 0, len(copyArr)-1)
		fmt.Printf("Quick Sort     : %v\n", copyArr)
	case "mergesort":
		result := methods.MergeSort(copyArr)
		fmt.Printf("Merge Sort     : %v\n", result)
	}
}





