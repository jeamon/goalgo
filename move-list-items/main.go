package main

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
			newIndex = newIndex + lg
			if newIndex < 0 {
				newIndex *= -1
			}
			output[newIndex] = value
		}
	}
	return output
}

// list of integers and shift them right by a given number of places n
// looping around to the end of the list.
// given the input: [1,2,3,4,5,6,7,8] n = 3 => output: [6,7,8,1,2,3,4,5]

func MoveToRight(input []int, n int) []int {
	lg := len(input)
	if lg <= 1 {
		return input
	}

	output := make([]int, lg)

	for currentIndex, value := range input {
		if newIndex := currentIndex + n; newIndex < lg {
			output[newIndex] = value
		} else {
			newIndex = newIndex % lg
			output[newIndex] = value
		}
	}
	return output
}
