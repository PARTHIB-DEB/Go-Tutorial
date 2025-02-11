package main

import "fmt"

/*
Comma,Ok Medium -
As we’ve seen, a map returns the zero value if you ask for the value associated with a key that’s not in
the map. Go provides the comma ok idiom to tell the difference between a key that’s associated with a
zero value and a key that’s not in the map.

Rather than assign the result of a map read to a single variable, with the comma ok idiom you assign
the results of a map read to two variables. The first gets the value associated with the key. The second
value returned is a bool.

*/

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, flag := m["hello"] // to check whether the key exists by accessing its value
	fmt.Println(v, flag)
	v, flag = m["world"] // to check whether the key exists by accessing its value
	fmt.Println(v, flag)
	v, flag = m["goodbye"] // to check whether the key exists by accessing its value
	fmt.Println(v, flag)

	delete(m, "hello")

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Printf("Len of Vals : %v and Len of intSet : %v\n", len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}
