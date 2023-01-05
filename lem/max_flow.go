package lem

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
				if !v.Start && (u.Flow[v] == 0 || v.Flow[u] == 1) && v.Predecessor == nil {
					q = append(q, v)
					v.Predecessor = u
				}
			}
		}
		if nest.End.Predecessor != nil {
			nest.End.Predecessor.Flow[nest.End] = 1
			for v := nest.End.Predecessor; !v.Start; {
				u := v.Predecessor
				u.Flow[v] = 1 ^ v.Flow[u]
				v.Flow[u] = 0
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
