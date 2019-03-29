package graph

import simpleSt "fsmgraph-lib/simplestructure"

func TopoSort(g *AbstractGraph) (sortVertexList []VertexInterface, err error) {
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
		edgeList := v.Edges()
		for _, edge := range edgeList {
			adjoinId := edge.Vertex().Id()
			indgreeMap[adjoinId]--
			if indgreeMap[adjoinId] == 0 {
				idQueue.Pushback(adjoinId)
			}

		}
	}

	return sortVertexList, nil
}
