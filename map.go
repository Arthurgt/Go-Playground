package main

import "fmt"

func createMapOfColors() map[string]string {
	fmt.Println("Creating map of colors")
	mapOfColors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"green": "#00FF00",
	}
	return mapOfColors
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("Color %s has hex value %s\n", key, value)
	}
}
