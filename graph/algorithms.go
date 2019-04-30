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

// Graph BFS
// search start from a root vertex
func BFS(g GraphInterface, executeFunc func(VertexInterface)) {
	hasVisted := make(map[string]bool)
	for _, v := range g.Verteces() {
		if _, ok := hasVisted[v.Name()]; !ok {
			BFSVertex(v, hasVisted, executeFunc)
		}
	}
}
func BFSVertex(root VertexInterface, hasVisted map[string]bool, executeFunc func(VertexInterface)) {
	vQueue := simpleSt.NewSimpleQueue()
	vQueue.Pushback(root)
	hasVisted[root.Name()] = true
	for {
		vi := vQueue.Popfront()
		if vi == nil {
			break
		}
		v := vi.(VertexInterface)
		executeFunc(v)

		for _, edge := range v.EdgesBackward() {
			adj := edge.To()
			if _, ok := hasVisted[adj.Name()]; ok {
				continue
			} else {
				vQueue.Pushback(adj)
				hasVisted[adj.Name()] = true
			}
		}
	}
}

// Graph BFS
// search start from a root vertex
func DFS(g GraphInterface, executeFunc func(VertexInterface)) {
	hasVisted := make(map[string]bool)
	for _, v := range g.Verteces() {
		if _, ok := hasVisted[v.Name()]; !ok {
			DFSVertex(v, hasVisted, executeFunc)
		}
	}
}

func DFSVertex(root VertexInterface, hasVisted map[string]bool, executeFunc func(VertexInterface)) {
	hasVisted[root.Name()] = true
	executeFunc(root)
	if len(root.EdgesBackward()) == 0 {
		return
	}

	for _, edge := range root.EdgesBackward() {
		adj := edge.To()
		if _, ok := hasVisted[adj.Name()]; !ok {
			hasVisted[adj.Name()] = true
			DFSVertex(adj, hasVisted, executeFunc)
		}
	}
}
