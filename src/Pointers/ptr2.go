package main

import "fmt"

/*
Go’s pointer syntax is partially borrowed from C and C++. Since Go has a garbage collector, most of
the pain of memory management is removed. Furthermore, some of the tricks that you can do with
pointers in C and C++, including pointer arithmetic, are not allowed in Go.

The Go standard library does have an unsafe package that lets you do some low-level operations on
data structures. While pointer manipulation is used in C for common operations, it is exceedingly rare
for Go developers to use unsafe.
*/

func main() {

	//A pointer type is a type that represents a pointer. It is written with a * before a type name. A pointer type
	//can be based on any type:

	x := 10
	var pointerToX *int
	pointerToX = &x
	fmt.Println(pointerToX)

	//The built-in function new creates a pointer variable. It returns a pointer to a zero value instance of the
	//provided type:

	k := new(int)         // Only takes new variable
	fmt.Println(k == nil) // prints false
	fmt.Println(*k)       // prints 0

	//The new function is rarely used. For structs, use an & before a struct literal to create a pointer
	//instance. You can’t use an & before a primitive literal (numbers, booleans, and strings) or a
	//constant because they don’t have memory addresses; they exist only at compile tim

	var y string
	z := &y
	fmt.Println(z)

}
