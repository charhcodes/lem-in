var requiredSteps int // minimum number of steps

// take a 2D slice of vertex pointers, return a 2D slice of integers, check if path is suitable
func FindCompatiblePaths(paths [][]*Vertex) [][]int {
	var compatiblePaths [][]int // a 2d int slice to store suitable paths

	// this loop will compare paths in the array
	for i, path1 := range paths {
		// make a new slice at that index for every new path
		// new slice = current path
		compatiblePaths = append(compatiblePaths, []int{i})

		// map to store names of rooms in current path
		roomMap := make(map[int]struct{})

		// loop thru rooms in the current path and add them to roomMap
		for _, room := range path1[1 : len(path1)-1] {
			// assign an empty struct value to a key in a map
			// we only need the key and not the value here
			roomMap[room.id] = struct{}{}
		}

		// loop thru the next path and compare it to current path
		// we need to check whether they are identical
		for j, path2 := range paths[i+1:] {
			isCompatible := true

			// iterate over rooms in current path
			for _, room := range path2[1 : len(path2)-1] {

				// if a room appears in both paths, the paths are not compatible
				if _, ok := roomMap[room.id]; ok {
					isCompatible = false
					break
				}
			}
			// if not
			if isCompatible {
				// append the index of the current path (i+1+j) to compatiblePaths at index i
				compatiblePaths[i] = append(compatiblePaths[i], i+1+j)

				// iterate over each room in path2, add all to roomMap
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
			path := paths[pathIndex]
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
		// split antPath string
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

func NewGraph() *Graph {
	return &Graph{
		Rooms: make(map[int]*Vertex),
	}
}

func PrintOutput(data []string) {
	antNbr, allRooms, allLinks := FilterData(data)
	if antNbr <= 0 {
		fmt.Println("ERROR: Invalid data format. Invalid number of ants. Must be > 0")
		return
	}

	graph := NewGraph()

	//add rooms to the graph
	roomIDs := make(map[string]int)
	for id, room := range allRooms {
		roomIDs[room] = id
		graph.AddRoom(id, room)
	}

	//add links to the graph
	for _, link := range allLinks {
		parts := strings.Split(link, "-")
		id1 := roomIDs[parts[0]]
		id2 := roomIDs[parts[1]]
		graph.AddLink(id1, id2)
	}

	//assign start and end points
	startRoom := graph.Rooms[0]
	endRoom := graph.Rooms[len(graph.Rooms)-1]

	paths := graph.FindPaths(startRoom, endRoom)
	if len(paths) == 0 {
		fmt.Println("ERROR: Invalid data format. No path found or text file is formatted incorrectly")
		return
	}

	validPaths := FindCompatiblePaths(paths)
	bestPath := PathAssign(paths, validPaths, antNbr)

	path := os.Args[1]
	bytes, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	content := string(bytes)
	fmt.Println(content)
	fmt.Println()
	PrintAntSteps(paths, bestPath)
}