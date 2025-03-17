package main

import (
	"errors"
	"fmt"
)

// LinkedList defines common operations for list implementations
type LinkedList interface {
	InsertAtBeginning(value int) error
	InsertAtEnd(value int) error
	Delete(value int) error
	Search(value int) bool
	PrintList()
	IsEmpty() bool
	Size() int
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

// IsEmpty checks if the list is empty
func (dll *DoublyLinkedList) IsEmpty() bool {
	return dll.head == nil
}

// Size returns the number of nodes in the list
func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (dll *DoublyLinkedList) InsertAtBeginning(value int) error {
	newNode := &Node{value: value, next: dll.head}
	if dll.head != nil {
		dll.head.prev = newNode
	} else {
		dll.tail = newNode // If the list was empty, update the tail
	}
	dll.head = newNode
	dll.size++
	return nil
}

// InsertAtEnd inserts a new node at the end of the list
func (dll *DoublyLinkedList) InsertAtEnd(value int) error {
	newNode := &Node{value: value, prev: dll.tail}
	if dll.tail != nil {
		dll.tail.next = newNode
	} else {
		dll.head = newNode // If the list was empty, update the head
	}
	dll.tail = newNode
	dll.size++
	return nil
}

// Delete removes a node by value and returns an error if not found
func (dll *DoublyLinkedList) Delete(value int) error {
	if dll.IsEmpty() {
		return errors.New("cannot delete from empty list")
	}

	current := dll.head
	for current != nil && current.value != value {
		current = current.next
	}

	if current == nil {
		return errors.New("value not found in the list")
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
	
	dll.size--
	return nil
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

// PrintList prints the entire list from head to tail
func (dll *DoublyLinkedList) PrintList() {
	if dll.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	
	current := dll.head
	for current != nil {
		fmt.Printf("%d <-> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	// Use the interface type for polymorphism
	var list LinkedList = &DoublyLinkedList{}

	// Test cases
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(5)
	list.InsertAtEnd(20)
	list.InsertAtEnd(25)
	fmt.Println("Doubly Linked List after insertions:")
	list.PrintList()

	if err := list.Delete(10); err == nil {
		fmt.Println("Doubly Linked List after deleting 10:")
		list.PrintList()
	} else {
		fmt.Println("Error:", err)
	}

	// Test error handling
	if err := list.Delete(100); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Searching for 20:", list.Search(20))
	fmt.Println("Searching for 100:", list.Search(100))
	fmt.Println("List size:", list.Size())
}
