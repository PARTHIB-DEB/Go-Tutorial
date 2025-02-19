package main

import (
	"fmt"
	"sort"
)

/*

Functions declared inside of functions are special; they are closures

Functions declared inside of functions are able to access and modify
variables declared in the outer function.

Often Anonymous Functions (Look inside Function/func05.go) as the inner function / closure function
*/

// Example - 1
func generatesum(i int, j int) {
	for k := i; k <= j; k++ {
		var summ int = 1
		func(num int) {
			summ = summ + num
			fmt.Printf("1 + %d = %d\n", num, summ)
		}(k)
	}
}

func main() {

	generatesum(5, 10)

	// Example - 2
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Print("Before Modifying:\n")
	fmt.Println(people)
	sort.Slice(people, func(i int, j int) bool { // Closure Function
		return people[i].LastName < people[j].LastName
	})
	fmt.Print("After Modifying using closure:\n")
	fmt.Println(people)

}
