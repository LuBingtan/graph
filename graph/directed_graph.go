package graph

import (
	"fmt"
)

/**********************************************************************************/
// directed graph
/**********************************************************************************/
type DirectedGraphInterface interface {
	GraphInterface
}

type DirectedGraph struct {
	*AbstractGraph
}

func NewDirectedGraph(name string) *DirectedGraph {
	return &DirectedGraph{
		AbstractGraph: NewGraph(name),
	}
}

func (g *DirectedGraph) InsertEdge(src, dst VertexInterface, ei EdgeInterface) error {
	if ei.Type() == UndirectedEdge {
		return fmt.Errorf("Edge type(%s) wrong! Edge in directed graph must be backward or forward.", ei.Type())
	}
	return g.AbstractGraph.InsertEdge(src, dst, ei)
}

func (g *DirectedGraph) InsertEdgeByName(srcName, dstName string, ei EdgeInterface) error {
	return g.InsertEdge(g.GetVertex(srcName), g.GetVertex(dstName), ei)
}

/**********************************************************************************/
// directed acyclic graph
/**********************************************************************************/

type DAGInterface interface {
	DirectedGraphInterface
	IsDag() bool
}

type DAG struct {
	*DirectedGraph
}

func NewDAG(name string) *DAG {
	return &DAG{
		DirectedGraph: NewDirectedGraph(name),
	}
}

func (g *DAG) IsDag() bool {
	sortVertexList, err := TopoSort(g)
	if err != nil {
		return false
	}

	if len(sortVertexList) < len(g.Verteces()) {
		return false
	}

	return true
}