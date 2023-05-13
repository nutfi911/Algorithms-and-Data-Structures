package recursion

func Factorial(number int32) int32 {
	if number == 1 {
		return number
	}
	return number * Factorial(number-1)
}
