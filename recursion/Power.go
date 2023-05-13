package recursion

// Write a function that recieves a number and the Power to
// Power(2,0) // 1
// Power(2,2) // 4
// Power(2,4) // 16
func Power(number int32, pow int32) int32 {
	if pow == 1 {
		return number
	}
	return number * Power(number, pow-1)
}
