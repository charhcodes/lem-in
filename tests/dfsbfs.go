func DFS(r *Room, g Graph) {
	args := ""
	if len(os.Args) == 2 {
		args = os.Args[1]
	} else {
		err := fmt.Errorf("ERROR: Invalid number of arguments")
		fmt.Println(err.Error())
		os.Exit(0)
	}

	sRoom := g.getRoom(StartRoom(readAntsFile(args)))
	// set the room being checked visited status to true
	if r.key != EndRoom(readAntsFile(args)) {
		r.visited = true
		// range through the neighbours of the r
		for _, nbr := range r.adjacent {
			if !nbr.visited {
				/* for each neighbour that hasn't been visited,
				- append their key to the visited slice,
				- then apply dfs to them recursively,
				- then append their key to their path value
				*/
				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndRoom(readAntsFile(args))) {
					dfsPaths = append(dfsPaths, nbr.path)
				}
				DFS(nbr, Graph{g.rooms})
			}
		}
	} else {
		if len(sRoom.adjacent) > 1 && !contains(sRoom.adjacent, EndRoom(readAntsFile(args))) {
			sRoom.adjacent = sRoom.adjacent[1:][:]
			DFS(sRoom, Graph{g.rooms})
		} else {
		}
	}
	dfsPaths = PathDupeCheck(dfsPaths)
}

func DFSBFS(r *Room, g Graph) bool {
	args := ""
	if len(os.Args) == 2 {
		args = os.Args[1]
	} else {
		err := fmt.Errorf("ERROR: Invalid number of arguments")
		fmt.Println(err.Error())
		os.Exit(0)
	}

	// set the room being checked visited status to true
	if r.key != EndRoom(readAntsFile(args)) {
		r.visited = true
		// range through the neighbours of the r
		for _, nbr := range r.adjacent {
			if !nbr.visited {
				/* for each neighbour that hasn't been visited,
				- append their key to the visited slice,
				- then apply dfs to them recursively,
				- then append their key to their path value
				*/
				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndRoom(readAntsFile(args))) {
					return true
				}
			}
		}
	}
	return false
}