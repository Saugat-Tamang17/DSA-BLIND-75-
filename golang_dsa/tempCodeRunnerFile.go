package main

import ("fmt")

func twoSum(nums []int, target int)[]int{
	seen:=make(map[int]int);

	for i,num:=range nums{
		complement:=target-num;

		if j,found:=seen[complement]; found{
			return []int{j,i};
		}
		seen[num]=i
	}
	return nil;
}


func main(){
	fmt.Println("this program will highlight the problem of 2 sum thing.")
	fmt.Println("\nEnter the number of elements you want to input inside the array.:");
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	fmt.Println("Enter the array elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	
	var target int
	fmt.Println("Enter the target sum:")
	fmt.Scan(&target)
	
	result := twoSum(nums, target)
	fmt.Println("Result:", result)
	

}