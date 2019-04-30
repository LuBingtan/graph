package graph

import (
	"testing"
)

// testing for create graph
func Test4Graph4Create(t *testing.T) {
	t.Logf("testing for graph initialize start.\n")
	// new graph
	g := createDirectedGraph4Test(t)

	graphSortPrint(t, g)
}

// testing for graph delete
func Test4Graph4Del(t *testing.T) {
	t.Logf("testing for graph delete start.\n")
	// new graph
	g := createDirectedGraph4Test(t)

	t.Log("----------------------------------------------------------------init----------------------------------------------------------------")
	graphSortPrint(t, g)

	t.Log("----------------------------------------------------------------remove node1----------------------------------------------------------------")
	g.RemoveVertex(g.GetVertex("node1"))

	graphSortPrint(t, g)

	t.Log("----------------------------------------------------------------remove edge 2->3----------------------------------------------------------------")
	err := g.RemoveEdge(g.GetVertex("node2"), g.GetVertex("node3"))
	if err != nil {
		t.Error(err)
		return
	}
	graphSortPrint(t, g)

	t.Log("----------------------------------------------------------------remove edge 0->1----------------------------------------------------------------")
	err = g.RemoveEdge(g.GetVertex("node0"), g.GetVertex("node1"))
	if err != nil {
		t.Log(err)
	}
}

/// [Inteface]
// print graph method
func graphSortPrint(t *testing.T, g GraphInterface) {
	vList, err := TopoSort(g)
	if nil != err {
		t.Error(err)
	}

	t.Log("*****************************sort*****************************")
	for _, v := range vList {
		printVertex(t, v)
	}
}

func graphPrint(t *testing.T, g GraphInterface) {
	t.Logf("AbstractGraph node number:%d\n", len(g.Verteces()))

	t.Log("*****************************original*****************************")
	for _, v := range g.Verteces() {
		printVertex(t, v)
	}
}

func printVertex(t *testing.T, v VertexInterface) {
	t.Logf("------- %s -------", v.Name())
	t.Logf("indegree:%d, outdegree:%d", v.Indegree(), v.Outdegree())
	for _, e := range v.EdgesForward() {
		t.Logf("edge_forward:%s -> %s", e.From().Name(), e.To().Name())
	}

	for _, e := range v.EdgesBackward() {
		t.Logf("edge_backward:%s -> %s", e.From().Name(), e.To().Name())
	}
}
