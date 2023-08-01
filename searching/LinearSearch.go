package searching

func Linearsearch(arr []int, key int) int {
	for i, item := range arr {
		if item == key {
			return i
		}
	}
	return -1
}
