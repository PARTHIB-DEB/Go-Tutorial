package main

import (
	"fmt"
	"math/rand"
)

/*
Apart from Switch-Case is another way to implement conditional Control Flow over a program

When we need to take decision/output based on specific types of inputs , we use SWITCH-CASE

*/

func main() {

	// Blank Switch (The input variable is directly judged in CASES)

	a := rand.Intn(10)
	switch {
	case a == 2:
		fmt.Println("a is 2")
	case a == 3:
		fmt.Println("a is 3")
	case a == 4:
		fmt.Println("a is 4")
	default:
		fmt.Println("a is ", a)
	}

	// Normal Switch
	switch b := rand.Intn(10); b {
	case 2:
		fmt.Println("b is 2")
	case 3:
		fmt.Println("b is 3")
	case 4:
		fmt.Println("b is 4")
	default:
		fmt.Println("b is ", b)
	}
}
