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

func NewDirectedGraph() *Graph {
	return &Graph{
		Vertices: map[int]*Vertex{},
		directed: true,
	}
}

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

func (g *Graph) DeleteVertex(key int) error {
	v := g.Vertices[key]
	if v == nil {
		return errors.New("the vertex does not exist")
	}

	if g.directed {
		for i := range g.Vertices {
			delete(g.Vertices[i].Vertices, v.Key)
		}
	} else {
		for i := range v.Vertices {
			delete(v.Vertices[i].Vertices, v.Key)
		}
	}

	delete(g.Vertices, v.Key)
	return nil
}

func (g *Graph) DeleteEdge(k1, k2 int) error {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	if v1 == nil || v2 == nil {
		return errors.New("one or more vertices do not exist")
	}

	if v1.Vertices[v2.Key] == nil {
		return nil
	}

	delete(v1.Vertices, v2.Key)
	if !g.directed {
		delete(v2.Vertices, v1.Key)
	}
	return nil
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

func (g Graph) DFS(startVertexKey int, visitedOrder map[int]bool, s *[]int) {

	startVertex := g.Vertices[startVertexKey]

	if startVertex == nil {
		return
	}

	visitedOrder[startVertex.Key] = true
	*s = append(*s, startVertex.Key)

	for _, v := range startVertex.Vertices {
		if !visitedOrder[v.Key] {
			g.DFS(v.Key, visitedOrder, s)
		}
	}
}

func (g *Graph) BFS(startVertex *Vertex, s *[]int) {
	vertexQueue := &queue{}
	visitedVertices := map[int]bool{}

	currentVertex := startVertex
	for {
		visitedVertices[currentVertex.Key] = true
		*s = append(*s, currentVertex.Key)

		for _, v := range currentVertex.Vertices {
			if !visitedVertices[v.Key] {
				vertexQueue.enqueue(v)
			}
		}

		currentVertex = vertexQueue.dequeue()

		if currentVertex == nil {
			break
		}
	}
}

type node struct {
	value *Vertex
	next  *node
}

type queue struct {
	head, tail *node
}

func (q *queue) enqueue(v *Vertex) {
	n := &node{value: v}

	if q.head == nil {
		q.head = n
		q.tail = n
		return
	}

	q.tail.next = n
	q.tail = n
}

func (q *queue) dequeue() *Vertex {
	n := q.head

	if n == nil {
		return nil
	}

	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return n.value
}
