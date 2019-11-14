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

//ExtractMin removes minimum element of heap
func (h *Heap) ExtractMin() *graphs.Node {
	min := h.Arr[1]
	h.Arr[1] = h.Arr[h.size]
	h.Arr[h.size] = nil
	h.heapifyDown(1)
	h.size--
	return min
}

//Delete removes element at location i in heap array
func (h *Heap) Delete(i int) {
	h.Arr[i] = h.Arr[h.size]
	h.Arr[h.size] = nil
	h.heapifyDown(i)
	h.size--
}

//Insert insert an element into the heap
func (h *Heap) Insert(elem *graphs.Node) {
	//TODO: check against capacity here
	h.size++
	h.Arr[h.size] = elem  //insert element at element size
	h.dict[elem] = h.size //store element and size in position dictionary
	h.heapifyUp(h.size)   //put heap in heap order
}

//FindMin finds minimum element in heap but doesn't remove it
func (h *Heap) FindMin() *graphs.Node {
	return h.Arr[1]
}

func (h *Heap) heapifyUp(i int) {
	if i > 1 {
		j := i / 2 //this works thanks to integer division

		if *h.Arr[j].Weight > *h.Arr[i].Weight { //TODO: will also need to swap dicitonary elements
			temp := h.Arr[j]
			h.Arr[j] = h.Arr[i]
			h.Arr[i] = temp

			h.heapifyUp(j)
		}
	}
}

func (h *Heap) heapifyDown(i int) {
	var j = i

	left := 2 * i
	right := 2*i + 1

	if left < h.size && *h.Arr[j].Weight > *h.Arr[left].Weight { //TODO: not sure about this
		j = left
	}
	if right < h.size && *h.Arr[j].Weight > *h.Arr[right].Weight {
		j = right
	}

	if j != i { //TODO: will need to update dictionary
		temp := h.Arr[j]
		h.Arr[j] = h.Arr[i]
		h.Arr[i] = temp

		h.heapifyDown(j)
	}
}
