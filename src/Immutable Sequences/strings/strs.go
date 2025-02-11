package main

import "fmt"

/*

String is a built-in type in Go .
The zero value for a string is the empty string (" ").
Strings in Go are immutable.
The rune type is an alias for the int32 type, just like byte is an alias for uint8.


string is composed of a sequence of UTF-8-encoded code points.
Just like you can extract a single value from an array or a slice, you can extract a sin‐ gle value from a
string by using an index expression :

runes are closest datatype to string

*/

func main() {
	var s string = "Hello there"
	var a byte = s[6]
	var p rune = 'x'
	var k string = string(a)
	var b byte = 'y'
	var s2 string = string(b)

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(k)
	fmt.Println(b)
	fmt.Println(s2)
}
