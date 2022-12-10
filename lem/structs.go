package lem

type Room struct {
	Name      string
	Neighbors []string
	X         int
	Y         int
	Start     bool
	End       bool
}
