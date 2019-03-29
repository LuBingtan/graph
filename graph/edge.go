package graph

/*****************************************  edge interface  *****************************************/

// define for edge type
type EdgeType string

// define for edge state
type EdgeState string

// define for edge interface
type EdgeInterface interface {
	/////// meta data ///////
	// update
	SetType(EdgeType)
	SetWeight(int)
	SetState(EdgeState)
	// read
	Type() EdgeType
	Weight() int
	State() EdgeState

	/////// relation data ///////
	// update
	SetVertex(VertexInterface)
	// read
	Vertex() VertexInterface
}

/*****************************************  edge struct  *****************************************/

type AbstractEdge struct {
	// meta data
	edgeType EdgeType
	weight   int
	state    EdgeState
	// graph data
	vertex VertexInterface
}

func NewEdge() *AbstractEdge {
	return &AbstractEdge{}
}

func (e *AbstractEdge) SetType(t EdgeType) {
	e.edgeType = t
}

func (e *AbstractEdge) SetWeight(w int) {
	e.weight = w
}

func (e *AbstractEdge) SetState(s EdgeState) {
	e.state = s
}

func (e *AbstractEdge) Type() EdgeType {
	return e.edgeType
}

func (e *AbstractEdge) Weight() int {
	return e.weight
}

func (e *AbstractEdge) State() EdgeState {
	return e.state
}

func (e *AbstractEdge) SetVertex(v VertexInterface) {
	e.vertex = v
}

func (e *AbstractEdge) Vertex() VertexInterface {
	return e.vertex
}
