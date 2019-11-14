package binaryheap

import (
	"PrimsAlgorithm/graphs"
)

//Heap a heap for implmenting a priority queue
type Heap struct {
	Arr      []*graphs.Node
	size     int
	Capacity int
	dict     map[*graphs.Node]int //position in the text
}

//StartHeap Initializes a heap of size N
func StartHeap(n int) *Heap {
	h := Heap{size: 0, Capacity: n + 1} //using 1 based indexing to make the math more managable
	h.Arr = make([]*graphs.Node, n+1)
	h.dict = make(map[*graphs.Node]int)

	return &h
}

//Insert insert an element into the heap
func (h *Heap) Insert(elem *graphs.Node) {
	//check against capacity here
	h.size++
	h.Arr[h.size] = elem  //insert element at element size
	h.dict[elem] = h.size //store element and size in position dictionary              //bump size by 1
	h.heapifyUp(h.size)   //put heap in heap order
}

func (h *Heap) heapifyUp(i int) {
	if i > 1 {
		j := i / 2 //this works thanks to integer division

		if *h.Arr[j].Weight > *h.Arr[i].Weight { //will also need to swap dicitonary elements
			temp := *h.Arr[j]
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

	if *h.Arr[j].Weight < *h.Arr[i].Weight { //will need to update dictionary
		temp := *h.Arr[j]
		*h.Arr[j] = *h.Arr[i]
		*h.Arr[i] = temp

		h.heapifyDown(j)
	}
}
