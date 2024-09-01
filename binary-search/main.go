package main

import "fmt"

// given a sorted array of integers, search for a given value
// time complexity O(log n) and space complexity O(1).

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 09 August 2021

// BinarySearch implements the approach of repeatedly dividing
// the list in half and search inside.
func BinarySearch(data []int, value int) (bool, int) {

	size := len(data)
	low := 0
	high := size - 1

	for low <= high {
		//=> same like (high + low) / 2
		mid := low + (high-low)/2

		if data[mid] == value {
			return true, mid
		} else {
			if data[mid] < value {
				// move the interval to right.
				low = mid + 1
			} else {
				// move interval to left.
				high = mid - 1
			}
		}
	}

	return false, -1
}

func main() {

	input := []int{10, 12, 45, 65, 78, 100, 124, 1230}

	value := 124

	found, index := BinarySearch(input, value)

	fmt.Printf("\n List of integers : %v \n Search results for value %d : (%v, index:[%d]) \n", input, value, found, index)
}

// go run main.go
// Output:
// List of integers : [10 12 45 65 78 100 124 1230]
// Search results for value 124 : (true, index:[6])
