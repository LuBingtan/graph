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
	SetType(VertexType)
	SetId(string)
	SetData(interface{})
	SetState(VertexState)
	// read
	Type() VertexType
	Id() string
	Data() interface{}
	State() VertexState
	// behavior
	SetExecutor(ExecutorFunc)
	Execute(...interface{}) (interface{}, error)

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
	id         string
	data       interface{}
	state      VertexState
	executor   ExecutorFunc
	// graph data
	edges     simpleSt.SimpleVector
	indegree  int
	outdegree int
	// mutex
	lock sync.RWMutex
}

func NewVertex(id string, data interface{}) *AbstractVertex {
	return &AbstractVertex{
		id:   id,
		data: data,
	}
}

// Update vertex type
func (v *AbstractVertex) SetType(t VertexType) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.vertexType = t
}

// Update vertex id
func (v *AbstractVertex) SetId(id string) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.id = id
}

// Update vertex data
func (v *AbstractVertex) SetData(data interface{}) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.data = data
}

// Update vertex state
func (v *AbstractVertex) SetState(state VertexState) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.state = state
}

// get vertex type
func (v *AbstractVertex) Type() VertexType {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.vertexType
}

// get vertex id
func (v *AbstractVertex) Id() string {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.id
}

// get vertex data
func (v *AbstractVertex) Data() interface{} {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.data
}

// get vertex state
func (v *AbstractVertex) State() VertexState {
	defer v.lock.RUnlock()
	v.lock.RLock()
	return v.state
}

// vertex behavior set
func (v *AbstractVertex) SetExecutor(executorFunc ExecutorFunc) {
	defer v.lock.Unlock()
	v.lock.Lock()
	v.executor = executorFunc
}

// vertex behavior execute
func (v *AbstractVertex) Execute(inputs ...interface{}) (interface{}, error) {
	if v.executor == nil {
		return nil, fmt.Errorf("no executor.")
	}

	return v.executor(inputs)
}

// update adjacent vertex
func (v *AbstractVertex) Adjoin(dst VertexInterface, ei EdgeInterface) {
	defer v.lock.Unlock()
	v.lock.Lock()

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

	defer v.lock.Unlock()
	v.lock.Lock()

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

	defer v.lock.Unlock()
	v.lock.Lock()

	v.edges.Remove(index)
	v.decOutdegree()
	vi.decIndegree()
}

// find adjacent vertex and return its binding edge's index
func (v *AbstractVertex) FindAdjoinVertex(vi VertexInterface) int {
	defer v.lock.RUnlock()
	v.lock.RLock()

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
	defer v.lock.RUnlock()
	v.lock.RLock()

	for _, d := range v.edges.Data() {
		ei = append(ei, d.(EdgeInterface))
	}

	return ei
}

// Indegree
func (v *AbstractVertex) Indegree() int {
	defer v.lock.RUnlock()
	v.lock.RLock()

	return v.indegree
}

// outdegree
func (v *AbstractVertex) Outdegree() int {
	defer v.lock.RUnlock()
	v.lock.RLock()

	return v.outdegree
}
