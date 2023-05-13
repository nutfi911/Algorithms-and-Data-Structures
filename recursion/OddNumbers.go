package recursion

import "fmt"

// Check if array has odd numbers
func OddNumbers(numbers []int32) bool {
	var ITERATION int = 0
	ITERATION++

	fmt.Println("Current iteration: ", ITERATION, ", array: ", numbers)

	if len(numbers) == 0 {
		return false
	}

	if numbers[0]%2 != 0 {
		return true
	} else {
		// Recursion with with the array of which current first index is being removed
		return OddNumbers(numbers[1:])
	}
}
