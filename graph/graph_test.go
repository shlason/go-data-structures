package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	t.Error()
	g := New()
	vertices := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	for _, v := range vertices {
		g.AddVertex(v)
	}

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("A", "D")
	g.AddEdge("C", "D")
	g.AddEdge("C", "G")
	g.AddEdge("D", "G")
	g.AddEdge("D", "H")
	g.AddEdge("B", "E")
	g.AddEdge("B", "F")
	g.AddEdge("E", "I")

	fmt.Println(g.ToString())

	distance, _ := g.BFS("A")

	fmt.Println(distance)

	g.PrintShortestPathByBFS("G")
}
