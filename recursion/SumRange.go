package recursion

//input 3 -> result: 3+2+1 = 6
func SumRange(number int32) int32 {
	if number == 1 {
		return number
	}

	return number + SumRange(number-1)
}
