package lem

type Nest struct {
	Rooms []*Room
	Start *Room
	End   *Room
}

type Room struct {
	Name        string
	Neighbors   []*Room
	Flow        map[*Room]bool
	X           int
	Y           int
	Start       bool
	End         bool
	Predecessor *Room
}

type Path struct {
	Second   *Room
	Rooms    []*Room
	Ants     int
	FirstAnt int
}
