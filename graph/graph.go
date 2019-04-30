package graph

import (
	"fmt"
	"sync"
)

/**********************************************************************************/
// graph interface
/**********************************************************************************/

// define for graph type
type GraphType string

// define for graph interface
type GraphInterface interface {
	/////// meta data ///////
	// update
	SetName(string)
	SetType(GraphType)
	// read
	Name() string
	Type() GraphType

	/////// relation data ///////
	// create
	InsertVertex(VertexInterface) error
	InsertEdge(src, dst VertexInterface, ei EdgeInterface) error
	// read
	GetVertex(name string) VertexInterface
	Verteces() map[string]VertexInterface
	// update
	UpdateVertex(v VertexInterface) error
	// delete
	RemoveVertex(VertexInterface)
	RemoveEdge(src, dst VertexInterface) error

	/////// copy ///////
	Clone() GraphInterface
}

/**********************************************************************************/
// graph struct
/**********************************************************************************/

type AbstractGraph struct {
	name     string
	gType    GraphType
	verteces map[string]VertexInterface
	// mutex
	mutex sync.RWMutex
}

func NewGraph(name string) *AbstractGraph {
	return &AbstractGraph{
		name:     name,
		verteces: make(map[string]VertexInterface),
	}
}

/**********************************************************************************/
// implementation for interface method
/**********************************************************************************/

// update graph name
func (g *AbstractGraph) SetName(n string) {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	g.name = n
}

// update graph type
func (g *AbstractGraph) SetType(t GraphType) {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	g.gType = t
}

// read graph name
func (g *AbstractGraph) Name() string {
	defer g.mutex.RUnlock()
	g.mutex.RLock()

	return g.name
}

// read graph type
func (g *AbstractGraph) Type() GraphType {
	defer g.mutex.Unlock()
	g.mutex.Lock()

	return g.gType
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
	if _, ok := g.verteces[src.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", src.Name())
	}

	if _, ok := g.verteces[dst.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", dst.Name())
	}

	defer g.mutex.Unlock()
	g.mutex.Lock()

	return Adjoin(src, dst, ei)
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
func (g *AbstractGraph) UpdateVertex(v VertexInterface) error {
	if _, ok := g.verteces[v.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", v.Name())
	}

	defer g.mutex.Unlock()
	g.mutex.Lock()

	g.verteces[v.Name()].SetData(v.Data())

	return nil
}

// remove a vertex in graph, and those verteces which point to this vertex will also remove it
func (g *AbstractGraph) RemoveVertex(v VertexInterface) {
	if v == nil {
		panic("Input is null, pleae do check!")
	}

	defer g.mutex.Unlock()
	g.mutex.Lock()

	for _, src := range g.verteces {
		RemoveAdjoin(v, src)
	}

	delete(g.verteces, v.Name())
}

// remove a edge in graph which is from src verte to dst vertex
func (g *AbstractGraph) RemoveEdge(src, dst VertexInterface) error {
	if nil == src || nil == dst {
		return fmt.Errorf("Input is null, pleae do check!")
	}

	if _, ok := g.verteces[src.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", src.Name())
	}

	if _, ok := g.verteces[dst.Name()]; !ok {
		return fmt.Errorf("vertex[name:%s] not exists, insert vertex first!", dst.Name())
	}

	defer g.mutex.Unlock()
	g.mutex.Lock()

	RemoveAdjoin(src, dst)

	return nil
}

func (g *AbstractGraph) Clone() GraphInterface {
	newG := NewGraph(g.name)
	for _, v := range g.verteces {
		newG.InsertVertex(v.Copy())
	}

	for _, v := range g.verteces {
		for _, ei := range v.EdgesBackward() {
			newG.InsertEdgeByName(ei.From().Name(), ei.To().Name(), ei.Copy())
		}
	}

	return newG
}

/**********************************************************************************/
// struct method
/**********************************************************************************/

func (g *AbstractGraph) InsertEdgeByName(srcName, dstName string, ei EdgeInterface) error {
	return g.InsertEdge(g.GetVertex(srcName), g.GetVertex(dstName), ei)
}
