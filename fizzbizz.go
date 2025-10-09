package main

import (
	"strconv"
)

// write a function that will accept a positive number n as input. It will
//       print numbers m from 1 to n. If m is divisible by 3, it will
//       print fizz. If it's divisible by 5, it will print bizz. If it's
//       divisible by 15, it will print fizzbizz. Otherwise, it will just
//       print the number.

func fizzbizz(n int) []string {

	var final []string

	for m := 1; m <= n; m++ {

		if m%15 == 0 {
			final = append(final, "fizzbizz")
		} else if m%5 == 0 {
			final = append(final, "bizz")
		} else if m%3 == 0 {
			final = append(final, "fizz")
		} else {
			final = append(final, strconv.Itoa(m))
		}
	}
	return final
}
