package lemin

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var requiredSteps int

func FindCompatiblePaths(paths [][]*Room) [][]int {
	var compatiblePaths [][]int
	for i, path1 := range paths {
		compatiblePaths = append(compatiblePaths, []int{i})
		roomMap := make(map[int]struct{})
		for _, room := range path1[1 : len(path1)-1] {
			roomMap[room.Id] = struct{}{}
		}
		for j, path2 := range paths[i+1:] {
			isCompatible := true
			for _, room := range path2[1 : len(path2)-1] {
				if _, ok := roomMap[room.Id]; ok {
					isCompatible = false
					break
				}
			}
			if isCompatible {
				compatiblePaths[i] = append(compatiblePaths[i], i+1+j)
				for _, room := range path2[1 : len(path2)-1] {
					roomMap[room.Id] = struct{}{}
				}
			}
		}
	}
	return compatiblePaths
}

func PathAssign(paths [][]*Room, validPaths [][]int, antNbr int) []string {
	var bestAssignedPath []string
	bestMaxStepLength := math.MaxInt32

	for _, validPath := range validPaths {
		var stepLength []int
		var assignedPath []string
		for _, pathIndex := range validPath {
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

func PrintAntSteps(filteredPaths [][]*Room, pathStrings []string) {
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
			temp := "L" + antStr + "-" + path.Name
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
