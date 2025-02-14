package main

import "fmt"

func main() {

	// This is a list of possible KEYS
	// It's an iterable sequence , So we need a FOR LOOP
	// We are fetching each value from the sequence
	// Then every word's length is being measured to create cases
	// Then based on each length , a suitable case is activated to give proper output
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	// In Go , 'break' in switch-cases is not necessary , based on result only one case got fired

	var n int
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &n)
	switch k := n % 2; {
	case k == 0:
		fmt.Printf("%d is even\n", n)
		// break
	case k > 0:
		fmt.Printf("%d is odd\n", n)
		//break
	default:
		fmt.Printf("HOHO")
	}

}
