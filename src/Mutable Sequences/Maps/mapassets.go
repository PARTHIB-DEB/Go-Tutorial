package main

import "fmt"

/*
SET - an unordered collection of UNIQUE elements
There is no built-in DS for SETs.

we can simulate a SET by using a MAP and a STRUCT (will discuss later)

LOOPS will be covered later
*/

func main() {
	intSet := map[int]bool{}                         // A map having INT key and BOOL value
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10} // A slice of all INT values
	for _, v := range vals {
		intSet[v] = true // Filling the map based on slice's values (this causes the map to act like a set)
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

}
