// structs and variables
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// struct concerning the ants and the route they must take
/*
number of ants
start
end
*/
type Ants struct {
	NumAnts   int
	StartRoom string
	EndRoom   string
}

// struct concerning the graph itself
// represents overall graph structure and manage relationships between vertices
// eg: storing list of vertices, managing edges between vertices, graph algos
type Graph struct {
	Vertices []*Vertex
}

// Vertex structure
// represents an individual vertex with all its own properties
// eg: key, value, its edges etc.
type Vertex struct {
	Name   string
	Xcoord int
	Ycoord int
	Edges  []*Vertex
	Key    int
}

// open file from terminal
func openFile() []string {
	// example input: go run . test01.txt
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}

	strings := strings.Split(string(file), "")
	return strings
}

// returns number of ants
func antCount(data []string) int {
	data = openFile()
	antsB := data[0]
	ants, _ := strconv.Atoi(string(antsB))
	return ants
}

// reads text file and add rooms to a slice
// no ant #, no comments, no links
func openRooms(data []string) []string {
	data = openFile()
	scanner := bufio.NewScanner(strings.NewReader(strings.Join(data, "")))
	var slice []string
	for scanner.Scan() {
		line := scanner.Text()
		if line != data[0] && !strings.Contains(line, "-") && !strings.Contains(line, "##") {
			slice = append(slice, line)
		}
	}
	return slice
}

// reads text file for edges, adds them to a slice
func openLinks(data []string) []string {
	data = openFile()
	scanner := bufio.NewScanner(strings.NewReader(strings.Join(data, "")))
	var slice []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			slice = append(slice, line)
		}
	}
	return slice
}

// Creates Vertices
// takes rooms and coordinates
// NOTE: data reads entire text file; we need to remove # of ants, links, and any comments (##start/##end)
func makeVertices(data, links []string) (vertices []*Vertex, mapVert map[string]*Vertex) {
	data = openFile()
	rooms := openRooms(data)

	mapVert = map[string]*Vertex{}
	for i := 0; i < len(rooms); i++ {
		split := strings.Split(rooms[i], " ")
		name := split[0]
		x, _ := strconv.Atoi(string(split[1]))
		y, _ := strconv.Atoi(string(split[2]))
		vertex := Vertex{
			Name:   name,
			Xcoord: x,
			Ycoord: y,
		}
		mapVert[name] = &vertex //returns memory address of vertex
		vertices = append(vertices, &vertex)
	}
	return vertices, mapVert
	// vertices = slice of pointers representing all vertices in the graph
	// mapvert = maps vertex names to pointers to vertices
}

func addVertices(mapVert map[string][]*Vertex, vertices []*Vertex) {
	for i := 0; i < len(vertices); i++ {
		vertex := vertices[i]
		mapVert[vertex.Name] = []*Vertex{}
	}
}

func (f *Graph) checkRoom(current string) bool {
	if contains(f.Vertices) {
		err := fmt.Errorf("room %v not added as room already exists", rname)
		fmt.Println(err.Error())
	}
}

func (f *antFarm) addRoom(rname string) {
	//first place statement to check whether a room in the antFarm has is a certain room with a certain name(
	///check a certain room is in the antFarm under a certian name:
	//if statement is true, means antFarm already has that certain room
	if contains(f.Rooms, rname) {
		err := fmt.Errorf("room %v not added as room already exists", rname)
		fmt.Println(err.Error())
	} else {
		//create a room that has k as the name --> "&room{name: k}"
		//append k to the rooms list in the antFarm (Rooms field) --> f.Rooms = append(f.Rooms, &room{name: k})
		f.Rooms = append(f.Rooms, &room{Name: rname})
	}
}

// obtain edges
func getEdges(data []string) []string {
	scanner := bufio.NewScanner(strings.NewReader(strings.Join(data, "")))
	var ifEnd bool
	var edges []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "##end" {
			ifEnd = true
			continue
		}
		if strings.Contains(line, "-") && ifEnd && len(line) > 0 {
			edges = append(edges, line)
		}
	}
	return edges
}

// add edges to graph
func addEdges(data []string) {

}

// --------------------------------------------------------------------------------------------------

// func main() {
// 	data := []string{"A 1 2", "B 3 4", "C 5 6"}
// 	links := []string{}
// 	vertices, mapVert := makeVertices(data, links)

// 	// Initialize the graph
// 	graph := map[string][]*Vertex{}

// 	// Add vertices to the graph
// 	addVerticesToGraph(graph, vertices)
// }

// Add Vertex
// Adds a Vertex to the graph
func (g *Graph) AddVertex(k int) { //takes pointer to graph object g and int k input
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
