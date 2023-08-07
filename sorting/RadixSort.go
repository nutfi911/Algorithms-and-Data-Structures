package sorting

import (
	"data-structures-algorithms/recursion"
	"math"
)

func RadixSort(arr []int) []int {
	maxDigitsCount := mostDigits(arr)

	for k := 0; k < maxDigitsCount; k++ {
		buckets := make([][]int, 10)
		for _, num := range arr {
			d := getDigit(num, k)
			buckets[d] = append(buckets[d], num)
		}
		arr = unpack(buckets)
	}
	return arr
}

func getDigit(num int, place int) int {
	return (int(math.Abs(float64(num))) / recursion.Power(10, place)) % 10
}

func digitCount(num int) int {
	if num == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(num))) + 1))
}

func mostDigits(nums []int) int {
	maxDigits := 0

	for _, n := range nums {
		dc := digitCount(n)
		if dc > maxDigits {
			maxDigits = dc
		}
	}
	return maxDigits
}

func unpack(sliceOfSlices [][]int) []int {

	var unpacked []int

	// Iterate through each sub-slice
	for _, subSlice := range sliceOfSlices {
		unpacked = append(unpacked, subSlice...)
	}

	return unpacked

}
