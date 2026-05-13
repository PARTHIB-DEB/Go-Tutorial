package main

import (
	"fmt"
	"math/rand"
)

/*
Like many other languages , If-else is also a form of conditional logic building in Go

Structure -
if(expression){
	// Logic
}else{
	// Logic
}
*/

func main() {
	n := rand.Intn(10)
	if n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	
	// Special Scoping in If-Else

func main(){
    var k int;
    fmt.Print("Enter number: ")
    fmt.Scan(&k);
	if k == 0{
		fmt.Println("0 is not allowed as a divisor");
		return;
	}
    if n := rand.Intn(10); k == n % k && k <= n { // n is a temporary used variable , scope limited to if-else ladder
        fmt.Println("That's too low")
    } else if k > n  {
        fmt.Println("That's too big:", n)
    } else {
        fmt.Println("That's a good number:", n)
    }
    // fmt.Print(n); // Undefined number 'n'
}
