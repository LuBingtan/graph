package graph

import "fmt"

/**********************************************************************************/
// directed graph interface
/**********************************************************************************/

// define for directed graph interface
type UndirectedGraphInterface interface {
	GraphInterface
}

type UndirectedGraph struct {
	AbstractGraph
}

// insert a new edge which is from src vertex to dst vertex
func (g *UndirectedGraph) InsertEdge(src, dst VertexInterface, ei EdgeInterface) error {
	if _, ok := g.verteces[src.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", src.Name())
	}

	if _, ok := g.verteces[dst.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", dst.Name())
	}

	ei.SetType(UndirectedEdge)

	defer g.mutex.Unlock()
	g.mutex.Lock()

	return src.Adjoin(dst, ei)
}
