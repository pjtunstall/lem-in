package lem

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PrintTurns(paths []*Path, nest *Nest, ants int) {
	paths[0].FirstAnt = 1
	if len(paths) > 1 {
		for i := 1; i < len(paths); i++ {
			paths[i].FirstAnt = paths[i-1].FirstAnt + paths[i-1].Ants
		}
	}
	for row := 1; ants > 0; row++ {
		for _, p := range paths {
			for i := p.FirstAnt; i <= p.Ants+p.FirstAnt-1; i++ {
				j := row - i + p.FirstAnt
				if j > 0 && j < len(p.Rooms) {
					fmt.Printf("L%v-%v ", i, p.Rooms[j].Name)
				}
				if j == len(p.Rooms)-1 {
					ants--
				}
			}
		}
		fmt.Println()
	}
}
