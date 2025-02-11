package main

import "fmt"

/*
3 functions are there to take user input -

1. Scan() - To take input of a single variable and store it by ACCESSING ITS ADDRESS
2. Scanln() - To take input of several variables and store them by ACCESSING THEIR ADDRESSES
3. Scanf() - To take input of several variables and store by ACCESS SPECIFIER and MEMEORY ADDRESS
*/

func main() {
	var p int
	fmt.Print("Enter number: ")
	fmt.Scan(&p)
	var s string
	fmt.Println("Enter text: ")
	fmt.Scan(&s)
	fmt.Printf("User Given Integer : %d\n", p)
	fmt.Printf("User Given String : %s\n", s)

	var a, b float32
	var bvar bool

	fmt.Print("Enter float and bool values: ")
	fmt.Scanln(&a, &b, &bvar)
	fmt.Printf("Float Values : %v %v , bool values: %v \n", a, b, bvar)

	var bytevar byte
	fmt.Print("Enter byte number: ")
	fmt.Scanf("%b", &bytevar)
	fmt.Printf("Byte Values : %v\n", bytevar) // boolean to deciman conversion
	fmt.Print("Enter character: ")
	fmt.Scanf("%c", &bytevar)
	fmt.Printf("Byte Values : %d\n", bytevar) // ASCII Of a character
}
