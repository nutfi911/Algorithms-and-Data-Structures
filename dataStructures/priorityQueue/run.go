package priorityqueue

import "fmt"

func Run() {
	pq := NewPriorityQueue()

	pq.Enqueue(1001, 19)
	pq.Enqueue(1002, 10)
	pq.Enqueue(1003, 10)
	pq.Enqueue(1004, 1)
	pq.Enqueue(1005, 5)

	fmt.Println(pq.Dequeue())
}
