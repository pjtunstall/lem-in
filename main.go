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
	firstNonCommentLine := 0
	for i := 0; text[i] == "" || text[i][0] == '#'; i++ {
		firstNonCommentLine++
	}
	ants, err := strconv.Atoi(text[firstNonCommentLine])
	if err != nil || ants < 1 {
		fmt.Println("ERROR: Invalid number of ants.")
		return
	}
	nest, problem := lem.ParseNest(text, firstNonCommentLine)
	if problem {
		return
	}
	flow, paths := lem.MaxFlow(&nest, ants)
	if flow == 0 {
		fmt.Println("ERROR: No paths found.")
		return
	}
	fmt.Printf("%v\n\n", textFile)
	lem.PrintTurns(paths, &nest, ants)
}
