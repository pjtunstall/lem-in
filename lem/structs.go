package lem

type Nest struct {
	Rooms []*Room
	Start *Room
	End   *Room
}

type Room struct {
	Name        string
	Neighbors   []*Room
	Residual    map[*Room]int
	Flow        map[*Room]int
	IsFlowing   bool
	X           int
	Y           int
	Start       bool
	End         bool
	Predecessor *Room
}

type Path struct {
	Penultimate *Room
	Steps       int
}
