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
	fmt.Print("\nPaths:\n")
	nest.Start.Predecessor = nest.Start
	for _, i := range nest.Start.Neighbors {
		if i.End {
			fmt.Printf("%v-%v\n", nest.Start.Name, nest.End.Name)
		} else {
			a, _ := lem.Scout(nest.Start, i, 0)
			if a != nil {
				fmt.Print(nest.Start.Name, "-")
				for j := len(*a) - 1; j > -1; j-- {
					fmt.Print((*a)[j].Name, "-")
				}
				fmt.Println(nest.End.Name)
			}
		}
	}
}
