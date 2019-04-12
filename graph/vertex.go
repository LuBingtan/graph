package graph

import (
	"fmt"
	"reflect"
	"sync"

	simpleSt "easyai/utils/simplestructure"
)

/**********************************************************************************/
// vertex interface
/**********************************************************************************/

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
	Adjoin(dst VertexInterface, ei EdgeInterface) error
	SetEdge(adj VertexInterface, ei EdgeInterface) error
	// delete
	RemoveAdjoin(VertexInterface)
	// read
	FindAdjoin(VertexInterface) int
	Edges() []EdgeInterface
	EdgesForward() []EdgeInterface
	EdgesBackward() []EdgeInterface
	Indegree() int
	Outdegree() int
	// inner method, degree operation
	incIndegree()
	incOutdegree()
	decOutdegree()
	decIndegree()

	/////// copy ///////
	Copy() VertexInterface
}

/**********************************************************************************/
// vertex struct
/**********************************************************************************/

/// [Define]
// structure for vertex
type AbstractVertex struct {
	// meta data
	name string
	data interface{}
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
		name: name,
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

// Insert adjacent vertex
// If edge type is forward edge, add input vertex as forward vertex and increase indegree
// If edge type is backward edge, add input vertex as backward vertex, and increase outdegree
// When Adjoin is done, it will continue 'Adjoin' itself to the 'next' vertex
func (v *AbstractVertex) Adjoin(next VertexInterface, ei EdgeInterface) error {
	if v.FindAdjoin(next) != -1 {
		return nil
	}

	var reverseEdgeType EdgeType
	switch ei.Type() {
	case BackwardEdge:
		v.incOutdegree()
		ei.SetVertex(v, next)
		reverseEdgeType = ForwardEdge
	case ForwardEdge:
		v.incIndegree()
		ei.SetVertex(next, v)
		reverseEdgeType = BackwardEdge
	case UndirectedEdge:
		v.incIndegree()
		v.incOutdegree()
		ei.SetVertex(v, next)
		reverseEdgeType = UndirectedEdge
	default:
		return fmt.Errorf("Unknown type[%s].", ei.Type())
	}

	v.mutex.Lock()

	v.edges.Pushback(ei)

	v.mutex.Unlock()

	reverseEdge := ei.Copy()
	reverseEdge.SetType(reverseEdgeType)

	return next.Adjoin(v, reverseEdge)
}

// Update edge
// update an edge's inner data such as weight
// the input vertex 'adj' is a adjacent vertex
func (v *AbstractVertex) SetEdge(adj VertexInterface, ei EdgeInterface) error {
	index := v.FindAdjoin(adj)
	if index == -1 {
		return fmt.Errorf("vertex(%v) not exists.", adj)
	}

	// update edge's weight
	defer v.mutex.Unlock()
	v.mutex.Lock()

	edgeI := v.edges.At(index)
	edge := edgeI.(EdgeInterface)
	edge.SetWeight(ei.Weight())

	return nil
}

// Delete adjacent vertex
// If not find target vertex, do nothing
// When find target vertex:
// if it's a backward vertex, decrease of self's outdegree and target's indegree
// if it's a forward vertex, decrease of self's indegree and target's outdegree
func (v *AbstractVertex) RemoveAdjoin(adj VertexInterface) {
	index := v.FindAdjoin(adj)
	if index == -1 {
		return
	}

	// remove edge
	v.mutex.Lock()

	edge := v.edges.At(index).(EdgeInterface)

	v.edges.Remove(index)

	// decrease degree
	if edge.Type() == ForwardEdge {
		v.decIndegree()
	} else {
		v.decOutdegree()
	}

	v.mutex.Unlock()

	// let adjacent vertex remove this vertex
	adj.RemoveAdjoin(v)
}

// Find adjacent vertex and return its binding edge's index
// If not find target vertex, then return -1
// If target vertex is equal with self, return -1
func (v *AbstractVertex) FindAdjoin(vi VertexInterface) int {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	if reflect.DeepEqual(v, vi) {
		return -1
	}

	for i, d := range v.edges.Data() {
		e := d.(EdgeInterface)
		if reflect.DeepEqual(e.From(), vi) || reflect.DeepEqual(e.To(), vi) {
			return i
		}
	}

	return -1
}

// Get all edges
func (v *AbstractVertex) Edges() (ei []EdgeInterface) {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	for _, d := range v.edges.Data() {
		edge := d.(EdgeInterface)
		ei = append(ei, edge)
	}

	return ei
}

// Get all forward edges
func (v *AbstractVertex) EdgesForward() (ei []EdgeInterface) {
	defer v.mutex.RUnlock()
	v.mutex.RLock()

	for _, d := range v.edges.Data() {
		edge := d.(EdgeInterface)
		if edge.Type() == ForwardEdge || edge.Type() == UndirectedEdge {
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
		if edge.Type() == BackwardEdge || edge.Type() == UndirectedEdge {
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

// copy
func (v *AbstractVertex) Copy() VertexInterface {
	return &AbstractVertex{
		name: v.name,
		data: v.data,
	}
}
