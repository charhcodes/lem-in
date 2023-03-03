// prints all paths

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	name    string
	links   []*Vertex // adjacent
	id      int
	visited bool
	ants    int
	path    []*Vertex
}

type Graph struct {
	vertices []*Vertex
}

// open file (os.Args[1]) and split into separate lines
func readFile() []string {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// return number of ants
func antCount() int {
	data := readFile()
	ants := data[0]
	if ants <= "0" {
		err := fmt.Errorf("ERROR: invalid number of ants")
		fmt.Println(err.Error())
	}
	count, _ := strconv.Atoi(string(ants))
	return count
}

// returns starting room
func getStart() string {
	var start string
	file := readFile()

	for i, _ := range file {
		if file[i] == "##start" {
			start = strings.Split(string(file[i+1]), " ")[0]
		}
	}
	return start
}

// returns ending room
func getEnd() string {
	var end string
	file := readFile()

	for i, _ := range file {
		if file[i] == "##end" {
			end = strings.Split(string(file[i+1]), " ")[0]
		}
	}
	return end
}

// return a pointer to the Vertex with its name
func (g *Graph) getVertex(name string) *Vertex {
	for i, v := range g.vertices {
		if v.name == name {
			return g.vertices[i]
		}
	}
	return nil
}

var (
	startroom = getStart()
	endroom   = getEnd()
)

// checks if there are repeated rooms
func contains(s []*Vertex, name string) bool {
	for _, v := range s {
		if name == v.name {
			return true
		}
	}
	return false
}

// add vertex to graph
func (g *Graph) AddVertex(name string) { //*Node
	if contains(g.vertices, name) {
		err := fmt.Errorf("Vertex %v not added because it is an existing key", name)
		fmt.Println(err.Error())
	} else {
		vertices := &Vertex{name: name} // creates a pointer to a new Vertex struct, initialises with a field called name
		// and sets its value to the value of the name variable
		g.vertices = append(g.vertices, vertices)
	}
}

// add edges to graph
func (g *Graph) AddEdge(from, to string) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil { // if edges are valid
		err := fmt.Errorf("ERROR: invalid edges")
		fmt.Println(err.Error())
	} else if contains(fromVertex.links, to) { // if vertex already exists
		err := fmt.Errorf("ERROR: edge already exists")
		fmt.Println(err.Error())
	} else if fromVertex == toVertex { // if edges are the same
		err := fmt.Errorf("ERROR: cannot connect room to itself")
		fmt.Println(err.Error())
	} else if fromVertex.name == endroom { // if 'from' room  = end
		toVertex.links = append(toVertex.links, fromVertex)
	} else if toVertex.name == startroom { // if 'to' room = start
		toVertex.links = append(toVertex.links, fromVertex)
	} else {
		fromVertex.links = append(fromVertex.links, toVertex)
	}
}

// print graph out
func (g *Graph) Print() {
	fmt.Printf("Number of Ants: %v", antCount())
	fmt.Printf("\nStarting room: %v", getStart())
	fmt.Printf("\nEnding room: %v", getEnd())
	fmt.Println()

	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.name)
		for _, v := range v.links {
			fmt.Printf("%v ", v.name)
		}
	}
}

var verticesMap = make(map[string]*Vertex)

// add vertices to a single map
func (g *Graph) addtoMap() map[string]*Vertex {
	// iterate over the vertices in the graph and add them to the map
	for _, v := range g.vertices {
		verticesMap[v.name] = v
	}
	return verticesMap
}

func DFS(start, end *Vertex, graph *Graph) [][]*Vertex {
	// Initialize visited set and path stack
	visited := make(map[*Vertex]bool)
	path := make([]*Vertex, 0)

	// Initialize result slice to store all valid paths from start to end
	result := [][]*Vertex{}

	// Call recursive DFS function to find all valid paths
	dfs(start, end, visited, path, &result)

	return result
}

