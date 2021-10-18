package main

import (
	"fmt"
)

type Node struct {
	value  string
	childs []*Node
}

// locateMe uses Depth-First Search (DFS) to
// check if a given node-value is present.
func locateMe(node *Node, me string) bool {
	fmt.Println("=>", node.value)
	if node.value == me {
		fmt.Println("=> found me.")
		return true
	}

	// recursive on childs nodes.
	if len(node.childs) > 0 {
		for _, child := range node.childs {
			if locateMe(child, me) {
				// found so stop here.
				return true
			}
		}
	}
	// not found.
	return false
}

func main() {

	// Leaf on branch A
	LA := Node{"LA", []*Node{}}
	// Fourth level node on branch A
	A4 := Node{"A4", []*Node{&LA}}
	// Third level node on branch A
	A3 := Node{"A3", []*Node{&A4}}
	// Second level node on branch A
	A2 := Node{"A2", []*Node{&A3}}
	// First level node on branch A
	A1 := Node{"A1", []*Node{&A2}}

	// Leaf on branch B
	LB := Node{"LB", []*Node{}}
	// Fourth level node on branch B
	B4 := Node{"B4", []*Node{&LB}}
	// Third level node on branch B
	B3 := Node{"B3", []*Node{&B4}}
	// Second level node on branch B
	B2 := Node{"B2", []*Node{&B3}}
	// First level node on branch B
	B1 := Node{"B1", []*Node{&B2}}

	// Leaf on branch C
	LC := Node{"LC", []*Node{}}
	// Fourth level node on branch C
	C4 := Node{"C4", []*Node{&LC}}
	// Third level node on branch C
	C3 := Node{value:"C3", childs:[]*Node{&C4}}
	// Second level node on branch C
	C2 := Node{"C2", []*Node{&C3}}
	// First level node on branch C
	C1 := Node{"C1", []*Node{&C2}}

	// Leaf on branch D
	LD := Node{"LD", []*Node{}}
	// Fourth level node on branch D
	D4 := Node{"D4", []*Node{&LD}}
	// Third level node on branch D
	D3 := Node{value:"D3", childs:[]*Node{&D4}}
	// Second level node on branch D
	D2 := Node{"D2", []*Node{&D3}}
	// First level node on branch D
	D1 := Node{"D1", []*Node{&D2}}

	// Root of the Tree.
	R := Node{"R", []*Node{&A1, &B1, &C1, &D1}}

	_ = locateMe(&R, "C3")
}

/*
[ 0:06:47] {nxos-geek}:~$ go run navigate-into-tree.go

=> R
=> A1
=> A2
=> A3
=> A4
=> LA
=> B1
=> B2
=> B3
=> B4
=> LB
=> C1
=> C2
=> C3
=> found me.

*/