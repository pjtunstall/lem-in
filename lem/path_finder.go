package lem

import "fmt"

var Paths []*Path

// func MakePath(nest *Nest, r *Room) *Path {
// 	var p Path
// 	p.Penultimate = r
// 	steps := 1
// 	for i := r; i.Predecessor != nil; i = i.Predecessor {
// 		steps++
// 	}
// 	p.Steps = steps
// 	return &p
// }

func MaxFlow(nest *Nest) int {
	for {
		q := []*Room{nest.Start}
		for _, r := range nest.Rooms {
			r.Predecessor = nil
		}
		for len(q) != 0 {
			u := q[0]
			if len(q) > 1 {
				q = q[1:]
			} else {
				q = q[:0]
			}
			for _, v := range u.Neighbors {
				if !v.Start && u.Residual[v] == 1 && v.Predecessor == nil {
					q = append(q, v)
					v.Predecessor = u
				}
			}
		}
		if nest.End.Predecessor != nil {
			nest.End.Residual[nest.End.Predecessor] = 1
			nest.End.Predecessor.Residual[nest.End] = 0
			nest.End.Predecessor.Flow[nest.End] = 1
			for v := nest.End.Predecessor; !v.Start; {
				fmt.Printf("%v<--", v.Name)
				u := v.Predecessor
				u.Flow[v] = (v.Flow[u] + 1) % 2
				u.Residual[v] = 0
				v.Flow[u] = 0
				v.Residual[u] = 1
				v = u
			}
			fmt.Println()
		} else {
			break
		}
	}
	flow := 0
	for _, n := range nest.End.Neighbors {
		flow += n.Flow[nest.End]
	}
	return flow
}
