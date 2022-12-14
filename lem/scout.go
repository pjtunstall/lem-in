package lem

func Scout(q, r *Room, steps int) (*[]Room, int) {
	steps++
	r.Predecessor = q
	for _, i := range r.Neighbors {
		if i.End {
			return Return(r, steps)
		}
	}
	for _, i := range r.Neighbors {
		if i.Predecessor == nil {
			return Scout(r, i, steps)
		}
	}
	return nil, 0
}

func Return(r *Room, steps int) (*[]Room, int) {
	path := []Room{}
	for i := r; !i.Start; i = i.Predecessor {
		path = append(path, *i)
	}
	return &path, steps + 1
}
