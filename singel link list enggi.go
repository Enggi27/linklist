// menghapus data diakhir list
package main

import "fmt"

// Node represents a node in the linked list
 type Node struct {
	data int
	next *Node
 }

 // LinkedList represents the linked list
 	type LinkedList struct {
		head *Node
	}

	// Append adds a new node at the end of the list
	func (list *LinkedList) Append(data int) {
		newNode := &Node {data: data}
		if list.head == nil {
			list.head == newNode 
			return
		}
		current := list.head
		for current.next != nil {
			current = current.next
		}
	}