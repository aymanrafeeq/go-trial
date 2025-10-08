package main

import (
	"fmt"
	"strings"
)

func pangram(s string) bool {
	s = strings.ToLower(s)
	for i := 'a'; i <= 'z'; i++ {
		if !strings.Contains(s, string(i)) {
			return false
		}
	}
	return true

}
func main() {
	str := "the quick brown fox jumps over the lazy dog"
	fmt.Println(pangram(str))
}
