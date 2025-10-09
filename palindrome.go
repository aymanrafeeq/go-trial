package main

func palindrome(s string) bool {

	var rev string

	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}

	return rev == s
}
