package main

import (
	"fmt"
)

func mapsExample() {

	// Declaration
	// var person map[string]string
	// OR
	// var mE := make(map[string]string)
	// fmt.Println(person)

	person := map[string]string{
		"name": "Ankit",
		"city": "Ranchi",
	}
	person["country"] = "IN"

	// Delete a key
	// delete(map, key)
	delete(person, "name")

	printMap(person)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Printf("%v: %v\n", key, val)
	}
}
