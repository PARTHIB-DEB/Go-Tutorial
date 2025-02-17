package main

import (
	"errors"
	"fmt"
)

/*

A subcase of Multiple Return Values - Named Return Values
Just like named parameter's implementation , you can now mention exactly what things are expected to get as return value/s
from the function

In this case - either specify all returns or don't specify anyone

Also there are functions which is expected to return anything

*/

func divAndRemainderNamed(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return // Can do this also , the compiler will do its job here (means always 3 values will be returned in the same order of param)
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

// Variadic Input Parameters and slices

/*
Variadic means Varies in Size. Go supports variadic parameters. The variadic parameter must be the last (or only) parameter in the
input parameter list.

You indicate it with three dots (…) before the type

This three dots reminds and does the exact same work of ARRAY DESTRUCTURING in JS/React
*/

// You must put three dots (…) after the variable or slice literal. If you do not, it is a compile-time
// error.

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func main() {
	res, rem, err := divAndRemainderNamed(20, 5)
	if err == nil {
		fmt.Printf("\n Result : %d , Remainder : %d \n", res, rem)
	} else {
		fmt.Println(err)
	}

	fmt.Println(addTo(3))             // No variadic params
	fmt.Println(addTo(3, 2))          // a single element as variadic param
	fmt.Println(addTo(3, 2, 4, 6, 8)) // A bunch of elements as variadic param
	a := []int{4, 3}
	fmt.Println(addTo(3, a...))                    // DE-STRUCTURING a slice as variadic parameter
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // sending a slice directly as variadic parameter
}
