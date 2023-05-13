package recursion

// Write a recursive function called someRecursive which accepts an array and a callback.
// The function returns true if a single value in the array returns true
// when passed to the callback. Otherwise it returns false.
func SomeRecursive(arr []int32, callback func(n int32) bool) bool {
	if len(arr) == 0 {
		return false
	}

	if callback(arr[0]) {
		return true
	} else {
		return SomeRecursive(arr[1:], callback)
	}
}
