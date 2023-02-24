// define a struct to represent a room in the graph
type Room struct {
	name      string
	neighbors []string // rooms next to current
}

// define a function to perform BFS and return the minimum number of steps required to reach the end room (int)
func bfs(startRoom string, endRoom string, rooms map[string]Room) int {

	// initialize a queue and a visited set
	queue := []string{startRoom}                // queue of rooms to be visited
	visited := map[string]bool{startRoom: true} // checks whether room in our map has been visited
	distance := map[string]int{startRoom: 0}    // distance from start to end

	// perform BFS
	for len(queue) > 0 {
		// dequeue the next room from the queue
		room := queue[0]  // current room
		queue = queue[1:] // queue starts from queue[1] aka room is taken out of the queue

		// if the current room is the end room, return the distance
		if room == endRoom {
			return distance[room] // = number of steps to get from start to end
		}

		// iterate over the neighbors of the current room
		for _, neighbor := range rooms[room].neighbors {
			// if the neighbor has not been visited, mark it as visited and enqueue it
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
				distance[neighbor] = distance[room] + 1 // distance of neighbour = current room + 1
			}
		}
	}
	// if the end room was not found, return -1 to indicate failure
	return -1
}
