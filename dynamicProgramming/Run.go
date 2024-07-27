package dynamicprogramming

import "fmt"

func Run() {

	memo := make(map[int8]int64)
	memo[1] = 1

	fmt.Println("Fib Dynamic of 14:", FibDynamic(14, memo))
	fmt.Println("Fib Dynamic of 77:", FibDynamic(77, memo))
}
