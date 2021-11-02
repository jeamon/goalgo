package main

// A short nice coding challenge I got. You need to find out the number of slices which don't contain another slice.
// You get an input a slice of anything where anything could be a slice which contains itself another slice.
// Then your code should return a number of slice without any slice. This lets think about tree of slices
// where you need to find out the number of leaves slices into this tree. Below is a simple strings-based approach.
// The idea is to convert the auto-generated tree to string and count the number of empty slice <[]> into it.

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 02 November 2021

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// countLeavesSlices compute the number of empty slice into the tree of slices.
func countLeavesSlices(input string) int {
	return strings.Count(input, "[]")
}

// buildTree constructs a tree with with n top level nodes based on random nodes models.
func buildTree(n int) string {
	// our node templates.
	models := map[int]interface{}{
		1: [1]interface{}{[0]int{}},
		2: [2]interface{}{[0]int{}, [0]int{}},
		3: [3]interface{}{[0]int{}, [0]int{}, [2][0]int{[0]int{}, [0]int{}}},
		4: [4]interface{}{[0]int{}, [0]int{}, [2][0]int{[0]int{}, [0]int{}}, [2][0]int{[0]int{}, [0]int{}}},
		5: [5]interface{}{[0]int{}, [0]int{}, [0]int{}, [2][0]int{[0]int{}, [0]int{}}, [2][0]int{[0]int{}, [0]int{}}},
	}

	// https://pkg.go.dev/math/rand#Seed
	rand.Seed(time.Now().UnixNano())

	// tree with min / max number of nodes into a branch.
	val, min, max := 0, 1, 5

	// build a tree where root has n childs.
	tree := make([]interface{}, n)
	for i := 0; i < n; i++ {
		// generate random value between [min, max].
		val = rand.Intn(max-min+1) + min
		branch := make([]interface{}, val)
		for j := 0; j < val; j++ {
			k := rand.Intn(max-min+1) + min
			branch[j] = models[k]
		}

		tree[i] = branch
	}
	// convert to json string and indent with 3 spaces.
	bytesTree, _ := json.MarshalIndent(tree, "", "   ")

	// replace all commas by newline for nice displaying.
	stringTree := fmt.Sprintf("%s\n", bytes.ReplaceAll(bytesTree, []byte(","), []byte("\n")))

	// display the Tree of slice.
	fmt.Println(stringTree)

	return stringTree
}

func main() {
	myTree := buildTree(2)
	fmt.Println("Number of empty/leaf slices:", countLeavesSlices(myTree))
}

/*
Output:

~$ go run number-of-leaves-slices-into-tree.go

[
   [
      [
         []
      ]
   ]

   [
      [
         []
      ]

      [
         []

         []
      ]
   ]
]

Number of empty/leaf slices: 4

*/
