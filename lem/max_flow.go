package lem

func MaxFlow(nest *Nest) int {
	for {
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
				if u.Flow[v] == 1 {
					uIsFlowing = true
				}
			}
			for _, v := range u.Neighbors {
				if !v.Start && (u.Flow[v] == 0 || v.Flow[u] == 1) && v.Predecessor == nil {
					if !uIsFlowing || v.Flow[u] == 1 || u.Flow[u.Predecessor] == 1 || u.Start {
						q = append(q, v)
						v.Predecessor = u
					}
				}
			}
		}
		if nest.End.Predecessor != nil {
			nest.End.Predecessor.Flow[nest.End] = 1
			for v := nest.End.Predecessor; !v.Start; {
				// Uncomment this and the Println below to see paths in the residual graph.
				// fmt.Printf("%v <--", v.Name)
				u := v.Predecessor
				u.Flow[v] = 1 ^ v.Flow[u]
				v.Flow[u] = 0
				v = u
			}
		} else {
			break
		}
		// fmt.Println()
	}
	flow := 0
	for _, n := range nest.End.Neighbors {
		flow += n.Flow[nest.End]
	}
	return flow
}
