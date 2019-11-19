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

	prim(graph)

}

func prim(graph *graphs.Graph) {
	heap := binaryheap.StartHeap(graph.VertCount)
	mst := make(map[string]bool)

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

	for _, v := range graph.AdjList { //going to break out of the loop if heap size drops to zero. This is more convenient for traversing the original graph
		if heap.Size == 0 {
			break
		}
		var min = heap.ExtractMin()
		mst[min.HeapLabel] = true //min.Val will be one more than what it should be in the graph so subtracting 1 to compensate

		v.Current = v.Head

		for v.Current != nil {
			index := heap.Dict[graph.Idict[v.Current.Val]]
			e := mst[graph.Idict[v.Current.Val]] //don't want to change the attachment costs of nodes already in the MST
			heapAtIndex := heap.Arr[index]

			if *v.Current.Weight < heapAtIndex.AttCost && !e {
				heap.ChangeKey(graph.Idict[v.Current.Val], *v.Current.Weight)
				graph.AdjList[v.Current.Val].Parent = min
			}

			v.Current = v.Current.Next
		}
		for k, v := range mst {
			if v {
				index := graph.Dict[k]
				parent := graph.AdjList[index].Parent
				if parent != nil {
					fmt.Print(parent.HeapLabel, ": ")
				} else {
					fmt.Print("П: ")
				}
				fmt.Print(k)
			}
			fmt.Println()
		}
		fmt.Println("BREAK")
	}

}
