package binaryheap

import (
	"Algorithms/Assignment_2/Golang/graphs"
)

//Heap a heap for implmenting a priority queue
type Heap struct {
	Arr  []*graphs.Node
	size int
	dict map[*graphs.Node]int //position in the text
}

//StartHeap Initializes a heap of size N
func StartHeap(n int) *Heap {
	h := Heap{size: 0}
	h.Arr = make([]*graphs.Node, n)
	h.dict = make(map[*graphs.Node]int)

	return &h
}

//Insert insert an element into the heap
func (h *Heap) Insert(elem *graphs.Node) {
	h.Arr[h.size] = elem    //insert element at index size since size is the length of the array before adding the new element
	h.dict[elem] = h.size   //store element and size in position dictionary
	h.size++                //bump size by 1
	h.heapifyUp(h.size - 1) //put heap in heap order
}

func (h *Heap) parent(i int) int {
	if i%2 == 0 {
		return i / 2
	}

	return (i - 2) / 2
}

func (h *Heap) heapifyUp(i int) {
	if i > 1 {
		j := h.parent(i)

		if *h.Arr[j].Weight < *h.Arr[i].Weight {
			var temp = *h.Arr[j]
			*h.Arr[j] = *h.Arr[i]
			*h.Arr[i] = temp

			h.heapifyUp(j)
		}
	}
}

func (h *Heap) heapifyDown(i int) {
	n := h.size
	var j = 0

	if 2*i > n {
		return
	}
	if 2*i < n {
		left := 2 * i
		right := 2*i + 1

		if *h.Arr[left].Weight <= *h.Arr[right].Weight { //not sure about this
			j = left
		} else {
			j = right
		}

	}
	if 2*i == n {
		j = i
	}

	if *h.Arr[j].Weight < *h.Arr[i].Weight {
		var temp = *h.Arr[j]
		*h.Arr[j] = *h.Arr[i]
		*h.Arr[i] = temp

		h.heapifyDown(j)
	}
}
