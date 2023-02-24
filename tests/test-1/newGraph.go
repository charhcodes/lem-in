package main

import (
	"fmt"
)

type Graph struct {
	Vertices []*Vertex
	Edges    []*Edge
}

type Vertex struct {
	Key int
}

type Edge struct {
	From *Vertex
	To   *Vertex
}

var (
	startRoom = 
)

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

func (g *Graph) AddVertex(k int) {
	if contains(g.Vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.Vertices = append(g.Vertices, &Vertex{Key: k})
	}
}

func (g *Graph) AddEdge(from, to *Vertex) {
	g.Edges = append(g.Edges, &Edge{From: from, To: to})
}
