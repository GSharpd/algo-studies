package main

import "fmt"

// Every collection of elements with paired links, even though one element might or not be connected to more than one other element
// for an example, linked list, trees, heaps, tries are all forms of graphs
// Graphs can be directed or undirected -> nodes or vertexes can only flow both ways or in the specified flow direction
// Can be weighted or unweighted -> edges can have a weight or not
// Cyclic or Acyclic -> edges can be connected to each other or not, forming a cycle if connected

// Ways of building a graph

//     2 --- 0
//	  / \
//   1 - 3

// Edge List
// graph := [][]int{ [0,2], [2,3], [2,1], [1,3] }

// Adjacency List
// graph2 := [][]int{ [2], [2,3], [0,1,3], [1,2] }

// Adjacency Matrix
// graph3 := [][]int{ [0,0,1,0], [0,0,1,1], [1,1,0,1], [0,1,1,0] }

type Graph struct {
	numberOfNodes int
	adjacencyList [][]int
}

func (g *Graph) AddEdge(from, to int) {
	g.adjacencyList[from] = append(g.adjacencyList[from], to)
	g.adjacencyList[to] = append(g.adjacencyList[to], from)
}

func (g *Graph) AddVertex() {
	g.numberOfNodes++
	g.adjacencyList = append(g.adjacencyList, []int{})
}

func (g *Graph) ShowConnections() {
	for i := 0; i < g.numberOfNodes; i++ {
		fmt.Printf("%d --> ", i)
		for _, value := range g.adjacencyList[i] {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	fmt.Println(g.numberOfNodes)
}

func main() {
	graph := Graph{}
	graph.AddVertex()
	graph.AddVertex()
	graph.AddVertex()
	graph.AddVertex()
	graph.adjacencyList = make([][]int, graph.numberOfNodes)
	graph.AddEdge(0, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 1)
	graph.AddEdge(1, 3)
	graph.ShowConnections()
}
