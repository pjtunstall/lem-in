package lem

type Nest struct {
	Rooms []*Room
	Start *Room
	End   *Room
}

type Room struct {
	Name        string
	Neighbors   []*Room
	X           int
	Y           int
	Start       bool
	End         bool
	Predecessor *Room
}

type PathNode struct {
	Node     *Room
	Previous *Room
	Next     *Room
}
