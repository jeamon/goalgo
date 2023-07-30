package main

// A short nice coding challenge to demonstrate Depth-First Search (DFS)
// and Breadth-First Search (BFS) to check existence of a given node key.

// Version  : 1.0
// Author   : Jerome AMON
// Created  : 24 October 2021

import (
	"fmt"
)

type Node struct {
	value  string
	childs []*Node
}

var Root *Node

func init() {
	BuildTree()
}

// BuildTree constructs default Tree.
func BuildTree() {
	// First Leaf on branch A
	L0A := Node{"L0A", []*Node{}}
	// Fourth level node on branch A
	A4 := Node{"A4", []*Node{&L0A}}
	// Third level node on branch A
	A3 := Node{"A3", []*Node{&A4}}
	// Second level node on branch A
	A2 := Node{"A2", []*Node{&A3}}
	// First level node on branch A
	A1 := Node{"A1", []*Node{&A2}}

	// Second Leaf on branch A
	L1A := Node{"L1A", []*Node{}}
	// Fourth level node on branch A
	A44 := Node{"A44", []*Node{&L1A}}
	// Third level node on branch A
	A33 := Node{"A33", []*Node{&A44}}
	// Second level node on branch A
	A22 := Node{"A22", []*Node{&A33}}
	// First level node on branch A
	A11 := Node{"A11", []*Node{&A22}}

	// Root level node of branch A
	A0 := Node{"A0", []*Node{&A1, &A11}}

	// First Leaf on branch B
	L0B := Node{"L0B", []*Node{}}
	// Fourth level node on branch B
	B4 := Node{"B4", []*Node{&L0B}}
	// Third level node on branch B
	B3 := Node{"B3", []*Node{&B4}}
	// Second level node on branch B
	B2 := Node{"B2", []*Node{&B3}}
	// First level node on branch B
	B1 := Node{"B1", []*Node{&B2}}

	// Second Leaf on branch B
	L1B := Node{"L1B", []*Node{}}
	// Fourth level node on branch B
	B44 := Node{"B44", []*Node{&L1B}}
	// Third level node on branch B
	B33 := Node{"B33", []*Node{&B44}}
	// Second level node on branch B
	B22 := Node{"B22", []*Node{&B33}}
	// First level node on branch B
	B11 := Node{"B11", []*Node{&B22}}

	// Root level node of branch B
	B0 := Node{"B0", []*Node{&B1, &B11}}

	// First Leaf on branch C
	L0C := Node{"L0C", []*Node{}}
	// Fourth level node on branch C
	C4 := Node{"C4", []*Node{&L0C}}
	// Third level node on branch C
	C3 := Node{"C3", []*Node{&C4}}
	// Second level node on branch C
	C2 := Node{"C2", []*Node{&C3}}
	// First level node on branch C
	C1 := Node{"C1", []*Node{&C2}}

	// Second Leaf on branch C
	L1C := Node{"L1C", []*Node{}}
	// Fourth level node on branch C
	C44 := Node{"C44", []*Node{&L1C}}
	// Third level node on branch C
	C33 := Node{"C33", []*Node{&C44}}
	// Second level node on branch C
	C22 := Node{"C22", []*Node{&C33}}
	// First level node on branch C
	C11 := Node{"C11", []*Node{&C22}}

	// Root level node of branch C
	C0 := Node{"C0", []*Node{&C1, &C11}}

	// First Leaf on branch D
	L0D := Node{"L0D", []*Node{}}
	// Fourth level node on branch D
	D4 := Node{"D4", []*Node{&L0D}}
	// Third level node on branch D
	D3 := Node{"D3", []*Node{&D4}}
	// Second level node on branch D
	D2 := Node{"D2", []*Node{&D3}}
	// First level node on branch D
	D1 := Node{"D1", []*Node{&D2}}

	// Second Leaf on branch D
	L1D := Node{"L1D", []*Node{}}
	// Fourth level node on branch D
	D44 := Node{"D44", []*Node{&L1D}}
	// Third level node on branch D
	D33 := Node{"D33", []*Node{&D44}}
	// Second level node on branch D
	D22 := Node{"D22", []*Node{&D33}}
	// First level node on branch D
	D11 := Node{"D11", []*Node{&D22}}

	// Root level node of branch D
	D0 := Node{"D0", []*Node{&D1, &D11}}

	// Root of the Tree.
	Root = &Node{"R0", []*Node{&A0, &B0, &C0, &D0}}
}

