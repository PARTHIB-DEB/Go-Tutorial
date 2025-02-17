package main

import "fmt"

/*
Although Go doesn't have any direct implementation of Named and Optional Parameters
We can use 'struct' for this purpose.
*/

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func tellage(data MyFuncOpts) string {
	res := "\n Fname : %s , Lname : %s , Age : %d"
	return res
}

func main() {
	obj := MyFuncOpts{FirstName: "Amal", LastName: "Saha"} // Age is optional here
	res := tellage(obj)
	fmt.Printf(res, obj.FirstName, obj.LastName, obj.Age)
	fmt.Println()
}
