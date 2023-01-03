package lem

import (
	"sort"
)

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
