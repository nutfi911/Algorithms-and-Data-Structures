package Queue

import (
	"errors"
)

type Node struct {
	Value int
	Next  *Node
}

type Queue struct {
	First *Node
	Last  *Node
	Size  int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(v int) {
	newNode := Node{Value: v, Next: nil}
	if q.Size == 0 {
		q.First = &newNode
		q.Last = &newNode
		q.Size++
		return
	}

	q.Last.Next = &newNode
	q.Last = &newNode
	q.Size++
}

func (q *Queue) Dequeue() (int, error) {
	if q.Size == 0 {
		return -1, errors.New("Empty Queue")
	}

	temp := q.First
	if q.First == q.Last {
		q.Last = nil
	}
	q.First = q.First.Next
	q.Size--
	return temp.Value, nil
}
