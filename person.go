package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFristName string) {
	(*pointerToPerson).firstName = newFristName
}
