package graph

type Graph struct {
	AdjacencyList map[string][]string
}

func New() *Graph {
	return &Graph{AdjacencyList: make(map[string][]string)}
}

func (v *Graph) AddVertex(vertex string) {
	if v.AdjacencyList[vertex] == nil {
		v.AdjacencyList[vertex] = []string{}
	}
}

func (v *Graph) AddEdge(vertex1, vertex2 string) {
	v.AdjacencyList[vertex1] = append(v.AdjacencyList[vertex1], vertex2)
	v.AdjacencyList[vertex2] = append(v.AdjacencyList[vertex2], vertex1)
}

func (v *Graph) RemoveEdge(vertex1, vertex2 string) {
	if neighbors, exists := v.AdjacencyList[vertex1]; exists {
		v.AdjacencyList[vertex1] = removeElement(neighbors, vertex2)
	}

	if neighbors, exists := v.AdjacencyList[vertex2]; exists {
		v.AdjacencyList[vertex2] = removeElement(neighbors, vertex1)
	}
}

func (v *Graph) RemoveVertex(vertex string) {
	for _, vert := range v.AdjacencyList[vertex] {
		v.RemoveEdge(vertex, vert)
	}

	delete(v.AdjacencyList, vertex)
}

func (v *Graph) DFS(start string) []string {
	var result []string
	visited := make(map[string]bool)

	var dfsHelper func(vertex string)
	dfsHelper = func(vertex string) {
		if vertex == "" {
			return
		}
		visited[vertex] = true
		result = append(result, vertex)
		for _, neighbor := range v.AdjacencyList[vertex] {
			if !visited[neighbor] {
				dfsHelper(neighbor)
			}
		}
	}

	dfsHelper(start)

	return result

}

func (v *Graph) BFS(start string) []string {
	var result []string
	visited := make(map[string]bool)

	queue := []string{start}
	visited[start] = true

	var currentVertex string

	for len(queue) > 0 {
		currentVertex = queue[0]
		queue = queue[1:]

		result = append(result, currentVertex)

		for _, neighbor := range v.AdjacencyList[currentVertex] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}

	}

	return result

}

// Helper function to remove an element from a slice
func removeElement(slice []string, element string) []string {
	newSlice := []string{}
	for _, v := range slice {
		if v != element {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}
