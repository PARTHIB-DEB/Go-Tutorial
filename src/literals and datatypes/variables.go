package main

/*
	There are 2 ways to create variables -
	1. using 'var' keyword
	2. using ':=' (walrus) operator

	':='  is for both declaration, assignment, and also for redeclaration;
	It guesses (infers) the variable's type automatically.

	------------------------COMPARISON-------------------------------

	'v:=32'						|				var v int
								|				v = 32
								|		OR
								|				var v int = 32
								|		OR
								|				var v = 32

	-------------------------------------------------------------------

	Can't declare same variable by := in same SCOPE.
	Can't use := Outside of package-name function (Scopes will be discussed later)

	More about := operator , click the link below -
	(Ans-365) => (https://stackoverflow.com/questions/17891226/difference-between-and-operators-in-go)
*/

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
