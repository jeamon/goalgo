package main

import "fmt"

// this is a nice coding challenge where you need to implement a function that takes as input a string containing
// a combinaison of common type of brackets like square brackets, curly braces and round brackets, then return a
// boolean stating if their order is valid or not. In short, if the string representation makes sense or not.

// The idea is to loop over the string and stack (append to a slice) any opening bracket then pop (remove the last)
// from that stack when the loop hits a closing bracket. Then compare if the popped element is the opening bracket
// of the closing bracket we did hit. If not, then the full string is not a valid sequence of brackets. Otherwise,
// continue until we have processed all element into the string and check at the end if we have something left into
// the stack. If we have, then the string is not valid sequence. If empty, then it is a valid sequence of brackets.
// A slice will be used as a stack. So the time & space complexity are both O(n) where n is total number of brackets.

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 22 October 2021

// Below are two global variable for basics checking.
// The 1st stores the valid brackets and the second provides
// closing & opening brackets correspondence.
// You can avoid them and use a manual approach too.

var brackets map[string]struct{} = map[string]struct{}{
	"[": struct{}{},
	"(": struct{}{},
	"{": struct{}{},
	"]": struct{}{},
	")": struct{}{},
	"}": struct{}{},
}

var mappings map[string]string = map[string]string{
	"[": "]",
	"(": ")",
	"{": "}",
}

// IsValidateSequenceOfBrackets implements above described algorithm.
func IsValidateSequenceOfBrackets(expression string) bool {
	// define our stack / empty slice.
	queue := []string{}
	// temporarily store converted rune to string value.
	var bracket string

	for _, char := range expression {

		if _, isValid := brackets[string(char)]; isValid == false {
			// not an open or close bracket symbol so abort here.
			return false
		}

		// we are sure it is a bracket (opening or closing).
		bracket = string(char)

		if _, isOpen := mappings[bracket]; isOpen {
			// hit opening bracket symbol. so add to stack.
			queue = append(queue, bracket)

		} else {

			// hit closing bracket, abort if the queue is empty.
			if len(queue) == 0 {
				return false
			}

			// we still have bracket(s), so retrieve/pop the latest added one.
			latestAddedOpeningBracket := queue[len(queue)-1]
			queue = queue[:len(queue)-1]

			// now check if it is or not the corresponding opening bracket.
			if correspondingClosingBracket, _ := mappings[latestAddedOpeningBracket]; correspondingClosingBracket != bracket {
				return false
			}
		}
	}

	// not more character into the expression to process.
	// check if we still have bracket(s) into the queue.
	if len(queue) > 0 {
		return false
	}

	return true
}

func main() {

	examples := []string{"[][]", "[()(){}]", "()", "[[]]", "[()()]{}", "(a)", "(()]", ")", "{}(", "{(})", "}"}
	for _, expression := range examples {

		if IsValidateSequenceOfBrackets(expression) {
			fmt.Println(expression, "is valid sequence.")
		} else {
			fmt.Println(expression, "is not valid sequence.")
		}
	}
}

/* Output:

[17:54:38] {nxos-geek}:~$ go run validate-brackets-sequence.go

[][] is valid sequence.
[()(){}] is valid sequence.
() is valid sequence.
[[]] is valid sequence.
[()()]{} is valid sequence.
(a) is not valid sequence.
(()] is not valid sequence.
) is not valid sequence.
{}( is not valid sequence.
{(}) is not valid sequence.
} is not valid sequence.

*/
