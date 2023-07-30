package main

import (
	"fmt"
)

func Factorial(n int) int {
	if n == 1 {
		return 1
	}

	return n * Factorial(n-1)
}

func main() {

	for _, n := range []int{3, 5, 7} {
		fmt.Printf("%d => %d\n", n, Factorial(n))
	}
}

/*
// add below code into a file name factorial_test.go
// to execute the test, see below example.

~$ go test factorial.go factorial_test.go -v
=== RUN   TestFactorial
--- PASS: TestFactorial (0.00s)
PASS
ok      command-line-arguments  0.202s

[23:05:58] {nxos-geek}:~$

// testing source code.

package main

import "testing"

func TestFactorial(t *testing.T) {
	casesTable := [][2]int{
		[2]int{1, 1},
		[2]int{2, 2},
		[2]int{3, 6},
		[2]int{4, 24},
		[2]int{5, 120},
	}

	for _, input := range casesTable {
		got := Factorial(input[0])
		if got != input[1] {
			t.Errorf("got %d, wanted: %d", got, input[1])
		}
	}
}

*/
