package graph

import (
	"fmt"
	"log"
	"testing"
)

func TestGraphStructure(t *testing.T) {
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

func TestGraphDFS(t *testing.T) {
	g := NewUndirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)
	g.AddVertex(9)
	g.AddVertex(10)

	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(1, 9)
	g.AddEdge(2, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)

	visitedOrder := make(map[int]bool)
	var s []int
	g.DFS(1, visitedOrder, &s)
	fmt.Println(visitedOrder)
	fmt.Println(s)
}

func TestGraphBFS(t *testing.T) {
	g := NewUndirectedGraph()

	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)
	g.AddVertex(8)
	g.AddVertex(9)
	g.AddVertex(10)

	g.AddEdge(1, 2)
	g.AddEdge(1, 5)
	g.AddEdge(1, 9)
	g.AddEdge(2, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)
	g.AddEdge(5, 8)
	g.AddEdge(6, 7)
	g.AddEdge(9, 10)

	var s []int
	g.BFS(g.Vertices[1], &s)

	fmt.Println(s)
}

func TestGraphDeleteVertex(t *testing.T) {
	g := NewDirectedGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	g.AddEdge(2, 1)
	g.AddEdge(3, 1)

	err := g.DeleteVertex(1)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(g)
}

func TestGraphDeleteEdge(t *testing.T) {
	g := NewUndirectedGraph()
	g.AddVertex(1)
	g.AddVertex(2)

	g.AddEdge(1, 2)

	fmt.Println(g)
	g.DeleteEdge(1, 2)
	fmt.Println(g)

}
