package Methods

func QuickSort(arr []int, low, high int) {
	if low < high {

		// pi is the partition return index of pivot
		pi := Partition(arr, low, high)

		// Recursion calls for smaller elements and greater or equal elements //
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)

	}
}
