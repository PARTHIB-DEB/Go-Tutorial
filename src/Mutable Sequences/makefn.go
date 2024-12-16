package main

import "fmt"

func main() {

	/*
		Make is a function which is used to create a SLICE DS
		The make function creates a ZEROED Array (Full of Zero) and returns a slice that refers
		to that Array.

		A slice doesn't store any data, it just describes a section of an underlying array.
		Changing the elements of a slice modifies corresponding elements of its underlying array.
		Other slices that share the same underlying array will also see those changes.

		The slice is just a reference to its underlying array.

		A slice has both LENGTH & CAPACITY

		LENGTH - Number of elements the SLICE has
		CAPACITY - Number of elements in Underlying ARRAY

		You can extend a slice's length by re-slicing it , provided it has sufficient capacity.

	*/

	/*
		The slice X has both Length and Capacity as 5. Means from X[0] to X[4]
		all elements are valid
	*/

	s := []int{1, 2, 3, 4, 5} // Length = Capacity = 5
	fmt.Printf("Length of s : %d , Capacity of s : %d\n", len(s), cap(s))

	/*
		Slicing slices makes SHALLOW COPY (new slice indicate to same memory of previous slice) .
		That's why , you can do the below steps.

		Whenever a slice of another SLICE , the SUBSLICE's capacity is set to the
		Capacity of Original Slice - Capacity of offset of the Subslice.
		This means , that any unused capacity in the original slice is also shared within any subslices.
	*/

	s = s[:0] // Shrinking
	fmt.Printf("Slice S : %v\n", s)
	fmt.Printf("Length of s : %d , Capacity of s : %d\n", len(s), cap(s))

	s = s[:4] // Expanding
	fmt.Printf("Slice S : %v\n", s)
	fmt.Printf("Length of s : %d , Capacity of s : %d\n", len(s), cap(s))

	x := make([]int, 5) // Length = 5 , Capacity = 5
	fmt.Printf("Slice x : %v\n", x)
	fmt.Printf("Length of X : %d , Capacity of A : %d\n", len(x), cap(x))

	a := make([]int, 0, 5) // Length = 0 , Capacity = 5
	fmt.Printf("Length of A : %d , Capacity of A : %d\n", len(a), cap(a))

	//ap := make([]int, 5, 3) NEVER MAKE CAPACITY LESSER THAN LENGTH - PANIC/ERROR
	//fmt.Printf("Length of A : %d , Capacity of A : %d\n", len(ap), cap(ap))

	a = append(a, 10)
	fmt.Printf("A : %d\n", a)
	fmt.Printf("Length of A : %d , Capacity of A : %d\n", len(a), cap(a))
	a = append(a, 10)
	a = append(a, 10)
	fmt.Printf("A : %d\n", a)
	fmt.Printf("Length of A : %d , Capacity of A : %d\n", len(a), cap(a))
	a = append(a, 10)
	a = append(a, 10)
	a = append(a, 10)
	/*
		At this moment , the slice has a length > Capacity - which is fundamentally wrong
		so the capacity becomes - capacity = 2 * capacity
	*/
	fmt.Printf("A : %d\n", a)
	fmt.Printf("Length of A : %d , Capacity of A : %d\n", len(a), cap(a))

}
