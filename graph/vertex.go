package graph

import (
	"fmt"
	"sync"

	simpleSt "graph/simplestructure"
)

/**********************************************************************************/
// vertex interface
/**********************************************************************************/

// define for vertex type
type VertexType string

// define for vertex interface
// vertex contains edge
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
	InsertEdge(ei EdgeInterface)
	// delete
	RemoveEdge(endpoint VertexInterface, edgeType EdgeType) EdgeInterface
	// read
	FindEdge(endpoint VertexInterface, edgeType EdgeType) EdgeInterface
	// list
	Edges() []EdgeInterface
	EdgesForward() []EdgeInterface
	EdgesBackward() []EdgeInterface
	Indegree() int
	Outdegree() int

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

func (v *AbstractVertex) InsertEdge(ei EdgeInterface) {
	defer v.mutex.Unlock()
	v.mutex.Lock()
	switch ei.Type() {
	case BackwardEdge:
		v.incOutdegree()
	case ForwardEdge:
		v.incIndegree()
	case UndirectedEdge:
		v.incIndegree()
		v.incOutdegree()
	default:
		panic(fmt.Sprintf("Unknown type[%s].", ei.Type()))
	}

	v.edges.Pushback(ei)
}

func (v *AbstractVertex) RemoveEdge(endpoint VertexInterface, edgeType EdgeType) EdgeInterface {
	index := v.findEdge(endpoint, edgeType)

	v.mutex.Lock()
	e := v.edges.Remove(index)
	v.mutex.Unlock()

	if nil == e {
		return nil
	}

	switch edgeType {
	case BackwardEdge:
		v.decOutdegree()
	case ForwardEdge:
		v.decIndegree()
	case UndirectedEdge:
		v.decIndegree()
		v.decOutdegree()
	}

	return e.(EdgeInterface)
}

// Update edge
func (v *AbstractVertex) FindEdge(endpoint VertexInterface, edgeType EdgeType) EdgeInterface {
	index := v.findEdge(endpoint, edgeType)
	if index == -1 {
		return nil
	}

	e := v.edges.At(index)
	if e == nil {
		return nil
	}

	return e.(EdgeInterface)
}

func (v *AbstractVertex) findEdge(endpoint VertexInterface, edgeType EdgeType) int {
	for i, e := range v.edges.Data() {
		edge := e.(EdgeInterface)
		if (edge.From().Name() == endpoint.Name() ||
		edge.To().Name() == endpoint.Name()) &&
		edge.Type() == edgeType {
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

/**********************************************************************************/
// vertex interface function
/**********************************************************************************/

// Insert adjacent vertex
// If edge type is forward edge, add input vertex as forward vertex and increase indegree
// If edge type is backward edge, add input vertex as backward vertex, and increase outdegree
// When Adjoin is done, it will continue 'Adjoin' itself to the 'next' vertex
func Adjoin(from, to VertexInterface, ei EdgeInterface) error {
	edge := from.FindEdge(to, ei.Type())
	if edge != nil {
		return nil
	}

	reverseEdge := ei.Copy()
	switch ei.Type() {
	case BackwardEdge:
		ei.SetVertex(from, to)
		reverseEdge.SetVertex(from, to)
		reverseEdge.SetType(ForwardEdge)
	case ForwardEdge:
		ei.SetVertex(to, from)
		reverseEdge.SetVertex(to, from)
		reverseEdge.SetType(BackwardEdge)
	case UndirectedEdge:
		ei.SetVertex(from, to)
		reverseEdge.SetVertex(from, to)
		reverseEdge.SetType(UndirectedEdge)
	default:
		return fmt.Errorf("Unknown type[%s].", ei.Type())
	}

	from.InsertEdge(ei)
	to.InsertEdge(reverseEdge)

	return nil
}

// Delete adjacent vertex
// If not find target vertex, do nothing
// When find target vertex:
// if it's a backward vertex, decrease of self's outdegree and target's indegree
// if it's a forward vertex, decrease of self's indegree and target's outdegree
func RemoveAdjoin(from, to VertexInterface) error {
	edgeTypeList := []EdgeType{BackwardEdge, ForwardEdge, UndirectedEdge}
	for _, edgeType := range edgeTypeList {
		fmt.Println(edgeType)
		var reverseEdgeType EdgeType
		switch edgeType {
		case BackwardEdge:
			reverseEdgeType = ForwardEdge
		case ForwardEdge:
			reverseEdgeType = BackwardEdge
		case UndirectedEdge:
			reverseEdgeType = UndirectedEdge
		}
		from.RemoveEdge(to, edgeType)
		to.RemoveEdge(from, reverseEdgeType)
	}

	return nil
}