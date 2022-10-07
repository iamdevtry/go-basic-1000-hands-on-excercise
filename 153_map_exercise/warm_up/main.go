package main

import "fmt"

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number

	phoneNumbersBLastNames := map[string]string{
		"Smith": "555-1234",
		"Jones": "555-5678",
		"Brown": "555-9012",
	}

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	productAvailability := map[int]bool{
		1: true,
		2: false,
		3: true,
	}

	// #3
	// Key        : Last name
	// Element    : Phone numbers

	phoneNumbersBLastNames2 := map[string][]string{
		"Smith": {"555-1234", "555-5678"},
		"Jones": {"555-9012", "555-3456"},
		"Brown": {"555-7890", "555-4321"},
	}

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	shopperCart := map[int]map[int]int{
		1: {
			1: 2,
			2: 1,
		},
		2: {
			1: 1,
			3: 3,
		},
	}

	fmt.Printf("phones     : %#v\n", phoneNumbersBLastNames)
	fmt.Printf("products   : %#v\n", productAvailability)
	fmt.Printf("multiPhones: %#v\n", phoneNumbersBLastNames2)
	fmt.Printf("basket     : %#v\n", shopperCart)
}
