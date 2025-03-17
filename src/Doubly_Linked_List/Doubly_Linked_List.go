package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// Insert a new node at the beginning of the list
func (dll *DoublyLinkedList) InsertAtBeginning(value int) {
	newNode := &Node{value: value, next: dll.head}
	if dll.head != nil {
		dll.head.prev = newNode
	} else {
		dll.tail = newNode // If the list was empty, update the tail
	}
	dll.head = newNode
}

// Insert a new node at the end of the list
func (dll *DoublyLinkedList) InsertAtEnd(value int) {
	newNode := &Node{value: value, prev: dll.tail}
	if dll.tail != nil {
		dll.tail.next = newNode
	} else {
		dll.head = newNode // If the list was empty, update the head
	}
	dll.tail = newNode
}

// Delete a node by value
func (dll *DoublyLinkedList) Delete(value int) {
	current := dll.head
	for current != nil && current.value != value {
		current = current.next
	}
	if current == nil {
		return // Value not found
	}
	if current.prev != nil {
		current.prev.next = current.next
	} else {
		dll.head = current.next // Update head if deleting first node
	}
	if current.next != nil {
		current.next.prev = current.prev
	} else {
		dll.tail = current.prev // Update tail if deleting last node
	}
}

// Search for a value in the list
func (dll *DoublyLinkedList) Search(value int) bool {
	current := dll.head
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

// Print the entire list from head to tail
func (dll *DoublyLinkedList) PrintList() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d <-> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	dll := &DoublyLinkedList{}

	// Test cases
	dll.InsertAtBeginning(10)
	dll.InsertAtBeginning(5)
	dll.InsertAtEnd(20)
	dll.InsertAtEnd(25)
	fmt.Println("Doubly Linked List after insertions:")
	dll.PrintList()

	dll.Delete(10)
	fmt.Println("Doubly Linked List after deleting 10:")
	dll.PrintList()

	fmt.Println("Searching for 20:", dll.Search(20))
	fmt.Println("Searching for 100:", dll.Search(100))

	// Exercises
	// 1. Implement a function to reverse the doubly linked list.
	// 2. Implement a function to find the middle element of the list.
	// 3. Implement a function to remove duplicates from the list.
	// 4. Try inserting and deleting elements, then print the list to observe changes.
}
