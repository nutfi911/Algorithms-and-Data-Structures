package DoublyLinkedList

import (
	"fmt"
)

type Node struct {
	Prev  *Node
	Next  *Node
	Value int
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func PrintList(list *LinkedList) {
	if list.Length == 0 {
		fmt.Println("Empty list")
		return
	}
	current := list.Head
	for current != nil {
		fmt.Print(" <- ", current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func New() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Push(n int) {
	new := &Node{
		Value: n,
	}

	l.Length++

	if l.Head == nil {
		new.Next = nil
		new.Prev = nil
		l.Head = new
		l.Tail = l.Head
		return
	}

	new.Prev = l.Tail
	l.Tail.Next = new
	l.Tail = new
}

// H                 T
// 1 <-> 2 <-> 5 <-> 3
// To become:
// H           T
// 1 <-> 2 <-> 5
func (l *LinkedList) Pop() {
	if l.Length == 0 {
		return
	}

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return
	}

	previousNode := l.Tail.Prev
	l.Tail.Prev = nil
	previousNode.Next = nil

	l.Tail = previousNode
	l.Length--
}

// H                 T
// 1 <-> 2 <-> 5 <-> 3
// To become:
// H           T
// 2 <-> 5 <-> 3
func (l *LinkedList) Shift() {
	if l.Length == 0 {
		return
	}

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return
	}

	nextNode := l.Head.Next
	l.Head.Next = nil

	nextNode.Prev = nil
	l.Head = nextNode
	l.Length--
}

func (l *LinkedList) Unshift(new int) {
	newNode := &Node{
		Value: new,
	}

	if l.Length == 0 {
		newNode.Prev = nil
		newNode.Next = nil
		l.Head = newNode
		l.Length++
		return
	}

	newNode.Prev = nil
	newNode.Next = l.Head

	l.Head.Prev = newNode
	l.Head = newNode
	l.Length++
}

func (l *LinkedList) Get(index int) *Node {
	if index < 0 || index > l.Length-1 {
		return nil
	}

	if index == l.Length-1 {
		return l.Tail
	}

	if index == 0 {
		return l.Head
	}

	if index > l.Length/2 {
		counter := l.Length - 1
		current := l.Tail

		for counter > index {
			current = current.Prev
			counter--
		}

		return current

	} else {
		counter := 0
		current := l.Head

		for counter < index {
			current = current.Next
			counter++
		}
		return current
	}
}

func (l *LinkedList) Set(index int, value int) {
	if index < 0 || index > l.Length-1 {
		return
	}

	l.Get(index).Value = value
}

func (l *LinkedList) Insert(index int, value int) {
	if index < 0 || index > l.Length-1 {
		return
	}

	currentNode := l.Get(index)

	newNode := &Node{
		Value: value,
		Prev:  currentNode.Prev,
		Next:  currentNode,
	}

	currentNode.Prev.Next = newNode
	currentNode.Prev = newNode

	l.Length++
}

func (l *LinkedList) Remove(index int) {
	if index < 0 || index > l.Length-1 {
		return
	}

	if index == 0 {
		l.Shift()
		return
	}

	if index == l.Length-1 {
		l.Pop()
		return
	}

	node := l.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

	l.Length--

}

func (l *LinkedList) Reverse() {
	currentNode := l.Head

	for currentNode != l.Tail {
		currentPrev := currentNode.Prev
		currentNode.Prev = currentNode.Next
		currentNode.Next = currentPrev

		currentNode = currentNode.Prev
	}

	currentNode.Next = currentNode.Prev
	currentNode.Prev = nil

	l.Tail = l.Head
	l.Head = currentNode

}
