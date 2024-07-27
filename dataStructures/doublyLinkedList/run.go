package DoublyLinkedList

import "fmt"

func Run() {
	l := New()

	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)

	l.Pop()

	l.Shift()

	l.Unshift(1)
	l.Unshift(6)

	fmt.Println("Position 0: ", l.Get(0).Value)
	fmt.Println("Position 2: ", l.Get(2).Value)

	l.Set(2, 44)

	l.Insert(2, 33)

	l.Remove(3)

	l.Reverse()
	l.Reverse()

	fmt.Println("Length ", l.Length)

	PrintList(l)
}
