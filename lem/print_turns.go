package lem

import "fmt"

func PrintTurns(paths []*Path, nest *Nest, ants int) {
	for row := 0; row < ants; row++ {
		for _, p := range paths {
			for i := 0; i < p.Ants; i++ {
				if row < i && row-i < len(p.Rooms) {
					fmt.Printf("L%v-%v", i+1, p.Rooms[row-i].Name)
				}
			}
		}
	}
}
