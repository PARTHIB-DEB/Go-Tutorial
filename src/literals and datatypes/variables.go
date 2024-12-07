package main

import "fmt"

var (
	A string = "String Data"
	B bool   = true
)

func main() {
	// Initialization and Store value in variable
	var x int
	x = 10
	fmt.Println(x)

	var y float64 = 190.12
	fmt.Println(y)

	var (
		Y string = "String Data"
		X bool   = true
	)

	fmt.Println(Y, X)

	// Using Walrus Operator
	k := "This is string"
	fmt.Println(k)

	k = "Modifying without walrus"
	fmt.Println(k)

	k, p := "Modified with walrus", 12 // Need atleast one more NEW variable - p here
	fmt.Println(k, p)

	// Using Const operator (Can't change the value of the variable)
	// Can't leave a variable Uninitialized if using Const
	// Also can't change the value of the variable

	const (
		q string = "String using Const"
		b bool   = false
	)
	fmt.Println(q, b)

}
