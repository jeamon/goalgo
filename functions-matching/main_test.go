package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMatches(t *testing.T) {
	cases := []struct {
		input     []Function
		argsTypes []string
		want      []string
	}{
		{
			input:     nil,
			argsTypes: []string{},
			want:      []string{},
		},
		{
			input: []Function{
				{name: "fnA", argumentTypes: []string{"bool", "int"}, isVariadic: false},
				{name: "fnB", argumentTypes: []string{"int"}, isVariadic: false},
				{name: "fnC", argumentTypes: []string{"int"}, isVariadic: true},
			},
			argsTypes: []string{"bool", "int"},
			want:      []string{"fnA"},
		},
		{
			input: []Function{
				{name: "fnA", argumentTypes: []string{"bool", "int"}, isVariadic: false},
				{name: "fnB", argumentTypes: []string{"int"}, isVariadic: false},
				{name: "fnC", argumentTypes: []string{"int"}, isVariadic: true},
			},
			argsTypes: []string{"int"},
			want:      []string{"fnB", "fnC"},
		},
		{
			input: []Function{
				{name: "fnA", argumentTypes: []string{"bool", "int"}, isVariadic: false},
				{name: "fnB", argumentTypes: []string{"int"}, isVariadic: false},
				{name: "fnC", argumentTypes: []string{"int"}, isVariadic: true},
			},
			argsTypes: []string{"int", "int", "int"},
			want:      []string{"fnC"},
		},
		{
			input: []Function{
				{name: "fnA", argumentTypes: []string{"bool", "int"}, isVariadic: true},
				{name: "fnB", argumentTypes: []string{"int"}, isVariadic: false},
				{name: "fnC", argumentTypes: []string{"bool", "int", "int"}, isVariadic: true},
			},
			argsTypes: []string{"bool", "int", "int", "int"},
			want:      []string{"fnA", "fnC"},
		},
	}

	for n, tc := range cases {
		t.Run(fmt.Sprint("case:", n+1), func(t *testing.T) {
			fl := NewFunctionLibrary(listFormat)
			fl.Register(tc.input...)
			got := fl.FindMatches(tc.argsTypes...)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got: %v but want: %v", got, tc.want)
			}
		})
	}
}

/*

go test -timeout 30s -run ^TestFindMatches$ github.com/jeamon/goalgo/functions-matching -count=1 -v
=== RUN   TestFindMatches
=== RUN   TestFindMatches/case:1
=== RUN   TestFindMatches/case:2
=== RUN   TestFindMatches/case:3
=== RUN   TestFindMatches/case:4
=== RUN   TestFindMatches/case:5
--- PASS: TestFindMatches (0.00s)
    --- PASS: TestFindMatches/case:1 (0.00s)
    --- PASS: TestFindMatches/case:2 (0.00s)
    --- PASS: TestFindMatches/case:3 (0.00s)
    --- PASS: TestFindMatches/case:4 (0.00s)
    --- PASS: TestFindMatches/case:5 (0.00s)
PASS
ok      github.com/jeamon/goalgo/functions-matching     0.169s

*/
