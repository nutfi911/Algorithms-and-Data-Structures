package recursion

// Write a function that adds up all numbers from 0 to num variable
// RecursiveRange(6) // 21
// RecursiveRange(10) // 55
func RecursiveRange(n int32) int32 {

	if n <= 1 {
		return 1
	}

	return n + RecursiveRange(n-1)
}
