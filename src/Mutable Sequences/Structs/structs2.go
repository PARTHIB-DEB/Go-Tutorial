package main

import "fmt"

/*
Comparison between structs depends on its fields.
Structs that are entirely composed of same types are comparable;
(those with slice or map fields are not comparable)

# In go , there is NO MAGIC METHODS for STRUCTS

Go doesn’t allow comparisons between variables that represent STRUCTS OF DIFFERENT TYPES
(exception - anonymous structs)

Go allows you to perform a type conversion from one struct type to another if the fields of both
structs have the same names, order, and types .
*/
type firstPerson struct {
	name string
	age  int
}
type secondPerson struct {
	name string
	age  int
}

func main() {

	mahup := firstPerson{name: "mahup", age: 12}
	fmt.Println(mahup)

	pohup := secondPerson{name: "pohup", age: 22}
	fmt.Println(pohup)

	//fmt.Println(mahup == pohup) mohup and pohup are differnt types yet their fields are same

	pihup := firstPerson(pohup)
	pihup.name = "pohup to pihup"
	fmt.Println(pihup) // Typecasted

	tohup := secondPerson(mahup)
	tohup.name = "mahup to tohup"
	fmt.Println(tohup) // Typecasted

}
