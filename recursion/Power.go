package recursion

// Write a function that recieves a number and the Power to
// Power(2,0) // 1
// Power(2,2) // 4
// Power(2,4) // 16
func Power(number int, pow int) int {
	switch pow {
	case 0:
		return 1
	case 1:
		return number
	default:
		return number * Power(number, pow-1)
	}
}
