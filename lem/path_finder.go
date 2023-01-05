package lem

import (
	"sort"
)

func PathFinder(nest *Nest) []*Path {
	paths := []*Path{}
	for _, n := range nest.Start.Neighbors {
		if nest.Start.Flow[n] == 1 {
			path := Path{
				Second: n,
				Rooms:  []*Room{nest.Start},
				Ants:   0,
			}
			if n != nest.End {
				for u := n; u != nest.End; {
					path.Rooms = append(path.Rooms, u)
					for _, r := range u.Neighbors {
						if u.Flow[r] == 1 {
							u = r
						}
					}
				}
			}
			path.Rooms = append(path.Rooms, nest.End)
			if nest.Start.Flow[n] == 1 {
				paths = append(paths, &path)
			}
		}
	}
	sort.Slice(paths, func(i, j int) bool { return len(paths[i].Rooms) < len(paths[j].Rooms) })
	return paths
}
