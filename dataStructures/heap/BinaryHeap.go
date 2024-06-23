package heap

import "fmt"

type Heap struct {
	Values []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) parentIndex(i int) (index int) {
	return (i - 1) / 2 // floored value
}

func (h *Heap) leftChildIndex(i int) (index int) {
	return (2 * i) + 1
}

func (h *Heap) rightChildIndex(i int) (index int) {
	return (2 * i) + 2
}

func (h *Heap) swap(index1 int, index2 int) {
	temp := h.Values[index1]
	h.Values[index1] = h.Values[index2]
	h.Values[index2] = temp
}

func (h *Heap) Insert(new int) {

	h.Values = append(h.Values, new)

	index := len(h.Values) - 1
	element := h.Values[index]

	for {
		parentIndex := h.parentIndex(index)
		parent := h.Values[parentIndex]
		if element <= parent {
			break
		}
		h.Values[parentIndex] = element
		h.Values[index] = parent
		index = parentIndex
	}

	fmt.Println("Inserted '", new, "' / Heap ", h.Values)

}

func (h *Heap) sinkDown(i int) {
	length := len(h.Values)
	element := h.Values[0]
	var leftChild, rightChild int

	for {
		swap := -1
		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		if leftChildIndex < length {
			leftChild = h.Values[leftChildIndex]
			if leftChild > element {
				swap = leftChildIndex
			}
		}

		if rightChildIndex < length {
			rightChild = h.Values[rightChildIndex]
			if (swap == -1 && rightChild > element) || (swap != -1 && rightChild > leftChild) {
				swap = rightChildIndex
			}
		}
		if swap == -1 {
			break
		}

		h.Values[i] = h.Values[swap]
		h.Values[swap] = element
		i = swap
	}

}

func (h *Heap) ExtractMax() int {
	lastIndex := len(h.Values) - 1

	h.swap(0, lastIndex)
	max := h.Values[len(h.Values)-1]
	if len(h.Values) < 2 {
		return max
	}
	h.Values = append(h.Values[:lastIndex])

	h.sinkDown(0)
	return max
}
