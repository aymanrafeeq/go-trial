package main

import (
	"fmt"
	"strings"
)

func frequency(s string) map[string]int {

	freq := make(map[string]int)

	s = strings.ToLower(s)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			freq[string(ch)] = freq[string(ch)] + 1
		}
	}
	return freq
}
func main() {
	fmt.Println(frequency("ayman"))
}
