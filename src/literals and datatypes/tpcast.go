package main

import "fmt"

/*
Go doesn’t allow automatic type promotion between variables. You must use a type conversion when
variable types do not match. Even different-sized integers and floats must be converted to the same
type to interact.

*/

func main() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = x + int(y)

	fmt.Printf("\nValue of x is : %v\n", x)
	fmt.Printf("\nValue of y is : %v\n", y)
	fmt.Printf("\nValue of z is : %v\n", z)
	fmt.Printf("\nValue of d is : %v\n", d)
}
