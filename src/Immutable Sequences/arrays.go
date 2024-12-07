package main

import "fmt"

/*

In Go , Arrays are rigid (Immutable) and homogenous

Each Mutable sequence has their ZERO / DEFAULT value.
For Arrays -> var arr [X] = [0,0,0,0.....upto X times]

Reasons behind Array is not so much in GO -
1. Go considers the array to be a part of array . That's why '[3]int arr{}' is different from
	'[100]int arr1{}'
2. This also means that you can't use a variable to specify the size of array , because types
	must be resolved at Compile Time , not at Runtime
3. What's more , you can't use a type conversion to convert arrays of different sizes to identical
	types.

*/

func main() {
	var arr = [3]int{1, 2, 4}                  // Defined Size
	var arr1 = [...]int{10, 11, -1, 0, 38, 10} // Undefined Size
	fmt.Println(arr)
	fmt.Println(arr1)
}
