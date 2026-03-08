package main

import ("fmt")


func main(){
	fmt.Println("this program will highlight the problem of 2 sum thing.")
	fmt.Println("\nEnter the number of elements you want to input inside the array.:");
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	fmt.Println("Enter the array elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	
	var target int
	fmt.Println("Enter the target sum:")
	fmt.Scan(&target)
	
	result := twoSum(arr, target)
	fmt.Println("Result:", result)
	
	

}