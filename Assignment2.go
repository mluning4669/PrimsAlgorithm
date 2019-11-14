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

	for _, v := range graph.AdjList {
		fmt.Print(*v.Head.Weight, " ")
	}
	fmt.Println()

	heap := binaryheap.StartHeap(len(graph.AdjList))

	for _, v := range graph.AdjList {
		heap.Insert(v.Head)
	}

	for i, v := range heap.Arr {
		if i > 0 {
			fmt.Print(*v.Weight, " ")
		}

	}

	fmt.Println()
	heap.ExtractMin()

	for i := 1; i < heap.Capacity-1; i++ {
		fmt.Print(*heap.Arr[i].Weight, " ")
	}

	fmt.Println()
}
