package lem

func SendAnts(paths []*Path, nest *Nest, ants int) {
	if len(paths) == 1 {
		paths[0].Ants = ants
		return
	}
	paths[0].Ants = 1
	ants--
	i := 0
	for ants > 0 {
		if len(paths[i].Rooms)+paths[i].Ants > len(paths[i+1].Rooms)+paths[i+1].Ants {
			paths[i+1].Ants++
			ants--
			i++
			if i+1 == len(paths) {
				i = 0
			}
		} else {
			paths[i].Ants++
			ants--
			i = 0
		}
	}
}
