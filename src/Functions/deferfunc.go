package main

import "fmt"

/*
In Go language, defer statements delay the execution of the function
or method or an anonymous method until the nearby functions returns. In other words, defer function or method call arguments evaluate instantly, but they don’t execute
until the nearby functions returns

Programs often create temporary resources, like files or network connections, that need to be
cleaned up

a defer method/function/module executes all its member or child function/methods/module until its left
*/

func sumtwonums(a int, b int) {
	fmt.Printf("Sum of %d and %d is: %d\n", 4, 5, 4+5)
}

func main() {
	defer sumtwonums(3, 4)
}
