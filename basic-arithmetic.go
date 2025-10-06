// Write a function that accepts two numbers a and b as input and
//
//	returns a+b, a-b and a*b
package main

import "fmt"

func arithmetic(a int, b int) (int, int, int) {
	sum := a + b
	diff := a - b
	prod := a * b
	return sum, diff, prod
}
func main() {
	a := 10
	b := 7
	sum, diff, prod := arithmetic(a, b)
	fmt.Println(sum)
	fmt.Println(diff)
	fmt.Println(prod)
}
