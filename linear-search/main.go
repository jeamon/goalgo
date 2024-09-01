package main

import "fmt"

// search into a list of integers for a given value
// time compexity O(n) and space complexity O(1).

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 09 August 2021

// SequentialSearch will blindly loop over the list and compare
// each element with the target until found or reach the end.
func SequentialSearch(data []int, value int) bool {

	size := len(data)
	for i := 0; i < size; i++ {
		if data[i] == value {
			return true
		}
	}

	return false
}

func main() {

	input := []int{12, 10, 2, 78, 45, 11}
	value := 78

	result := SequentialSearch(input, value)

	fmt.Println("\n List of integers :", input, "\n Search result :", result)
}

// go run main.go
// Output:
// List of integers : [12 10 2 78 45 11]
// Search result : true
