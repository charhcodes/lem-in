package main

import (
	"bufio"
	"fmt"
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

// 5. check for text file errors: # of ants, if there are ##start and ##end rooms
func isError(fileOpen string) {
	var errCheck bool // if errCheck == true, then there is an error and the program exits

	if antCount(fileOpen) > 1 && antCount(fileOpen) < 1000 {
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
}

func main() {
	lines := openFile()
	isError(lines)

	antCount(lines)
	fmt.Println("Number of ants:", antCount(lines))
	findStart(lines)
	fmt.Println("Starting room:", findStart(lines))
	findEnd(lines)
	fmt.Println("Ending room:", findEnd(lines))
}
