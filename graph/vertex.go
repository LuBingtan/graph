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

// define for vertex interface
type VertexInterface interface {
	/////// meta data ///////
	// update
	SetName(string)
	SetData(interface{})
	// read
	Name() string
	Data() interface{}

	/////// relation data ///////
	// update
	AdjoinForward(dst VertexInterface, ei EdgeInterface)
	AdjoinBackward(dst VertexInterface, ei EdgeInterface)
	SetEdge(adj VertexInterface, ei EdgeInterface) error
	// delete
	RemoveAdjoin(VertexInterface)
	// read
	FindAdjoin(VertexInterface) int
	EdgesForward() []EdgeInterface
	EdgesBackward() []EdgeInterface
	Indegree() int
	Outdegree() int
	// inner method, degree operation
	incIndegree()
	incOutdegree()
	decOutdegree()
	decIndegree()
}

/*****************************************  vertex struct  *****************************************/

/// [Define]
// structure for vertex
type AbstractVertex struct {
	// meta data
	name         string
	data       interface{}
	// graph data
	edges     simpleSt.SimpleVector
	indegree  int
	outdegree int
	// mutex
	mutex sync.RWMutex
}

// create new vertex with a name and data
func NewVertex(name string, data interface{}) *AbstractVertex {
	return &AbstractVertex{
		name:   name,
		data: data,
	}
}

// Update vertex name
func (v *AbstractVertex) SetName(name string) {
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

// Get vertex name
func (v *AbstractVertex) Name() string {
	defer v.mutex.RUnlock()
	v.mutex.RLock()
	return v.name
}

// Get vertex data
func (v *AbstractVertex) Data() interface{} {
	defer v.mutex.RUnlock()
	v.mutex.RLock()
	return v.data
}

// Insert forward adjacent vertex
func (v *AbstractVertex) AdjoinForward(pre VertexInterface, ei EdgeInterface) {
	defer v.mutex.Unlock()
	v.mutex.Lock()

	ei.SetType(ForwardEdge)
	ei.SetVertex(pre, v)
	v.edges.Pushback(ei)
	v.incIndegree()
	pre.incOutdegree()
}

// Insert backward adjacent vertex
func (v *AbstractVertex) AdjoinBackward(next VertexInterface, ei EdgeInterface) {
	defer v.mutex.Unlock()
	v.mutex.Lock()

	ei.SetType(BackwardEdge)
	ei.SetVertex(v, next)
	v.edges.Pushback(ei)
	v.incOutdegree()
	next.incIndegree()
}

// Update edge
func (v *AbstractVertex) SetEdge(adj VertexInterface, ei EdgeInterface) error {
	index := v.FindAdjoin(adj)
	if index == -1 {
		return fmt.Errorf("vertex(%v) not exists.", adj)
	}

	defer v.mutex.Unlock()
	v.mutex.Lock()

	edgeI := v.edges.At(index)
	edge := edgeI.(EdgeInterface)
	ei.CopyRelateFrom(edge)

	return v.edges.Replace(index, ei)
}

// Delete adjacent vertex
// If not find target vertex, do nothing
// When find target vertex:
// if it's a backward vertex, decrease of self's outdegree and target's indegree
// if it's a forward vertex, decrease of self's indegree and target's outdegree
func (v *AbstractVertex) RemoveAdjoin(vi VertexInterface) {
	index := v.FindAdjoin(vi)
	if index == -1 {
		return
	}

	edge := v.edges.At(index).(EdgeInterface)

	defer v.mutex.Unlock()
	v.mutex.Lock()

	if edge.Type() == ForwardEdge {
		v.decIndegree()
		vi.decOutdegree()
	} else {
		v.decOutdegree()
		vi.decIndegree()
	}
	v.edges.Remove(index)
}

// Find adjacent vertex and return its binding edge's index
// If not find target vertex, then return -1
// If target vertex is equal with self, return -1
func (v *AbstractVertex) FindAdjoin(vi VertexInterface) int {
	if reflect.DeepEqual(v, vi) {
		return -1
	}

	defer v.mutex.RUnlock()
	v.mutex.RLock()

	for i, d := range v.edges.Data() {
		e := d.(EdgeInterface)
		if reflect.DeepEqual(e.From(), vi) || reflect.DeepEqual(e.To(), vi) {
			return i
		}
	}

	return -1
}

// Get all forward edges
func (v *AbstractVertex) EdgesForward() (ei []EdgeInterface) {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	for _, d := range v.edges.Data() {
		edge := d.(EdgeInterface)
		if edge.Type() == ForwardEdge {
			ei = append(ei, edge)
		}
	}

	return ei
}

// Get all backward edges
func (v *AbstractVertex) EdgesBackward() (ei []EdgeInterface) {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	for _, d := range v.edges.Data() {
		edge := d.(EdgeInterface)
		if edge.Type() == BackwardEdge {
			ei = append(ei, edge)
		}
	}

	return ei
}

// Get indegree
func (v *AbstractVertex) Indegree() int {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	return v.indegree
}

// Get outdegree
func (v *AbstractVertex) Outdegree() int {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	return v.outdegree
}

// degree operation: increase indegree
func (v *AbstractVertex) incIndegree() {
	v.indegree++
}

// degree operation: decrease indegree
func (v *AbstractVertex) decIndegree() {
	v.indegree--
}

// degree operation: increase outdegree
func (v *AbstractVertex) incOutdegree() {
	v.outdegree++
}

// degree operation: decrease outdegree
func (v *AbstractVertex) decOutdegree() {
	v.outdegree--
}