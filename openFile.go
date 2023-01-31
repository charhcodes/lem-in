package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//graph "lem-in-practice/control.go"
)

// struct concerning the ants and the route they must take
type Ants struct {
	numAnts   int
	startRoom string
	endRoom   string
	roomName  []string
	adjacent  []string
}

// struct concerning the graph itself
type Vertex struct {
	vertices []*Vertex
	key      int
	adjacent []*Vertex // neighbouring vertices
}

// 1. take text file and convert into useable information
func openFile() string {
	// example input: go run . test01.txt
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}

	fileOpen := string(file)
	return fileOpen
}

// 2. find number of ants
func antCount(fileOpen string) int {
	antsB := fileOpen[0]
	ants, _ := strconv.Atoi(string(antsB))
	return ants
}

// 3. find starting room
func findStart(fileOpen string) string {
	scanner := bufio.NewScanner(strings.NewReader(fileOpen))

	// Find the line after "start".
	var start string
	var foundStart bool
	for scanner.Scan() {
		line := scanner.Text()
		if foundStart {
			start = line
			break
		}
		if line == "##start" {
			foundStart = true
		}
	}
	//fmt.Println("starting room:", start)
	return start
}

// 4. find ending room
func findEnd(fileOpen string) string {
	scanner := bufio.NewScanner(strings.NewReader(fileOpen))

	var end string
	var foundEnd bool
	for scanner.Scan() {
		line := scanner.Text()
		if foundEnd {
			end = line
			break
		}
		if line == "##end" {
			foundEnd = true
		}
	}
	//fmt.Println("ending room:", end)
	return end
}

// 5. get room names
func getNames(fileOpen string) []string {
	scanner := bufio.NewScanner(strings.NewReader(fileOpen))
	var names []string

	for scanner.Scan() {
		line := scanner.Text()
		if line == "##start" {
			// Start collecting the first words of each line
			for scanner.Scan() {
				line := scanner.Text()
				if line == "##end" {
					// Stop collecting when we reach the end keyword
					break
				}
				// Split the line into words and                                                                                            3  the first word to the slice
				name := strings.Split(line, " ")[0]
				names = append(names, name)
			}
			break
		}
	}
	return names
}

// 6. get room links
func getAdjacents(fileOpen string) []string {
	scanner := bufio.NewScanner(strings.NewReader(fileOpen))
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

// 7. check for text file errors
func isError(fileOpen string) {
	var errCheck bool // if errCheck = true, then there is an error and the program exits

	// check for # of ants
	if antCount(fileOpen) > 1 && antCount(fileOpen) < 1001 {
		errCheck = false
	} else if antCount(fileOpen) < 1 {
		errCheck = true
	} else {
		errCheck = true
	}
	if errCheck {
		fmt.Println("ERROR: invalid data format, invalid number of Ants")
		os.Exit(0)
	}

	// check for start and end rooms
	if !strings.Contains(fileOpen, "##start") {
		errCheck = true
		fmt.Println("ERROR: invalid data format, no start room found")
		os.Exit(0)
	}
	if !strings.Contains(fileOpen, "##end") {
		errCheck = true
		fmt.Println("ERROR: invalid data format, no end room found")
		os.Exit(0)
	}

	// check if there's more than one start/end room
	Scount := strings.Count(fileOpen, "start")
	Ecount := strings.Count(fileOpen, "end")
	if Scount > 1 {
		errCheck = true
		fmt.Println("ERROR: more than one start room found")
		os.Exit(0)
	} else {
		errCheck = false
	}
	if Ecount > 1 {
		errCheck = true
		fmt.Println("ERROR: more than one end room found")
		os.Exit(0)
	} else {
		errCheck = false
	}

	// check for duplicate rooms
	seen := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(fileOpen))
	for scanner.Scan() {
		line := scanner.Text()
		if seen[line] {
			errCheck = true
			fmt.Printf("ERROR: duplicate rooms found")
			os.Exit(0)
		} else {
			seen[line] = true
			errCheck = false
		}
	}
}

// redefine struct values of type Ant
func main() {
	file := openFile()
	isError(file)

	a := new(Ants)

	a.numAnts = antCount(file)
	a.startRoom = findStart(file)
	a.endRoom = findEnd(file)
	a.roomName = getNames(file)
	// room links
	a.adjacent = getAdjacents(file)

	fmt.Println("Number of ants:", a.numAnts)
	fmt.Println("Starting room:", a.startRoom)
	fmt.Println("Ending room:", a.endRoom)
	fmt.Println("Room names:", a.roomName)
	fmt.Println("Room links:", a.adjacent)
	//fmt.Println("Room links:", a.adjacent[1])
}
