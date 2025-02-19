package main

import "fmt"

/*
you can also define new functions within a function and assign them to variables.
These inner functions are anonymous functions

They don’t have a name.
You can write them inline and call them immediately

*/

func main() {

	for i := 0; i < 10; i++ {
		// Inner Function
		func(j int) {
			fmt.Printf("Inside Function Printing : %d\n", j)
		}(i)
	}

}
