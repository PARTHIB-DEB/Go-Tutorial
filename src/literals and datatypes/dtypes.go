/*
	4 Types of Datatypes/Literals are there - Integer , Floating Point , Rune & raw-string , booleans
	5th one is rare - Complex Number
*/

package main

import "fmt"

// Integer
var num int = 18
var num1 int8 = 100
var num2 int16 = -200
var num3 int32 = 1249
var num4 int64 = 100000

// Float
var fnum float32 = 13.12
var fnum1 float64 = -12243.90

// Runes & Strings
var rstr rune = 'a' // Always single quote - Single character
var str string = "Hello"

// Boolean
var bvar bool = true

// Complex Number
var cvar complex64 = 4 + 5i

// ZERO / DEFAULT Values of each datatype

var ab int
var bc float32
var cd rune
var ef string
var gh bool
var op complex64

func main() {
	// %v is common for all datatypes (%v means verbose - as it is)
	fmt.Printf("\nIntegers : %d %d %d %d %d", num, num1, num2, num3, num4)
	fmt.Printf("\nFloats : %f %f", fnum, fnum1)
	fmt.Printf("\nRune : %c", rstr)
	fmt.Printf("\nString : %s", str)
	fmt.Printf("\nBool : %v", bvar)
	fmt.Printf("\nComplex64 : %v", cvar)

	fmt.Printf("\nZERO(Integer) : %d", ab)
	fmt.Printf("\nZERO(Float) : %f", bc)
	fmt.Printf("\nZERO(Rune) : %c", cd)
	fmt.Printf("\nZERO(String) : %s", ef)
	fmt.Printf("\nZERO(Bool) : %v", gh)
	fmt.Printf("\nZERO(Complex No) : %v", op)
}
