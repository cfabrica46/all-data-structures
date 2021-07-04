package graph

import (
	"fmt"
	"log"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewUndirectedGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	fmt.Println(g)
	fmt.Println(g.Vertices[1])

	err := g.AddEdge(1, 2)
	if err != nil {
		log.Fatal(err)
	}
	err = g.AddEdge(1, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(g)
	fmt.Println(g.Vertices[1])
	fmt.Println(g.Vertices[2])

	err = g.AddEdge(3, 2)
	if err != nil {
		log.Fatal(err)
	}
}
