package lem

type Nest struct {
	Rooms []*Room
	Start *Room
	End   *Room
}

type Room struct {
	Name         string
	Neighbors    []*Room
	X            int
	Y            int
	Start        bool
	End          bool
	Visited      bool
	PathNode     bool
	CurrentLevel bool
	Predecessor  *Room
	Level        int
	CoLevel      int
}

type Path struct {
	Penultimate *Room
	Steps       int
}
