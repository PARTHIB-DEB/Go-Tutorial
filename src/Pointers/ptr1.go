package main

import "fmt"

/*
A pointer is simply a variable that holds the location in memory where a value is stored. Every variable
is stored in one or more contiguous memory locations, called addresses.

The smallest amount of memory that can be independently addressed is a byte. A pointer is simply a
variable whose contents are the address where another variable is stored.

The zero value for a pointer is nil.

nil is an untyped identifier that represents the lack of a
value for certain types.

The * is the INDIRECTION OPERATOR . It precedes a variable of pointer type and returns the pointed-to value.


*/

func main() {
	var x int32 = 10
	var y bool = true
	pointerX := &x
	pointerY := &y
	fmt.Println("Value of x:", x, "Address of x:", pointerX)
	fmt.Println("Value of y:", y, "Address of y:", pointerY)

	// This is called dereferencing
	// Its done with * operator
	x = 10
	pointerToX := &x
	fmt.Println(pointerToX) // or  fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z) // prints 15

	var k *int
	fmt.Println(k == nil) // prints true
	//fmt.Println(*k)       // panics

}
