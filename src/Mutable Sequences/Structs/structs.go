package main

import "fmt"

/*
Features of Struct (From Limitations of Map) -
1. They can't define an API since there’s no way to constrain a map to only allow certain keys
2. All of those values in a map must be of the same type.

So, When you have related data that you want to group together,
you should define a struct.

So , Using Struct , we can create a composite data-type.

A zero value struct has every field set to the field’s zero value.
Unlike maps, there is no difference between assigning an empty struct literal and not assigning a value
at all. In both scenarios , all of the fields in the struct to their zero values.

*/

type person struct {
	name string
	age  int
	pet  string
}

var fred person // Null Struct

func main() {

	// By using key

	bob := person{}
	bob.name = "Bob"
	bob.age = 47
	bob.pet = "cat"

	fmt.Printf("Value of person : %v\n", bob)

	// Declaration using struct literal form

	julia := person{
		"Julia",
		40,
		"cat",
	}

	fmt.Printf("Value of person : %v\n", julia)

	beth := person{
		age:  30,
		name: "Beth",
	}

	fmt.Printf("Value of person : %v\n", beth)

	// Anonymous Structs

	// using var

	var person1 struct {
		name string
		age  int
		pet  string
	}

	person1.name = "Alice"
	person1.age = 32
	person1.pet = "dog"
	fmt.Printf("(Anonymous Struct) Value of person : %v\n", person1)

	//  using walrus

	pet1 := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Printf("Pet : %v\n", pet1)

	// Comparison , Comparing and Typecasting in structs are not allowed,
	// However , we can Compare a 'Named' Struct to an 'Anonymous' Struct , IFF their fields are exactly SAME (Order , Number , Types)

	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)

}
