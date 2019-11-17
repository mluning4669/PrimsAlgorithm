package main

import (
	"PrimsAlgorithm/binaryheap"
	"PrimsAlgorithm/graphs"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Prim's Algorithm")
	graph := graphs.ReadFile("Prim_000.gl")

	graphs.PrintGraph(graph)

	testVals := []float64{1.0, 5.0, 6.0, 9.0, 11.0, 8.0, 15.0, 17.0, 21.0}

	testHeap(testVals)

	mst := prim(graph)

	graphs.PrintGraph(mst)
}

func prim(graph *graphs.Graph) *graphs.Graph {
	heap := binaryheap.StartHeap(graph.VertCount)
	var mst *graphs.Graph

	//populate heap with nodes from graph
	for i, v := range graph.AdjList {
		if i != 0 {
			*v.Head.AttCost = math.MaxFloat64
		} else {
			*v.Head.AttCost = 0
		}
		heap.Insert(v.Head)
	}

	var min *graphs.Node
	for i, v := range graph.AdjList { //going to break out of the loop if heap size drops to zero. This is more convenient for traversing the original graph
		min = heap.ExtractMin()
		mst.InsertVertex(min.Val) //insert min into MST

		v.Current = v.Current.Next //compute attachment cost for each neighbor to the min node
		for v.Current != v.Tail {
			if v.Current.Weight < v.Current.AttCost {
				v.Current.AttCost = v.Current.Weight
				v.Current.Parent = min
			}

			v.Current = v.Current.Next
		}
	}

	return graph //returning graph to get rid of red lines until I finish algorithm
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
