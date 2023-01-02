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
	fmt.Printf("%v", textFile)

	// lem.PrintFormattedNest(&nest, ants)

	// The functions MaxFlow, PathFinder, and PrintPaths need to be called in this order.

	maxFlow := lem.MaxFlow(&nest)
	// fmt.Printf("\n\nMax Flow: %v\n", maxFlow)
	if maxFlow == 0 {
		fmt.Println("No paths found.")
		return
	}

	// lem.PrintFormattedNest(&nest, ants)

	fmt.Print("\n\nPaths:\n")
	paths := lem.PathFinder(&nest)
	lem.PrintPaths(paths, &nest)

	fmt.Println()
	lem.SendAnts(paths, &nest, ants)
	for _, p := range paths {
		fmt.Println(p.Rooms[1].Name, len(p.Rooms), p.Ants)
	}
	// lem.PrintTurns(paths, &nest, ants)
}
