package graph

import (
	"fmt"
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

// testing for graph bfs
func Test4BFS(t *testing.T) {
	t.Logf("testing for graph bfs start.\n")
	g := createUnirectedGraph4Test(t)
	sum := 0
	f := func(v VertexInterface) {
		fmt.Println(v.Name())
		sum += v.Data().(int)
	}

	BFS(g, f)

	fmt.Println("data sum:", sum)
}

// testing for graph dfs
func Test4DFS(t *testing.T) {
	t.Logf("testing for graph dfs start.\n")
	// new graph
	g := createUnirectedGraph4Test(t)

	sum := 0
	f := func(v VertexInterface) {
		fmt.Println(v.Name())
		sum += v.Data().(int)
	}
	DFS(g, f)
	fmt.Println("data sum:", sum)
}

/// create directed graph for test
//             0
//            /  \
//           1    7
//          / \   /\
//         2  4  5  8
//          \ | / \
//            3    6
func createDirectedGraph4Test(t *testing.T) *DirectedGraph {
	// new graph
	g := NewDirectedGraph("DirectedGraph")
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
	if g.InsertEdge(n0, n1, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n0, n7, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n1, n2, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n1, n4, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n7, n5, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n7, n8, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n2, n3, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n4, n3, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n5, n3, NewEdge(0, BackwardEdge)) != nil ||
	g.InsertEdge(n5, n6, NewEdge(0, BackwardEdge)) != nil {
		t.Error("InserEdge error")
	}

	return g
}

/// create undrected graph for test
//             0
//            /  \
//           1    7
//          / \   /\
//         2  4  5  8
//          \ | / \
//            3    6
func createUnirectedGraph4Test(t *testing.T) *UndirectedGraph {
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

/// [Inteface]
// print graph method
func graphSortPrint(t *testing.T, g GraphInterface) {
	t.Logf("AbstractGraph node number:%d\n", len(g.Verteces()))

	t.Log("*****************************original*****************************")
	for _, v := range g.Verteces() {
		printVertex(t, v)
	}

	vList, err := TopoSort(g)
	if nil != err {
		t.Error(err)
	}

	t.Log("*****************************sort*****************************")
	for _, v := range vList {
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
