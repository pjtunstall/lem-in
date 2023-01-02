package lem

import (
	"fmt"
)

// For testing. Not used in finished program.

func PrintFormattedNest(nest *Nest, ants int) {
	fmt.Printf("\nNumber of ants: %v\n", ants)
	for i := 0; i < ants; i++ {
		fmt.Print(string([]rune{0x1F41C, 32}))
	}
	fmt.Print("\n\n")
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
		fmt.Print("Capacity to:\t")
		for i := range j.Neighbors {
			if j.Residual[j.Neighbors[i]] == 1 {
				fmt.Print(j.Neighbors[i].Name, ", ")
			}
		}
		fmt.Println()
		fmt.Println("IS FLOWING?", j.IsFlowing)
		fmt.Print("SENDING FLOW TO:\t")
		for i := range j.Neighbors {
			if j.Flow[j.Neighbors[i]] == 1 {
				fmt.Print(j.Neighbors[i].Name, ", ")
			}
		}
		fmt.Println()
		fmt.Print("RECEIVING FLOW FROM:\t")
		for _, i := range j.Neighbors {
			if i.Flow[j] == 1 {
				fmt.Print(i.Name, ", ")
			}
		}
		fmt.Println()
		fmt.Printf("Coordinates:\t(%v, %v)\n", j.X, j.Y)
		if i < len(nest.Rooms)-1 {
			fmt.Println()
		}
	}
}
