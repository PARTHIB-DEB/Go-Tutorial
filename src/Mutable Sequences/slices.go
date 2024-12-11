package main

import "fmt"

/*
The restrictions of array is removed by SLICE - a new DS , making it more useful

Slices has no defined size

Like Array we can't also read/write elements index by index in a Slice until final element comes

But we can Append elements in a slice

Zero Value of Slices - [0,0,0...]
*/

func main() {
	// Single Dimesional Slice for all datatypes (demo is shown using 'int' datatype)
	myslice := []int{1, 3, 4, 5} // using [] instead of [...] makes slices
	fmt.Printf("Integer : %v\n", myslice)

	//myslice[4] = 10	Error
	//myslice[5] = 50	Error

	fmt.Printf("slice of Slice : %v\n", myslice[1:3]) // Prints [1] [2]th elements

	// Sparse Slice
	sparseslice := []int{0: 5, 2: 4, 5: 8}
	fmt.Printf("Sparse Slice : %v\n", sparseslice)

	// Length of slice
	fmt.Printf("slice of Slice : %v\n", len(sparseslice))

	// Append element
	a_slice := append(myslice, 1000000)
	fmt.Printf("New Slice : %v\n", a_slice)
	fmt.Printf("slice of Slice : %v\n", myslice)

}
