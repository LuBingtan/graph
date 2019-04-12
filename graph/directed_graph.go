package graph

import (
	"fmt"
)

/**********************************************************************************/
// directed graph struct
/**********************************************************************************/

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
