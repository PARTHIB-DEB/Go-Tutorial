package main

import "fmt"

/*
Not only SLICES but ARRAYS can also be sliced . One can take a slice of an array using SLICE EXPRESSION
This is very fruitful to connect an array to a function which only takes slices.
*/

func main() {
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10 // changes in array reflects in slices (shallow copy)
	fmt.Printf("x is of Type :%T\n", x)
	fmt.Printf("y is of Type :%T\n", y)
	fmt.Printf("z is of Type :%T\n", z)
	fmt.Println("x :", x)
	fmt.Println("y :", y) // shows the changes
	fmt.Println("z :", z)
}
