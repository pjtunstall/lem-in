package lem

func CountTurns(paths []*Path, nest *Nest, ants int) int {
	return len(paths[0].Rooms) + paths[0].Ants - 2
}

func MaxFlow(nest *Nest, ants int) (int, []*Path) {
	numberOfTurns := len(nest.Rooms) + ants - 2
	var paths []*Path
	for i := 1; i <= ants; i++ {
		q := []*Room{nest.Start}
		for _, r := range nest.Rooms {
			r.Predecessor = nil
		}
		for len(q) != 0 {
			// Uncomment to see what's in the queue at each iteration.
			// for _, r := range q {
			// 	fmt.Print(r.Name)
			// }
			// fmt.Println()
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
		if nest.End.Predecessor != nil {
			nest.End.Predecessor.Flow[nest.End] = true
			for v := nest.End.Predecessor; !v.Start; {
				// Uncomment this and the Println below to see paths in the residual graph as they're found.
				// fmt.Printf("%v <--", v.Name)
				u := v.Predecessor
				u.Flow[v] = !v.Flow[u]
				v.Flow[u] = false
				v = u
			}
			newPaths := PathCollector(nest)
			SendAnts(newPaths, nest, ants)
			newNumberOfTurns := CountTurns(newPaths, nest, ants)
			if newNumberOfTurns > numberOfTurns {
				break
			}
			numberOfTurns = newNumberOfTurns
			paths = newPaths
		} else {
			break
		}
		// fmt.Println()
	}
	flow := 0
	for _, n := range nest.End.Neighbors {
		if n.Flow[nest.End] {
			flow++
		}
	}
	return flow, paths
}
