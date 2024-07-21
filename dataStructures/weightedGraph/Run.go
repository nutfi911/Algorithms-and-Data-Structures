package weightedgraph

import "fmt"

func Run() {
	wg := New()

	wg.AddNode("A")
	wg.AddNode("B")
	wg.AddNode("C")
	wg.AddNode("D")
	wg.AddNode("E")
	wg.AddNode("F")

	wg.AddEdge("A", "B", 4)
	wg.AddEdge("A", "C", 2)
	wg.AddEdge("B", "E", 3)
	wg.AddEdge("C", "D", 2)
	wg.AddEdge("C", "F", 4)
	wg.AddEdge("D", "E", 3)
	wg.AddEdge("D", "F", 1)
	wg.AddEdge("E", "F", 1)

	// Shortest path A -> E = A C D F E

	wg.Print()
	sp := wg.ShortestPath("A", "E")
	fmt.Println("Shortest path", sp)
}
