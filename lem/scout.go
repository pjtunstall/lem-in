package lem

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func LevelFinder(nest *Nest, from *Room) {
	ClearVisited(nest)
	for _, r := range nest.Rooms {
		switch from {
		case nest.Start:
			r.Level = len(nest.Rooms)
		case nest.End:
			r.CoLevel = len(nest.Rooms)
		}
	}
	switch from {
	case nest.Start:
		from.Level = 0
	case nest.End:
		from.CoLevel = 0
	}
	from.Visited = true
	unvisited := len(nest.Rooms) - 1
	for unvisited > 0 {
		for _, r := range nest.Rooms {
			if !r.Visited {
				flag := false
				for _, s := range r.Neighbors {
					if s.Visited {
						flag = true
						switch from {
						case nest.Start:
							r.Level = Min(r.Level, s.Level+1)
						case nest.End:
							r.CoLevel = Min(r.CoLevel, s.CoLevel+1)
						}
					}
				}
				if flag {
					unvisited--
				}
			}
		}
		for _, r := range nest.Rooms {
			switch from {
			case nest.Start:
				if r.Level < len(nest.Rooms) {
					r.Visited = true
				}
			case nest.End:
				if r.CoLevel < len(nest.Rooms) {
					r.Visited = true
				}
			}
		}
	}
}

func ClearVisited(nest *Nest) {
	for _, i := range nest.Rooms {
		if !i.PathNode {
			i.Visited = false
		}
	}
}

func MakePath(nest *Nest, r *Room) Path {
	var p Path
	p.Penultimate = r
	steps := 1
	for i := r; i.Predecessor != nil; i = i.Predecessor {
		i.PathNode = true
		steps++
	}
	p.Steps = steps
	return p
}

func Scout(nest *Nest) []Path {
	paths := []Path{}
	p := Path{}
	for _, n := range nest.Start.Neighbors {
		if n == nest.End {
			p.Steps = 1
		} else {
			ClearVisited(nest)
			nest.Start.Visited = true
			n.Visited = true
			n.Predecessor = nest.Start
			p = PathFinder(nest, n)
		}
		if p.Steps != 0 {
			paths = append(paths, p)
		}
	}
	return paths
}

func PathFinder(nest *Nest, n *Room) Path {
	var change bool
loop:
	for r := n; ; {
		r.Visited = true
		change = false
		for _, s := range r.Neighbors {
			if !s.Visited && !s.PathNode && r.Level <= s.Level {
				change = true
				if s.End {
					return MakePath(nest, r)
				}
				s.Predecessor = r
				r = s
				break
			}
		}
		if !change {
			break loop
		}
	}
	return Path{nil, 0}
}

// func Scout(nest *Nest) []Path {
// 	paths := []Path{}
// 	// Insert something to append a path from start directly to end.
// 	// First, find a way to represent such a path. This is example02.
// 	for _, n := range nest.Start.Neighbors {
// 		ClearVisited(nest)
// 		nest.Start.Visited = true
// 		n.Visited = true
// 		n.Predecessor = nest.Start
// 		p := PathFinder(nest, n)
// 		if p.Steps != 0 {
// 			paths = append(paths, p)
// 		}
// 	}
// 	return paths
// }

// // func Scout(q, r *Room, steps int, nest *Nest) (*[]Room, int) {
// // 	steps++
// // 	r.Predecessor = q
// // 	if r.End {
// // 		path := []Room{}
// // 		for i := r; !i.Start; {
// // 			path = append(path, *i)
// // 			k := i.Predecessor
// // 			i.Predecessor = nil
// // 			i = k
// // 		}
// // 		path = append(path, *nest.Start)
// // 		return &path, steps
// // 	}
// // 	for _, i := range r.Neighbors {
// // 		if i.Predecessor == nil {
// // 			return Scout(r, i, steps, nest)
// // 		}
// // 	}
// // 	return nil, 0
// // }
