package main

import (
	"bufio"
	"fmt"

	//graphs "lem-in-practice/graph"
	"os"
	"strconv"
	"strings"
)

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

// // redefine struct values of type Ant
func main() {
	file := openFile()
	isError(file)

	//a := new(graphs.Ants)

	// a.NumAnts = antCount(file)
	// a.StartRoom = findStart(file)
	// a.EndRoom = findEnd(file)
	// a.RoomName = getNames(file)
	// a.Neighbour = getAdjacents(file)

	fmt.Println("Number of ants:", antCount(file))
	fmt.Println("Starting room:", findStart(file))
	fmt.Println("Ending room:", findEnd(file))
	fmt.Println("Room names:", getNames(file))
	fmt.Println("Room links:", getAdjacents(file))

	// create graph called test
	// test := &graphs.Graph{}

	// add nodes
	// for i := 0; i < 5; i++ {
	// 	test.AddVertex(i)
	// }

	//fmt.Println(test) // prints out addresses
	// test.AddEdge(1, 2)
	// test.AddEdge(3, 2)

	// test.Print()
}
