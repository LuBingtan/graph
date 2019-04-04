package graph

import "sync"

/**********************************************************************************/
// edge interface
/**********************************************************************************/

// define for edge type
type EdgeType string

const (
	ForwardEdge    EdgeType = "forward"
	BackwardEdge   EdgeType = "backward"
	UndirectedEdge EdgeType = "undirected"
)

// define for edge interface
type EdgeInterface interface {
	/////// meta data ///////
	// update
	SetType(EdgeType)
	SetWeight(int)
	// read
	Type() EdgeType
	Weight() int

	/////// relation data ///////
	// update
	SetVertex(from, to VertexInterface)
	// read
	From() VertexInterface
	To() VertexInterface

	/////// copy ///////
	Copy() EdgeInterface
}

/**********************************************************************************/
// edge struct
/**********************************************************************************/

type AbstractEdge struct {
	// meta data
	edgeType EdgeType
	weight   int
	// graph data
	fromVertex VertexInterface
	toVertex   VertexInterface
	// mutex
	mutex sync.RWMutex
}

func NewEdge() *AbstractEdge {
	return &AbstractEdge{}
}

func (e *AbstractEdge) SetType(t EdgeType) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.edgeType = t
}

func (e *AbstractEdge) SetWeight(w int) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.weight = w
}

func (e *AbstractEdge) Type() EdgeType {
	defer e.mutex.RUnlock()
	e.mutex.RLock()
	return e.edgeType
}

func (e *AbstractEdge) Weight() int {
	defer e.mutex.RUnlock()
	e.mutex.RLock()
	return e.weight
}

func (e *AbstractEdge) SetVertex(from, to VertexInterface) {
	defer e.mutex.Unlock()
	e.mutex.Lock()
	e.fromVertex = from
	e.toVertex = to
}

func (e *AbstractEdge) From() VertexInterface {
	defer e.mutex.RUnlock()
	e.mutex.RLock()
	return e.fromVertex
}

func (e *AbstractEdge) To() VertexInterface {
	defer e.mutex.RUnlock()
	e.mutex.RLock()
	return e.toVertex
}

func (e *AbstractEdge) Copy() EdgeInterface {
	return &AbstractEdge{
		weight: e.weight,
	}
}
