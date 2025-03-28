package main

import "fmt"

func changeNums(num1 int, num2 int) {
	var temp int = 1000
	num1 = temp
	fmt.Printf("(changeNums function) a : %d , b : %d\n", num1, num2)
}

func changeNumsbyPtr(num1 *int, num2 int) {
	var temp int = 1000
	*num1 = temp
	// The copy of the address of a is stored in 'num1 *int' . but its a copy and the variable which is holding
	// the address of the pointer variable 'num1 *int' , is '&num1'
	fmt.Printf("Address of a OR value of the pointer variable (changeNumbyPtr) : %x\n", num1)
	fmt.Printf("Address of the pointer variable which holds the address of a (changeNumbyPtr) : %x\n", &num1)
	fmt.Printf("(changeNumsbyPtr function) a : %d , b : %d\n", *num1, num2)
}

func main() {
	a := 4
	b := 5
	fmt.Printf("Original a : %d , b : %d\n", a, b)
	changeNums(a, b)
	fmt.Printf("After changeNums , a : %d , b : %d\n", a, b)

	fmt.Printf("Original Address of a : %x\n", &a)
	var ad *int = &a
	changeNumsbyPtr(&a, b)
	fmt.Printf("Original Address of the pointer variable '&a' : %x\n", &ad)
	fmt.Printf("After changeNumsbyPtr , a : %d , b : %d\n", a, b)
}
