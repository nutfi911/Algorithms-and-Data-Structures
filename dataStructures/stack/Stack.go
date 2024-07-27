package Stack

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	First *Node
	Last  *Node
	Size  int
}

func New() *Stack {
	return &Stack{}
}

func (this *Stack) Print() {
	fmt.Println("First: ", this.First.Value, "Last: ", this.Last.Value, "Size: ", this.Size)

	fmt.Print("End -> ")
	current := this.First
	for i := 0; i < this.Size; i++ {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Print("<- Start")
}

// Push = Unshift
// better performance for singly linked list to unshift from the beginning
func (this *Stack) Push(v int) {
	newNode := Node{Value: v}

	if this.First == nil {
		this.First = &newNode
		this.Last = &newNode
	} else {
		temp := this.First
		this.First = &newNode
		this.First.Next = temp
	}
	this.Size++
}

// Pop = Shift
func (this *Stack) Pop() (int, error) {
	if this.Size == 0 {
		return -1, errors.New("Empty Stack")
	}

	shifted := this.First
	if this.First == this.Last {
		this.Last = nil
	}
	this.First = this.First.Next
	this.Size--
	return shifted.Value, nil
}
