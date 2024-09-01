package main

import "fmt"

// this is a nice coding challenge where you need to implement a function that will return all pair of indexes where
// the sum of their values equals a target number. The input is an unsorted slice of integers. We need to use a slice
// for temporarily storing indexes with same value - so another for loop for is needed. TC is like O(nk) & SC is O(n+k).

// the idea is to loop over the slice and compute the difference between the target sum and the current
// value then check if the value is already into a map. Because if (a + b = x) then (x - b = a).
// The map uses the slice's (input slice) value as key and its indexes as value. In order to not miss any
// duplicated value's index, the map could use a slice of integers (indexes) as value for each key.

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 09 August 2021

// Pair defines the format of a found result.
type Pair struct {
	firstIndex  int
	firstValue  int
	secondIndex int
	secondValue int
}

func findAllPairs(input []int, sum int, results *[]Pair) {
	var diff int
	// to store each number and all corresponding indexes.
	tempMap := make(map[int][]int)
	for index, number := range input {
		// compute the difference and check if already found.
		diff = sum - number
		if indexList, exist := tempMap[diff]; exist {
			// matching pair found so build the pair struct for each index
			// mapped to that value (diff) and add to the results store.
			for _, i := range indexList {
				*results = append(*results, Pair{firstIndex: i, firstValue: diff, secondIndex: index, secondValue: number})
			}
		}
		// add into temporary map for tracking.
		tempMap[number] = append(tempMap[number], index)
	}
}

func main() {
	// slice of pair to store found pairs.
	var results []Pair
	var input = []int{12, 100, 10, -7, 45, -1, 10, 90, 412, 55, 0, -45, 6, 107, 101, 100}
	target := 100
	// results slice passed by reference.
	findAllPairs(input, target, &results)
	// routine to display all the results.
	for _, p := range results {
		fmt.Printf("found pair of indexes : (%d, %d) => %d + %d = %d\n", p.firstIndex, p.secondIndex, p.firstValue, p.secondValue, (p.firstValue + p.secondValue))
	}
}

// go run main.go
// Output:
// found pair of indexes : (2, 7) => 10 + 90 = 100
// found pair of indexes : (6, 7) => 10 + 90 = 100
// found pair of indexes : (4, 9) => 45 + 55 = 100
// found pair of indexes : (1, 10) => 100 + 0 = 100
// found pair of indexes : (3, 13) => -7 + 107 = 100
// found pair of indexes : (5, 14) => -1 + 101 = 100
// found pair of indexes : (10, 15) => 0 + 100 = 100
