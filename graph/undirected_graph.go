package graph

import (
	"fmt"
)

/**********************************************************************************/
// undirected graph struct
/**********************************************************************************/

type UndirectedGraph struct {
	*AbstractGraph
}

func NewUndirectedGraph(name string) *UndirectedGraph {
	return &UndirectedGraph{
		AbstractGraph: NewGraph(name),
	}
}

func (g *UndirectedGraph) InsertEdge(src, dst VertexInterface, ei EdgeInterface) error {
	if ei.Type() != UndirectedEdge {
		return fmt.Errorf("Edge type(%s) wrong! Edge in undirected graph must be undirected.", ei.Type())
	}
	return g.AbstractGraph.InsertEdge(src, dst, ei)
}
