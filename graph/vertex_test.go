package graph

import (
	"testing"
)

/// test for vertex
//             0
//            /  \
//           1    7
//          / \   /\
//         2  4  5  8
//          \ | / \
//            3    6
func Test4Vertex(t *testing.T) {
	t.Logf("testing for vertex start.\n")
	n0 := NewVertex("node0", 0)
	n1 := NewVertex("node1", 1)
	n2 := NewVertex("node2", 2)
	n3 := NewVertex("node3", 3)
	n4 := NewVertex("node4", 4)
	n5 := NewVertex("node5", 5)
	n6 := NewVertex("node6", 6)
	n7 := NewVertex("node7", 7)
	n8 := NewVertex("node8", 8)

	Adjoin(n0, n1, NewEdge(0, BackwardEdge))
	Adjoin(n7, n0, NewEdge(0, ForwardEdge)) // same as: n0.Adjoin(n7, NewEdge(0, BackwardEdge))
	Adjoin(n1, n2, NewEdge(0, BackwardEdge))
	Adjoin(n1, n4, NewEdge(0, BackwardEdge))
	Adjoin(n7, n5, NewEdge(0, BackwardEdge))
	Adjoin(n7, n8, NewEdge(0, BackwardEdge))
	Adjoin(n2, n3, NewEdge(0, BackwardEdge))
	Adjoin(n4, n3, NewEdge(0, BackwardEdge))
	Adjoin(n5, n3, NewEdge(0, BackwardEdge))
	Adjoin(n5, n6, NewEdge(0, BackwardEdge))

	printVertex := func(v VertexInterface) {
		t.Logf("========= v_name:%s, indegree:%d, outdegree:%d =========\n", v.Name(), v.Indegree(), v.Outdegree())
		for _, e := range v.EdgesForward() {
			t.Logf("edge_forward:%s -> %s", e.From().Name(), e.To().Name())
		}

		for _, e := range v.EdgesBackward() {
			t.Logf("edge_backward:%s -> %s", e.From().Name(), e.To().Name())
		}
	}

	t.Logf("-----------------------------------BFS-----------------------------------")
	BFSVertex(n0, make(map[string]bool), printVertex)

	t.Logf("-----------------------------------DFS-----------------------------------")
	DFSVertex(n0, make(map[string]bool), printVertex)

	t.Logf("-----------------------------------delete 0 -> 7-----------------------------------")
	RemoveAdjoin(n0, n7)
	BFSVertex(n0, make(map[string]bool), printVertex)

	t.Logf("-----------------------------------delete 1 -> 4-----------------------------------")
	RemoveAdjoin(n4, n1) // same as RemoveAdjoin(n4, n1)
	BFSVertex(n0, make(map[string]bool), printVertex)
}
