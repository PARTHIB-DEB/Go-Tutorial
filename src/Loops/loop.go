package main

import "fmt"

/*
When you have to access each element of an Iterable Sequence ,
You need Loops

There are several kind of Loops , but in GO - Only For Loop and its variants are allowed

There are 4 kinds of For-Loop in Go -
1. Complete For Loop (C-style)
2. Condition-Only For Loop (Simulation of WHILE loop)
3. Infinite For Loop  (Naked Loop)
4. For-Range Loop (Most Useful , Helpful over SLICES , ARRAYS , MAPS etc)

*/

func main() {
	// C style
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()

	// Condition-Only Loop
	i := 1
	for i < 100 {
		fmt.Printf("%d\t", i)
		i = i * 2
	}

	// Infinite Loop
	//for {
	//	fmt.Println("Hello")
	//}

	// For-Range Loop
	evenVals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range evenVals {
		fmt.Printf("Key :%d\t %d\n", i, v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k) // Fetching only Keys
	}

}
