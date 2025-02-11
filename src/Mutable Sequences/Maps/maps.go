package main

import "fmt"

/*
Slices are useful when you have sequential data . Those datas have no individual properties and they are sorted by some default order.
But when in a sequence , each element has some individual worth/propeties/values then we need some DS for it
Maps are here to help you in that situation.

Map is a mutable DS like slices.

The map type is written as - map[keyType]valueType.

Zero value for a map is - Nil. If we create a nilmap , we can't modify it.

Features of Map -
1. Maps automatically grow as you add key-value pairs to them .
2. If you know how many key-value pairs you plan to insert into a map, you can use make to
create a map with a specific initial size.
3. Passing a map to the len function tells you the number of key-value pairs in a map.

4. Maps are NOT COMPARABLE. You can check if they are equal to nil, but you cannot check if
two maps have identical keys and values using == or differ using !=.


*/

var nilMap map[string]int

var ages = make(map[int][]string, 10)

func main() {
	// nilMap["a"] = 1  // This map is of length 0 , FOREVER
	//fmt.Printf("Value of Key %v : %v\n", "a", nilMap["a"]) // Cannot assign values to a 'Nil' or 'Empty' Map

	totalWins := map[string]int{} // This is empty map Literal but not same as Nil-Map. R/W can be done with this declarations
	totalWins["a"] = 1            // Inserting INT value to a STR key
	fmt.Printf("Value of Key %v : %v\n", "a", totalWins["a"])

	// Otherwise we can set key-value pairs by this way. Each key-value pair is seperated by comma .
	// This Map tells - For each STR key there will a SLICE-OF-STR values
	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}

	fmt.Printf("\nTeams : %v\n", teams)

	// This Map tells - For each STR key there will a STR value
	userdetails := map[string]string{
		"Name": "Parthib",
		"ID":   "1234",
	}

	fmt.Printf("\nUserDetails : %v\n", userdetails)

	// This Map tells that ANY type of key can be assigned to ANY type of value
	anymap := map[any]any{
		1:      "a",
		"ancd": 1234,
	}

	fmt.Printf("\nAnyMap : %v\n", anymap)

	// Delete from Map - Deletion is done by mentioning the key which removes the whole key-value pair
	delete(userdetails, "Name")
	fmt.Printf("\nNow UserDetails : %v\n", userdetails)
}
