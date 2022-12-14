package lem

var steps int

func PathFinder(nest *Nest) (*[]Room, int) {
	steps = 0
	// for _, i := range nest.Start.Neighbors {
	// 	return Scout(nest.Start, i, nest)
	// }
	return Scout(nest.Start, nest.Start.Neighbors[0], nest, steps)
}

func Scout(q, r *Room, nest *Nest, steps int) (*[]Room, int) {
	steps++
	r.Predecessor = q
	for _, i := range r.Neighbors {
		if i == nest.End {
			return Return(r, steps)
		}
	}
	for _, i := range r.Neighbors {
		if i.Predecessor == nil {
			return Scout(r, i, nest, steps)
		}
	}
	return nil, 0
}

func Return(r *Room, steps int) (*[]Room, int) {
	path := []Room{}
	i := r
	for i.Predecessor != nil {
		path = append(path, *i.Predecessor)
		i = i.Predecessor
	}
	return &path, steps + 1
}
