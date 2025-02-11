package main

import "fmt"

/*
Rather than use the slice and index expressions with strings, you should extract sub‐ strings and code
points from strings using the functions in the strings and unicode/ utf8 packages in the standard library.

*/

func main() {
	var s string = "Hello, "
	var bs []byte = []byte(s) // slice of bytes
	var rs []rune = []rune(s) // slice of runes
	fmt.Printf("Value : %s and Type : %T\n", bs, bs)
	fmt.Printf("Value : %x and Type : %T\n", rs, rs)
}
