package main

import (
	"sort"
)

// given a list of words and a search word then display at most k words with common prefix from the list in sorted order
// each time a character from word is typed. so a list of list is expected as output. list of size n and word of m length.

func SuggestedProducts(wordsList []string, searchWord string, maxSuggestions int) [][]string {
	var allSuggestions [][]string
	// 1. lets sort the productsList. O(nlogn)
	sort.Strings(wordsList)

	// 2. lets use the 2 pointers technique: indexes left & right track words from the list
	// we iterate over each character of the search word and ensure pointed word has this
	// character or not. If not then move the pointer to the next until both pointers reach.
	leftIndex, rightIndex := 0, len(wordsList)-1
	for i := 0; i < len(searchWord); i++ {
		char := searchWord[i] // type of rune
		// left index should not go above the right and once reached it means no more words to check
		// move the left index up if the word into the list does not have character at position i or
		// has a non-matching character at that position i
		for leftIndex <= rightIndex && (len(wordsList[leftIndex]) <= i || wordsList[leftIndex][i] != char) {
			leftIndex++
		}

		// right index should not go behind the left and once reached it means no more words to check
		// move the right index down if the word into the list does not have character at position i
		// or has a non-matching character at that position i
		for leftIndex <= rightIndex && (len(wordsList[rightIndex]) <= i || wordsList[rightIndex][i] != char) {
			rightIndex--
		}

		// all words from the list which are between left and right indexes are matching (common prefix with search word)
		// as we need at most maxSuggestions lets compute the exact available words and append them to our final results.
		availableSuggestions := maxSuggestions
		if n := (rightIndex - leftIndex) + 1; n < availableSuggestions {
			availableSuggestions = n
		}
		suggestions := make([]string, availableSuggestions)
		for k := 0; k < availableSuggestions; k++ {
			suggestions[k] = wordsList[leftIndex+k]
		}

		allSuggestions = append(allSuggestions, suggestions)
	}

	return allSuggestions
}
