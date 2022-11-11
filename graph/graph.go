package graph

import (
	"fmt"
	"strings"

	"github.com/shlason/go-data-structures/queue"
	"github.com/shlason/go-data-structures/stack"
)

type nodes []string

type graph struct {
	vertices nodes
	adjList  map[string]nodes
}

func New() *graph {
	g := new(graph)
	g.vertices = make(nodes, 0)
	g.adjList = make(map[string]nodes)
	return g
}

func (g *graph) AddVertex(v string) {
	g.vertices = append(g.vertices, v)
	g.adjList[v] = make(nodes, 0)
}

func (g *graph) AddEdge(v, w string) {
	g.adjList[v] = append(g.adjList[v], w)
	g.adjList[w] = append(g.adjList[w], v)
}

func (g *graph) ToString() string {
	r := ""
	for _, v := range g.vertices {
		r += fmt.Sprintf("%s -> ", v)
		for _, vv := range g.adjList[v] {
			r += fmt.Sprintf("%s ", vv)
		}
		r += "\n"
	}
	return r
}

func (g *graph) BFS(v string) (distance map[string]int, predecessors map[string]string) {
	q := queue.New()
	colors := make(map[string]string)

	distance = make(map[string]int)
	predecessors = make(map[string]string)

	for _, vertex := range g.vertices {
		distance[vertex] = 0
		colors[vertex] = "white"
		predecessors[vertex] = ""
	}

	colors[v] = "grey"
	q.Enqueue(v)

	for !q.IsEmpty() {
		vertex := q.Dequeue().(string)
		for _, gv := range g.adjList[vertex] {
			if colors[gv] == "white" {
				colors[gv] = "grey"
				predecessors[gv] = vertex
				distance[gv] = distance[vertex] + 1
				q.Enqueue(gv)
			}
		}
		colors[vertex] = "black"
	}

	return distance, predecessors
}

func (g *graph) PrintShortestPathByBFS(fromVertex string) {
	_, predecessors := g.BFS(fromVertex)
	for _, toVertext := range g.vertices {
		path := stack.New()
		for v := toVertext; v != fromVertex; v = predecessors[v] {
			path.Push(v)
		}
		path.Push(fromVertex)
		r := []string{}
		for !path.IsEmpty() {
			r = append(r, path.Pop().(string))
		}
		fmt.Println(strings.Join(r, " -> "))
	}
}
