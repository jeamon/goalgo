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
