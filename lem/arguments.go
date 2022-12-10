package lem

import (
	"fmt"
)

func RightNumberOfArguments(n int) bool {
	switch {
	case n == 2:
		return true
	case n < 2:
		fmt.Println("ERROR: Not enough arguments.")
		return false
	default:
		fmt.Println("ERROR: Too many arguments.")
		return false
	}
}
