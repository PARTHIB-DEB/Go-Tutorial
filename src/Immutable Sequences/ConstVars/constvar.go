package main

import "fmt"

/*
const - keyword is used to make a constant
*/

const x int = 20
const z = 20 * 10

func main() {
	const y = "hello"
	// x = x + y 	// causes error
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
