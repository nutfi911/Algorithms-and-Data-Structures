package sorting

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			temp := arr[min]
			arr[min] = arr[i]
			arr[i] = temp

		}
	}

	return arr
}