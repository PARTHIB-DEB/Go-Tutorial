package main

import "fmt"

/*
Just like custom datatype , we create custom Function Type

Treat it like a Typecasting Function , where you manually design the description which will translate
datas of certain existing datatype to a new datatype

What’s the advantage of declaring a function type? One use is documentation. It’s use‐
ful to give something a name if you are going to refer to it multiple times.

The syntax of a FUNCTION TYPE is -
type ftype func(arg dtype) dtype

There is no function description for a FUNCTION TYPE . Only Declaration
*/
var res string

type toNewDatatype func(arg string) string

func showstatus(st string) {
	if st == "employee" {
		res = "IT"
	} else {
		res = "ADMIN"
	}

	data := toNewDatatype(func(arg string) string {
		return res
	})
	fmt.Println(data("")) // Call the function with an argument (even if unused)
}

func main() {

	showstatus("employee")

}
