package graph

import "fmt"

func Run() {

	g := New()

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	g.AddEdge("E", "F")
	/*
				A
			  /	  \
		    B		C
			|		|
			D  ---	E
			  \   /
			  	F

		DFS: A B D E C F
	*/

	fmt.Printf("g: %v\n", g)

	fmt.Println("DFS start A:", g.DFS("A"))
	fmt.Println("DFS start B:", g.DFS("B"))
	fmt.Println("BFS start A:", g.BFS("A")) // A B

}
