package recursion

// Write a function that returns the multiplication of an array
// productOfArray([1,2,3,10]) // 60
func ProductOfArray(numbers []int32) int32 {
	if len(numbers) == 1 {
		return numbers[0]
	}

	return numbers[0] * ProductOfArray(numbers[1:])
}
