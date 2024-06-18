package Queue

import "fmt"

func Run() {
	q := New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.First.Value, q.First.Next.Value, q.First.Next.Next.Value)

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())

	fmt.Println("Left in Q", q.First.Value)
}
