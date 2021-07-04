package graph

import (
	"errors"
	"fmt"
)

type Vertex struct {
	Key      int
	Vertices map[int]*Vertex
}

type Graph struct {
	Vertices map[int]*Vertex
	directed bool
}

func NewVertex(key int) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[int]*Vertex{},
	}
}

//func NewDirectedGraph() *Graph {
//	return &Graph{
//		Vertices: map[int]*Vertex{},
//		directed: true,
//	}
//}

func NewUndirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
	}
}

func (v Vertex) String() string {

	s := fmt.Sprintf("Vertex {key:%d\tContains:[", v.Key)
	x := 0
	for i := range v.Vertices {
		if x != 0 {
			s += ","
		}
		s += fmt.Sprintf("%v", v.Vertices[i].Key)
		x++
	}
	s += "]}"
	return s
}

func (g Graph) String() string {
	s := "Graph: ["
	for i := range g.Vertices {
		s += fmt.Sprint(g.Vertices[i])
		s += " | "
	}
	s += "]"

	return s
}

func (g *Graph) AddVertex(key int) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2 int) error {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		return errors.New("one or more vertices do not exist")
	}

	if v1.Vertices[v2.Key] != nil {
		return nil

	}

	v1.Vertices[v2.Key] = v2
	if !g.directed && v1.Key != v2.Key {
		v2.Vertices[v1.Key] = v1
	}

	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
	return nil
}
