// structs and variables
package graph

// struct concerning the ants and the route they must take
/*
number of ants
start
end
room names
room links
*/
type Ants struct {
	numAnts   int
	startRoom string
	endRoom   string
	roomName  string
	adjacent  []*Vertex
}

// struct concerning the graph itself
type Vertex struct {
	vertices []*Vertex
	key      int
	adjacent []*Vertex // neighbouring vertices
}
