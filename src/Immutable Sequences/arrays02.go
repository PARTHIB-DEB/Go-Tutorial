package main

import "fmt"

/*
We can create different types and dimensions of arrays based on datatypes
*/

func main() {
	// Defined Size
	// Integer ( X.0 is also considered as Integer )
	x := [5]int{1, -1, 1.0, -9.0, 10000}
	fmt.Printf("\n Integer : %v", x)

	// Rune
	y := [10]rune{'a', 'b', 'c', 'd', 'e', 'f'}
	fmt.Printf("\n Rune : %v", y)

	// byte
	z := [4]byte{'a', 'b', 'c', 'd'}
	fmt.Printf("\n Byte : %v", z)

	// More types of arrays possible

	// Sparse Array {index:value}
	k := [8]int{0: 4, 1: 1, 2: 2, 4, 6}
	fmt.Printf("\n Sparse : %v", k)

	// Multidimensional Array

	var matrix [2][2]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[1][0] = 3
	matrix[1][1] = 4

	fmt.Printf("\n matrix : %v", matrix)

	// String

	var op = [2]string{"abc", "wer"}
	fmt.Printf("\n op : %v\n", op)

}
