package priorityqueue

import (
	"fmt"
	"time"
)

type Node struct {
	Value      int
	Priority   int
	InsertedAt time.Time
}

type PriorityQueue struct {
	Queue []Node
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func (pq *PriorityQueue) swap(index1 int, index2 int) {
	temp := pq.Queue[index1]
	pq.Queue[index1] = pq.Queue[index2]
	pq.Queue[index2] = temp
}

func (pq *PriorityQueue) Enqueue(new int, priority int) {

	newNode := Node{
		Value:      new,
		Priority:   priority,
		InsertedAt: time.Now().UTC(),
	}
	pq.Queue = append(pq.Queue, newNode)

	index := len(pq.Queue) - 1
	element := pq.Queue[index]

	for {
		parentIndex := (index - 1) / 2
		parent := pq.Queue[parentIndex]
		if element.Priority >= parent.Priority {
			break
		}
		pq.Queue[parentIndex] = element
		pq.Queue[index] = parent
		index = parentIndex
	}

	fmt.Println("Inserted '", new, "' / Heap ", pq.Queue)

}

func (pq *PriorityQueue) sinkDown(i int) {
	length := len(pq.Queue)
	element := pq.Queue[0]
	var leftChild, rightChild int

	for {
		swap := -1
		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		if leftChildIndex < length {
			leftChild = pq.Queue[leftChildIndex].Priority
			if leftChild < element.Priority {
				swap = leftChildIndex
			}
		}

		if rightChildIndex < length {
			rightChild = pq.Queue[rightChildIndex].Priority
			if (swap == -1 && rightChild < element.Priority) || (swap != -1 && rightChild < leftChild) {
				swap = rightChildIndex
			}
		}
		if swap == -1 {
			break
		}

		pq.Queue[i] = pq.Queue[swap]
		pq.Queue[swap] = element
		i = swap
	}

}

func (pq *PriorityQueue) Dequeue() int {
	lastIndex := len(pq.Queue) - 1

	pq.swap(0, lastIndex)
	max := pq.Queue[len(pq.Queue)-1]
	if len(pq.Queue) < 2 {
		return max.Value
	}
	pq.Queue = append(pq.Queue[:lastIndex])

	pq.sinkDown(0)
	return max.Value
}
