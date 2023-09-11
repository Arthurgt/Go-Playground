package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

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

	sergio := spanishBot{}
	marcus := englishBot{}

	printGreeting(sergio)
	printGreeting(marcus)

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)

	square := square{sideLength: 20}
	triangle := triangle{height: 10, base: 12}
	printArea(square)
	printArea(triangle)

	// makeChannels()
	unpackJson()
	divider()
}
