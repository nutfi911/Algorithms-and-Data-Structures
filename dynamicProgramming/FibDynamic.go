package dynamicprogramming

// Fibonacci sequence using dynamic programming
// Time Complexity = O(n)
func FibDynamic(n int8, memo map[int8]int64) int64 {
	if val, exists := memo[n]; exists {
		return val
	}

	if n <= 2 {
		return 1
	}

	res := FibDynamic(n-1, memo) + FibDynamic(n-2, memo)
	memo[n] = res
	return res
}
