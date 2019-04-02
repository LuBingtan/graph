package graph

import (
	"testing"
)

/// [Inteface]
// print graph method
func GraphSortPrint(t *testing.T, g *AbstractGraph) {
	t.Logf("AbstractGraph node number:%d\n", len(g.Verteces()))

	t.Log("*****************************original*****************************")
	for k, v := range g.Verteces() {
		t.Log("=========Node=========:")
		t.Log("key:", k)
		t.Log("id:", v.Name())
		t.Log("Indegree:", v.Indegree())
		t.Log("Outdegree:", v.Outdegree())

		for _, edge := range v.EdgesForward() {
			t.Logf("edge_forward: %v -> %v", edge.From().Name(), edge.To().Name())
		}
		for _, edge := range v.EdgesBackward() {
			t.Logf("edge_backward: %v -> %v", edge.From().Name(), edge.To().Name())
		}
	}

	vList, err := TopoSort(g)
	if nil != err {
		t.Error(err)
	}
	t.Log("*****************************sort*****************************")
	for _, v := range vList {
		t.Log("=========Node=========:")
		t.Log("id:", v.Name())
		t.Log("Indegree:", v.Indegree())
		t.Log("Outdegree:", v.Outdegree())

		for _, edge := range v.EdgesForward() {
			t.Logf("edge_forward: %v -> %v", edge.From().Name(), edge.To().Name())
		}

		for _, edge := range v.EdgesBackward() {
			t.Logf("edge_backward: %v -> %v", edge.From().Name(), edge.To().Name())
		}
	}
}

func Test4Graph4Create(t *testing.T) {
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
	g.InsertEdge(n3, n1, NewEdge())
	g.InsertEdge(n3, n2, NewEdge())
	g.InsertEdge(n2, n1, NewEdge())
	g.InsertEdge(n1, n0, NewEdge())

	GraphSortPrint(t, g)
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

	t.Log("----------------------------------------------------------------init----------------------------------------------------------------")
	GraphSortPrint(t, g)

	t.Log("----------------------------------------------------------------remove vertex----------------------------------------------------------------")
	g.RemoveVertex(n1)

	GraphSortPrint(t, g)

	t.Log("----------------------------------------------------------------remove edge 2->3----------------------------------------------------------------")
	err := g.RemoveEdge(n2, n3)
	if err != nil {
		t.Error(err)
		return
	}
	GraphSortPrint(t, g)

	t.Log("----------------------------------------------------------------remove edge 0->1----------------------------------------------------------------")
	err = g.RemoveEdge(n0, n1)
	if err != nil {
		t.Log(err)
	}
}
