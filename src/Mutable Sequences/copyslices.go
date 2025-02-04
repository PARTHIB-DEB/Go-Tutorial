package main

/*
We can perform both shallow and deep copy with slices .
Shallow copy refers to a copy of a slice with same memory address . This causes to reflect changes in one slice to another
Deep copy refers to a copy of a slice with different memory address .  This does not change one slice's identity by other slice (containing same data)
*/

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 15}
	y := make([]int, len(x))
	num := copy(y, x) // make deepcopy and return No of values copied
	k := x[:]
	fmt.Printf("XValue : %v , XHex-Id : %v\n", x, &x[0])
	fmt.Printf("YValue : %v , YHex-Id : %v\n", y, &y[0])
	fmt.Printf("KValue : %v , KHex-Id : %v\n", k, &k[0]) // same memory address with X
	fmt.Printf("NUMValue(intermediatory memory) : %v , NUMHex-Id : %v\n", num, &num)
}
