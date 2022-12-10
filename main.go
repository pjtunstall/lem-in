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
	fmt.Printf("%v\n\nNumber of ants: %v\n\n", textFile, ants)
	for i, j := range nest {
		fmt.Printf("Name of room:\t%v", j.Name)
		if j.Start {
			fmt.Printf(" (start)")
		}
		if j.End {
			fmt.Printf(" (end)")
		}
		fmt.Println()
		fmt.Printf("Connected to:\t%v\n", j.Neighbors)
		fmt.Printf("Coordinates:\t(%v, %v)\n", j.X, j.Y)
		if i < len(nest)-1 {
			fmt.Println()
		}
	}
}
