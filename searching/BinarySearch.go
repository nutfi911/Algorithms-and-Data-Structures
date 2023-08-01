package searching

func BinarySearch(arr []int, n int) int {

	j := 0
	right := len(arr) - 1

	for j <= right {

		i := (j + right) / 2

		if arr[i] == n {
			return i
		}

		if arr[i] > n {
			right = i - 1
		} else {
			j = i + 1
		}

	}

	return -1
}
