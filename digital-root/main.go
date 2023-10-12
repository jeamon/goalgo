package main

import (
	"fmt"
	"strconv"
)

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 13 August 2021

// this is a nice coding challenge that implements two approach to compute the digital root of a given
// number. It consist of summing the digits of a number recursively until getting a number of digit.
// example : digital_root(12) => 1 + 2 = 3 // digital_root(12.4) = 1 + 2 + 3 + 4 = 10 = 1 + 0 = 1

// DigitalRootWithString is a function that compute digital root of a given number
// it works by using only string-based approach to extract digits and sum them.
func DigitalRootWithString(number int) int {

	// convert number from integer to string.
	s := strconv.Itoa(number)

	// corner case - if number of one digit.
	if len(s) == 1 {
		return number
	}

	// once here means - number has more than 1 digit inside.
	// variables to store conversion and to avoid redeclaring
	// new one for each iteration.
	var sum, digit int

	// loop over the string and add each character to the sum.
	for _, c := range s {
		// c is a rune type so convert into string type.
		// then convert from string to integer.
		digit, _ = strconv.Atoi(string(c))
		// add the digit to the sum.
		sum = sum + digit
	}

	// recursively perform on the sum value.
	return DigitalRootWithString(sum)
}

// DigitalRootWithModulo is a function that compute digital root of a given number
// it works by using only mathematical approach based on modulo to extract digits.
func DigitalRootWithModulo(number int) int {

	//corner case - check if number with one digit.
	if 10-number > 0 {
		return number
	}

	// once here - means number has more than 1 digit.
	// variables to store computation - default to 0 as integer.
	var sum, remainder int
	// perform euclidian division until quotient which
	// is the number equals 0.
	for number != 0 {
		// get the rest and add it to the sum.
		remainder = number % 10
		sum = sum + remainder
		// result of euclidean division by 10 will be used for next iteration.
		number = number / 10
	}

	// recursively perform on the sum value.
	return DigitalRootWithModulo(sum)
}

func main() {
	// sample numbers for testing purpose. each will be computed by each function.
	numbers := []int{0, 123, 1234, 12, 45785, 1, 45}
	for _, number := range numbers {
		fmt.Printf(" [string func] digital root of %d is : %d\n", number, DigitalRootWithString(number))
		fmt.Printf(" [modulo func] digital root of %d is : %d\n", number, DigitalRootWithModulo(number))
		fmt.Println() // make one line space.
	}
}

/* Output:

~$ go run digital-root-algorithms.go
 [string func] digital root of 0 is : 0
 [modulo func] digital root of 0 is : 0

 [string func] digital root of 123 is : 6
 [modulo func] digital root of 123 is : 6

 [string func] digital root of 1234 is : 1
 [modulo func] digital root of 1234 is : 1

 [string func] digital root of 12 is : 3
 [modulo func] digital root of 12 is : 3

 [string func] digital root of 45785 is : 2
 [modulo func] digital root of 45785 is : 2

 [string func] digital root of 1 is : 1
 [modulo func] digital root of 1 is : 1

 [string func] digital root of 45 is : 9
 [modulo func] digital root of 45 is : 9

*/
