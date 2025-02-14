package main

import "fmt"

/*
Any go script has a single major endpoint - main()
This function holds every logic of that script.
Yet we can create our own functions .

A prototype of a function would be look like -

func func-name(params type) return-type{body}
*/

func div(numerator int, denominator int) float64 {
	if denominator == 0 {
		return 0
	}
	return float64(numerator) / float64(denominator)
}

// Multiple Return Values
// In go , a function can return multiple values. Its helpful in designing APIs where function can even throw error if required inputs have not come

func anything(x int, y int, z string) (int, string) {

	if z == "add" {
		return x + y, "nil"
	} else if z == "subtract" {
		return x - y, "nil"
	} else {
		return 0, "impossible"
	}
}

func main() {
	var i, j int
	fmt.Print("Numerator & Denominator : ")
	fmt.Scanln(&i, &j)
	fmt.Printf("%d / %d = %.3f\n", i, j, div(i, j))

	var x, y int
	var z string
	fmt.Print("2 Numbers : ")
	fmt.Scanln(&x, &y)
	fmt.Print("Command : ")
	fmt.Scanln(&z)
	res, restr := anything(x, y, z)
	fmt.Printf("x:%d , y:%d , z:%s , result :%d , message:%s \n", x, y, z, res, restr)
}
