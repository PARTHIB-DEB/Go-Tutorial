package main

/*
Go is a call by value language .
It means that when you supply a variable
for a parameter to a function, Go always makes a copy of the value of the variable.

When you pass a struct as a function parameter, the entire struct is copied. This means that a new,
independent copy of the struct is created.
To modify the original struct within a function, you can pass a pointer to the struct.
When you pass a pointer, Go still copies the pointer value. However, the pointer value itself holds
the memory address of the original struct.

For the map, it’s easy to explain what happens: any changes made to a map parameter are
reflected in the variable passed into the function. For a slice, it’s more complicated. You can modify
any element in the slice, but you can’t lengthen the slice. This is true for maps and slices that are
passed directly into functions as well as map and slice fields in structs.

*/

import "fmt"

type person struct {
	age  int
	name string
}

func modMap(m map[int]string) {
	// It’s because maps and slices are both implemented with pointers.
	// Their address are COPIED
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}
func modSlice(s []int) []int {
	// It’s because maps and slices are both implemented with pointers.
	// Their address are COPIED
	for k, v := range s {
		s[k] = v * 2 // value changes by
	}
	s = append(s, 10)
	return s
}

func modifyFails(i int, s string, p person) {
	i = i * 2      // i is taking a copy of the value and modifying it
	s = "Goodbye"  // s is taking a copy of the value and modifying it
	p.name = "Bob" // p is taking a copy of the STRUCT value and modifying within that memory
}

func modifyStructbyAddress(p *person) {
	/*
		What's Happening - A copy of the ADDRESS of the struct is made
		But ITS ADDRESS :)
		SO changes are happening
	*/
	p.name = "Alice"
}
func main() {
	p := person{}
	i := 2
	s := "Hello"
	fmt.Println(i, s, p)
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
	q := person{}
	modifyStructbyAddress(&q)
	fmt.Println(q)

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)
	fmt.Println(modSlice([]int{1, 2, 3}))
}
