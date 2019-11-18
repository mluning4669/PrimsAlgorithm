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

	mst := prim(graph)

	fmt.Println(mst.Idict)
	fmt.Println(mst.AdjList)

	// for i, v := range mst.AdjList {
	// 	fmt.Print(mst.Idict[i])
	// 	fmt.Print(" : ")
	// 	fmt.Print(mst.Idict[v.Parent.Val])
	// 	fmt.Println()
	// }
}

func prim(graph *graphs.Graph) *graphs.Graph {
	heap := binaryheap.StartHeap(graph.VertCount)
	mst := graphs.NewGraph(false, true)

	//populate heap with nodes from graph
	for i, v := range graph.AdjList {

		if i != 0 {
			v.Head.AttCost = math.MaxFloat64
		} else {
			v.Head.AttCost = 0.0
		}

		heap.Insert(v.Head)
	}

	for _, v := range graph.AdjList { //going to break out of the loop if heap size drops to zero. This is more convenient for traversing the original graph
		if heap.Size == 0 {
			break
		}
		var min = heap.ExtractMin()
		mst.InsertVertex(graph.Idict[min.Val]) //insert min into MST

		v.Current = v.Head

		for v.Current != nil {
			index := heap.Dict[v.Current]
			_, e := mst.Idict[v.Current.Val] //don't want to change the attachment costs of nodes already in the MST

			if *v.Current.Weight < heap.Arr[index].AttCost && !e {
				fmt.Println("heap.Dict: ", heap.Dict)
				fmt.Println("&v.Current: ", &v.Current)
				fmt.Println("index: ", index)
				fmt.Println("&heap.Arr[index]: ", &heap.Arr[index])
				heap.ChangeKey(v.Current, *v.Current.Weight)
				v.Parent = min
			}

			v.Current = v.Current.Next
		}

		v.Current = v.Head //reset current to head
	}

	return mst //returning graph to get rid of red lines until I finish algorithm
}

// func testHeap(testVals []float64) {
// 	heap := binaryheap.StartHeap(len(testVals))

// 	values := [9]float64{1.0, 5.0, 6.0, 9.0, 11.0, 8.0, 15.0, 17.0, 21.0}

// 	for i := range values {
// 		heap.Insert(&graphs.Node{Val: 0, AttCost: values[i]})
// 	}

// 	for i, v := range heap.Arr {
// 		if i > 0 {
// 			fmt.Print(v.AttCost, " ")
// 		}

// 	}

// 	fmt.Println()
// 	newWeight := 16.0
// 	heap.ChangeKey(heap.Arr[2], newWeight)

// 	for i := 1; i <= heap.Size; i++ {
// 		fmt.Print(heap.Arr[i].AttCost, " ")
// 	}

// 	fmt.Println()
// }
