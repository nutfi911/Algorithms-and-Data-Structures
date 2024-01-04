package SinglyLinkedList

import "fmt"

func Run(){
	sls := New()

	Push(sls, 1)
	Push(sls, 11)
	Push(sls, 12)
	Push(sls, 13)
	Push(sls, 14)
	Push(sls, 20)
	
	PrintList(sls)

	fmt.Println(Pop(sls))
	PrintList(sls)

	fmt.Println(Shift(sls))
	PrintList(sls)
	
	Unshift(sls, 0)
	PrintList(sls)

	fmt.Println(Get(sls,sls.Length-1).Value)
	Set(sls,sls.Length-1, 22)
	fmt.Println(Get(sls,sls.Length-1).Value)
	PrintList(sls)

	Insert(sls, 2, 10)
	PrintList(sls)

	fmt.Println("Removed value : ", Remove(sls,3))
	PrintList(sls)

	Reverse(sls)
	PrintList(sls)

	Reverse(sls)
	PrintList(sls)
}