// LocateMe uses Depth-First Search (DFS) to
// check if a given node-value is present.
// It traverses the tree depth.
func LocateMeByDFS(node *Node, me string) bool {
	fmt.Printf("=> %s\t", node.value)
	if len(node.childs) == 0 {
		// next branch so move to next line.
		fmt.Println()
	}

	if node.value == me {
		fmt.Println("\n\n:: found me by DFS.")
		return true
	}

	// recursive on childs nodes.
	if len(node.childs) > 0 {
		for _, child := range node.childs {
			if LocateMeByDFS(child, me) {
				// found so stop here.
				return true
			}
		}
	}

	return false
}

// LocateMeByBFS uses Breadth-First Search (BFS) to
// check if a given node-value is present.
// It traverses the tree width.
func LocateMeByBFS(node *Node, me string) bool {
	// slice to store nodes for processing.
	nodesToProcess := []*Node{node}
	// until there is a node - check.
	for len(nodesToProcess) > 0 {
		// pick up the first node.
		currentNode := nodesToProcess[0]
		// remove that node from the list.
		nodesToProcess = nodesToProcess[1:]
		fmt.Printf("=> %s\n", currentNode.value)
		if currentNode.value == me {
			// found me.
			fmt.Println("\n:: found me by BFS.")
			return true
		}

		// me not found, so add & check currentNode childs.
		if len(currentNode.childs) > 0 {
			for _, child := range currentNode.childs {
				nodesToProcess = append(nodesToProcess, child)
			}
		}
	}
	return false
}

func main() {
	if LocateMeByDFS(Root, "C3") == false {
		fmt.Println("\n\n:: not found me by DFS.")
	}

	fmt.Println("\n")

	if LocateMeByBFS(Root, "C3") == false {
		fmt.Println("\n\n:: not found me by BFS.")
	}
}

/*// full dot language transcript of the tree used into this code.

digraph Tree {

	{R0 [label="Root", shape=doublecircle, fontcolor=red]} -> {A0; B0; C0; D0;}

	A0 -> {A1; A11;}
	A1 -> A2 -> A3 -> A4 -> L0A;
	A11 -> A22 -> A33 -> A44 -> L1A;

	B0 -> {B1; B11;}
	B1 -> B2 -> B3 -> B4 -> L0B;
	B11 -> B22 -> B33 -> B44 -> L1B;

	C0 -> {C1; C11;}
	C1 -> C2 -> C3 -> C4 -> L0C;
	C11 -> C22 -> C33 -> C44 -> L1C;

	D0 -> {D1; D11;}
	D1 -> D2 -> D3 -> D4 -> L0D;
	D11 -> D22 -> D33 -> D44 -> L1D;
}

// install graphviz and use below to generate jpg.
~$ dot -Tjpg -otree.jpg tree.dot


// Output:

{nxos-geek}:~$ go run navigate-search-into-non-binary-trees.go
=> R0   => A0   => A1   => A2   => A3   => A4   => L0A
=> A11  => A22  => A33  => A44  => L1A
=> B0   => B1   => B2   => B3   => B4   => L0B
=> B11  => B22  => B33  => B44  => L1B
=> C0   => C1   => C2   => C3

:: found me by DFS.


=> R0
=> A0
=> B0
=> C0
=> D0
=> A1
=> A11
=> B1
=> B11
=> C1
=> C11
=> D1
=> D11
=> A2
=> A22
=> B2
=> B22
=> C2
=> C22
=> D2
=> D22
=> A3
=> A33
=> B3
=> B33
=> C3

:: found me by BFS.

*/
