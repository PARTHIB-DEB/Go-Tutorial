/*
	Any Go File has this default structure -
	1. A package name , which we are creating in this script (default is 'main'). A package is a way to group total go code
	2. A single endpoint function in the name of package
	3. Any imported packages (atleast 'fmt' - a standard library package)
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
