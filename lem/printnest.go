package lem

import (
	"fmt"
)

func PrintFormattedNest(nest *Nest, ants int) {
	fmt.Printf("\nNumber of ants: %v\n\n", ants)
	for i, j := range nest.Rooms {
		fmt.Printf("Name of room:\t%v", j.Name)
		if j.Start {
			fmt.Printf(" (start)")
		}
		if j.End {
			fmt.Printf(" (end)")
		}
		fmt.Println()
		fmt.Print("Connected to:\t")
		for i := range j.Neighbors {
			fmt.Print(j.Neighbors[i].Name)
			if i+1 != len(j.Neighbors) {
				fmt.Print(", ")
			}
		}
		fmt.Println()
		fmt.Printf("Coordinates:\t(%v, %v)\n", j.X, j.Y)
		if i < len(nest.Rooms)-1 {
			fmt.Println()
		}
	}
}
