package SinglyLinkedList

import (
	"fmt"
)

type Node struct{
	Value int
	Next *Node
}

type LinkedList struct{
	Head *Node
	Tail *Node
	Length int
}

func New() *LinkedList{
	return &LinkedList{}
}

func PrintList(list *LinkedList){
	if list.Length == 0 {
		fmt.Println("Empty list")
		return
	}
	current := list.Head

	for current != nil {
		fmt.Print(current.Value, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func Push(list *LinkedList, value int){
	newNode := Node{
		Value: value,
		Next: nil,
	}

	if list.Length == 0 {
		list.Head = &newNode
		list.Tail = &newNode
		list.Length = 1
		return 
	}

	list.Tail.Next = &newNode
	list.Tail = &newNode
	list.Length += 1
}

func Pop(list *LinkedList) *Node{
	if list.Head == nil{
		return nil
	}

	if list.Head == list.Tail {
		n := list.Head
		list.Head = nil
		list.Tail = nil
		list.Length = 0
		return n
	}

	current := list.Head
	
	for current.Next != list.Tail {
		current = current.Next
	}
	n := current.Next
	list.Tail = current
	list.Tail.Next = nil
	list.Length -=1
	return n
}

func Shift(list *LinkedList)int{
	if list.Head == nil {
		return -1
	}

	if list.Head == list.Tail {
		n := list.Head.Value
		list.Head = nil
		list.Tail = nil
		return n
	}
	
	n := list.Head.Value
	list.Head = list.Head.Next
	list.Length -= 1
	return n
}

func Unshift(list *LinkedList, value int){
	newNode := Node{
		Value: value,
		Next: list.Head,
	}

	list.Head = &newNode
	list.Length+=1
}

func Get(list *LinkedList, index int) *Node{
	if index < 0 || index >= list.Length{
		return nil
	}
	counter := 0
	current := list.Head
	for counter < index {
		current = current.Next
		counter++
	}
	return current
}

func Set(list *LinkedList, index int, value int){
	node := Get(list, index)
	if node != nil{
		node.Value = value
	}
}

func Insert(list *LinkedList, index int, value int)bool{
	if index < 0 || index > list.Length{
		return false
	}

	if index == list.Length{
		Push(list, value)
		return true
	}

	if index == 0 {
		Unshift(list, value)
		return true
	}

	currentNode := Get(list, index)
	newNode := Node{
		Value: value,
		Next: currentNode,
	}
	previousNode := Get(list, index-1)
	previousNode.Next = &newNode

	return true
}

func Remove(list *LinkedList, index int) int{
	if index < 0 || index > list.Length-1{
		return -1
	}
	if index == 0 {
		return Shift(list)
	}

	if index == list.Length -1 {
		return Pop(list).Value
	}

	prevNode := Get(list, index-1)
	removedNode := prevNode.Next
	prevNode.Next = removedNode.Next

	return removedNode.Value
}

func Reverse(list *LinkedList){
	if list.Length <= 1 {
		return
	}
	current := list.Head
	list.Head = list.Tail
	list.Tail = current

	// prev is initially nil
	var prev *Node

	for current != nil {
		next := current.Next
		current.Next = prev

		// Continuing
		prev = current
		current = next
	}
	fmt.Println("Head: ", list.Head.Value, " -> " , list.Head.Next.Value, " Tail: ", list.Tail.Value, " -> ", list.Tail.Next)
}