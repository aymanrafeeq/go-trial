package main

import (
	"reflect"
	"testing"
)

func TestPalindrome(t *testing.T) {

	input := "level"
	output := palindrome(input)

	if !output {
		t.Errorf("NExpected to be a Palindrome")
	}

	input = "hello"
	output = palindrome(input)

	if output {
		t.Errorf("Expected not to be a Palindrome")
	}
}

func TestAbbreviate(t *testing.T) {

	output := abbreviate("Central Processing Unit")

	if output != "CPU" {
		t.Errorf("Expected CPU but got %s\n", output)
	}

	output = abbreviate("Central processing unit")

	if output == "CPU" {
		t.Errorf("Expected %s but got CPU, Please check your program\n", output)
	}
}
func TestPangram(t *testing.T) {
	intput := "The quick brown fox jumps over the lazy dog"
	output := pangram(intput)

	if !output {
		t.Errorf("Expected to be Pangram")
	}

	intput = "The quick brown fox jumps over the lazy do"
	output = pangram(intput)

	if output {
		t.Errorf("Expected not to be Pangram, some error in program")
	}
}
func TestFrequency(t *testing.T) {
	input := "bobby"
	expected := map[string]int{"b": 3, "o": 1, "y": 1}
	output := frequency(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v as frequency but got %v", expected, output)
	}
}
func TestFizzBizz(t *testing.T) {
	input := 15
	expected := []string{"1", "2", "fizz", "4", "bizz", "fizz", "7", "8", "fizz", "bizz", "11", "fizz", "13", "14", "fizzbizz"}
	output := fizzbizz(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v  but got %v", expected, output)
	}
}
