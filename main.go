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

	fmt.Println()
	maxFlow := lem.MaxFlow(&nest)
	fmt.Printf("\nMax Flow: %v\n", maxFlow)
	if maxFlow == 0 {
		return
	}

	lem.PrintFormattedNest(&nest, ants)

	fmt.Print("\nPaths:\n")
	paths := [][]string{}
	for _, n := range nest.Start.Neighbors {
		path := []string{nest.Start.Name}
		if n != nest.End {
			count := 0
			for u := n; u != nest.End && count < len(nest.Rooms); {
				path = append(path, u.Name)
				for _, r := range u.Neighbors {
					if u.Flow[r] == 1 {
						u = r
					} else {
						count++
					}
				}
			}
		}
		path = append(path, nest.End.Name)
		if nest.Start.Flow[n] == 1 {
			paths = append(paths, path)
		}
	}

	for _, i := range paths {
		for k, j := range i {
			fmt.Print(j)
			if k+1 != len(i) {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}
