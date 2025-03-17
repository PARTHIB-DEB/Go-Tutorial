package main

import (
	"errors"
	"fmt"
)

// List interface to enable polymorphism for different list implementations
type List interface {
	InsertAtBeginning(value int) error
	InsertAtEnd(value int) error
	Delete(value int) error
	Search(value int) bool
	PrintList()
}

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

// InsertAtBeginning inserts a new node at the beginning of the list
func (ll *LinkedList) InsertAtBeginning(value int) error {
	newNode := &Node{value: value, next: ll.head}
	ll.head = newNode
	return nil
}

// InsertAtEnd inserts a new node at the end of the list
func (ll *LinkedList) InsertAtEnd(value int) error {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		return nil
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return nil
}

// Delete removes a node with specified value
func (ll *LinkedList) Delete(value int) error {
	if ll.head == nil {
		return errors.New("cannot delete from empty list")
	}
	
	if ll.head.value == value {
		ll.head = ll.head.next
		return nil
	}
	
	current := ll.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}
	
	if current.next == nil {
		return errors.New("value not found in list")
	}
	
	current.next = current.next.next
	return nil
}

// Search looks for a value in the list
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

// PrintList displays the entire list
func (ll *LinkedList) PrintList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	var list List = &LinkedList{}

	// Test cases
	list.InsertAtBeginning(10)
	list.InsertAtBeginning(5)
	list.InsertAtEnd(20)
	list.InsertAtEnd(25)
	fmt.Println("Linked List after insertions:")
	list.PrintList()

	if err := list.Delete(10); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Linked List after deleting 10:")
	list.PrintList()

	// Test error handling
	if err := list.Delete(100); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Searching for 20:", list.Search(20))
	fmt.Println("Searching for 100:", list.Search(100))
}
