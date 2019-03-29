package graph

import (
	"testing"
)

func GraphPrint(t *testing.T, g *AbstractGraph) {
	t.Logf("AbstractGraph node number:%d\n", len(g.Verteces()))

	t.Log("=======original=======")
	for k, v := range g.Verteces() {
		t.Log("=========Node=========:")
		t.Log("key:", k)
		t.Log("id:", v.Id())
		t.Log("Indegree:", v.Indegree())
		t.Log("Outdegree:", v.Outdegree())
		edges := v.Edges()
		for _, edge := range edges {
			t.Logf("edge:%v,", edge.Vertex())
		}
	}

	vList, err := TopoSort(g)
	if nil != err {
		t.Error(err)
	}
	t.Log("=======sort=======")
	for _, v := range vList {
		t.Log("=========Node=========:")
		t.Log("id:", v.Id())
		t.Log("Indegree:", v.Indegree())
		t.Log("Outdegree:", v.Outdegree())
		edges := v.Edges()
		for _, edge := range edges {
			t.Logf("edge:%v,", edge.Vertex())
		}
	}
}

func Test4Graph4Init(t *testing.T) {
	t.Logf("testing for graph initialize start.\n")
	// new graph
	g := NewGraph("AbstractGraph")
	// new vertex
	n0 := NewVertex("node0", 0)
	n1 := NewVertex("node1", 1)
	n2 := NewVertex("node2", 2)
	n3 := NewVertex("node3", 3)
	// add vertex
	g.InsertVertex(n0)
	g.InsertVertex(n1)
	g.InsertVertex(n2)
	g.InsertVertex(n3)
	// add edge
	g.InsertEdge(n0, n1, NewEdge())
	g.InsertEdge(n0, n2, NewEdge())
	g.InsertEdge(n1, n2, NewEdge())
	g.InsertEdge(n2, n3, NewEdge())

	GraphPrint(t, g)
}

func Test4Graph4Del(t *testing.T) {
	t.Logf("testing for graph delete start.\n")
	// new graph
	g := NewGraph("AbstractGraph")
	// new vertex
	n0 := NewVertex("node0", 0)
	n1 := NewVertex("node1", 1)
	n2 := NewVertex("node2", 2)
	n3 := NewVertex("node3", 3)
	// add vertex
	g.InsertVertex(n0)
	g.InsertVertex(n1)
	g.InsertVertex(n2)
	g.InsertVertex(n3)
	// add edge
	g.InsertEdge(n0, n1, NewEdge())
	g.InsertEdge(n0, n2, NewEdge())
	g.InsertEdge(n1, n2, NewEdge())
	g.InsertEdge(n2, n3, NewEdge())

	t.Log("**********************init**************************")
	GraphPrint(t, g)

	t.Log("**********************remove vertex**************************")
	g.RemoveVertex(n1)

	GraphPrint(t, g)

	t.Log("**********************remove edge**************************")
	err := g.RemoveEdge(n0, n1)
	if err != nil {
		t.Log(err)
	}
}
