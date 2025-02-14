package main

import "fmt"

/*
A complex program showing capabilities of break and continue , Decode it by yourself :)
*/

func main() {
	var k int = 0
	mylist := []int{12, 14, 16, 11}
	for i := 0; i < len(mylist); i++ {
		if k == len(mylist) {
			break
		}
		fmt.Printf("Now i : %d , k : %d\t", i, k)
		fmt.Println(mylist[i])
		if mylist[i]%2 == 0 {
			mylist = append(mylist, mylist[k])
			k += 1
			continue
		}
	}
}
