// https://www.youtube.com/watch?v=bSZ57h7GN2w

package main

import "fmt"

// Graph structure
// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex structure
// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add Vertex
// Adds a Vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		// if graph does not already contain key, add it to graph
		g.vertices = append(g.vertices, &Vertex{key: k}) // vertex is created with key as the id
	}
}

// AddEdge adds an edge to the graph
// this is for a directed graph
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		// checks if edge is invalid
		err := fmt.Errorf("invalid edge (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		// checks if edge already exists
		err := fmt.Errorf("edge already exists (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// Contains checks if map contains a key
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex on the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v :", v.key)
		for _, v := range v.adjacent {
			fmt.Printf("%v", v.key)
		}
	}
	fmt.Println()
}

func main() {
	// create graph called test
	test := &Graph{}

	// add nodes
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	//fmt.Println(test) // prints out addresses
	test.AddEdge(1, 2)
	test.AddEdge(6, 2)
	test.AddEdge(3, 2)
	test.AddEdge(3, 2)

	test.Print()
}
