package recursion

// Write a Fibonacci sequence function
// Fn = X = Fn-1 + Fn-2
// given n (position in the sequence), find X (number on that position)

// Fib(4) // 3
// Fib(10) // 55
// Fib(28) // 317811
// Fib(35) // 9227465
func Fib(n int32) int32 {

	if n <= 2 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}
