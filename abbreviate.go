package main

import (
	"fmt"
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

func main() {
	str := "Central Processing Unit"
	fmt.Println(abbreviate(str))
}
