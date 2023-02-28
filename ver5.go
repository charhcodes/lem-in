// prints all paths

package main

import (
	"bufio"
	"fmt"
	"math"
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

// var verticesMap = make(map[string]*Vertex)

// func (g *Graph) addtoMap() map[string]*Vertex {
// 	// iterate over the vertices in the graph and add them to the map
// 	for _, v := range g.vertices {
// 		verticesMap[v.name] = v
// 	}
// 	return verticesMap
// }

func (v *Vertex) dfs(end *Vertex, path []*Vertex, paths map[int][]*Vertex, visited map[*Vertex]bool) {
	// v == current
	visited[v] = true      // marks current vertex as visited
	path = append(path, v) // append current to path

	if v == end {
		// Found a path from start to end
		length := len(paths)
		paths[length] = path
	} else {
		for _, link := range v.links {
			if !visited[link] {
				link.dfs(end, path, paths, visited)
			}
		}
	}
	// Remove v from the current path and visited set to backtrack
	delete(visited, v)
}

func (g *Graph) findAllPaths(start *Vertex, end *Vertex) map[int][]*Vertex {
	paths := make(map[int][]*Vertex)
	visited := make(map[*Vertex]bool)

	start.dfs(end, []*Vertex{}, paths, visited)

	return paths
}

var requiredSteps int

// takes a 2D slice of vertex pointers, returns a 2D slice of integers
func FindCompatiblePaths(paths [][]*Vertex) [][]int {
	var compatiblePaths [][]int

	// loops thru each path
	for i, path1 := range paths {
		// initialize a new slice of integers
		// new slice = current path
		compatiblePaths = append(compatiblePaths, []int{i})
		// map to store names of rooms in path1
		roomMap := make(map[int]struct{})

		// loops thru remaining rooms and add them to roomMap
		for _, room := range path1[1 : len(path1)-1] {
			roomMap[room.id] = struct{}{}
		}
		// loops thru each remaining path
		for j, path2 := range paths[i+1:] {
			isCompatible := true
			// iterates over rooms in current path
			for _, room := range path2[1 : len(path2)-1] {
				// checks if each room is present
				if _, ok := roomMap[room.id]; ok {
					isCompatible = false
					break
				}
			}
			if isCompatible {
				// appends the index of the current path (i+1+j) to a compatiblePaths at index i
				compatiblePaths[i] = append(compatiblePaths[i], i+1+j)
				// iterates over rooms in current path
				for _, room := range path2[1 : len(path2)-1] {
					// add each room to the roomMap
					roomMap[room.id] = struct{}{}
				}
			}
		}
	}
	return compatiblePaths
}

// pathAssign is a function that takes the 2D array of paths and the compatible paths, the number of ants
// and assigns a path to each ant.
func PathAssign(paths [][]*Vertex, validPaths [][]int, antNbr int) []string {
	var bestAssignedPath []string
	bestMaxStepLength := math.MaxInt32 // bestmax = max value of int32

	// iterate over valid paths
	for _, validPath := range validPaths {
		var stepLength []int      // store the step lengths of each path in the current valid path
		var assignedPath []string // store the assigned path for each ant

		// loop thru each valid path
		for _, pathIndex := range validPath {
			// add length of the slice to stepLength
			path := paths[pathIndex]
			stepLength = append(stepLength, len(path)-1)
		}
		for i := 1; i <= antNbr; i++ {
			minStepsIndex := 0
			for j, steps := range stepLength {
				if steps <= stepLength[minStepsIndex] {
					minStepsIndex = j
				}
			}
			assignedPath = append(assignedPath, fmt.Sprintf("%d-%d", i, validPath[minStepsIndex]))
			stepLength[minStepsIndex]++
		}
		maxStepLength := 0
		for _, steps := range stepLength {
			if steps > maxStepLength {
				maxStepLength = steps
			}
		}
		if maxStepLength < bestMaxStepLength {
			bestAssignedPath = assignedPath
			bestMaxStepLength = maxStepLength
		}
	}
	requiredSteps = bestMaxStepLength
	return bestAssignedPath
}

func PrintAntSteps(filteredPaths [][]*Vertex, pathStrings []string) {
	var antSteps [][]string
	arrayLen := requiredSteps - 1
	orderedSteps := make([][]string, arrayLen)

	for _, antPath := range pathStrings {
		var steps []string
		parts := strings.SplitN(antPath, "-", 2)
		antStr := parts[0]
		antPath, _ := strconv.Atoi(string(parts[1]))
		for i := 1; i < len(filteredPaths[antPath]); i++ {
			path := filteredPaths[antPath][i]
			temp := "L" + antStr + "-" + path.name
			steps = append(steps, temp)
		}
		antSteps = append(antSteps, steps)
	}
	for i := 0; i < len(antSteps); i++ {
		slice := antSteps[i]
		var row int
		for j := 0; j < len(slice); j++ {
			temp := slice[j]
			if j == 0 {
				parts := strings.SplitN(temp, "-", 2)
				row = getRow(orderedSteps, "-"+parts[1])
			}
			orderedSteps[j+row] = append(orderedSteps[j+row], temp)
		}
		row = 0
	}
	for i, printline := range orderedSteps {
		fmt.Println(strings.Trim(fmt.Sprint(printline), "[]"))
		if i == len(orderedSteps)-1 {
			fmt.Println()
			fmt.Printf("Number of turns: %v\n", i+1)
		}
	}
}

func getRow(tocheck [][]string, value string) int {
	for i, row := range tocheck {
		found := false
		for _, item := range row {
			if strings.Contains(item, value) {
				found = true
				break
			}
		}
		if !found {
			return i
		}
	}
	return 0
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

	start := verticesMap[startroom]
	end := verticesMap[endroom]
	allPaths := test.findAllPaths(start, end)
	fmt.Println("All available paths from start to end:")
	for _, path := range allPaths {
		for _, vertex := range path {
			fmt.Printf("%v ", vertex.name)
		}
		fmt.Println()
	}

}

// https://www.youtube.com/watch?v=bSZ57h7GN2w
