// prints one path

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

// -------------- MAKE GRAPH ------------------

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

func (g *Graph) addtoMap() map[string]*Vertex {
	// iterate over the vertices in the graph and add them to the map
	for _, v := range g.vertices {
		verticesMap[v.name] = v
	}
	return verticesMap
}

// -------------- DFS ------------------

// apply dfs to find all valid paths
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

// apply dfs algo to a specific path
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

// -------------- GUIDE ------------------

var requiredSteps int // minimum number of steps

// take a 2D slice of vertex pointers, return a 2D slice of integers, check if path is suitable
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

// pathAssign is a function that takes the 2D array of paths and the compatible paths, the number of ants
// and assigns a path to each ant.
func PathAssign(paths [][]*Vertex, validPaths [][]string, antNbr int) []string {
	var bestAssignedPath []string      // most suitable path
	bestMaxStepLength := math.MaxInt32 // bestmax = max value of int32

	// iterate over valid paths
	// we are finding the shortest available paths and assigning them to our ants
	for _, validPath := range validPaths {
		var stepLength []int      // store the step lengths of each path in the current valid path
		var assignedPath []string // store the assigned path for each ant

		// loop thru each valid path
		for _, pathIndex := range validPath {
			// add length of the slice to the end of slice stepLength
			// pathIndex tracks number of steps
			pathint, _ := strconv.Atoi(pathIndex)
			path := paths[pathint]
			stepLength = append(stepLength, len(path)-1)
		}

		// loop thru number of ants
		for i := 1; i <= antNbr; i++ {
			minStepsIndex := 0 // used to store the index of the shortest path in the slice

			for j, steps := range stepLength {
				// if we find a shorter path minstepsIndex is updated
				if steps <= stepLength[minStepsIndex] {
					minStepsIndex = j
				}
			}
			// assign shortest path to our ant by appending to assignedPath
			// format: 'ant number - validpath at minStepsIndex'
			assignedPath = append(assignedPath, fmt.Sprintf("%d-%d", i, validPath[minStepsIndex]))
			stepLength[minStepsIndex]++
		}
		// calculate maximum step length
		maxStepLength := 0
		// range thru stepLength to find the longest path possible
		for _, steps := range stepLength {
			if steps > maxStepLength {
				maxStepLength = steps
			}
		}
		// if maxStepLength is less than bestMaxStep
		if maxStepLength < bestMaxStepLength {
			bestAssignedPath = assignedPath   // assigned path for this ant = best
			bestMaxStepLength = maxStepLength // update bestMaxStep to be equal to maxStep
		}
	}
	requiredSteps = bestMaxStepLength // update requiredSteps with bestMaxStep
	return bestAssignedPath
}

// print ant steps in the correct format
func PrintAntSteps(filteredPaths [][]*Vertex, pathStrings []string) {
	var antSteps [][]string                    // store the steps taken by each ant in order
	arrayLen := requiredSteps - 1              // the number of turns required to get from start to end
	orderedSteps := make([][]string, arrayLen) // make a slice to store all the steps taken by the ants, in order

	// loop thru each ant's path
	for _, antPath := range pathStrings {
		// store steps taken by the ant
		var steps []string
		// split the antPath string into its ant number and path index components
		parts := strings.SplitN(antPath, "-", 2)     // separate string 'antPath' by a hyphen into two pieces
		antStr := parts[0]                           // ant number
		antPath, _ := strconv.Atoi(string(parts[1])) // path index (what path this specific ant will take)
		// loop thru each room in the ant's chosen path, add a step string to the steps slice for every room
		for i := 1; i < len(filteredPaths[antPath]); i++ {
			path := filteredPaths[antPath][i]      // the current room of the ant's path
			temp := "L" + antStr + "-" + path.name // format "L1-richard"
			steps = append(steps, temp)
		}
		// add steps (current ant) to antSteps (all ants)
		antSteps = append(antSteps, steps)
	}
	// loop thru antSteps slice, add every step in order to orderedSteps
	for i := 0; i < len(antSteps); i++ { // i = row index
		slice := antSteps[i] // slice = antSteps at row i
		var row int

		// split the antSteps string to get the room name and use getRow to find the first row in orderedSteps
		// that does not contain the room name
		// for j := 0; j < len(slice); j++ { // j tracks the columns of the row
		// 	temp := slice[j] // temp = j at row i
		// 	// if j = 0, j is first
		// 	if j == 0 {
		// 		parts := strings.SplitN(temp, "-", 2)    // split by a hyphen to get room name (parts[1])
		// 		row = getRow(orderedSteps, "-"+parts[1]) // row = row number that doesn't contain the room name
		// 	}
		// 	// [range error here]
		// 	orderedSteps[j+row] = append(orderedSteps[j+row], temp) // add temp to orderedSteps at column j row
		// }

		for j := 0; j < len(slice); j++ { // j tracks the columns of the row
			temp := slice[j] // temp = j at row i
			if j == 0 {      // if j = 0, j is first
				parts := strings.SplitN(temp, "-", 2)    // split by a hyphen to get room name (parts[1])
				row = getRow(orderedSteps, "-"+parts[1]) // row = row number that doesn't contain the room name
			}
			idx := j + row
			if idx >= len(orderedSteps) { // if idx is greater than or equal to the length of orderedSteps, it is out of range
				orderedSteps = append(orderedSteps, make([]string, 0)) // append a new slice of strings to orderedSteps
			}
			orderedSteps[idx] = append(orderedSteps[idx], temp) // add temp to orderedSteps at index idx
		}
		// reset row to
		row = 0
	}
	// loop thru every step of the orderedSteps slice, print number of steps
	for i, printline := range orderedSteps {
		fmt.Println(strings.Trim(fmt.Sprint(printline), "[]"))
		if i == len(orderedSteps)-1 {
			fmt.Println()
			fmt.Printf("Number of turns: %v\n", i+1)
		}
	}
}

// takes a 2D slice and a value to search for and returns the index
// of the first row in the slice that does not contain the value
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
	test := &Graph{}
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

	//start, end *Vertex, graph *Graph
	startvert := test.getVertex(startroom)
	endvert := test.getVertex(endroom)
	paths := DFS(startvert, endvert, test)
	if len(paths) == 0 {
		fmt.Println("ERROR: Invalid data format. No path found or text file is formatted incorrectly")
		return
	}
	fmt.Println("dfs success")
	for _, path := range paths {
		for _, v := range path {
			fmt.Printf("%s ", v.name)
		}
		fmt.Println()
	}

	validPaths := FindCompatiblePaths(paths)
	bestPath := PathAssign(paths, validPaths, antCount())

	fmt.Println("get here")
	path := os.Args[1]
	_, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("print content")
	PrintAntSteps(paths, bestPath)
}

// https://www.youtube.com/watch?v=bSZ57h7GN2w
