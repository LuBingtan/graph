package graph

import (
	"fmt"
	simpleSt "fsmgraph-lib/simplestructure"
	"reflect"
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
	SetEdge(dst VertexInterface, ei EdgeInterface)
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
}

func NewVertex(id string, data interface{}) *AbstractVertex {
	return &AbstractVertex{
		id:   id,
		data: data,
	}
}

// Update vertex type
func (v *AbstractVertex) SetType(t VertexType) {
	v.vertexType = t
}

// Update vertex id
func (v *AbstractVertex) SetId(id string) {
	v.id = id
}

// Update vertex data
func (v *AbstractVertex) SetData(data interface{}) {
	v.data = data
}

// Update vertex state
func (v *AbstractVertex) SetState(state VertexState) {
	v.state = state
}

// get vertex type
func (v *AbstractVertex) Type() VertexType {
	return v.vertexType
}

// get vertex id
func (v *AbstractVertex) Id() string {
	return v.id
}

// get vertex data
func (v *AbstractVertex) Data() interface{} {
	return v.data
}

// get vertex state
func (v *AbstractVertex) State() VertexState {
	return v.state
}

// vertex behavior set
func (v *AbstractVertex) SetExecutor(executorFunc ExecutorFunc) {
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
	ei.SetVertex(dst)
	v.edges.Pushback(ei)
	v.incOutdegree()
	dst.incIndegree()
}

// update edge
func (v *AbstractVertex) SetEdge(dst VertexInterface, ei EdgeInterface) {
	v.RemoveAdjoin(dst)
	v.Adjoin(dst, ei)
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

	v.edges.Remove(index)
	v.decOutdegree()
	vi.decIndegree()
}

// find adjacent vertex and return its binding edge's index
func (v *AbstractVertex) FindAdjoinVertex(vi VertexInterface) int {
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
	for _, d := range v.edges.Data() {
		ei = append(ei, d.(EdgeInterface))
	}

	return ei
}

// Indegree
func (v *AbstractVertex) Indegree() int {
	return v.indegree
}

// outdegree
func (v *AbstractVertex) Outdegree() int {
	return v.outdegree
}
