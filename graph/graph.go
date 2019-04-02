package graph

import (
	"fmt"
	"sync"
)

/*****************************************  graph interface  *****************************************/

// define for graph type
type GraphType string

// define for graph interface
type GraphInterface interface {
	/////// meta data ///////
	// update
	SetName(string)
	// read
	Name() string

	/////// relation data ///////
	// create
	InsertVertex(VertexInterface) error
	InsertEdge(src, dst VertexInterface, ei EdgeInterface) error
	// read
	GetVertex(name string) VertexInterface
	Verteces() map[string]VertexInterface
	// update
	SetVertex(v VertexInterface) error
	SetEdge(src, dst VertexInterface, ei EdgeInterface) error
	// delete
	RemoveVertex(VertexInterface)
	RemoveEdge(src, dst VertexInterface) error
}

/*****************************************  graph struct  *****************************************/
type AbstractGraph struct {
	name      string
	verteces  map[string]VertexInterface
	// mutex
	mutex sync.RWMutex
}

func NewGraph(name string) *AbstractGraph {
	return &AbstractGraph{
		name:     name,
		verteces: make(map[string]VertexInterface),
	}
}

// update graph name
func (g *AbstractGraph) SetName(n string) {
	defer g.mutex.Unlock()
	g.mutex.Lock()
	g.name = n
}

// read graph name
func (g *AbstractGraph) Name() string {
	defer g.mutex.RUnlock()
	g.mutex.RLock()
	return g.name
}

// insert a new vertex
func (g *AbstractGraph) InsertVertex(v VertexInterface) error {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	if _, ok := g.verteces[v.Name()]; ok {
		return fmt.Errorf("vertex[name:%s] already exists!", v.Name())
	}

	g.verteces[v.Name()] = v

	return nil
}

// insert a new edge which is from src vertex to dst vertex
func (g *AbstractGraph) InsertEdge(src, dst VertexInterface, ei EdgeInterface) error {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	if _, ok := g.verteces[src.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", src.Name())
	}

	if _, ok := g.verteces[dst.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", dst.Name())
	}

	src.AdjoinBackward(dst, ei)

	return nil
}

// get a vertex by name
func (g *AbstractGraph) GetVertex(name string) VertexInterface {
	defer g.mutex.RUnlock()
	g.mutex.RLock()

	v, ok := g.verteces[name]
	if !ok {
		return nil
	}

	return v
}

// get name-vertex map
func (g *AbstractGraph) Verteces() map[string]VertexInterface {
	defer g.mutex.RUnlock()
	g.mutex.RLock()

	return g.verteces
}

// update a vertex with specified name, and return an error if the name not exist
func (g *AbstractGraph) SetVertex(v VertexInterface) error {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	if _, ok := g.verteces[v.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", v.Name())
	}

	g.verteces[v.Name()] = v

	return nil
}

// update a edge which is from src verte to dst vertex
func (g *AbstractGraph) SetEdge(src, dst VertexInterface, ei EdgeInterface) error {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	if _, ok := g.verteces[src.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", src.Name())
	}

	if _, ok := g.verteces[dst.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", dst.Name())
	}

	src.SetEdge(dst, ei)

	return nil
}

// remove a vertex in graph, and those verteces which point to this vertex will also remove it
func (g *AbstractGraph) RemoveVertex(v VertexInterface) {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	for _, src := range g.verteces {
		v.RemoveAdjoin(src)
		src.RemoveAdjoin(v)
	}
	delete(g.verteces, v.Name())
}

// remove a edge in graph which is from src verte to dst vertex
func (g *AbstractGraph) RemoveEdge(src, dst VertexInterface) error {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	if _, ok := g.verteces[src.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", src.Name())
	}

	if _, ok := g.verteces[dst.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", dst.Name())
	}

	src.RemoveAdjoin(dst)

	return nil
}
