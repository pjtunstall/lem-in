package lem

import (
	"strings"
)

func FirstTunnel(text []string) int {
	for i := 5; i < len(text); i++ {
		if strings.Contains(text[i], "-") {
			return i
		}
	}
	return 0
}
