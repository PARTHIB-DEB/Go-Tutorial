package main

import (
	"fmt"
	"math/rand"
)

/*
Like many other languages , If-else is also a form of conditional logic building in Go

Structure -
if(expression){
	// Logic
}else{
	// Logic
}
*/

func main() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	// Special Scoping in If-Else

	if name := "Parthib"; name == "" { // Temporary used variable - name (its scope is limited to this if-else ladder)
		fmt.Printf("No Name \n")
	} else {
		fmt.Printf("Name %s\n", name)
	}

	if nk := rand.Intn(10); nk == 0 { // n is a temporary used variable , scope limited to if-else ladder
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	//fmt.Print(nk)   // undefined here so panick
	//fmt.Println(name)  // undefined here so panick

	// But watch this example below (It may be an error in this version of Go)

	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else {
		fmt.Println("That's a good number:", n)
	}

	fmt.Println(n) // This doesn't cause any error , No books cover this - It only happens with character 'n'

}
