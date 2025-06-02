package main

import "fmt"

// list of integers and shift them left by a given number of places n
// looping around to the end of the list.
// given the input: [1,2,3,4,5,6,7,8] n = 3 => output: [4,5,6,7,8,1,2,3]

func MoveToLeft(input []int, n int) []int {
	lg := len(input)
	if lg <= 1 {
		return input
	}

	output := make([]int, lg)

	for currentIndex, value := range input {
		if newIndex := currentIndex - n; newIndex >= 0 {
			output[newIndex] = value
		} else {
			newIndex = lg + newIndex
			output[newIndex] = value
		}
	}
	return output
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	n := 3
	output := MoveToLeft(input, n)
	fmt.Println("output:", output)

	input = []int{1}
	output = MoveToLeft(input, n)
	fmt.Println("output:", output)

	input = []int{}
	output = MoveToLeft(input, n)
	fmt.Println("output:", output)

	input = nil
	output = MoveToLeft(input, n)
	fmt.Println("output:", output)
}
