package main

import (
	"PrimsAlgorithm/binaryheap"
	"PrimsAlgorithm/graphs"
	"fmt"
)

func main() {
	fmt.Println("Prim's Algorithm")
	graph := graphs.ReadFile("test.gl")

	graphs.PrintGraph(graph)

	testVals := []float64{1.0, 5.0, 6.0, 9.0, 11.0, 8.0, 15.0, 17.0, 21.0}

	testHeap(testVals)
}

func testHeap(testVals []float64) {
	heap := binaryheap.StartHeap(len(testVals))

	values := [9]float64{1.0, 5.0, 6.0, 9.0, 11.0, 8.0, 15.0, 17.0, 21.0}

	for i := range values {
		heap.Insert(&graphs.Node{Val: 0, Weight: &values[i]})
	}

	for i, v := range heap.Arr {
		if i > 0 {
			fmt.Print(*v.Weight, " ")
		}

	}

	fmt.Println()
	newWeight := 16.0
	heap.ChangeKey(heap.Arr[2], &graphs.Node{Val: 0, Weight: &newWeight})

	for i := 1; i <= heap.Size; i++ {
		fmt.Print(*heap.Arr[i].Weight, " ")
	}

	fmt.Println()
}
