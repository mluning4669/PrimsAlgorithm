package main

import (
	"PrimsAlgorithm/binaryheap"
	"PrimsAlgorithm/graphs"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Prim's Algorithm")
	graph := graphs.ReadFile("Prim_001.gl")

	mst := prim(graph)

	mstGraphToFile(mst)
}

func mstGraphToFile(graph *graphs.Graph) {
	f, err := os.Create("mst.gl")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	vertDict := make(map[string]bool)

	fmt.Fprintln(f, "undirected weighted")
	for i, v := range graph.AdjList {
		v.Current = v.Head

		for v.Current != nil {
			if !vertDict[graph.Idict[v.Current.Val]] {
				fmt.Fprint(f, graph.Idict[i], "=")
				fmt.Fprint(f, graph.Idict[v.Current.Val], "=")
				fmt.Fprintln(f, *v.Current.Weight)
			}
			v.Current = v.Current.Next
		}

		vertDict[graph.Idict[i]] = true
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file written successfully")
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
			_, e := mstGraph.Dict[graph.Idict[v.Current.Val]] //don't want to change the attachment costs of nodes already in the MST

			if *v.Current.Weight < heap.Arr[index].AttCost && !e {
				heap.ChangeKey(graph.Idict[v.Current.Val], *v.Current.Weight, min)
			}

			v.Current = v.Current.Next
		}
		min = heap.ExtractMin()
		mstGraph.InsertVertex(min.HeapLabel)
		mstGraph.InsertEdge(min.Parent.HeapLabel, min.HeapLabel, &min.AttCost)
	}

	return &mstGraph
}
