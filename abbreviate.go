package main

import (
	"strings"
)

func abbreviate(s string) string {
	words := strings.Fields(s)
	abbr := ""

	for _, word := range words {
		firstChar := word[0]
		if firstChar >= 'A' && firstChar <= 'Z' {
			abbr += string(firstChar)
		}
	}
	return abbr
}
