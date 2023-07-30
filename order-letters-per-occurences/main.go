package main

import (
	"fmt"
	"strings"
)

// This is a small coding challenge where we need to rebuilt a given string based on the number of occurence of
// each character it contains. The new string should follow the first order of occurence. See below example :
//
// Input  :  "LLKMLKRTYSHKFFKVLMFEWPRLGFSPRTVV"
// Output :  "LLLLLKKKKMMRRRTTYSSHFFFFVVVEWPPG"

// The approach used here is based on map. The idea is to loop over the given string - so we have at least O(n) as time
// complexity and keep each character / letter into a map as key (so we avoid duplicate), then increment the value of that
// letter each time we meet this character during the loop. Map structure is not ordered by default so another slice is
// used to maintain character occurence order and will facilitate the assembling of new string. Extra space will make O(n+k).

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 09 August 2021

// core function that rebuild a given string.
func reorder(inputString string) string {
	// store each unique letter with its occurence number.
	mapOccurrences := make(map[string]int)
	// store each letter inside the string in same order.
	inputListWithoutDuplicates := make([]string, 0)
	// temporary variable for conversion.
	var charIntoString string
	// host new final string.
	var result string

	for _, char := range inputString {
		// char is rune type so lets convert into string type.
		charIntoString = string(char)
		// check if this character was already seen.
		if occurrence, exists := mapOccurrences[charIntoString]; exists {
			// increment that character occurences.
			mapOccurrences[charIntoString] = occurrence + 1
		} else {
			// initialize occurence and add to list.
			mapOccurrences[charIntoString] = 1
			inputListWithoutDuplicates = append(inputListWithoutDuplicates, charIntoString)
		}
	}

	// loop over new list without duplicates to build result string
	// for each character - repeat it with its occurrences number.
	for _, character := range inputListWithoutDuplicates {
		result = result + strings.Repeat(character, mapOccurrences[character])
	}

	return result
}

func main() {

	str := "LLKMLKRTYSHKFFKVLMFEWPRLGFSPRTVV"
	fmt.Println(reorder(str))
}

// go run ordering-letters-per-occurences.go
// Output:
// LLLLLKKKKMMRRRTTYSSHFFFFVVVEWPPG
