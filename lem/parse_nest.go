package lem

import (
	"fmt"
	"strconv"
	"strings"
)

func findFirstLineWithTunnel(text []string) int {
	for i := 5; i < len(text); i++ {
		if strings.Contains(text[i], "-") && !strings.Contains(text[i], "#") {
			return i
		}
	}
	return 0
}

func nameRoom(name string, room *Room) *Room {
	room.Name = name
	return room
}

func newRoom(row string) (*Room, bool) {
	a := strings.Split(row, " ")
	var problem bool
	if len(a) != 3 {
		fmt.Println("ERROR: Invalid coordinates.")
		problem = true
	}
	x, er1 := strconv.Atoi(a[1])
	y, er2 := strconv.Atoi(a[2])
	if er1 != nil || er2 != nil {
		fmt.Println("ERROR: Invalid coordinates.")
		problem = true
	}
	var room Room
	nameRoom(a[0], &room)
	room.X = x
	room.Y = y
	room.Flow = make(map[*Room]bool)
	return &room, problem
}

func findRoom(name string, counter int, nest *Nest) *Room {
	if counter == 0 {
		return nil
	}
	if nest.Rooms[counter-1].Name == name {
		return nest.Rooms[counter-1]
	}
	counter--
	return findRoom(name, counter, nest)
}

func ParseNest(text []string, firstNonCommentLine int) (Nest, bool) {
	var nest Nest
	n := findFirstLineWithTunnel(text)
	for i := firstNonCommentLine + 1; i < n; i++ {
		switch {
		case strings.Contains(text[i], "#"):
		case text[i] == "":
		default:
			r, problem := newRoom(text[i])
			if problem {
				return nest, true
			}
			if text[i-1] == "##start" {
				r.Start = true
				nest.Start = r
			}
			if text[i-1] == "##end" {
				r.End = true
				nest.End = r
			}
			nest.Rooms = append(nest.Rooms, r)
		}
	}
	if nest.Start == nil {
		fmt.Println("ERROR: No start room found.")
		return nest, true
	}
	if nest.End == nil {
		fmt.Println("ERROR: No end room found.")
		return nest, true
	}
	for i := 0; i < len(nest.Rooms); i++ {
		for j := n; j < len(text); j++ {
			if !(text[j] == "") && !strings.Contains(text[j], "#") {
				pair := strings.Split(text[j], "-")
				if nest.Rooms[i].Name == pair[0] {
					v := findRoom(pair[1], len(nest.Rooms), &nest)
					nest.Rooms[i].Neighbors = append(nest.Rooms[i].Neighbors, v)
				}
				if nest.Rooms[i].Name == pair[1] {
					u := findRoom(pair[0], len(nest.Rooms), &nest)
					nest.Rooms[i].Neighbors = append(nest.Rooms[i].Neighbors, u)
				}
			}
		}
	}
	for i, ii := range nest.Rooms {
		for j, jj := range nest.Rooms {
			if ii.Name == jj.Name && j != i {
				fmt.Println("ERROR: Duplicated room.")
				return nest, true
			}
		}
		for _, jj := range nest.Rooms[i].Neighbors {
			if ii == nil || jj == nil {
				fmt.Println("ERROR: Link to unknown room.")
				return nest, true
			}
			if jj.Name == ii.Name {
				fmt.Println("ERROR: Room links to itself.")
				return nest, true
			}
		}
	}
	return nest, false
}
