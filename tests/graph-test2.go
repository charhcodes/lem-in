// makes graph, also checks for duplicate rooms, links, and rooms with identical coordinates
// bfs in go
package main

import (
	"fmt"
)

type Node struct {
	Name     string
	Children []*Node
}

func BFS(root *Node) {
	queue := []*Node{root}
	visited := make(map[*Node]bool)
	visited[root] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.Name)

		for _, child := range current.Children {
			if !visited[child] {
				queue = append(queue, child)
				visited[child] = true
			}
		}
	}
}

func main() {
	root := &Node{
		Name: "A",
		Children: []*Node{
			{
				Name: "B",
				Children: []*Node{
					{
						Name: "C",
						Children: []*Node{
							{Name: "D"},
							{Name: "E"},
						},
					},
					{
						Name: "F",
						Children: []*Node{
							{Name: "G"},
							{Name: "H"},
						},
					},
				},
			},
			{
				Name: "I",
				Children: []*Node{
					{Name: "J"},
					{Name: "K"},
					{Name: "L"},
				},
			},
		},
	}

	BFS(root)
}
