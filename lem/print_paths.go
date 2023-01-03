package lem

import (
	"fmt"
)

func PrintPaths(paths []*Path, nest *Nest) {
	for _, i := range paths {
		for k, j := range i.Rooms {
			fmt.Print(j.Name)
			if k+1 != len(i.Rooms) {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
