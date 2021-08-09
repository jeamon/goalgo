package main

import "fmt"

// classic challenge to check if a reversed word
// or number or sentence equals or not to itself.

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 09 August 2021

func checkPalindrome(input string) bool {
	var reverseInput string
	// reversed a word with extra temporary variable.
	// character is rune type so conver to string.
	for _, character := range input {
		reverseInput = string(character) + reverseInput
	}
	// compare the two words and return the result.
	if reverseInput == input {
		return true
	} else {
		return false
	}
}

func main() {

	// an input value. you can check with radar or stats or wow or mom.
	word := "10001"

	if checkPalindrome(word) == true {
		fmt.Println("This word ", word, " is a palindrome")
	} else {
		fmt.Println("This word ", word, " is not a palindrome")
	}
}

// Output:
// This word  10001  is a palindrome
