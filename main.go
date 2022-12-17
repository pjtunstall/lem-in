package main

import (
	"fmt"
	"lem-in/lem"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if !lem.RightNumberOfArguments(len(args)) {
		return
	}
	textBytes, _ := os.ReadFile(args[1])
	textFile := string(textBytes)
	text := strings.Split(string(textFile), "\n")
	if len(text) == 1 && text[0] == "" {
		fmt.Println("ERROR: Empty file.")
		return
	}
	ants, err := strconv.Atoi(text[0])
	if err != nil || ants < 1 {
		fmt.Println("ERROR: Invalid number of ants.")
		return
	}
	nest, problem := lem.Rooms(text)
	if problem {
		return
	}
	fmt.Printf("%v\n", textFile)
	lem.PrintFormattedNest(&nest, ants)

	lem.LevelFinder(&nest)
	fmt.Println()

	for i := 0; i < len(nest.Rooms)+1; i++ {
		for _, r := range nest.Rooms {
			if r.Level == i {
				fmt.Printf("%v, distance from start: %v\n", r.Name, r.Level)
			}
		}
	}

	fmt.Print("\nPaths:\n")

	paths := lem.Scout(&nest)

	for _, path := range paths {
		if path.Steps != 0 {
			pathString := ""
			for room := path.Penultimate; room.Predecessor != nil; room = room.Predecessor {
				pathString = room.Name + "-" + pathString
			}
			fmt.Printf("%v-%v%v, steps: %v\n", nest.Start.Name, pathString, nest.End.Name, path.Steps)
		}
	}

	// nest.Start.Predecessor = nest.Start
	// for _, i := range nest.Start.Neighbors {
	// 	a, steps := lem.Scout(nest.Start, i, 0, &nest)
	// 	if a != nil {
	// 		for j := len(*a) - 1; j > -1; j-- {
	// 			fmt.Print((*a)[j].Name)
	// 			if j > 0 {
	// 				fmt.Print("-")
	// 			} else {
	// 				fmt.Printf(" (steps: %v)\n", steps)
	// 			}
	// 		}
	// 	}
	// }
}
