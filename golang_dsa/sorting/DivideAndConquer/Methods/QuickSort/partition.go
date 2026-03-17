package Methods

func Partition(arr []int, low int, high int) int {

	//we choosing the last element as pivot in this one //
	pivot := arr[high]

	// Index of smaller element and indicates the right position of pivot found so far//
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)

		}
	}
	swap(arr, i+1, high)
	return i + 1
}
