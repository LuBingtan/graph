package graph

import (
	"testing"
)

func Test4UndirectedGraph_Positive(t *testing.T) {
	g := createUndirectedGraph4Test(t)
	graphPrint(t, g)
}

func Test4UndirectedGraph_Negative(t *testing.T) {
	g := createUndirectedGraph4Test(t)
	err := g.InsertEdge(g.GetVertex("node0"), g.GetVertex("node3"), NewEdge(0, BackwardEdge))
	if err != nil {
		t.Log(err)
	} else {
		t.Error("edge type should not matched.")
	}
}

/// create undrected graph for test
//             0
//            /  \
//           1    7
//          / \   /\
//         2  4  5  8
//          \ | / \
//            3    6
func createUndirectedGraph4Test(t *testing.T) *UndirectedGraph {
	// new graph
	g := NewUndirectedGraph("UndirectedGraph")
	// new vertex
	n0 := NewVertex("node0", 0)
	n1 := NewVertex("node1", 1)
	n2 := NewVertex("node2", 2)
	n3 := NewVertex("node3", 3)
	n4 := NewVertex("node4", 4)
	n5 := NewVertex("node5", 5)
	n6 := NewVertex("node6", 6)
	n7 := NewVertex("node7", 7)
	n8 := NewVertex("node8", 8)
	// add vertex
	if g.InsertVertex(n0) != nil ||
	g.InsertVertex(n1) != nil ||
	g.InsertVertex(n2) != nil ||
	g.InsertVertex(n3) != nil ||
	g.InsertVertex(n4) != nil ||
	g.InsertVertex(n5) != nil ||
	g.InsertVertex(n6) != nil ||
	g.InsertVertex(n7) != nil ||
	g.InsertVertex(n8) != nil {
		t.Error("InsertVertex error")
	}
	// add edge
	if g.InsertEdge(n0, n1, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n0, n7, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n1, n2, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n1, n4, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n7, n5, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n7, n8, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n2, n3, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n4, n3, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n5, n3, NewEdge(0, UndirectedEdge)) != nil ||
	g.InsertEdge(n5, n6, NewEdge(0, UndirectedEdge)) != nil {
		t.Error("InsertEdge error")
	}

	return g
}