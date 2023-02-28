func DFS(r *Room, g Graph) {
	vList := []string{}
	sRoom := g.getRoom(StartR)

	// set the room being checked visited status to true
	if r.key != EndR {
		r.visited = true
		// append the r key to the visited list
		vList = append(vList, r.key)
		// range through the neighbours of the r
		for _, nbr := range r.adjacent {
			if !nbr.visited {
				nbr.path = append(r.path, nbr)
				if contains(nbr.path, EndR) {
					pathSlice = append(pathSlice, nbr.path)
				}
				vList = append(vList, nbr.key)
				DFS(nbr, Graph{g.rooms})
			}
		}
	} else {
		if len(sRoom.adjacent) > 1 && !contains(sRoom.adjacent, EndR) {
			vList = append(vList, r.key)
			sRoom.adjacent = sRoom.adjacent[1:][:]
			DFS(sRoom, Graph{g.rooms})
		} else {
			vList = append(vList, r.key)
		}
	}
}
