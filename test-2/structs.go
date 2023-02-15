// structs and variables
package graphs

import (
	"fmt"
	"strconv"
	"strings"
)

// struct concerning the ants and the route they must take
/*
number of ants
start
end
room names
room links
*/
type Ants struct {
	NumAnts   int
	StartRoom string
	EndRoom   string
	//RoomName  []string
	//Neighbour []string
}

// struct concerning the graph itself
type Graph struct {
	Vertices []*Vertex
}

// Vertex structure
// Vertex represents a graph vertex
type Vertex struct {
	Name   string
	Xcoord int
	Ycoord int
	Edges  []*Vertex
	Key    int
}

// Creates Vertices
// takes rooms and coordinates
func CreateRooms(data, links []string) (vertices []*Vertex, mapVert map[string]*Vertex) {
	mapVert = map[string]*Vertex{}
	for i := 0; i < len(data); i++ {
		split := strings.Split(data[i], " ")
		name := split[0]
		x, _ := strconv.Atoi(string(split[1]))
		y, _ := strconv.Atoi(string(split[2]))
		vertex := Vertex{
			Name:   name,
			Xcoord: x,
			Ycoord: y,
		}
		mapVert[name] = &vertex
		vertices = append(vertices, &vertex)
	}
	return vertices, mapVert
}

// Add Vertex
// Adds a Vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.Vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		// if graph does not already contain key, add it to graph
		g.Vertices = append(g.Vertices, &Vertex{Key: k}) // vertex is created with key as the id
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
	} else if contains(fromVertex.Edges, to) {
		// checks if edge already exists
		err := fmt.Errorf("edge already exists (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.Edges = append(fromVertex.Edges, toVertex)
	}
}

// getVertex returns a pointer to the Vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.Vertices {
		if v.Key == k {
			return g.Vertices[i]
		}
	}
	return nil
}

// Contains checks if map contains a key
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex on the graph
func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\nVertex %v: ", v.Key)
		for _, v := range v.Edges {
			fmt.Printf("%v", v.Key)
		}
	}
	fmt.Println()
}

// func executeGraph() {
// 	// create graph called test
// 	test := &Graph{}

// 	// add nodes
// 	for i := 0; i < 5; i++ {
// 		test.AddVertex(i)
// 	}

// 	//fmt.Println(test) // prints out addresses
// 	test.AddEdge(1, 2)
// 	test.AddEdge(6, 2)
// 	test.AddEdge(3, 2)
// 	test.AddEdge(3, 2)

// 	test.Print()
// }
