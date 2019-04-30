package graph

import (
	"testing"
)
func Test4TopoSort(t *testing.T) {
	g := createDirectedGraph4Test(t)
	graphSortPrint(t, g)
}

// testing for graph bfs
func Test4BFS(t *testing.T) {
	sum := 0
	f := func(v VertexInterface) {
		t.Log(v.Name())
		sum += v.Data().(int)
	}

	t.Logf("testing for undirected graph bfs start.\n")
	ug := createUndirectedGraph4Test(t)
	sum = 0
	BFS(ug, f)
	t.Log("data sum:", sum)

	t.Logf("testing for directed graph bfs start.\n")
	dg := createDirectedGraph4Test(t)
	sum = 0
	BFS(dg, f)
	t.Log("data sum:", sum)
}

// testing for graph dfs
func Test4DFS(t *testing.T) {
	t.Logf("testing for graph dfs start.\n")
	// new graph
	g := createUndirectedGraph4Test(t)

	sum := 0
	f := func(v VertexInterface) {
		t.Log(v.Name())
		sum += v.Data().(int)
	}
	DFS(g, f)
	t.Log("data sum:", sum)
}