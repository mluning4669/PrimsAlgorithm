package main

import (
	"Algorithms/Assignment_2/Golang/binaryheap"
	"Algorithms/Assignment_2/Golang/graphs"
	"fmt"
)

func main() {
	fmt.Println("Prim's Algorithm")
	graph := graphs.ReadFile("Prim_000.gl")

	graphs.PrintGraph(graph)

	for _, v := range graph.AdjList {
		fmt.Print(*v.Head.Weight, " ")
	}
	fmt.Println()

	heap := binaryheap.StartHeap(len(graph.AdjList))

	for _, v := range graph.AdjList {
		heap.Insert(v.Head)
	}

	for _, v := range heap.Arr {
		fmt.Print(*v.Weight, " ")
	}

	fmt.Println()
}
