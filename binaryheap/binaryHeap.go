package binaryheap

import (
	"PrimsAlgorithm/graphs"
	"errors"
)

//Heap a heap for implmenting a priority queue
type Heap struct {
	Arr      []*graphs.Node
	Size     int
	Capacity int
	Dict     map[string]int //position in the text
}

//StartHeap Initializes a heap of Size N
func StartHeap(n int) *Heap {
	h := Heap{Size: 0, Capacity: n + 1} //using 1 based indexing to make the math more managable
	h.Arr = make([]*graphs.Node, n+1)
	h.Dict = make(map[string]int)

	return &h
}

//ExtractMin removes minimum element of heap
func (h *Heap) ExtractMin() *graphs.Node {
	min := h.Arr[1]
	h.Arr[1] = h.Arr[h.Size]
	h.Arr[h.Size] = nil
	h.heapifyDown(1)
	h.Size--
	return min
}

//Delete removes element at location i in heap array. Also acts as Delete(elem) when you pass Delete(dict[elem])
func (h *Heap) Delete(i int) {
	h.Arr[i] = h.Arr[h.Size]
	h.Arr[h.Size] = nil
	h.heapifyDown(i)
	h.Size--
}

//Insert insert an element into the heap
func (h *Heap) Insert(elem *graphs.Node) error {
	if h.Size+1 >= h.Capacity {
		return errors.New("Heap at capacity")
	}

	h.Size++
	h.Arr[h.Size] = elem            //insert element at element Size
	h.Dict[elem.HeapLabel] = h.Size //store element and Size in position dictionary
	h.heapifyUp(h.Size)             //put heap in heap order

	return nil
}

//FindMin finds minimum element in heap but doesn't remove it
func (h *Heap) FindMin() *graphs.Node {
	return h.Arr[1]
}

//ChangeKey change the key (in this case attachment cost) of the current element to a new value
func (h *Heap) ChangeKey(currentLabel string, newKey float64) {
	currentIndex := h.Dict[currentLabel]
	h.Arr[currentIndex].AttCost = newKey
	h.heapifyUp(currentIndex)
}

func (h *Heap) swap(i int, j int) {
	h.Dict[h.Arr[i].HeapLabel] = j //update dictionary
	h.Dict[h.Arr[j].HeapLabel] = i

	temp := h.Arr[j]
	h.Arr[j] = h.Arr[i]
	h.Arr[i] = temp
}

func (h *Heap) heapifyUp(i int) {
	if i > 1 {
		j := i / 2 //this works thanks to integer division

		if h.Arr[j].AttCost > h.Arr[i].AttCost {
			h.swap(i, j)

			h.heapifyUp(j)
		}
	}
}

func (h *Heap) heapifyDown(i int) {
	var j = i

	left := 2 * i
	right := 2*i + 1

	if left < h.Size && h.Arr[j].AttCost > h.Arr[left].AttCost {
		j = left
	}
	if right < h.Size && h.Arr[j].AttCost > h.Arr[right].AttCost {
		j = right
	}

	if j != i {
		h.swap(i, j)

		h.heapifyDown(j)
	}
}
