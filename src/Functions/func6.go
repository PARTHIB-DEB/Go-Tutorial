package main

import "fmt"

/*
You can return a function (a reference of a function) from a function
*/

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := makeMult(2)   // returning a reference of a function
	threeBase := makeMult(3) // returning a reference of a function
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}

}
