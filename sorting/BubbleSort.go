package sorting

import "fmt"

func BubbleSort(arr []int) []int {

	noSwap := true

	for i := len(arr); i > 0; i-- {
		noSwap = true

		for j := 0; j < i-1; j++ {

			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp

				noSwap = false

			}

			fmt.Println("Main iteration ", (i-len(arr))*(-1), " / Sub iteration ", j+1, " / Array: ", arr)

		}
		if noSwap {
			break
		}
	}

	return arr
}
