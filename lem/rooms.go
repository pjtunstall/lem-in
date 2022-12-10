package lem

import (
	"fmt"
	"strconv"
	"strings"
)

func newRoom(row string) (Room, bool) {
	a := strings.Split(row, " ")
	var problem bool
	x, er1 := strconv.Atoi(a[1])
	y, er2 := strconv.Atoi(a[2])
	if er1 != nil || er2 != nil {
		fmt.Println("ERROR: Invalid coordinates.")
		problem = true
	}
	return Room{
		Name: a[0],
		X:    x,
		Y:    y,
	}, problem
}

func Rooms(text []string) ([]Room, bool) {
	nest := []Room{}
	n := FirstTunnel(text)
	startExists := false
	endExists := false
loop:
	for i := 1; i < n; i++ {
		switch {
		case strings.Contains(text[i], "-"):
			break loop
		case strings.Contains(text[i], "#"):
		default:
			r, problem := newRoom(text[i])
			if problem {
				return nest, true
			}
			if text[i-1] == "##start" {
				r.Start = true
				startExists = true
			}
			if text[i-1] == "##end" {
				r.End = true
				endExists = true
			}
			nest = append(nest, r)
		}
	}
	if !startExists {
		fmt.Println("ERROR: No start room found.")
		return nest, true
	}
	if !endExists {
		fmt.Println("ERROR: No end room found.")
		return nest, true
	}
	for i := 0; i < len(nest); i++ {
		for j := n; j < len(text); j++ {
			pair := strings.Split(text[j], "-")
			if nest[i].Name == pair[0] {
				nest[i].Neighbors = append(nest[i].Neighbors, pair[1])
			}
			if nest[i].Name == pair[1] {
				nest[i].Neighbors = append(nest[i].Neighbors, pair[0])
			}
		}
	}
	for i, ii := range nest {
		for j, jj := range nest {
			if ii.Name == jj.Name && j != i {
				fmt.Println("ERROR: Duplicated room.")
				return nest, true
			}
		}
		for j, jj := range nest[i].Neighbors {
			if jj == ii.Name {
				fmt.Println("ERROR: Room links to itself.")
				return nest, true
			}
			match := false
			for _, kk := range nest {
				if jj == kk.Name {
					match = true
				}
			}
			if !match {
				fmt.Println("ERROR: Link to unknown room.")
				return nest, true
			}
			if j < len(nest[i].Neighbors)-1 {
				for _, kk := range nest[i].Neighbors[j+1:] {
					if jj == kk {
						fmt.Println("ERROR: Two links between same two rooms.")
						return nest, true
					}
				}
			}
		}
	}
	return nest, false
}
