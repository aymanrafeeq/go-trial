package main

func frequency(s string) map[string]int {

	freq := make(map[string]int)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		freq[string(ch)] = freq[string(ch)] + 1
	}
	return freq
}
