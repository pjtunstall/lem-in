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
	if lem.MaxFlow(&nest) == 0 {
		fmt.Println("ERROR: No paths found.")
		return
	}
	paths := lem.PathCollector(&nest)
	lem.SendAnts(paths, &nest, ants)
	fmt.Printf("%v\n\n", textFile)
	lem.PrintTurns(paths, &nest, ants)
}
