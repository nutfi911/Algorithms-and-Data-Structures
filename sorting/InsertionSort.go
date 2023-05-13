package sorting

import "fmt"

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		var j int
		current := arr[i]
		for j = i - 1; j >= 0 && arr[j] > current; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = current
		fmt.Println("Insertion sort / Iteration #", i , arr)
	}

	return arr
}
