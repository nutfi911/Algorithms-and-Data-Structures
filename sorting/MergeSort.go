package sorting

import "fmt"


func MergeSort() {
	fmt.Println(merge([]int{9, 10, 50}, []int{2, 14, 99, 100}))
}

// a[9,10,50] b[2,14,99,100]
func merge(a []int, b []int) []int {
	merged := []int{}
	// while element of array 1 is bigger than element of array 2
	// incriease array 2 index and insert into new array
	// else while

	p1 := 0
	p2 := 0

	for i := 0; i < len(a)+len(b); i++ {
		if a[p1] < b[p2] {
			merged = append(merged, a[p1])
			if p1 + 1 > len(a) - 1{
				merged = append(merged, b[p2:]...)
				return merged
			} else{
				p1 = p1 + 1
			}
			continue
		} else {
			merged = append(merged, b[p2])
			if p2 + 1 > len(b) - 1{
				merged = append(merged, a[p1:]...)
				return merged

			} else{
				p2 = p2 + 1
			}
			continue
		}

	}

	return merged
}