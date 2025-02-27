package main

import "fmt"

/*

Go developers use pointers to indicate that a parameter is mutable.
Since Go is a 'call by value' language, the values passed to functions are copies.
For nonpointer types like primitives, structs, and arrays, this means that the called function cannot modify the original.
Since the called function has a copy of the original data, the immutability of the original data is guaranteed.

However, if a pointer is passed to a function, the function gets a copy of the pointer. This still points
to the original data, which means that the original data can be modified by the called function.

*/

// In this case you are a COPY of the value of the pointer variable , not the address of a normal variable
func failedUpdate(g *int) {
	x := 10
	g = &x
} // have to return the value of the pointer , signifies why go is a call-by-value lang

// In this case , you are directly passing the address of a normal variable
func failedUpdate2(px *int) {
	x2 := 20
	px = &x2
}
func update(px *int) {
	*px = 20
}
func main() {
	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f) // prints nil

	x := 10
	failedUpdate2(&x)
	fmt.Println(x) // prints 10
	update(&x)
	fmt.Println(x) // prints 20
}
