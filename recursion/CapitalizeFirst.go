package recursion

import "strings"

// Write a recursive function called capitalizeFirst.
// Given an array of strings, capitalize the first letter of each string in the array.
func CapitalizeFirst(arr []string) []string {
	var result []string
	if len(arr) == 0 {
		return result
	}

	result = append(result, strings.ToUpper(string(arr[0][0]))+string(arr[0][1:]))

	return append(result, CapitalizeFirst(arr[1:])...)
}
