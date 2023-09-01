package main

import "fmt"

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 5001,
		},
	}
	alex.print()
	fmt.Println()
	alex.firstName = "Alexy"
	alex.lastName = "Warlord"
	alex.contact.email = "warlord@gmail.com"
	alex.contact.zipCode = 2341
	alex.print()
	fmt.Println()
	alex.updateName("Tomasz")
	alex.print()

	mapOfColors := createMapOfColors()
	printMap(mapOfColors)
}
