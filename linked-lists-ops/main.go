package main

import "fmt"

type Node struct {
	value any
	next  *Node
}

type LinkedList struct {
	head *Node
}

/*
1 -> 2 -> 3 -> 4

previous = nil
current = 1
next = curr.next = 2

--

cur.next = prev
current = next
next = current.next


*/

func Reverse(root *LinkedList) *LinkedList {
	var previous *Node
	var next *Node
	current := root.head

	for current != nil {
		next = current.next
		current.next = previous
		previous = current
		current = next
	}

	return &LinkedList{head: previous}
}

func Construct(items []int) *LinkedList {
	lg := len(items)
	switch lg {
	case 0:
		return &LinkedList{head: nil}
	case 1:
		return &LinkedList{head: &Node{value: items[0], next: nil}}
	}
	head := &Node{value: items[0], next: &Node{}}
	current := head
	for i := 1; i < lg; i++ {
		node := &Node{value: items[i], next: nil}
		current.next = node
		current = node
	}
	return &LinkedList{head: head}
}

func Display(root *LinkedList) {
	current := root.head
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	root := Construct(items)
	Display(root)

	reversed := Reverse(root)
	Display(reversed)
}
