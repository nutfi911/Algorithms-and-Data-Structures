package sorting

import "fmt"

// args[0] - left
// args[1] - right
func QuickSort(arr []int, args ...int) []int {
	var left int
	var right int

	if len(args) == 0 {
		left = 0
		right = len(arr) - 1
	} else {
		left = args[0]
		right = args[1]
	}

	if left < right {
		pivotIndex := Pivot(arr, left, right)

		fmt.Println("Pivot Index:", pivotIndex)
		fmt.Println("Pivot Value:", arr[0])
		fmt.Println("Current Array:", arr)

		// left
		QuickSort(arr, left, pivotIndex-1)

		// right
		QuickSort(arr, pivotIndex+1, right)
		// QuickSort(arr, pivotIndex+1, right)
	}
	return arr
}

// Returns the index of where the Pivot should go
// in our case the pivot is always the first element
// [3, 2, 4, 1, 5]
// => [2, 1, 3, 5, 5]
// => index 2
func Pivot(arr []int, start int, end int) int {
	pivotValue := arr[start]
	pivotIndex := start

	for i := start + 1; i <= end; i++ {
		if pivotValue > arr[i] {
			pivotIndex++
			swap(arr, pivotIndex, i)
		}
	}
	swap(arr, start, pivotIndex)

	return pivotIndex
}

func swap(arr []int, val1 int, val2 int) {

	temp := arr[val1]
	arr[val1] = arr[val2]
	arr[val2] = temp

}
