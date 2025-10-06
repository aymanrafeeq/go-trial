package main

import "fmt"

func palindrome(str string) {
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			fmt.Println("It is not a palindrome")
			return
		}
	}
	fmt.Println("it is a palindrome")
}

func main() {
	str := "hello"
	palindrome(str)
}
