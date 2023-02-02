package lem

func CountTurns(paths []*Path, nest *Nest, ants int) int {
	return len(paths[0].Rooms) + paths[0].Ants - 2
}

func FindPaths(nest *Nest, ants int) (int, []*Path) {
	numberOfTurns := len(nest.Rooms) + ants - 2
	var paths []*Path
	for len(paths) < ants {
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
			uIsFlowing := false
			for _, v := range u.Neighbors {
				if u.Flow[v] {
					uIsFlowing = true
				}
			}
			for _, v := range u.Neighbors {
				neighboringRoomIsUnvisited := v.Predecessor == nil
				tunnelCapacityPermits := v.Flow[u] || !u.Flow[v]
				roomCapacityPermits := !uIsFlowing || v.Flow[u] || u.Start || u.Flow[u.Predecessor]
				if !v.Start && neighboringRoomIsUnvisited && tunnelCapacityPermits && roomCapacityPermits {
					q = append(q, v)
					v.Predecessor = u
				}
			}
		}
		if nest.End.Predecessor == nil {
			break
		}
		nest.End.Predecessor.Flow[nest.End] = true
		for v := nest.End.Predecessor; !v.Start; {
			u := v.Predecessor
			u.Flow[v] = !v.Flow[u]
			v.Flow[u] = false
			v = u
		}
		newPaths := GatherPaths(nest)
		SendAnts(newPaths, nest, ants)
		newNumberOfTurns := CountTurns(newPaths, nest, ants)
		if newNumberOfTurns > numberOfTurns {
			break
		}
		numberOfTurns = newNumberOfTurns
		paths = newPaths
	}
	flow := 0
	for _, n := range nest.End.Neighbors {
		if n.Flow[nest.End] {
			flow++
		}
	}
	return flow, paths
}
