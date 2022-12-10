package lem

type Nest struct {
	Rooms []*Room
}

type Room struct {
	Name      string
	Neighbors []*Room
	X         int
	Y         int
	Start     bool
	End       bool
}
