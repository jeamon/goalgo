package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Reference struct {
	id  int // document id its belongs to
	pos int // position into the document
}

var store map[string][]Reference

func indexDocuments(documents [][]string) {
	store = make(map[string][]Reference)
	for id, doc := range documents {
		for pos, word := range doc {
			word = strings.ToLower(word)
			if _, exists := store[word]; exists {
				store[word] = append(store[word], Reference{id: id, pos: pos})
			} else {
				store[word] = []Reference{{id: id, pos: pos}}
			}
		}
	}
}

// searchOne returns list of documents ids containing word.
func searchOne(word string) []int {
	results := make([]int, 0)
	refs, found := store[strings.ToLower(word)]
	if !found {
		return []int{}
	}
	for _, ref := range refs {
		results = append(results, ref.id)
	}
	// remove duplicate document ids
	m := make(map[int]struct{}, len(results))
	for _, v := range results {
		m[v] = struct{}{}
	}
	return slices.Sorted(maps.Keys(m))
}

func searchOneRefs(word string) []Reference {
	refs, found := store[word]
	if !found {
		return []Reference{}
	}
	return refs
}

// search returns list of documents ids containing in order all words.
func search(words ...string) []int {
	switch len(words) {
	case 1:
		return searchOne(words[0])
	case 0:
		return []int{}
	}

	m := make(map[int][]int) // map of document ids to positions
	for _, word := range words {
		word = strings.ToLower(word)
		refs := searchOneRefs(word)
		for _, ref := range refs {
			if positions, found := m[ref.id]; found {
				m[ref.id] = append(positions, ref.pos)
			} else {
				m[ref.id] = []int{ref.pos}
			}
		}
	}

	results := make([]int, 0)
	for id, positions := range m {
		if len(positions) >= len(words) && slices.IsSorted(positions) {
			results = append(results, id)
		}
	}

	slices.Sort(results)
	return results
}

func main() {
	documents := [][]string{
		{"my", "name", "is", "jerome"},
		{"hello", "how", "are", "you"},
		{"I", "worked", "at", "Cisco"},
		{"I", "worked", "at", "BNP Paribas"},
		{"also", "I", "worked", "at", "SauceLabs"},
		{"and", "I", "worked", "at", "AppDynamics"},
		{"and", "I", "worked", "at", "Remote"},
		{"and", "I", "worked", "at", "Neurones Technologies"},
	}
	indexDocuments(documents)
	for word, refs := range store {
		fmt.Printf("%s: %v\n", word, refs)
	}
	/*
		name: [{0 1}]
		jerome: [{0 3}]
		hello: [{1 0}]
		at: [{2 2} {3 2} {4 3} {5 3} {6 3} {7 3}]
		appdynamics: [{5 4}]
		remote: [{6 4}]
		how: [{1 1}]
		i: [{2 0} {3 0} {4 1} {5 1} {6 1} {7 1}]
		worked: [{2 1} {3 1} {4 2} {5 2} {6 2} {7 2}]
		and: [{5 0} {6 0} {7 0}]
		neurones technologies: [{7 4}]
		my: [{0 0}]
		is: [{0 2}]
		you: [{1 3}]
		cisco: [{2 3}]
		bnp paribas: [{3 3}]
		also: [{4 0}]
		saucelabs: [{4 4}]
		are: [{1 2}]
	*/

	results := searchOne("name")
	fmt.Println(results) // => [0]

	results = search("I", "worked", "at")
	fmt.Println(results) // => [2 3 4 5 6 7]
}
