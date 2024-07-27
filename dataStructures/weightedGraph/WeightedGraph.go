package weightedgraph

import (
	priorityqueue "data-structures-algorithms/dataStructures/priorityQueue"
	"fmt"
	"math"
)

type Edge struct {
	Value  string
	Weight int
}

type Graph struct {
	AdjacencyList map[string][]Edge
}

func New() *Graph {
	return &Graph{AdjacencyList: make(map[string][]Edge)}
}

func (g *Graph) AddNode(node string) {
	if _, exists := g.AdjacencyList[node]; !exists {
		g.AdjacencyList[node] = []Edge{}
	}
}

func (g *Graph) AddEdge(node1, node2 string, weight int) {
	g.AdjacencyList[node1] = append(g.AdjacencyList[node1], Edge{Value: node2, Weight: weight})
	g.AdjacencyList[node2] = append(g.AdjacencyList[node2], Edge{Value: node1, Weight: weight})
}

func (g *Graph) Print() {
	for node, edges := range g.AdjacencyList {
		fmt.Printf("%s -> ", node)
		for _, edge := range edges {
			fmt.Printf("%s(%d) ", edge.Value, edge.Weight)
		}
		fmt.Println()
	}
}

// Dijsktra's Algorithm
func (g *Graph) ShortestPath(start, finish string) []string {
	type Distance map[string][]Edge

	nodes := priorityqueue.NewPriorityQueue()
	distances := make(map[string]int)
	previous := make(map[string]string)
	path := []string{}

	for vertex := range g.AdjacencyList {
		if vertex == start {
			distances[vertex] = 0
			nodes.Enqueue(vertex, 0)
		} else {
			distances[vertex] = math.MaxInt
			nodes.Enqueue(vertex, math.MaxInt)
		}

		previous[vertex] = ""

	}

	for nodes.Length > 0 {
		smallest := nodes.Dequeue()
		if smallest == finish {
			// done
			for previous[smallest] != "" {
				path = append(path, smallest)
				smallest = previous[smallest]
			}
			break
		}
		if smallest != "" || distances[smallest] != math.MaxInt {
			for neighbour := range g.AdjacencyList[smallest] {
				nextNode := g.AdjacencyList[smallest][neighbour]

				// calc distance to neighbouring node
				candidate := distances[smallest] + nextNode.Weight

				if candidate < distances[nextNode.Value] {
					// updating new smallest distance to neighbour
					distances[nextNode.Value] = candidate

					// updating previous, how we got to neighbour (nextNode)
					previous[nextNode.Value] = smallest

					// enqueue in priority queue
					nodes.Enqueue(nextNode.Value, candidate)
				}

			}
		}
	}

	// Reverse the path to start from the source
	if len(path) > 0 || start == finish {
		path = append(path, start)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
