package main

import "fmt"

/*
scope is the region where a variable/function or any component becomes active

scopes are determined by blocks.
majorly you can divide blocks into 2 sections - Universal , Functional

In Go , the universal block is called 'Package Block' - because in go , every script is under a package (default 'main')
Variables are called Global Variables here

In package block , the scope of the components are within the total script.
Also You don't need to use any 'fmt' function on those components
But in any functional scope , we MUST USE THE COMPONENT (BY fmt FUNCTIONS)

Using the power of scope , we can do SHADOWING

SHADOWING means hiding a component's activity by another component of same properties but different scope
Generally the functional scope SHADOWS global scopes if variables (or others) are exactly same

To detect Shadow Variables , Install this package -

golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest

Also we can perform shadowing , in nested scopes , where the child scope does the same with immediate parent scope components.

*/

var x string = "Hello World"

func main() {

	fmt.Printf("X in Package block - %s\n", x)
	x = "Hello World123"
	fmt.Printf("X in main(functional) block - %s\n", x)
	x, y := 10, 20
	fmt.Printf("Updated X and Y in main(functional) block - %d %d\n", x, y)
}
