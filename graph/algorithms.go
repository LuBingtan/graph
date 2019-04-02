package graph

import (
	"fmt"

	simpleSt "graph/simplestructure"
)

// Determine if a graph is acyclic
// Using topological sort
func IsAcyclic(g GraphInterface) error {
	vertices, err := TopoSort(g)
	if err != nil {
		return err
	}

	if len(vertices) < len(g.Verteces()) {
		return fmt.Errorf("not acyclic.")
	}

	return nil
}

// Graph Topological Sort
// If this graph is cyclic, the sorted vertices' number is less than the total vertices in grah
func TopoSort(g GraphInterface) (sortVertexList []VertexInterface, err error) {
	indgreeMap := make(map[string]int)
	idQueue := simpleSt.NewSimpleQueue()

	// put all indegree in map
	verteces := g.Verteces()
	for k, v := range verteces {
		indgreeMap[k] = v.Indegree()
	}

	// find 0 indgree id
	for k, d := range indgreeMap {
		if 0 == d {
			idQueue.Pushback(k)
		}
	}

	// 0 indgree adjoin vertex degree minus 1
	for {
		// get id
		idInterface := idQueue.Popfront()
		if idInterface == nil {
			break
		}
		id := idInterface.(string)
		sortVertexList = append(sortVertexList, g.GetVertex(id))
		// ge vertex
		v := g.GetVertex(id)
		edgeList := v.EdgesBackward()
		for _, edge := range edgeList {
			adjoinId := edge.To().Name()
			indgreeMap[adjoinId]--
			if indgreeMap[adjoinId] == 0 {
				idQueue.Pushback(adjoinId)
			}

		}
	}

	return sortVertexList, nil
}
