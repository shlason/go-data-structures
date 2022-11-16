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

func (g *graph) initailizeColor() map[string]string {
	colors := make(map[string]string)
	for _, v := range g.vertices {
		colors[v] = "white"
	}
	return colors
}

func (g *graph) dfsVisit(
	vertex string,
	colors map[string]string,
	time *int,
	dt map[string]int,
	ft map[string]int,
	pred map[string]string,
) {
	*time++
	colors[vertex] = "grey"
	dt[vertex] = *time
	for _, v := range g.adjList[vertex] {
		if (colors[v]) == "white" {
			pred[v] = vertex
			g.dfsVisit(v, colors, time, dt, ft, pred)
		}
	}
	*time++
	colors[vertex] = "black"
	ft[vertex] = *time
}

func (g *graph) DFS() (discoverT map[string]int, finishT map[string]int, pred map[string]string) {
	colors := g.initailizeColor()
	time := new(int)
	discoverT = make(map[string]int)
	finishT = make(map[string]int)
	pred = make(map[string]string)

	for _, v := range g.vertices {
		discoverT[v] = 0
		finishT[v] = 0
		pred[v] = ""
	}

	for _, vertex := range g.vertices {
		if colors[vertex] == "white" {
			g.dfsVisit(vertex, colors, time, discoverT, finishT, pred)
		}
	}
	return discoverT, finishT, pred
}

func (g *graph) BFS(v string) (distance map[string]int, predecessors map[string]string) {
	q := queue.New()
	colors := g.initailizeColor()

	distance = make(map[string]int)
	predecessors = make(map[string]string)

	for _, vertex := range g.vertices {
		distance[vertex] = 0
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
