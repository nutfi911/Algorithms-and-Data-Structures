package main

import (
	"data-structures-algorithms/dataStructures/hashmap"
	recursion "data-structures-algorithms/recursion"
	searching "data-structures-algorithms/searching"
	sorting "data-structures-algorithms/sorting"
	"fmt"
)

func recursionSet() {

	fmt.Println("There are odd numbers :", recursion.OddNumbers([]int32{3124, 5798, 6550, 5914}))

	fmt.Println("Sum range 4: ", recursion.SumRange(4)) // 10
	fmt.Println("Sum range 3: ", recursion.SumRange(3)) // 6

	fmt.Println("Factorial 10: ", recursion.Factorial(10)) // 24

	fmt.Println("2 to the power of 10: ", recursion.Power(2, 10)) // 1024

	fmt.Println("Product of array 1, 2, 3, 10: ", recursion.ProductOfArray([]int32{1, 2, 3, 10})) // 60

	fmt.Println("Recursive range to 6: ", recursion.RecursiveRange(6))   // 21
	fmt.Println("Recursive range to 10: ", recursion.RecursiveRange(10)) // 55

	fmt.Println("Fibonacci 35: ", recursion.Fib(35)) //9228465

	fmt.Println("Reverse string `pizza`: ", recursion.Reverse("pizza"))

	fmt.Println("Is palindrome 'abba':", recursion.IsPalindrome("abba"))

	fmt.Println("Pass callback function to check for even values in array 1,3,5,2: ",
		recursion.SomeRecursive([]int32{1, 3, 5, 2}, func(n int32) bool { return n%2 == 0 }))

	fmt.Println("Capitalize first: ", recursion.CapitalizeFirst([]string{"burger", "king"}))
}

func searchingSet() {

	fmt.Println(searching.BinarySearch([]int{1, 1, 2, 3, 5, 8, 13, 21, 44, 65}, 13))

	fmt.Println(searching.NaiveStringSearch("pizosasomgasdomgssomg", "omg"))

}

func sortingSet() {
	fmt.Println(sorting.BubbleSort([]int{5, 3, 2, 1, 4}))

	fmt.Println(sorting.SelectionSort([]int{5, 3, 2, 1, 4}))

	fmt.Println(sorting.InsertionSort([]int{5, 3, 2, 1, 4}))

	fmt.Println(sorting.MergeSort([]int{9, 1, 30, 5, 2, 14, 19, 100}))

	fmt.Println(sorting.QuickSort([]int{9, 4, 8, 2, 1, 5, 7, 6, 3}))

	fmt.Println(sorting.RadixSort([]int{9, 1, 30, 2023, 12, 333, 510, 2, 55, 5, 2, 1412, 19, 100}))

}

func main() {
	// fmt.Println("----- ALGORITHMS -----")
	// fmt.Println("----- Recursion -----")
	// recursionSet()

	// fmt.Println("\n----- Searching -----")
	// searchingSet()

	// fmt.Println("\n----- Sorting -----")
	// sortingSet()

	fmt.Println("\n----- DATA STRUCTURES -----")
	// fmt.Println("----- Singly Linked List -----")
	// SinglyLinkedList.Run();

	// fmt.Println("----- Doubly Linked List -----")
	// DoublyLinkedList.Run()

	// fmt.Println("----- Stack -----")
	// Stack.Run()

	// fmt.Println("----- Queue -----")
	// Queue.Run()

	// fmt.Println("----- Binary Search Tree -----")
	// bst.Run()

	// fmt.Println("----- Binary Heap -----")
	// heap.Run()

	// fmt.Println("----- Priority Queue -----")
	// priorityqueue.Run()

	fmt.Println("----- Hash Table -----")
	hashmap.Run()
}
