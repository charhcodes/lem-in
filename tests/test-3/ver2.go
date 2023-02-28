// make graph

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
	visited bool
	ants    int
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

	for i := range file {
		if file[i] == "##start" {
			start = strings.Split(string(file[i+1]), " ")[0] //splits file at the line after file[i]
		}
	}
	return start
}

// returns ending room
func getEnd() string {
	var end string
	file := readFile()

	for i := range file {
		if file[i] == "##end" {
			end = strings.Split(string(file[i+1]), " ")[0]
		}
	}
	return end
}

var (
	startroom = getStart()
	endroom   = getEnd()
)

// return a pointer to the Vertex with its name
func (g *Graph) getVertex(name string) *Vertex { //to be called on a graph object (vertices)
	for i, v := range g.vertices {
		if v.name == name {
			return g.vertices[i]
		}
	}
	return nil
}

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
		err := fmt.Errorf("ERROR: vertex already exists")
		fmt.Println(err.Error())
	} else {
		vertices := &Vertex{name: name}
		g.vertices = append(g.vertices, vertices)
	}
}

// add edges to graph
func (g *Graph) AddEdge(f, n string) {
	from := g.getVertex(f)
	next := g.getVertex(n)

	if from == nil || next == nil { // if there are no edges before/after
		err := fmt.Errorf("ERROR: invalid edges")
		fmt.Println(err.Error())
	} else if contains(from.links, n) { // if edge already exists
		err := fmt.Errorf("ERROR: edge already exists")
		fmt.Println(err.Error())
	} else if from == next { // if edges are the same
		err := fmt.Errorf("ERROR: cannot connect room to itself")
		fmt.Println(err.Error())
	} else if from.name == endroom { // if 'from' room  = end
		next.links = append(next.links, from)
		//adds a link from 'next' to 'from' to next.links
	} else if next.name == startroom { // if 'next' room = start
		next.links = append(next.links, from)
	} else {
		from.links = append(from.links, next)
		// adds a link from 'from' to 'next' to from.links
	}
}

// print graph out
func (g *Graph) Print() {
	fmt.Printf("Number of Ants: %v", antCount())
	fmt.Printf("\nStarting room: %v", getStart())
	fmt.Printf("\nEnding room: %v", getEnd())
	fmt.Println()

	for _, v := range g.vertices { // print vertices
		if v.name == endroom {
			fmt.Printf("\nVertex %v", v.name)
		} else {
			fmt.Printf("\nVertex %v: ", v.name)
		}
		for _, v := range v.links { // print each vertex's links
			fmt.Printf(" %v ", v.name)
		}
	}
}

func main() {
	test := Graph{} // create empty graph object
	for i, line := range readFile() {
		// look for comments (non-start/end)
		// ignore the first start/end rooms
		var startFound bool
		var endFound bool

		if strings.Contains(string(line), "##start") && startFound || strings.Contains(string(line), "##end") && endFound {
			err := fmt.Errorf("ERROR: more than one start/end found")
			fmt.Println(err.Error())
		} else if strings.Contains(string(line), "##start") && !startFound {
			startFound = true
		} else if line == "##end" && !endFound {
			endFound = true
		} else if strings.HasPrefix(string(line), "##") && line != "##start" && line != "##end" {
			continue
		}

		if strings.Contains(string(line), " ") { // if string contains a space
			test.AddVertex(strings.Split(readFile()[i], " ")[0]) // split string by spaces, add to test
		}
		if strings.Contains(string(line), "-") { // if string contains a hyphen
			test.AddEdge(strings.Split(readFile()[i], "-")[0], strings.Split(readFile()[i], "-")[1])
			// split string by hyphen, add (link) to test
		}
		test.Print()
	}
}

// https://www.youtube.com/watch?v=bSZ57h7GN2w
