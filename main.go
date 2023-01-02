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

	// The functions MaxFlow, PathFinder, and PrintPaths need to be called in this order.

	maxFlow := lem.MaxFlow(&nest)
	if maxFlow == 0 {
		fmt.Println("No paths found.")
		return
	}
	paths := lem.PathFinder(&nest)
	fmt.Print("\n\n")
	lem.SendAnts(paths, &nest, ants)
	lem.PrintTurns(paths, &nest, ants)
}
