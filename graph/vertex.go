package graph

import (
	"fmt"
	"reflect"
	"sync"

	simpleSt "graph/simplestructure"
)

/*****************************************  vertex interface  *****************************************/

// define for vertex type
type VertexType string

// define for vertex state stype
type VertexState string

// define for vertex executor function type
type ExecutorFunc func(...interface{}) (interface{}, error)

// define for vertex interface
type VertexInterface interface {
	/////// meta data ///////
	// update
	SetId(string)
	SetData(interface{})
	// read
	Name() string
	Data() interface{}

	/////// relation data ///////
	// update
	Adjoin(dst VertexInterface, ei EdgeInterface)
	SetEdge(dst VertexInterface, ei EdgeInterface) error
	incIndegree()
	decIndegree()
	incOutdegree()
	decOutdegree()
	// delete
	RemoveAdjoin(VertexInterface)
	// read
	FindAdjoinVertex(VertexInterface) int
	Edges() []EdgeInterface
	Indegree() int
	Outdegree() int
}

/*****************************************  vertex struct  *****************************************/

/// [Define]
// structure for vertex
type AbstractVertex struct {
	// meta data
	vertexType VertexType
	name         string
	data       interface{}
	// graph data
	edges     simpleSt.SimpleVector
	indegree  int
	outdegree int
	// mutex
	mutex sync.RWMutex
}

func NewVertex(name string, data interface{}) *AbstractVertex {
	return &AbstractVertex{
		name:   name,
		data: data,
	}
}

// Update vertex type
func (v *AbstractVertex) SetType(t VertexType) {
	defer v.mutex.Unlock()
	v.mutex.Lock()
	v.vertexType = t
}

// Update vertex name
func (v *AbstractVertex) SetId(name string) {
	defer v.mutex.Unlock()
	v.mutex.Lock()
	v.name = name
}

// Update vertex data
func (v *AbstractVertex) SetData(data interface{}) {
	defer v.mutex.Unlock()
	v.mutex.Lock()
	v.data = data
}

// get vertex type
func (v *AbstractVertex) Type() VertexType {
	defer v.mutex.RUnlock()
	v.mutex.RLock()
	return v.vertexType
}

// get vertex name
func (v *AbstractVertex) Name() string {
	defer v.mutex.RUnlock()
	v.mutex.RLock()
	return v.name
}

// get vertex data
func (v *AbstractVertex) Data() interface{} {
	defer v.mutex.RUnlock()
	v.mutex.RLock()
	return v.data
}

// update adjacent vertex
func (v *AbstractVertex) Adjoin(dst VertexInterface, ei EdgeInterface) {
	defer v.mutex.Unlock()
	v.mutex.Lock()

	ei.SetVertex(dst)
	v.edges.Pushback(ei)
	v.incOutdegree()
	dst.incIndegree()
}

// update edge
func (v *AbstractVertex) SetEdge(dst VertexInterface, ei EdgeInterface) error {
	index := v.FindAdjoinVertex(dst)
	if index == -1 {
		return fmt.Errorf("vertex(%v) not exists.", dst)
	}

	defer v.mutex.Unlock()
	v.mutex.Lock()

	edgeI := v.edges.At(index)
	edge := edgeI.(EdgeInterface)
	ei.SetVertex(edge.Vertex())

	return v.edges.Replace(index, ei)
}

// increase indegree
func (v *AbstractVertex) incIndegree() {
	v.indegree++
}

// decrease indegree
func (v *AbstractVertex) decIndegree() {
	v.indegree--
}

// increase outdegree
func (v *AbstractVertex) incOutdegree() {
	v.outdegree++
}

// decrease outdegree
func (v *AbstractVertex) decOutdegree() {
	v.outdegree--
}

// delete adjacent vertex
func (v *AbstractVertex) RemoveAdjoin(vi VertexInterface) {
	index := v.FindAdjoinVertex(vi)
	if index == -1 {
		return
	}

	defer v.mutex.Unlock()
	v.mutex.Lock()

	v.edges.Remove(index)
	v.decOutdegree()
	vi.decIndegree()
}

// find adjacent vertex and return its binding edge's index
func (v *AbstractVertex) FindAdjoinVertex(vi VertexInterface) int {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	for i, d := range v.edges.Data() {
		e := d.(EdgeInterface)
		if reflect.DeepEqual(e.Vertex(), vi) {
			return i
		}
	}

	return -1
}

// get all edges
func (v *AbstractVertex) Edges() (ei []EdgeInterface) {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	for _, d := range v.edges.Data() {
		ei = append(ei, d.(EdgeInterface))
	}

	return ei
}

// Indegree
func (v *AbstractVertex) Indegree() int {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	return v.indegree
}

// outdegree
func (v *AbstractVertex) Outdegree() int {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	return v.outdegree
}
