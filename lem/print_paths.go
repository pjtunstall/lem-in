package lem

import (
	"fmt"
)

// For testing. Not used in finished program.

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
