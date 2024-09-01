package main

import "fmt"

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 09 August 2021

// SumArray is a function that return the sum of all elements of a given list of intergers.
func SumArray(data []int) int {

	total := 0
	size := len(data)
	for index := 0; index < size; index++ {
		total = total + data[index]
	}
	return total
}

// implementation use case. O(n) time complexity and no extra space.
func main() {
	input := []int{14, 45, -4, 10, 23}
	result := SumArray(input)

	fmt.Println("\n List of integers :", input, "\n Sum of elements :", result)
}

// go run main.go
// Output:
// List of integers : [14 45 -4 10 23]
// Sum of elements : 88
