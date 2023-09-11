package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Name      string `json:"Name"`
	MovieType string `json:"MovieType"`
	Length    int    `json:"Length"`
}

func unpackJson() {
	myJson := `
[
	{
		"Name" : "Json Born",
		"MovieType" : "Action",
		"Length" : 68
	},
	{
		"Name" : "Barbie",
		"MovieType" : "Comedy",
		"Length" : 120
	}
]`
	var unmarshalled []Movie

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	var mySlice []Movie
	var marshall1 Movie
	marshall1.Name = "Private Ryan"
	marshall1.MovieType = "War action"
	marshall1.Length = 230

	mySlice = append(mySlice, marshall1)

	newJson, err := json.Marshal(mySlice)
	if err != nil {
		log.Println("Error marshalling json", err)
	}
	fmt.Println(string(newJson))
}