func dfs(current, end *Vertex, visited map[*Vertex]bool, path []*Vertex, result *[][]*Vertex) {
	// Mark current vertex as visited and add it to the path
	visited[current] = true
	path = append(path, current)

	// If current vertex is the end vertex, add the current path to the result
	if current == end {
		// Create a copy of the path slice to avoid modifying the original slice
		newPath := make([]*Vertex, len(path))
		copy(newPath, path)
		*result = append(*result, newPath)
	} else {
		// Recursively call DFS on all adjacent vertices
		for _, neighbor := range current.links {
			if !visited[neighbor] {
				dfs(neighbor, end, visited, path, result)
			}
		}
	}

	// Remove current vertex from path and unmark it as visited for backtracking
	path = path[:len(path)-1]
	visited[current] = false
}

// func (g *Graph) FindPaths(start, end *Vertex) [][]*Room {
// 	// Create a list to keep track of the rooms that have been visited
// 	visited := make([]bool, len(g.Rooms))

// 	// Initialize an empty path
// 	path := []*Room{}

// 	// Initialize an empty list of paths
// 	paths := [][]*Room{}

// 	// Call the depth-first search function to find all paths from the start room to the end room
// 	g.dfs(start, end, visited, path, &paths)

// 	// Return the list of all paths
// 	return paths
// }

var requiredSteps int // minimum number of steps

// take a 2D slice of vertex pointers, return a 2D string slice, check if path is suitable
func FindCompatiblePaths(paths [][]*Vertex) [][]string {
	var compatiblePaths [][]string // a 2d int slice to store suitable paths

	// this loop will compare paths in the array
	for i, path1 := range paths {
		// make a new slice at that index for every new path
		// new slice = current path
		compatiblePaths = append(compatiblePaths, []string{string(i)})

		// map to store names of rooms in current path
		roomMap := make(map[string]struct{})

		// loop thru rooms in the current path and add them to roomMap
		for _, room := range path1[1 : len(path1)-1] {
			// assign an empty struct value to a key in a map
			// we only need the key and not the value here
			roomMap[room.name] = struct{}{}
		}

		// loop thru the next path and compare it to current path
		// we need to check whether they are identical
		for j, path2 := range paths[i+1:] {
			isCompatible := true

			// iterate over rooms in current path
			for _, room := range path2[1 : len(path2)-1] {

				// if a room appears in both paths, the paths are not compatible
				if _, ok := roomMap[room.name]; ok {
					isCompatible = false
					break
				}
			}
			// if not
			if isCompatible {
				// append the index of the current path (i+1+j) to compatiblePaths at index i
				compatiblePaths[i] = append(compatiblePaths[i], string(i+1+j))

				// iterate over each room in path2, add all to roomMap
				for _, room := range path2[1 : len(path2)-1] {
					// add each room to the roomMap
					roomMap[room.name] = struct{}{}
				}
			}
		}
	}
	return compatiblePaths
}

func main() {
	test := Graph{}
	for i, line := range readFile() {
		if strings.Contains(string(line), " ") {
			test.AddVertex(strings.Split(readFile()[i], " ")[0])
		}
		if strings.Contains(string(line), "-") {
			test.AddEdge(strings.Split(readFile()[i], "-")[0], strings.Split(readFile()[i], "-")[1])
		}
	}

	test.Print()
	fmt.Println()

	test.addtoMap()
	fmt.Println()
	fmt.Println(verticesMap)
	fmt.Println()

	// apply dfs
	start := verticesMap[startroom]
	end := verticesMap[endroom]
	allPaths := test.findAllPaths(start, end)
	fmt.Println("All available paths from start to end:")
	// range thru all paths
	for _, path := range allPaths {
		// range thru all vertices
		for _, vertex := range path {
			fmt.Printf("%v ", vertex.name)
		}
		fmt.Println()
	}

	FindCompatiblePaths()
}

// https://www.youtube.com/watch?v=bSZ57h7GN2w
