package lem

import (
	"fmt"
	"sort"
)

var Paths []*Path

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
					if u.Start || !u.IsFlowing || v.Flow[u] == 1 {
						if v.Flow[u] == 1 {
							v.IsFlowing = false
						}
						q = append(q, v)
						v.Predecessor = u
					}
				}
			}
		}
		if nest.End.Predecessor != nil {
			nest.End.Residual[nest.End.Predecessor] = 1
			nest.End.Predecessor.Residual[nest.End] = 0
			nest.End.Predecessor.Flow[nest.End] = 1
			nest.End.Predecessor.IsFlowing = true
			nest.End.IsFlowing = true
			for v := nest.End.Predecessor; !v.Start; {
				// fmt.Printf("%v<--", v.Name)
				u := v.Predecessor
				u.Flow[v] = (v.Flow[u] + 1) % 2
				if u.Flow[v] == 1 {
					u.IsFlowing = true
					v.IsFlowing = true
				} else {
					u.IsFlowing = false
				}
				u.Residual[v] = 0
				v.Flow[u] = 0
				v.Residual[u] = 1
				v = u
			}
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

func PathFinder(nest *Nest) []*Path {
	paths := []*Path{}
	for _, n := range nest.Start.Neighbors {
		path := Path{
			Second: n,
			Rooms:  []*Room{nest.Start},
			Ants:   0,
		}
		if n != nest.End {
			count := 0
			for u := n; u != nest.End && count < len(nest.Rooms); {
				path.Rooms = append(path.Rooms, u)
				for _, r := range u.Neighbors {
					if u.Flow[r] == 1 {
						u = r
					} else {
						count++
					}
				}
			}
		}
		path.Rooms = append(path.Rooms, nest.End)
		if nest.Start.Flow[n] == 1 {
			paths = append(paths, &path)
		}
	}
	sort.Slice(paths, func(i, j int) bool { return len(paths[i].Rooms) < len(paths[j].Rooms) })
	return paths
}

func PrintPaths(paths []*Path, nest *Nest) {
	for _, i := range paths {
		for k, j := range i.Rooms {
			fmt.Print(j.Name)
			if k+1 != len(i.Rooms) {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
