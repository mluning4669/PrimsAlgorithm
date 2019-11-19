package main

import (
	"PrimsAlgorithm/binaryheap"
	"PrimsAlgorithm/graphs"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Prim's Algorithm")
	graph := graphs.ReadFile("Prim_002.gl")

	graphs.PrintGraph(graph)

	fmt.Println()

	mst := prim(graph)

	graphToFile(mst)
}

func graphToFile(graph *graphs.Graph) {
	vertDict := make(map[string]bool)

	fmt.Println("undirected weighted")
	for i, v := range graph.AdjList {
		v.Current = v.Head

		for v.Current != nil {
			if !vertDict[graph.Idict[v.Current.Val]] {
				fmt.Print(graph.Idict[i], "=")
				fmt.Print(graph.Idict[v.Current.Val], "=")
				fmt.Println(*v.Current.Weight)
			}
			v.Current = v.Current.Next
		}

		vertDict[graph.Idict[i]] = true
	}
}

func prim(graph *graphs.Graph) *graphs.Graph {
	heap := binaryheap.StartHeap(graph.VertCount)
	mstGraph := *graphs.NewGraph(false, true)

	//populate heap with nodes from graph
	for i := range graph.AdjList {
		var newNode *graphs.Node

		if i != 0 {
			newNode = &graphs.Node{HeapLabel: graph.Idict[i], AttCost: math.MaxFloat64, Val: i + 1}
		} else {
			newNode = &graphs.Node{HeapLabel: graph.Idict[i], AttCost: 0.0, Val: i + 1}
		}

		heap.Insert(newNode)
	}

	var min = heap.ExtractMin()
	mstGraph.InsertVertex(min.HeapLabel)

	for heap.Size > 0 {
		var v = graph.AdjList[graph.Dict[min.HeapLabel]]

		v.Current = v.Head

		for v.Current != nil { //
			index := heap.Dict[graph.Idict[v.Current.Val]]
			_, e := mstGraph.Idict[v.Current.Val] //don't want to change the attachment costs of nodes already in the MST

			if *v.Current.Weight < heap.Arr[index].AttCost && !e {
				heap.ChangeKey(graph.Idict[v.Current.Val], *v.Current.Weight, min)
			}

			v.Current = v.Current.Next
		}
		min = heap.ExtractMin()
		mstGraph.InsertVertex(min.HeapLabel)
		mstGraph.InsertEdge(min.HeapLabel, min.Parent.HeapLabel, &min.AttCost)
	}

	return &mstGraph
}
