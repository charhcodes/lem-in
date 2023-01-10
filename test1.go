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
		fmt.Println("File error! Cannot run command")
		os.Exit(0)
	}

	fileOpen := string(file)
	return fileOpen
}

// 2. find number of ants
func antCount(fileOpen string) int {
	antsB := fileOpen[0]
	ants, _ := strconv.Atoi(string(antsB))
	fmt.Println("number of ants:", ants)
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
	fmt.Println("starting room:", start)
	return start
}

// 4. find ending room
func findEnd(fileOpen string) string {
	scanner := bufio.NewScanner(strings.NewReader(fileOpen))

	// Find the line after "start".
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
	fmt.Println("ending room:", end)
	return end
}

func main() {
	lines := openFile()
	antCount(lines)
	findStart(lines)
	findEnd(lines)
}
