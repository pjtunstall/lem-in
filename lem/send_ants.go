package lem

func SendAnts(paths []*Path, nest *Nest, ants int) {
	if len(paths) == 1 {
		paths[0].Ants = ants
		return
	}
	for a := ants; a > 0; {
		for i := 0; i < len(paths); i++ {
			if i == len(paths)-1 {
				paths[i].Ants++
				a--
				if a == 0 {
					return
				}
				break
			}
			counter := 0
			for j := i + 1; j < len(paths); j++ {
				if paths[i].Ants+len(paths[i].Rooms) <= paths[j].Ants+len(paths[j].Rooms) {
					counter++
				}
			}
			if counter == len(paths)-1-i {
				paths[i].Ants++
				a--
				if a == 0 {
					return
				}
				break
			}
		}
	}
}
