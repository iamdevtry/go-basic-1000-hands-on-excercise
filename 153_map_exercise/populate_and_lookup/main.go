package main

import "fmt"

func main() {
	phones := map[string]string{
		"John": "555-1234",
		"Mary": "555-5678",
		"Bob":  "555-4321",
	}

	products := map[int]bool{
		617841573: true,
		617841574: false,
		617841575: true,
	}

	multiPhones := map[string][]string{
		"John": {"555-1234", "555-5678"},
		"Mary": {"555-9012", "555-3456"},
		"Bob":  {"555-7890", "555-4321"},
	}

	basket := map[int]map[int]int{
		100: {
			617841573: 2, 577841574: 1,
		},
		101: {
			617841573: 1, 617841575: 3,
		},
		102: {},
	}

	//Print Mary's phone number
	who, phone := "May", "N/A"
	if v, ok := phones[who]; ok {
		phone = v
	}
	fmt.Printf("%s's phone number is %s\n", who, phone)

	//Is product 617841574 available?
	id, status := 617841574, "available"
	if !products[id] {
		status = "not " + status
	}
	fmt.Printf("Product ID %d is %s\n", id, status)

	// What is Greco's second phone number?
	who, phone = "Mary", "N/A"
	if phones := multiPhones[who]; len(phones) > 2 {
		phone = phones[1]
	}
	fmt.Printf("%s's 2nd phone number: %s\n", who, phone)

	// How many of 617841575 the customer 101 is going to buy?
	cid, pid := 101, 617841575
	fmt.Printf("Customer #%d is going to buy %d from Product ID #%d.\n", cid, basket[cid][pid], pid)
}
