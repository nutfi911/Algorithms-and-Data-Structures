package heap

import "fmt"

func Run() {

	heap := NewHeap()

	heap.Insert(41)
	heap.Insert(39)
	heap.Insert(33)
	heap.Insert(18)
	heap.Insert(27)
	heap.Insert(12)
	heap.Insert(55)

	fmt.Println("Extracted Max: ", heap.ExtractMax(), " / Heap ", heap.Values)
	fmt.Println("Extracted Max: ", heap.ExtractMax(), " / Heap ", heap.Values)
}
