package sorting

// [9, 1, 30, 5]
// [9, 1] [30, 5]
// [9] [1] [30] [5]
// Merge(Merge([9], [1]), Merge([30], [5]))

func MergeSort(arr []int) []int {
	if len(arr) > 1 {
		return Merge(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
	}
	return arr
}

// Function for merging arrays - used in Merge Sort
// while element of array 1 is bigger than element of array 2
// incriease array 2 index and insert into new array
// else while
//
// a[9,10,50] b[2,14,99,100]
// => [2, 9, 10, 14, 50, 99, 100]
func Merge(a []int, b []int) []int {
	var merged []int = []int{}

	// Pointer for first array
	var p1 int = 0

	// Pointer for second array
	var p2 int = 0

	for p1 < len(a) && p2 < len(b) {
		if a[p1] < b[p2] {
			merged = append(merged, a[p1])
			p1 = p1 + 1
		} else {
			merged = append(merged, b[p2])
			p2 = p2 + 1
		}
	}

	if p1 < len(a) {
		merged = append(merged, a[p1:]...)
	}
	if p2 < len(b) {
		merged = append(merged, b[p2:]...)
	}

	return merged
}
