package main

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
}

// Insert a new node at the beginning of the list
func (ll *LinkedList) InsertAtBeginning(value int) {
	newNode := &Node{value: value, next: ll.head}
	ll.head = newNode
}

// Insert a new node at the end of the list
func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Delete a node by value
func (ll *LinkedList) Delete(value int) {
	if ll.head == nil {
		return
	}
	if ll.head.value == value {
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

// Search for a value in the list
func (ll *LinkedList) Search(value int) bool {
	current := ll.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

// Print the entire list
func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}

	// Test cases
	ll.InsertAtBeginning(10)
	ll.InsertAtBeginning(5)
	ll.InsertAtEnd(20)
	ll.InsertAtEnd(25)
	fmt.Println("Linked List after insertions:")
	ll.PrintList()

	ll.Delete(10)
	fmt.Println("Linked List after deleting 10:")
	ll.PrintList()

	fmt.Println("Searching for 20:", ll.Search(20))
	fmt.Println("Searching for 100:", ll.Search(100))

	// Exercises
	// 1. Try inserting multiple elements and printing the list.
	// 2. Delete an element that doesn’t exist and check the result.
	// 3. Implement a function to reverse the linked list.
	// 4. Implement a function to find the middle element of the linked list.
	// 5. Implement a function to remove duplicates from the linked list.
}
